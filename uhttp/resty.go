package uhttp

import (
	"fmt"
	"io"
	"time"

	"resty.dev/v3"
)

const (
	ContentTypeForm  = "application/x-www-form-urlencoded"
	ContentTypeJSON  = "application/json"
	ContentTypeOctet = "application/octet-stream"
)

// defaultTimeout 默认超时时间
const defaultTimeout = 30 * time.Second

// DownloadFile 使用resty下载文件
func DownloadFile(url string, contentType string, to ...time.Duration) (file []byte, err error) {
	client := resty.New()
	if len(to) > 0 && to[0] > 0 {
		client.SetTimeout(to[0])
	} else {
		client.SetTimeout(defaultTimeout)
	}

	req := client.R()
	req.SetHeader("Content-Type", contentType)

	resp, err := req.Get(url)
	if err != nil {
		return nil, fmt.Errorf("下载文件失败: %w", err)
	}

	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("download failed, status code: %d", resp.StatusCode())
	}

	bodyContent, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	return bodyContent, nil
}

// Get 发送GET请求
func Get(url string, headers map[string]string, to ...time.Duration) ([]byte, error) {
	client := resty.New()

	if len(to) > 0 && to[0] > 0 {
		client.SetTimeout(to[0])
	} else {
		client.SetTimeout(defaultTimeout)
	}

	req := client.R()
	for k, v := range headers {
		req.SetHeader(k, v)
	}

	resp, err := req.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET request failed: %w", err)
	}

	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("GET failed, status code: %d", resp.StatusCode())
	}

	return io.ReadAll(resp.Body)
}

// PostJSON 发送POST JSON请求
// data 接受obj-str/struct/map
func PostJSON(url string, data interface{}, headers map[string]string, to ...time.Duration) ([]byte, error) {
	client := resty.New()
	if len(to) > 0 && to[0] > 0 {
		client.SetTimeout(to[0])
	} else {
		client.SetTimeout(defaultTimeout)
	}

	req := client.R()
	req.SetBody(data)
	req.SetHeader("Content-Type", ContentTypeJSON)

	for k, v := range headers {
		req.SetHeader(k, v)
	}

	resp, err := req.Post(url)
	if err != nil {
		return nil, fmt.Errorf("POST JSON request failed: %w", err)
	}

	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("POST JSON failed, status code: %d", resp.StatusCode())
	}

	return io.ReadAll(resp.Body)
}

// PostForm 发送POST表单请求
func PostForm(url string, formData map[string]string, headers map[string]string, to ...time.Duration) ([]byte, error) {
	client := resty.New()
	if len(to) > 0 && to[0] > 0 {
		client.SetTimeout(to[0])
	} else {
		client.SetTimeout(defaultTimeout)
	}

	req := client.R()
	req.SetFormData(formData)
	req.SetHeader("Content-Type", ContentTypeForm)

	for k, v := range headers {
		req.SetHeader(k, v)
	}

	resp, err := req.Post(url)
	if err != nil {
		return nil, fmt.Errorf("POST form request failed: %w", err)
	}

	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("POST form failed, status code: %d", resp.StatusCode())
	}

	return io.ReadAll(resp.Body)
}
