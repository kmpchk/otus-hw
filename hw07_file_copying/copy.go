package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrLimitExceedsFileSize  = errors.New("limit exceeds file size")
	ErrCopyFile              = errors.New("smth wrong with file copying")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	inFileInfo, err := os.Stat(fromPath)
	if err != nil {
		return err
	}

	if !inFileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	inFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}

	defer func() {
		if err := inFile.Close(); err != nil {
			panic(err)
		}
	}()

	inFileSize := inFileInfo.Size()

	if offset > inFileSize {
		return ErrOffsetExceedsFileSize
	}

	if offset+limit > inFileSize {
		return ErrLimitExceedsFileSize
	}

	if limit == 0 || limit > inFileSize {
		limit = inFileSize
	}

	outFile, err := os.Create(toPath)
	if err != nil {
		return err
	}

	defer func() {
		if err := outFile.Close(); err != nil {
			panic(err)
		}
	}()

	inFile.Seek(offset, io.SeekStart)

	progressBar := pb.Full.Start64(limit)
	barReader := progressBar.NewProxyReader(inFile)
	_, err = io.CopyN(outFile, barReader, limit)
	if err != nil {
		return ErrCopyFile
	}
	progressBar.Finish()

	return nil
}
