package finder

import (
	"github.com/Terif1/config-file-validator/pkg/filetype"
)

// The File Metadata object stores the
// name and the path of the file and the type
// of file that it is, example: json, yml, etc
type FileMetadata struct {
	Name     string
	Path     string
	FileType filetype.FileType
}

// FileFinder is the interface that wraps the Find method

// Find will return an array of FileMetadata objects from
// a provided path and array of FileTypes. Any files in
// subdirectories defined in excludeDirs will not be returned
type FileFinder interface {
	Find() ([]FileMetadata, error)
}
