package fanTool

import (
	"fanTool/utils/curlUtils"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"strings"
)

func GetClientIp(c *gin.Context) string {
	clientIp := c.GetHeader("X-Real-IP")
	if clientIp == "" {
		clientIp = c.GetHeader("X-Forwarded-For")
	}
	if clientIp == "" {
		clientIp = c.GetHeader("Proxy-Client-IP")
	}
	if clientIp == "" {
		clientIp = c.GetHeader("WL-Proxy-Client-IP")
	}
	if clientIp == "" {
		clientIp = c.GetHeader("HTTP_CLIENT_IP")
	}
	if clientIp == "" {
		clientIp = c.GetHeader("HTTP_X_FORWARDED_FOR")
	}
	if clientIp == "" {
		clientIp = c.GetHeader("HTTP_X_FORWARDED_FOR")
	}
	if clientIp == "" {
		clientIp = c.ClientIP()
	}
	return strings.TrimSpace(clientIp)
}

func Get(url string) ([]byte, error) {
	return curlUtils.Create().Get(url)
}

func GetByHeader(url string, headers map[string]string) ([]byte, error) {
	return curlUtils.Create().GetByHeader(url, headers)
}

func Post(url string, params map[string]string) ([]byte, error) {
	return curlUtils.Create().Post(url, params)
}

func PostByHeader(url string, params map[string]string, headers map[string]string) ([]byte, error) {
	return curlUtils.Create().PostByHeader(url, params, headers)
}

func PostJson(url string, jsonString string) ([]byte, error) {
	return curlUtils.Create().PostJson(url, jsonString)
}

func PostJsonByHeader(url string, jsonString string, headers map[string]string) ([]byte, error) {
	return curlUtils.Create().PostJsonByHeader(url, jsonString, headers)
}

func PostMultipart(url string, params map[string]string, paramName string, tmpFile *multipart.FileHeader) ([]byte, error) {
	return curlUtils.Create().PostMultipart(url, params, paramName, tmpFile)
}

func SetTimeOut(seconds int64) *http.Client {
	return curlUtils.Create().SetTimeOut(seconds)
}
