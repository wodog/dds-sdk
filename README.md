dds-sdk
=======

Usage
-----

```
// init instance
var d = dds.New(map[string]string{
	"host":   "http://xxx.com",
	"bucket": "xxx",
})

// upload
fileid, err := d.Upload(filename, reader)

// info
file, err := d.Stat(fileid)

// get
reader, err := d.Get(fileid)

// delete
err := d.Delete(fileid)
```
