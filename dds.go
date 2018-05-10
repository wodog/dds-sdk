package dds

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type dds struct {
	host   string
	bucket string
}

func New(m map[string]string) *dds {
	return &dds{
		host:   m["host"],
		bucket: m["bucket"],
	}
}

// 获取url
func (d *dds) Url(id string) string {
	return d.host + "/buckets/" + d.bucket + "/files/" + id
}

// 上传文件
func (d *dds) Upload(name string, r io.Reader) (*File, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	fw, err := w.CreateFormFile("file", name)
	if err != nil {
		return nil, err
	}
	io.Copy(fw, r)
	w.Close()

	req, err := http.NewRequest("POST", d.Url(""), &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
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

// 获取文件内容
func (d *dds) Get(url string) ([]byte, error) {
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

	return b, nil
}

// 删除文件
func (d *dds) Delete(url string) error {
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
