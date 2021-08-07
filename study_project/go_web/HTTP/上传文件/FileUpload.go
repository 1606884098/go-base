package main

import (
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"time"
)

func main() {
	http.HandleFunc("/", FileUploadHandler)
	http.ListenAndServe(":8080", nil)
}

func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		fmt.Println("get upload file fail, err = ", err)
		w.WriteHeader(500)
		return
	}
	defer file.Close()
	//获取文件大小
	size, err := getUploadFileSize(file)
	fmt.Println("size = ", size)

	//获取文件后缀名
	lastname := path.Ext(fileHeader.Filename)
	//为上传的文件生成新的文件名
	newFile := GetRandomString(8) + lastname
	f, err := os.OpenFile(newFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("save upload file, err = ", err)
		w.WriteHeader(500)
		return
	}
	defer f.Close()
	if _, err := io.Copy(f, file); err != nil {
		fmt.Println("save upload fail, err = ", err)
		w.WriteHeader(500)
		return
	}
	w.Write([]byte("upload file:" + fileHeader.Filename + "-- saveTo:" + newFile))
}

//2019-05-17
type fileSizer interface {
	Size() int64
}

func getUploadFileSize(f multipart.File) (int64, error) {
	if s, ok := f.(fileSizer); ok {
		return s.Size(), nil
	}

	if fp, ok := f.(*os.File); ok {
		fi, err := fp.Stat()
		if err != nil {
			return 0, err
		}
		return fi.Size(), nil
	}
	return 0, nil
}

//生成指定长度的字符串
func GetRandomString(lens int) string {
	str := "0123456789abcdefehijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
