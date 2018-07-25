package dds

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

type DDS struct {
	host   string
	bucket string
}

// 传入map配置返回DDS实例
func New(m map[string]string) *DDS {
	return &DDS{
		host:   m["host"],
		bucket: m["bucket"],
	}
}

// 根据id获取url地址
func (d *DDS) Url(id string) string {
	return fmt.Sprintf("%s/api/files/%s", d.host, id)
}

// 解析url,  返回id
func (d *DDS) ParseUrl(url string) string {
	ss := strings.Split(url, "/")
	return ss[len(ss)-1]
}

// 上传文件
func (d *DDS) Upload(name string, r io.Reader) (string, error) {
	url := fmt.Sprintf("%s/api/buckets/%s/files", d.host, d.bucket)

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	fw, err := w.CreateFormFile("file", name)
	if err != nil {
		return "", err
	}
	io.Copy(fw, r)
	w.Close()

	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", errors.New(string(b))
	}

	return string(b), nil
}

// 打开文件流
func (d *DDS) Open(id string) (io.ReadCloser, error) {
	url := d.Url(id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

// 删除文件
func (d *DDS) Delete(id string) error {
	url := fmt.Sprintf("%s/api/buckets/%s/files/%s", d.host, d.bucket, id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New(string(b))
	}
	return nil
}

func (d *DDS) Stat(id string) (*File, error) {
	url := fmt.Sprintf("%s/api/buckets/%s/files/%s", d.host, d.bucket, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(string(b))
	}

	var file File
	err = json.Unmarshal(b, &file)
	if err != nil {
		return nil, err
	}

	return &file, nil
}
