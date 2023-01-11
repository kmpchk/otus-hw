package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrCopyFile              = errors.New("smth wrong with file copying")
)

func getNewFileSize(fileSize int64, offset int64, limit int64) int64 {
	n := fileSize
	if offset > 0 {
		n -= offset
	}
	if limit > 0 && n > limit {
		n = limit
	}
	return n
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	//println("Hello from Copy!")

	inFileInfo, err := os.Stat(fromPath)
	if err != nil {
		return err
	}

	if !inFileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	inFile, err := os.Open(fromPath)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := inFile.Close(); err != nil {
			panic(err)
		}
	}()

	inFileSize := inFileInfo.Size()
	fmt.Println(inFile.Name())
	fmt.Println(inFileSize)

	if offset > inFileSize {
		panic(ErrOffsetExceedsFileSize)
	}

	if limit == 0 || limit > inFileSize {
		limit = inFileSize
	}

	outFile, err := os.Create(toPath)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := outFile.Close(); err != nil {
			panic(err)
		}
	}()

	limit = getNewFileSize(inFileSize, offset, limit)
	if offset > 0 {
		_, err = inFile.Seek(offset, io.SeekStart)
		if err != nil {
			return err
		}
	}

	//var seek io.ReadSeeker = inFile
	//seek.Seek(offset, io.SeekStart)
	buf := bufio.NewReaderSize(inFile, int(inFileSize))
	inFile.Seek(offset, io.SeekStart)

	progressBar := pb.Full.Start64(limit)
	barReader := progressBar.NewProxyReader(buf)
	_, err = io.CopyN(outFile, barReader, limit)
	if err != nil {
		panic(ErrCopyFile)
	}
	progressBar.Finish()

	return nil
}
