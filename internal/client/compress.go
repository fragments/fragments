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
	tarf := tar.NewWriter(gzf)

	for _, f := range files {
		err := t.addFile(tarf, f)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to add %s", f)
		}
	}

	if err := tarf.Close(); err != nil {
		return nil, err
	}
	if err := gzf.Close(); err != nil {
		return nil, err
	}

	return io.Reader(buffer), nil
}

func (t *targz) addFile(w *tar.Writer, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return errors.Wrap(err, "unable to create tar header")
	}

	if err = w.WriteHeader(header); err != nil {
		return errors.Wrap(err, "unable to write header")
	}

	if _, err := io.Copy(w, file); err != nil {
		return errors.Wrap(err, "unable to copy data to tarball")
	}
	if err := file.Close(); err != nil {
		return errors.Wrap(err, "could not close file")
	}

	return nil
}
