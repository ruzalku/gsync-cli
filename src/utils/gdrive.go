package utils

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"google.golang.org/api/drive/v3"
)


func getMimeType(path string) string {
	mimeTypeDict := map[string]string{
		".docx": "application/vnd.google-apps.document",
		".doc": "application/vnd.google-apps.document",
		".xlsx": "application/vnd.google-apps.spreadsheet",
		".xls": "application/vnd.google-apps.spreadsheet",
		".csv": "application/vnd.google-apps.spreadsheet",
		".ods": "application/vnd.google-apps.spreadsheet",
		".xlsm" : "application/vnd.google-apps.spreadsheet",
		".ppt": "application/vnd.google-apps.presentation",
		".pptx": "application/vnd.google-apps.presentation",
	}

	formatFile := filepath.Ext(path)

	mimeType, ok := mimeTypeDict[formatFile]

	if !ok {
		return ""
	}

	return mimeType
}

func NormalizePath(p string) (string, error) {
	p = path.Clean(p)

	if runtime.GOOS != "windows" {
		p = strings.ReplaceAll(p, `\`, "")
	}

	if !strings.HasPrefix(p, "~") {
			return p, nil
		}

	home, err := os.UserHomeDir()
	if err != nil {
		return p, err
	}

	if p == "~" {
		return home, nil
	}

	if strings.HasPrefix(p, "~/") || strings.HasPrefix(p, "~\\") {
		return filepath.Join(home, p[2:]), nil
	}

	return p, nil
}

func (f GFile) SaveFile (service *drive.Service) error {
	file, err := os.Open(f.Path)
	
	if err != nil {
		return err
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fileInfo.Size() == 0 {
		return errors.New("This file empty or corrupted, please add data to file")
	}

	mimeType := getMimeType(f.Path)

	fileName := filepath.Base(f.Path)
	
	
	fileMetadata := &drive.File{
		Name: fileName,
		MimeType: mimeType,
	}

	if mimeType == "" {
		fileMetadata = &drive.File{
			Name: fileName,
		}

	}

	has, err := f.ExistingInRegistry()
	if has {
		_, err = service.Files.Update(f.Id, fileMetadata).Media(file).Do()
		return err
	}
	driveFile, err := service.Files.Create(fileMetadata).Media(file).Do()
	if err != nil {
		return err
	}
	f.Id = driveFile.Id

	err = f.AddToRegistry()

	return err
}
