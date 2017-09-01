package gof

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Goff is main struct type containing all the functions,
// exposed for documentation generation purposes
type Goff struct{}

// Gof is the main exposed Goff instance containing all the functions
var Gof Goff

// ErrFileNotExists raises when the file does not exists
type ErrFileNotExists struct {
	path string
}

func (e ErrFileNotExists) Error() string {
	return fmt.Sprintf("File %s does not exists", e.path)
}

// ErrFileAlreadyExists raises when the file does not exists
type ErrFileAlreadyExists struct {
	path string
}

func (e ErrDirectoryNotExists) Error() string {
	return fmt.Sprintf("Directory %s does not exists", e.path)
}

// ErrDirectoryNotExists raises when the directory does not exists
type ErrDirectoryNotExists struct {
	path string
}

func (e ErrFileAlreadyExists) Error() string {
	return fmt.Sprintf("File %s already exists", e.path)
}

// ErrDirectoryAlreadyExists raises when the directory does not exists
type ErrDirectoryAlreadyExists struct {
	path string
}

func (e ErrDirectoryAlreadyExists) Error() string {
	return fmt.Sprintf("Directory %s already exists", e.path)
}

// Read returns content of a file
func (g *Goff) Read(path string) (string, error) {
	exists, err := g.Exists(path)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", ErrFileNotExists{path}
	}

	bytesData, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	data := string(bytesData[:])
	return data, nil
}

// Write writes content to a file
func (g *Goff) Write(path, content string) error {
	exists, err := g.Exists(path)
	if err != nil {
		return err
	}
	if !exists {
		return ErrFileNotExists{path}
	}
	err = ioutil.WriteFile(path, []byte(content), 0644)
	return err
}

// Append appends content to the end of a file
func (g *Goff) Append(path, content string) error {
	data, err := g.Read(path)
	if err != nil {
		return err
	}
	data = data + content
	return g.Write(path, data)
}

// Prepend prepends content to the start of a file
func (g *Goff) Prepend(path, content string) error {
	data, err := g.Read(path)
	if err != nil {
		return err
	}
	data = content + data
	return g.Write(path, data)
}

// Exists returns true if the given path (file or directory) exists
func (g *Goff) Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

// FileExists returns true if the given file path exists
func (g *Goff) FileExists(path string) (bool, error) {
	return g.IsFile(path)
}

// DirectoryExists returns true if the path exists
func (g *Goff) DirectoryExists(path, content string) (bool, error) {
	return g.IsDirectory(path)
}

// Mkdir creates a directory
func (g *Goff) Mkdir(path string) error {
	exists, err := g.Exists(path)
	if err != nil {
		return err
	}
	if exists {
		return ErrDirectoryAlreadyExists{path}
	}

	err = os.Mkdir(path, os.ModeDir)
	return err
}

// Create creates a file
func (g *Goff) Create(path string) error {
	exists, err := g.Exists(path)
	if err != nil {
		return err
	}
	if exists {
		return ErrFileAlreadyExists{path}
	}

	file, err := os.OpenFile(path, os.O_CREATE, 0644)
	defer file.Close()
	return err
}

// Touch touches a file
func (g *Goff) Touch(path string) error {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	defer file.Close()
	return err
}

// Rm deletes a file or directory
func (g *Goff) Rm(path string) error {
	exists, err := g.Exists(path)
	if err != nil {
		return err
	}
	if !exists {
		return ErrFileNotExists{path}
	}

	return os.Remove(path)
}

// IsDirectory returns true if the path is a directory
func (g *Goff) IsDirectory(path string) (bool, error) {
	exists, err := g.Exists(path)
	if err != nil || !exists {
		return false, err
	}

	stat, err := g.Stat(path)
	if err != nil {
		return false, err
	}
	if stat.IsDir() {
		return true, nil
	}
	return false, nil
}

// IsFile returns true if the path is a file
func (g *Goff) IsFile(path string) (bool, error) {
	exists, err := g.Exists(path)
	if err != nil || !exists {
		return false, err
	}

	stat, err := g.Stat(path)
	if err != nil {
		return false, err
	}
	if stat.IsDir() {
		return false, nil
	}
	return true, nil
}

// Rename renames a file
func (g *Goff) Rename(oldPath, newPath string) error {
	oldExists, err := g.Exists(oldPath)
	if err != nil {
		return err
	}
	if !oldExists {
		return ErrFileNotExists{oldPath}
	}

	newExists, err := g.Exists(newPath)
	if err != nil {
		return err
	}
	if newExists {
		return ErrFileAlreadyExists{newPath}
	}

	return os.Rename(oldPath, newPath)
}

// Copy copies a file
func (g *Goff) Copy(sourcePath, destinationPath string) error {
	sourceExists, err := g.Exists(sourcePath)
	if err != nil {
		return err
	}
	if !sourceExists {
		return ErrFileNotExists{sourcePath}
	}

	destinationExists, err := g.Exists(destinationPath)
	if err != nil {
		return err
	}
	if destinationExists {
		return ErrFileAlreadyExists{destinationPath}
	}

	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}
	return destination.Close()
}

// Move moves a file
func (g *Goff) Move(sourcePath, destinationPath string) error {
	err := g.Copy(sourcePath, destinationPath)
	if err != nil {
		return err
	}

	return g.Rm(sourcePath)
}

// Stat returns stat of a file
func (g *Goff) Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}
