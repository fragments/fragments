package client

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"

	"github.com/pkg/errors"
)

type targz struct{}

// Compress compresses a list of files to a tar.gz archive
func Compress(files []string) (io.Reader, error) {
	if len(files) == 0 {
		return nil, errors.New("no files specified")
	}

	archive := &targz{}
	return archive.Compress(files)
}

func (t *targz) Compress(files []string) (io.Reader, error) {
	buffer := &bytes.Buffer{}

	gzf := gzip.NewWriter(buffer)
	defer func() {
		_ = gzf.Close()
	}()

	tarf := tar.NewWriter(gzf)
	defer func() {
		_ = tarf.Close()
	}()

	for _, f := range files {
		err := t.addFile(tarf, f)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to add %s", f)
		}
	}

	return io.Reader(buffer), nil
}

func (t *targz) addFile(w *tar.Writer, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return errors.Wrap(err, "unable to create tar header")
	}

	err = w.WriteHeader(header)
	if err != nil {
		return errors.Wrap(err, "unable to write header")
	}

	_, err = io.Copy(w, file)
	if err != nil {
		return errors.Wrap(err, "unable to copy data to tarball")
	}

	return nil
}
