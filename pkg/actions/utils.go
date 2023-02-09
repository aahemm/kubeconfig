package actions

import (
	"fmt"
	"io"
	"os"
)

func areFilesEqual(fileOnePath, fileTwoPath string) bool {
	fileOneData, _ := os.ReadFile(fileOnePath)
	fileTwoData, _ := os.ReadFile(fileTwoPath)
	return string(fileOneData) == string(fileTwoData)
}

func copyFile(srcPath, dstPath string) error {
	srcFile, err := os.Open(srcPath)
    if err != nil {
        return fmt.Errorf("could not open kubeconfig file %s: %w", srcPath, err)
    }
    defer srcFile.Close()

    dstFile, err := os.Create(dstPath)
    if err != nil {
        return fmt.Errorf("[copyFile] could not create src file %s: %w", dstPath, err)
    }
    defer dstFile.Close()

    byteNum, err := io.Copy(dstFile, srcFile)
    if err != nil {
        return fmt.Errorf("[copyFile] could not write dst file %s: %w", dstPath, err)
    }
	fmt.Printf("Copied %d bytes to %s \n", byteNum, dstPath)
	return nil 
}