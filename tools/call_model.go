package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// 要上传的图片路径
	filePath := "D:/_Code/python/lab/Introduction to Artificial Intelligence/cats_vs_dogs_small/validation/dogs/dog.1005.jpg"

	// 服务器地址
	serverURL := "http://localhost:8000/predict"

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("无法打开文件:", err)
	}
	defer file.Close()

	// 创建 multipart 请求体
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 创建文件表单字段
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		log.Fatal("创建表单字段失败:", err)
	}

	// 将文件内容写入表单
	_, err = io.Copy(part, file)
	if err != nil {
		log.Fatal("写入文件内容失败:", err)
	}

	// 关闭 multipart writer（会自动添加结束边界）
	err = writer.Close()
	if err != nil {
		log.Fatal("关闭 multipart writer 失败:", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", serverURL, body)
	if err != nil {
		log.Fatal("创建请求失败:", err)
	}

	// 设置 Content-Type 头（包含 boundary）
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("请求失败:", err)
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("读取响应失败:", err)
	}

	// 输出结果
	fmt.Printf("响应状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应内容: %s\n", respBody)
}
