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
		tmpfile, err := os.CreateTemp("", "test_file-")
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
		err = Copy(tmpfile.Name(), "/tmp/", 10000, 100)
		if err != nil {
			log.Println(err)
		}
		require.EqualError(t, err, "offset exceeds file size")
	})
}
