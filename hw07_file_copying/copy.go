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
	ErrCopyFile              = errors.New("smth wrong with file copying")
	ErrSetOffset             = errors.New("failed to set offset")
	ErrPathCmp               = errors.New("in path equal to out path")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if fromPath == toPath {
		return ErrPathCmp
	}

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

	if _, err := inFile.Seek(offset, io.SeekStart); err != nil {
		return ErrSetOffset
	}

	if limit == 0 || limit > inFileSize {
		limit = inFileSize
	}

	if limit+offset > inFileSize {
		limit = inFileSize - offset
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

	progressBar := pb.Full.Start64(limit)
	barReader := progressBar.NewProxyReader(inFile)
	_, err = io.CopyN(outFile, barReader, limit)
	progressBar.Finish()

	return err
}
