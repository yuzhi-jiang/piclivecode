package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {

	http.HandleFunc("/upload-qr", handleUploadQR)
	http.HandleFunc("/images/", serveImage)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./upload.html")
	})
	fmt.Println("服务器启动在 :8080 端口")
	http.ListenAndServe(":8080", nil)
}
func handleUploadQR(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "只允许 POST 请求", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 检查文件类型
	if !isImage(header.Filename) {
		http.Error(w, "只允许上传图片文件", http.StatusBadRequest)
		return
	}

	// 创建上传目录
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 创建目标文件
	dst, err := os.Create(filepath.Join("./uploads", header.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// 保存文件
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 生成图片URL
	imageURL := fmt.Sprintf("http://%s/images/%s", r.Host, header.Filename)

	// 生成二维码
	qr, err := qrcode.New(imageURL, qrcode.Medium)
	if err != nil {
		http.Error(w, "生成二维码失败", http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "image/png")

	// 将二维码写入响应流
	err = qr.Write(256, w)
	if err != nil {
		http.Error(w, "写入二维码失败", http.StatusInternalServerError)
		return
	}
}

func serveImage(w http.ResponseWriter, r *http.Request) {
	imagePath := filepath.Join("./uploads", filepath.Base(r.URL.Path))
	http.ServeFile(w, r, imagePath)
}

func isImage(filename string) bool {
	ext := filepath.Ext(filename)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	}
	return false
}
