package alipay

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// logger 日志接口，可自己实现日志记录
type logger interface {
	recordLog(resp *http.Response)
}

// FileLogger 文件日志
type FileLogger struct {
	file *os.File
}

func (f *FileLogger) recordLog(resp *http.Response) {
	reqBody, err := resp.Request.GetBody()
	if err != nil {
		return
	}

	reqData, err := ioutil.ReadAll(reqBody)
	if err != nil {
		return
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	_, _ = f.file.WriteString(fmt.Sprintf(
		"[%s] req_url: %s, req_params: %s, resp_status: %d, resp_data: %s \n\r",
		time.Now().Format("2006-01-02 15:04:05"),
		resp.Request.URL,
		string(reqData),
		resp.StatusCode,
		string(respData),
	))
}

// NewFileLogger 获取 FileLogger 实例
func NewFileLogger(file *os.File) *FileLogger {
	return &FileLogger{file: file}
}
