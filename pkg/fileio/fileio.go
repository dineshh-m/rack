package fileio

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func IsDirExists(dirPath string) error {
  if _, err := os.Stat(dirPath); os.IsNotExist(err) {
    return errors.New("The specified source directory doesn't exists")    
  }

  return nil
}

func copyFile(srcPath, destPath string) error {
  srcFile, pErr := os.Open(srcPath)
  if pErr != nil {
    return errors.New("Error opening the file" + srcPath)
  }
  // Closing the source file
  defer srcFile.Close()

  // Creating the destination directory if it doesn't exists
  if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
    return errors.New("Cannot create the directory at" + filepath.Dir(destPath))
  }
  
  // creating the destination file 
  destFile, err := os.Create(destPath)
  if err != nil {
    return errors.New("Cannot the create file " + srcFile.Name() + " at " + filepath.Dir(destPath))
  }
  defer destFile.Close()

  _, copyErr := io.Copy(destFile, srcFile)


  if copyErr != nil {
    fmt.Println("Copy error", copyErr.Error())
    return errors.New("Cannot copy file " + srcFile.Name())
  }

  return nil
}

func CopyFiles(sourceDir, destDir, fileType string) error {
  files, err := ioutil.ReadDir(sourceDir)
  if err != nil {
    return err
  }

  for _, file := range files {
    if !file.IsDir() {
      filePath := filepath.Join(sourceDir, file.Name())
      destPath := filepath.Join(destDir, file.Name())
      ext := strings.Split(file.Name(), ".")
      if len(ext) >= 2 {
        last := len(ext) - 1
        if ext[last] == fileType {
          err := copyFile(filePath, destPath)
          if err != nil {
            fmt.Fprintln(os.Stderr, err.Error())
          }
        }
      }
    }
  }

  return nil
}

func getFileType(fileName string) string {
  split := strings.Split(fileName, ".")
  if len(split) >= 2 {
    length := len(split)
    ext := split[length - 1]
    return ext
  }

  return ""
}

func OrganizeFiles(sourceDir, destDir string) error {
  files, err := ioutil.ReadDir(sourceDir)
  if err != nil {
    return errors.New("Cannon read the directory" + sourceDir)
  }

  for _, fileInfo := range files {
    if !fileInfo.IsDir() {  
      fileType := getFileType(fileInfo.Name())
      if fileType != "" {
        srcPath := filepath.Join(sourceDir, fileInfo.Name())
        destPath := filepath.Join(destDir, fileType + "/")    
        os.Mkdir(destPath, 0644)

        destPath = filepath.Join(destPath, fileInfo.Name())
        err := copyFile(srcPath, destPath)
        if err != nil {
          fmt.Fprintln(os.Stderr, "Cannot copy file ", srcPath)
        }
      }
    } 
  }

  return nil
}
