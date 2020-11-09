package lib

import (
	"io"
	"io/ioutil"
	"mall_go/global"
	"mall_go/pkg/utils"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const (
	TypeImage FileType = iota + 1
	TypePng   FileType = iota + 2
)

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)
	return fileName + ext
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func GetServerUrl() string {
	return global.AppSetting.UploadServerUrl
}

func CheckSavePath(dst string) bool {
	_, e := os.Stat(dst)
	return os.IsNotExist(e)
}

func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}

	}

	return false
}

func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

func CheckPermission(dst string) bool {
	_, e := os.Stat(dst)

	return os.IsPermission(e)
}

func CreateSavePath(dst string, perm os.FileMode) (e error) {
	e = os.MkdirAll(dst, perm)
	if e != nil {
		return
	}

	return
}

func SaveFile(file *multipart.FileHeader, dst string) (e error) {
	src, e := file.Open()
	if e != nil {
		return
	}
	defer src.Close()

	out, e := os.Create(dst)
	if e != nil {
		return
	}
	defer out.Close()

	_, e = io.Copy(out, src)
	return
}
