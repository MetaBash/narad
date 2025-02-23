package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// # Upload Image
func UploadImage(image string) (string, error) {
	var err error

	var errMsg string

	base64ImgStr := image

	folder := "public"

	imageURL, err := SaveBase64Image(base64ImgStr, folder)
	if err != nil {
		errMsg = "error: failed to save image: " + err.Error()
		err = errors.New(errMsg)
		return "", err
	}

	return imageURL, nil
}

// # Save Base64 Image
func SaveBase64Image(base64ImgStr, folderPath string) (string, error) {
	var err error

	data, err := base64.StdEncoding.DecodeString(base64ImgStr)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("image_%d.png", time.Now().UnixNano())

	filePath := filepath.Join(folderPath, fileName)

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return "", err
	}

	imageURL := fmt.Sprintf("http://localhost:3001/images/%s", fileName)

	return imageURL, nil
}
