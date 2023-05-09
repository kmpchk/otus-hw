package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("offset is larger than file size", func(t *testing.T) {
		// create temp file
		content := []byte("TestFile")
		tmpfile, err := os.CreateTemp("", "test_file.bin")
		if err != nil {
			require.NoError(t, err)
		}

		// close temp file at the end of the test
		defer tmpfile.Close()
		// close temp file at the end of the test
		defer os.Remove(tmpfile.Name())

		// write temp file content
		if _, err := tmpfile.Write(content); err != nil {
			require.NoError(t, err)
		}

		// do copying
		err = Copy(tmpfile.Name(), "/tmp/test_file.bin", 10000, 100)
		if err != nil {
			log.Println(err)
		}
		require.EqualError(t, err, "offset exceeds file size")
	})

	t.Run("unsupported source", func(t *testing.T) {
		// do copying
		err := Copy("/dev/urandom", "/tmp/test_file.bin", 10000, 100)
		if err != nil {
			log.Println(err)
		}
		require.EqualError(t, err, ErrUnsupportedFile.Error())
	})

	t.Run("input path equal to output path", func(t *testing.T) {
		// do copying
		err := Copy("testdata/input.txt", "testdata/input.txt", 0, 100000)
		if err != nil {
			log.Println(err)
		}
		require.EqualError(t, err, ErrPathCmp.Error())
	})

	t.Run("input path equal to relative output path", func(t *testing.T) {
		// do copying
		err := Copy("testdata/input.txt", "testdata/../testdata/input.txt", 0, 100000)
		if err != nil {
			log.Println(err)
		}
		require.EqualError(t, err, ErrPathCmp.Error())
	})

	t.Run("positive full file copying", func(t *testing.T) {
		// do copying
		err := Copy("testdata/input.txt", "testdata/test_offset0_limit0.txt", 0, 0)
		require.NoError(t, err)
		expectedFile, err := os.ReadFile("testdata/test_offset0_limit0.txt")
		require.NoError(t, err)
		actualFile, err := os.ReadFile("testdata/out_offset0_limit0.txt")
		require.NoError(t, err)
		require.Equal(t, expectedFile, actualFile)
	})

	t.Run("positive partitial file copying", func(t *testing.T) {
		// do copying
		err := Copy("testdata/input.txt", "testdata/test_offset6000_limit1000.txt", 6000, 1000)
		require.NoError(t, err)
		expectedFile, err := os.ReadFile("testdata/test_offset6000_limit1000.txt")
		require.NoError(t, err)
		actualFile, err := os.ReadFile("testdata/out_offset6000_limit1000.txt")
		require.NoError(t, err)
		require.Equal(t, expectedFile, actualFile)
	})
}
