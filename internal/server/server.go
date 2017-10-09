package server

import (
	"context"

	"github.com/fragments/fragments/internal/backend"
	"github.com/fragments/fragments/internal/filestore"
	"github.com/fragments/fragments/internal/state"
	"github.com/pkg/errors"
)

// Server is the fragments server that accepts models and keeps them in the
// store.
type Server struct {
	StateStore    backend.KV
	SecretStore   backend.KV
	SourceStore   filestore.SourceTarget
	GenerateToken func() string
}

// New creates a new server.
// Upload tokens are generated by server.GenerateToken.
func New(state backend.KV, secrets backend.KV, sourceTarget filestore.SourceTarget) *Server {
	return &Server{
		StateStore:    state,
		SecretStore:   secrets,
		SourceStore:   sourceTarget,
		GenerateToken: GenerateToken,
	}
}

// UploadRequest is a request for source code, returned from the server.
type UploadRequest struct {
	// Token is the token to use when confirming the upload.
	Token string
	// URL is where to upload the source code. The source must be uploaded as a
	// tar.gz. After completion the upload should be confirmed.
	URL string
}

// PutFunction creates or updates a function. In case the function already
// exists it is updated. If not, source upload is requested.
func (s *Server) PutFunction(ctx context.Context, input *state.Function) (*UploadRequest, error) {
	if input == nil {
		return nil, errors.New("no function supplied")
	}
	name := input.Meta.Name
	if name == "" {
		return nil, errors.New("function has no meta or name")
	}
	// Get existing function
	existing, err := state.GetFunction(ctx, s.StateStore, name)
	if err != nil {
		return nil, errors.Wrap(err, "could not check existing function")
	}

	if existing == nil || existing.Checksum != input.Checksum {
		// nolint: vetshadow
		res, err := s.requestUpload(ctx, input, existing)
		if err != nil {
			return nil, errors.Wrap(err, "could not request source upload")
		}
		return res, nil
	}

	if err = s.updateFunctionConfiguration(ctx, input); err != nil {
		return nil, errors.Wrap(err, "could not update function configuration")
	}

	return nil, nil
}

// ConfirmUpload is called by the client when the source has been uploaded
func (s *Server) ConfirmUpload(ctx context.Context, token string) error {
	if token == "" {
		return errors.New("token not set")
	}

	upload, err := state.GetPendingUpload(ctx, s.StateStore, token)
	if err != nil {
		return errors.Wrap(err, "could not get pending upload")
	}

	if upload == nil {
		return errors.New("not found")
	}

	if err := s.SourceStore.Persist(ctx, token); err != nil {
		return errors.Wrap(err, "could not persist source")
	}

	function := upload.Function
	function.SourceFilename = upload.Filename

	if err := s.updateFunctionConfiguration(ctx, function); err != nil {
		return errors.Wrap(err, "error storing function update")
	}

	if err := state.DeletePendingUpload(ctx, s.StateStore, token); err != nil {
		return errors.Wrap(err, "could not delete pending upload")
	}

	// TODO(akupila): if set; archive upload.PreviousFilename from source store
	return nil
}

// EnvironmentInput is the input to specify when creating an environment. The
// credentials are stored in the secret store so the underlying
// state.Environment does not contain these fields.
type EnvironmentInput struct {
	// Name is the name that identifies the environment.
	Name string
	// Labels are used to map a deployment to the environment.
	Labels map[string]string
	// Infrastructure is the type of infrastructure the environment is for
	Infrastructure state.InfrastructureType
	// Username is the username used to authenticate to the infrastructure
	// provider.
	Username string
	// Password is the password used to authenticate to the infrastructure
	// provider.
	Password string
	// AWS contains AWS environment specific information
	AWS *state.InfrastructureAWS
}

// CreateEnvironment creates a new target deployment environment. Returns an
// error if an environment with the same name already exists.
func (s *Server) CreateEnvironment(ctx context.Context, input *EnvironmentInput) error {
	if input == nil {
		return errors.New("no environment supplied")
	}
	name := input.Name
	if name == "" {
		return errors.New("environment has no name")
	}

	existing, err := state.GetEnvironment(ctx, s.StateStore, name)
	if err != nil {
		return errors.Wrap(err, "could not check existing environment")
	}
	if existing != nil {
		return errors.Errorf("an environment with the name %q already exists", name)
	}

	// Store environment credentials
	if err := putUserCredentials(ctx, s.SecretStore, name, input.Username, input.Password); err != nil {
		return errors.Wrap(err, "error storing credentials")
	}

	env := &state.Environment{
		Meta: state.Meta{
			Name:   name,
			Labels: input.Labels,
		},
		Infrastructure: input.Infrastructure,
		AWS:            input.AWS,
	}

	if err := state.PutModel(ctx, s.StateStore, state.ModelTypeEnvironment, env); err != nil {
		return errors.Wrap(err, "could not store environment")
	}

	return nil
}

// PutDeployment creates or updates a deployment. In case the deployment already
// exists it is updated.
func (s *Server) PutDeployment(ctx context.Context, input *state.Deployment) error {
	if input == nil {
		return errors.New("no deployment supplied")
	}
	name := input.Name()
	if name == "" {
		return errors.New("deployment has no name")
	}
	if err := state.PutModel(ctx, s.StateStore, state.ModelTypeDeployment, input); err != nil {
		return errors.Wrap(err, "could not store deployment")
	}
	return nil
}

// requestUpload creates a url the client can upload source code to. The upload
// request is stored as a PendingUpload in the store so it can be retrieved
// when the client confirms the upload.
func (s *Server) requestUpload(ctx context.Context, input *state.Function, existing *state.Function) (*UploadRequest, error) {
	token := s.GenerateToken()

	url, err := s.SourceStore.NewUploadURL(token)
	if err != nil {
		return nil, errors.New("could not create upload url")
	}

	pendingUpload := &state.PendingUpload{
		Filename: token,
		Function: input,
	}
	if existing != nil {
		pendingUpload.PreviousFilename = existing.SourceFilename
	}

	if err := state.PutPendingUpload(ctx, s.StateStore, token, pendingUpload); err != nil {
		return nil, errors.Wrap(err, "could not store pending upload in backend")
	}

	uploadRequest := &UploadRequest{
		Token: token,
		URL:   url,
	}

	return uploadRequest, nil
}

// updateFunctionConfiguration updates the function's configuration
// (performance specs, environment variables etc) without updating the source
// code.
func (s *Server) updateFunctionConfiguration(ctx context.Context, input *state.Function) error {
	if err := state.PutModel(ctx, s.StateStore, state.ModelTypeFunction, input); err != nil {
		return errors.Wrap(err, "could not update function configuration")
	}
	return nil
}
