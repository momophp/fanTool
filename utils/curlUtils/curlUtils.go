package curlUtils

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var client *http.Client
var curlClientInstance *CurlClient
var curlClientOnce sync.Once

func Create() *CurlClient {
	curlClientOnce.Do(func() {
		client = &http.Client{}
		client.Timeout = time.Duration(10) * time.Second
		curlClientInstance = &CurlClient{}
	})
	return curlClientInstance
}

type CurlClient struct {
}

func (u *CurlClient) Get(url string) ([]byte, error) {
	request, _ := http.NewRequest("GET", url, nil)
	return doGet(request, url)
}

func (u *CurlClient) GetByHeader(url string, headers map[string]string) ([]byte, error) {
	request, _ := http.NewRequest("GET", url, nil)
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	return doGet(request, url)
}

func doGet(request *http.Request, url string) ([]byte, error) {
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) { _ = Body.Close() }(response.Body)
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *CurlClient) Post(url string, params map[string]string) ([]byte, error) {
	request, _ := http.NewRequest("POST", url, strings.NewReader(encodeParams(params)))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return doPost(request)
}

func (u *CurlClient) PostByHeader(url string, params map[string]string, headers map[string]string) ([]byte, error) {
	request, _ := http.NewRequest("POST", url, strings.NewReader(encodeParams(params)))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	return doPost(request)
}

func (u *CurlClient) PostJson(url string, jsonString string) ([]byte, error) {
	request, _ := http.NewRequest("POST", url, strings.NewReader(jsonString))
	request.Header.Set("Content-Type", "application/json")
	return doPost(request)
}

func (u *CurlClient) PostJsonByHeader(url string, jsonString string, headers map[string]string) ([]byte, error) {
	request, _ := http.NewRequest("POST", url, strings.NewReader(jsonString))
	request.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	return doPost(request)
}

func (u *CurlClient) PostMultipart(url string, params map[string]string, paramName string, tmpFile *multipart.FileHeader) ([]byte, error) {
	fp, err := tmpFile.Open()
	if err != nil {
		return nil, err
	}
	defer func(fp multipart.File) { _ = fp.Close() }(fp)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	part, err := writer.CreateFormFile(paramName, filepath.Base(tmpFile.Filename))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, fp)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) { _ = Body.Close() }(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func doPost(request *http.Request) ([]byte, error) {
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) { _ = Body.Close() }(response.Body)
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func encodeParams(params map[string]string) string {
	paramsData := url.Values{}
	for k, v := range params {
		paramsData.Set(k, v)
	}
	return paramsData.Encode()
}

func (u *CurlClient) SetTimeOut(seconds int64) *http.Client {
	client.Timeout = time.Duration(seconds) * time.Second
	return client
}
