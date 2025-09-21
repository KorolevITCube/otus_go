package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3" //nolint:depguard
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, err := os.OpenFile(fromPath, os.O_RDONLY, 0o644)
	if err != nil {
		return fmt.Errorf("error opening file %w", err)
	}

	toFile, err := os.OpenFile(toPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	fileInfo, err := fromFile.Stat()
	if err != nil {
		return fmt.Errorf("error getting file information: %w", err)
	}

	log.Printf("File size %d\n", fileInfo.Size())

	if limit == 0 {
		limit = fileInfo.Size()
	}

	log.Printf("Offset %d, limit %d\n", offset, limit)

	if offset > fileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}

	fromFile.Seek(offset, io.SeekStart)

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(fromFile)

	n, err := io.CopyN(toFile, barReader, limit)
	log.Printf("Copyed bytes %d\n", n)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	bar.Finish()
	return nil
}
