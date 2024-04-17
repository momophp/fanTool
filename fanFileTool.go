package fanTool

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ZipFile 压缩文件
func ZipFile(inputFile string, outputFile string) error {
	// 创建输出文件
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer func(output *os.File) { _ = output.Close() }(output)
	// 创建zip写入器
	zipWriter := zip.NewWriter(output)
	defer func(zipWriter *zip.Writer) { _ = zipWriter.Close() }(zipWriter)
	// 打开输入文件
	input, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer func(input *os.File) { _ = input.Close() }(input)
	// 将输入文件添加到zip压缩包中
	fileInfo, err := input.Stat()
	if err != nil {
		return err
	}
	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}
	header.Name = inputFile
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, input)
	if err != nil {
		return err
	}
	return nil
}

// ZipFolder 压缩文件夹
func ZipFolder(folderPath, zipPath string) error {
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer func(zipFile *os.File) { _ = zipFile.Close() }(zipFile)
	zipWriter := zip.NewWriter(zipFile)
	defer func(zipWriter *zip.Writer) { _ = zipWriter.Close() }(zipWriter)
	err = filepath.Walk(folderPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Create a new file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		// Update the file name to be relative to the folder path
		relPath, err := filepath.Rel(folderPath, filePath)
		if err != nil {
			return err
		}
		header.Name = relPath
		if info.IsDir() {
			header.Name += "/"
		}
		// Write the file header
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		// If it's a file, write the file content
		if !info.IsDir() {
			file, err2 := os.Open(filePath)
			if err2 != nil {
				return err2
			}
			defer func(file *os.File) { _ = file.Close() }(file)
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// UnzipFile 解压文件
func UnzipFile(zipFile, destDir string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer func(reader *zip.ReadCloser) { _ = reader.Close() }(reader)
	for _, file := range reader.File {
		zippedFile, err2 := file.Open()
		if err2 != nil {
			return err2
		}
		targetPath := fmt.Sprintf("%s/%s", destDir, file.Name)
		if file.FileInfo().IsDir() {
			err = os.MkdirAll(targetPath, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			writer, err3 := os.Create(targetPath)
			if err3 != nil {
				return err3
			}
			if _, err = io.Copy(writer, zippedFile); err != nil {
				return err
			}
			_ = writer.Close()
		}
		_ = zippedFile.Close()
	}
	return nil
}

// DirExists 判断目录是否存在
func DirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// FileExists 判断文件是否存在
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// ReadFile 读取文件内容
func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// GetFolderPath 获取文件所在目录
func GetFolderPath(filePath string) string {
	dir := filepath.Dir(filePath)
	return dir
}

// SaveFile 保存文件
func SaveFile(filename, content string) error {
	folderPath := GetFolderPath(filename)
	if _, err := os.Stat(folderPath); err != nil {
		if err = os.MkdirAll(folderPath, os.ModePerm); err != nil {
			return err
		}
	}
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}
