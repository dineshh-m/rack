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

func copyFile(srcPath, destFilePath string) error {
  srcFile, pErr := os.Open(srcPath)
  if pErr != nil {
    return errors.New("Error opening the file" + srcPath)
  }
  // Closing the source file
  defer srcFile.Close()

  // Creating the destination directory if it doesn't exists
  if err := os.MkdirAll(filepath.Dir(destFilePath), 0755); err != nil {
    return errors.New("Cannot create the directory at" + filepath.Dir(destFilePath))
  }

  // creating the destination file 
  destFile, err := os.Create(destFilePath)
  if err != nil {
    return errors.New("Cannot the create file " + srcFile.Name() + " at " + filepath.Dir(destFilePath))
  }
  defer destFile.Close()

  _, copyErr := io.Copy(destFile, srcFile)


  if copyErr != nil {
    fmt.Println("Copy error", copyErr.Error())
    return errors.New("Cannot copy file " + srcFile.Name())
  }

  return nil
}

func moveFile(srcFilePath, destFilePath string) error {
  destDirectory := filepath.Dir(destFilePath)
  if err := os.MkdirAll(destDirectory, 644); err != nil {
    return fmt.Errorf("Cannot create the destination directory: %s", destDirectory)
  }
  err := os.Rename(srcFilePath, destFilePath)
  if err != nil {
    return fmt.Errorf("Cannot move file: %s to %s", srcFilePath, destFilePath)
  }
  return nil
}

func CopyFiles(sourceDir, destDir, fileType string, moveEnabled bool) error {
  files, err := ioutil.ReadDir(sourceDir)
  if err != nil {
    return err
  }

  for _, file := range files {
    if !file.IsDir() {
      filePath := filepath.Join(sourceDir, file.Name())
      destPath := filepath.Join(destDir, file.Name())
      fileExtension := getFileType(file.Name())
      if fileExtension != "" {
        if fileExtension == fileType {
          var err error
          if moveEnabled {
            err = moveFile(filePath, destPath)
          } else {
            err = copyFile(filePath, destPath)
          }
          if err != nil {
            fmt.Fprintf(os.Stderr, err.Error())
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
    fileType := split[length - 1]
    return fileType
  }

  return ""
}

func OrganizeFiles(sourceDir, destDir string, moveEnabled bool) error {
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
        var err error
        if moveEnabled {
          err = moveFile(srcPath, destPath)
        } else {
          err = copyFile(srcPath, destPath)
        }
        if err != nil {
          fmt.Fprintf(os.Stderr, err.Error())
        }
      }
    } 
  }

  return nil
}

