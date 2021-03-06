package dds_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	dds "github.com/wodog/dds-sdk"
)

var d = dds.New("http://dds.deepblue.com", "test", "")

func TestUrl(t *testing.T) {
	id := "123456"
	url := d.Url(id)
	parsedId := d.ParseUrl(url)
	if id != parsedId {
		t.Fatal("id != parsedId")
	}
}

func Test(t *testing.T) {
	file, err := os.Open("helloworld.txt")
	if err != nil {
		t.Fatal(err)
	}
	_, err = d.Upload(file.Name(), file)
	if err != nil {
		t.Fatal(err)
	}

	f, err := d.Stat(file.Name())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("文件信息:")
	fmt.Println(f)

	r, err := d.Open(file.Name())
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	r.Close()
	fmt.Println("文件内容:")
	fmt.Println(string(b))

	time.Sleep(10 * time.Second)

	err = d.Delete(file.Name())
	if err != nil {
		t.Fatal(err)
	}
}
