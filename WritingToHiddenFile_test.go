package main

import (
	"io/ioutil"
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

var someData = []byte("Hello, World!")

func TestWriteFile(t *testing.T) {

	var err error

	tempFile, err := ioutil.TempFile("", "write_to_hidden_file_test_")
	tempFile.Close()
	assert.NoError(t, err, "Creating a temporary file.")

	fileName := tempFile.Name()

	err = ioutil.WriteFile(fileName, someData, 0600)
	assert.NoError(t, err, "Writing to the file before it's hidden.")

	err = makeVisible(fileName, true)
	assert.NoError(t, err, "Hiding the file.")

	bytes, err := ioutil.ReadFile(fileName)
	assert.NoError(t, err, "Reading the file after it's been hidden.")
	assert.Equal(t, someData, bytes, "Read content is the same as written content.")

	err = ioutil.WriteFile(fileName, someData, 0600)
	assert.NoError(t, err, "Writing to the file after it's been hidden.")

	err = makeVisible(fileName, false)
	assert.NoError(t, err, "Unhiding the file.")

	err = ioutil.WriteFile(fileName, someData, 0600)
	assert.NoError(t, err, "Writing to the file after it's been made visible again.")

	err = makeVisible(fileName, true)
	assert.NoError(t, err, "Hiding the file.")

	err = os.Remove(fileName)
	assert.NoError(t, err, "Deleting the file after it's been hidden.")
}

func makeVisible(path string, b bool) error {
	ptr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}
	attributes, err := syscall.GetFileAttributes(ptr)
	if err != nil {
		return err
	}

	if b {
		attributes |= syscall.FILE_ATTRIBUTE_HIDDEN
	} else {
		attributes &^= syscall.FILE_ATTRIBUTE_HIDDEN
	}
	return syscall.SetFileAttributes(ptr, attributes)
}
