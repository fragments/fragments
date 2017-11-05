package filestore

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Local stores files locally on disk
type Local struct {
	UploadDirectory string
	SourceDirectory string
	httpServer      *http.Server
	address         string
}

// NewLocal creates a local filestore and starts listening on a port assigned
// by the operating system. The files are stored on local disk on the server.
// The uploadDir is used for storing uploads, files are moved to the sourceDir
// when persisted. The directories are created if they don't already exist.
func NewLocal(uploadDir, sourceDir string) (*Local, error) {
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, errors.Wrap(err, "could not make upload directory")
	}
	if err := os.MkdirAll(sourceDir, 0755); err != nil {
		return nil, errors.Wrap(err, "could not make source directory")
	}

	srv := &http.Server{Addr: "127.0.0.1:0"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimPrefix(r.URL.Path, "/")
		file, err := os.Create(fmt.Sprintf("%s/%s", uploadDir, token))
		if err != nil {
			http.Error(w, errors.Wrap(err, "could not create uploaded file").Error(), http.StatusInternalServerError)
			return
		}
		_, err = io.Copy(file, r.Body)
		if err != nil {
			http.Error(w, errors.Wrap(err, "could not save uploaded file").Error(), http.StatusInternalServerError)
			return
		}
	})

	addrc := make(chan string)
	go func() {
		lis, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			panic(err)
		}
		addrc <- lis.Addr().String()
		close(addrc)
		_ = srv.Serve(lis)
	}()

	l := &Local{
		UploadDirectory: uploadDir,
		SourceDirectory: sourceDir,
		httpServer:      srv,
		address:         <-addrc,
	}

	return l, nil
}

// NewUploadURL creates a new upload url that the local filestore will handle.
func (l *Local) NewUploadURL(name string) (string, error) {
	return fmt.Sprintf("http://%s/%s", l.address, name), nil
}

// Persist moves the file from the upload directory to the source directory.
func (l *Local) Persist(ctx context.Context, name string) error {
	from := fmt.Sprintf("%s/%s", l.UploadDirectory, name)
	to := fmt.Sprintf("%s/%s", l.SourceDirectory, name)
	if err := os.Rename(from, to); err != nil {
		return errors.Wrap(err, "could not move to source directory")
	}
	return nil
}

// GetFile returns a source file from the local filestore.
func (l *Local) GetFile(name string) (*os.File, error) {
	filename := filepath.Join(l.SourceDirectory, name)
	return os.Open(filename)
}

// Shutdown gracefully closes the local filestore. New connections are not
// accepted after Close() and existing connections are drained before shutdown.
func (l *Local) Shutdown() error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := l.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "failed to gracefully shut down server")
	}
	return nil
}
