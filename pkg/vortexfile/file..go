package vortexfile

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"os"
	"strconv"
	"time"
)

func Md5Name() string{
	t := time.Now().UnixNano()
	str := strconv.FormatInt(t, 10)
	h := md5.New()
	h.Write([]byte(str))
	b := h.Sum(nil)
	return hex.EncodeToString(b)
}

func PathNotExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func MkDir(path string) error {
	if ! PathNotExists(path){
		return nil
	}
	err := os.MkdirAll(path, os.ModePerm)
	if permission := os.IsPermission(err); permission {
		return errors.New("无操作权限！")
	}
	if err != nil {
		return errors.New("目录创建失败！")
	}
	return nil
}