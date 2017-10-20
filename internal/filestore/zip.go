package filestore

import (
	"archive/tar"
	"archive/zip"
	"io"

	"github.com/pkg/errors"
)

// Zip reads all files from a tar reader and compresses them to a zip byte
// array.
func Zip(input *tar.Reader, output io.Writer) error {
	z := zip.NewWriter(output)

	for {
		h, err := input.Next()
		if err != nil {
			if err == io.EOF {
				// Done
				break
			}
			return errors.Wrap(err, "could not read tar")
		}

		info := h.FileInfo()
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if info.IsDir() {
			header.Name += "/"
			header.Method = zip.Store
		}

		writer, err := z.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			continue
		}

		if header.Mode().IsRegular() {
			_, err := io.CopyN(writer, input, info.Size())
			if err != nil {
				return err
			}
		}
	}

	if err := z.Close(); err != nil {
		return errors.Wrap(err, "could not close zip stream")
	}

	return nil
}
