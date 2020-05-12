package vortexfile

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"mime/multipart"
	"path"
	"strconv"
	"time"
)

func Path() string {
	return viper.GetString("file.path")+"/"+time.Now().Format("20060102")
}

func FullPath(fileHeader *multipart.FileHeader) string {
	filePath := Path()
	err := MkDir(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	ext := path.Ext(fileHeader.Filename)
	return filePath+"/"+Md5Name()+ext
}


func SizeLimit(fileHeader *multipart.FileHeader, fileType string) error {
	size := fileHeader.Size / (1024 * 1024)
	key := "file."+fileType+"_max_size"
	max := viper.GetInt(key)
	if int(size) > max {
		return errors.New("超过最大"+strconv.Itoa(max)+"M上传限制")
	}
	return nil
}