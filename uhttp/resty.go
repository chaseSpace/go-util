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

	body, readerr := io.ReadAll(resp.Body)
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("GET failed, status code: %d, BODY: %s", resp.StatusCode(), body)
	}

	return body, readerr
}

// PostRequestOption 定义POST请求的选项
type PostRequestOption struct {
	URL         string
	Headers     map[string]string
	Timeout     time.Duration
	ContentType string
}

// postRequest 执行通用的POST请求
func postRequest(option PostRequestOption, setData func(*resty.Request)) ([]byte, error) {
	client := resty.New()
	if option.Timeout > 0 {
		client.SetTimeout(option.Timeout)
	} else {
		client.SetTimeout(defaultTimeout)
	}

	req := client.R()

	// 设置数据
	setData(req)

	// 设置Content-Type
	if option.ContentType != "" {
		req.SetHeader("Content-Type", option.ContentType)
	}

	// 设置头部
	for k, v := range option.Headers {
		req.SetHeader(k, v)
	}

	resp, err := req.Post(option.URL)
	if err != nil {
		return nil, fmt.Errorf("POST request failed: %w", err)
	}

	body, readerr := io.ReadAll(resp.Body)
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("POST failed, status code: %d, BODY: %s", resp.StatusCode(), body)
	}

	return body, readerr
}

// PostJSON 发送POST JSON请求
// data 接受obj-str/struct/map
func PostJSON(url string, data interface{}, headers map[string]string, to ...time.Duration) ([]byte, error) {
	timeout := defaultTimeout
	if len(to) > 0 && to[0] > 0 {
		timeout = to[0]
	}

	option := PostRequestOption{
		URL:         url,
		Headers:     headers,
		Timeout:     timeout,
		ContentType: ContentTypeJSON,
	}

	return postRequest(option, func(req *resty.Request) {
		req.SetBody(data)
	})
}

// PostForm 发送POST表单请求
func PostForm(url string, formData map[string]string, headers map[string]string, to ...time.Duration) ([]byte, error) {
	timeout := defaultTimeout
	if len(to) > 0 && to[0] > 0 {
		timeout = to[0]
	}

	option := PostRequestOption{
		URL:         url,
		Headers:     headers,
		Timeout:     timeout,
		ContentType: ContentTypeForm,
	}

	return postRequest(option, func(req *resty.Request) {
		req.SetFormData(formData)
	})
}
