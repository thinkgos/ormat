package utils

import (
	"os"
	"path"
	"text/template"
)

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(paths string) bool {
	_, err := os.Stat(paths)
	return err == nil || os.IsExist(err)
}

// WriteFile writes data to a file named by filename.
// If the file does not exist, WriteFile creates it
// and its upper level paths.
func WriteFile(filename string, data []byte) error {
	if err := os.MkdirAll(path.Dir(filename), os.ModePerm); err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0655)
}

// WriteFileWithTemplate write file with template.
func WriteFileWithTemplate(filename string, t *template.Template, data any) error {
	if err := os.MkdirAll(path.Dir(filename), os.ModePerm); err != nil {
		return err
	}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0655)
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, data)
}
