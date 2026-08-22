package utils

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type GFile struct {
	Id string `json:"id"`
	Path string `json:"path"`
	Hash string `json:"hash"`
	LastSave time.Time `json:"last_save"`
	Autosave bool `json:"autosave"`
}


// Get the path to the registry file
func getRegistryPath() string {
	builder := &strings.Builder{}

	rootPath, _ := os.UserConfigDir()

	builder.WriteString(rootPath)
	builder.WriteRune(os.PathSeparator)
	builder.WriteString("gsynccli")
	builder.WriteRune(os.PathSeparator)
	builder.WriteString("gsync_reg.json")

	return builder.String()
}

// Create registry file
func createRegistryFile(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}

	dir := filepath.Dir(path)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}


// Add an entry to the registry file
func (f GFile) AddToRegistry() error {
	regPath := getRegistryPath()
	err := createRegistryFile(regPath)

	if err != nil {
		return err
	}

	file, err := os.OpenFile(
		regPath,
		os.O_RDWR|os.O_CREATE,
		0666,
	)

	if err != nil {
		return nil
	}
	defer file.Close()


	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if content[0] != 91 || content[len(content) - 1] != 93 {
		return errors.New("Error with json registry, its must be array")
	}
	if !json.Valid(content) {
		return errors.New("Error with syntax in json registry")
	}

	byteCode, err := json.Marshal(f)
	if err != nil {
		return err
	}
	byteCode = append(byteCode, 93)

	if len(content) == 0 {
		addByteCode := []byte("[")
		addByteCode = append(addByteCode, byteCode...)

		_, err = file.Write(addByteCode)
		
		return err
	} else if len(content) == 2 {
		_, err = file.Seek(-1, io.SeekEnd)
		if err != nil {
			return err
		}

		_, err = file.Write(byteCode)
		
		return err
	}

	_, err = file.Seek(-1, io.SeekEnd)
	if err != nil {
		return err
	}

	addByteCode := []byte(", ")
	addByteCode = append(addByteCode, byteCode...)

	_, err = file.Write(addByteCode)
	
	return err
}

// Update the entry in the registry file
func (f GFile) UpdateFileRegistry() error {
	return nil
}


// Check the file of existing in registry
func (f GFile) CheckInRegistry() (bool, error) {
	return true, nil
}
