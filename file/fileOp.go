package main

import (
	"errors"
	"fmt"
	"os"
)

// 创建目录
func CreateDir(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				return nil, err
			} else {
				return os.Open(path)
			}
		}
	}
	return file, err
}

// 创建文件
func CreateFile(path string, file string) (*os.File, error) {
	// 文件是否存在
	filePath := fmt.Sprintf("%s%c%s", path, os.PathSeparator, file)
	_, err := os.Stat(filePath)
	if err == nil {
		return nil, errors.New(fmt.Sprintf("file(%s) is exist", filePath))
	} else {
		if os.IsExist(err) {
			return nil, errors.New(fmt.Sprintf("file(%s) is exist, but throw detect error, cause: %v", filePath, err))
		}
	}

	// 文件夹
	if _, err := CreateDir(path); err != nil {
		return nil, err
	}

	// 文件操作
	return os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
}
