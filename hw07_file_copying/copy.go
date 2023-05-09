package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrCopyFile              = errors.New("smth wrong with file copying")
	ErrSetOffset             = errors.New("failed to set offset")
	ErrPathCmp               = errors.New("input path equal to output path")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromPath, err := filepath.Abs(fromPath)
	if err != nil {
		return err
	}

	toPath, err = filepath.Abs(toPath)
	if err != nil {
		return err
	}

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
			fmt.Printf("Failed to close input file: %s\n", err)
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
			fmt.Printf("Failed to close output file: %s\n", err)
		}
	}()

	progressBar := pb.Full.Start64(limit)
	barReader := progressBar.NewProxyReader(inFile)
	_, err = io.CopyN(outFile, barReader, limit)
	progressBar.Finish()

	return err
}
