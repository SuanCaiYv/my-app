package db

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGridFSDaoService_UploadFile(t *testing.T) {
	instance := NewGridFSDaoService()
	filePath, _ := filepath.Abs("/Users/cwb/Desktop/Home/IMG_8559.JPG")
	file, _ := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	data, _ := ioutil.ReadAll(file)
	err := instance.UploadFile(data, "IMG_8559.JPG", primitive.M{"aaa": "bbb"})
	fmt.Println(err)
}

func TestGridFSDaoService_DownloadFile(t *testing.T) {
	instance := NewGridFSDaoService()
	file, m, err := instance.DownloadFile("IMG_8559.JPG")
	fmt.Println(err)
	fmt.Println(m)
	fmt.Println(len(file))
}

func TestGridFSDaoService_DeleteFile(t *testing.T) {
	instance := NewGridFSDaoService()
	err := instance.DeleteFile("IMG_8559.JPG")
	fmt.Println(err)
}
