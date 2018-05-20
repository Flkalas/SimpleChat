// Code generated by go-bindata.
// sources:
// public/app.js
// public/index.html
// public/style.css
// DO NOT EDIT!

package main

import (
	"github.com/elazarl/go-bindata-assetfs"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _publicAppJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xdf\x6f\xdb\xb6\x13\x7f\xf7\x5f\x71\x55\x03\x90\x82\x6d\x3a\xdf\x16\x7d\x71\xbe\x2e\x90\xb5\x01\xda\xa2\x59\x80\x26\xdb\x9e\x29\xf1\x24\xb1\xa5\x48\x81\xa4\xac\x1a\x43\xfe\xf7\x81\xfa\x65\x39\x96\xb1\x0d\xe3\x8b\x69\xde\xdd\xe7\xee\x3e\x47\xde\x49\x63\x03\xbf\xd7\x48\xff\x5c\x00\x00\xa0\xda\x02\x79\xcd\xab\x8a\xac\x16\xed\x81\xe0\x9e\x6f\xa1\x13\x86\xd5\xb8\x2d\xe8\x5a\xa9\x15\x6c\x36\xf0\x50\x5b\x68\x30\x71\x26\xfd\x81\x7e\x54\xd1\xd8\xdc\xbb\x7c\x0b\x84\xb4\x4a\x9f\x8c\x12\x2e\x1c\x42\x89\xce\xf1\x1c\x1d\x78\x03\x09\x82\x43\xed\xc3\xd6\x17\x61\x6f\xf7\x68\x47\x88\xb4\xe0\xfe\x83\xd1\x1e\xb5\x1f\x71\x6e\xc1\xd6\x5a\x4b\x9d\x83\x92\xce\x83\xc9\x5a\xad\x23\xa8\x90\xae\x52\xfc\x80\x02\x8c\xee\x30\x53\x8b\xa8\x47\x4c\x2c\xb9\x54\x93\xe0\xef\xc2\x7f\xe0\x42\x58\x74\x0e\x6a\x87\x02\x32\x63\x21\xb7\x3c\x49\x82\x17\xae\x81\xef\xb9\xe7\xc7\xa8\x6a\x87\x56\xf3\x12\x5f\x30\x30\x1c\x8f\x7a\xdf\x8d\xd4\x28\xb6\x90\x71\xe5\x30\x68\x3d\xd9\x1a\x41\x66\x5d\x08\xc0\xb5\x18\x6d\xa0\xe0\x7b\x84\x04\x51\x43\x26\x95\x42\x01\xb2\x8b\xf8\xb9\xe7\x3f\xb5\xc8\x7d\x0b\x56\xeb\xd4\x4b\xa3\x69\x3c\xa9\xc6\x9e\x5b\x70\xa8\x32\xd8\x81\x2f\xa4\xbb\x19\x05\xe1\x1f\x6b\x1c\xec\x5a\xe2\xff\xc0\xe4\xb1\xad\x11\x25\x8d\xdb\x6e\x36\x04\x96\xd0\x48\x2d\x4c\xc3\x94\x49\x79\x80\x65\x85\x71\x1e\x96\x40\x36\x8d\x23\xf1\x19\x10\xe3\x42\xdc\xed\x51\xfb\xaf\xd2\x79\xd4\x68\x29\xe9\x89\x27\xab\x63\x68\x38\x8d\x6d\x88\xaf\x74\x39\xec\xe0\xcb\xe3\xc3\xaf\xac\xe2\xd6\x21\x45\x16\x2e\xd5\xc4\x47\x58\x21\x0b\x36\x29\x3b\x2c\x77\x27\xf2\xb0\xc8\xff\x95\x84\x54\x71\xe7\x76\x91\xc2\xcc\x43\xaa\x90\xdb\x4c\xfe\x04\x27\x75\xae\x70\xdd\x87\x14\xbd\x27\xcb\x19\x63\x57\x71\x3d\x98\x07\x4f\x6b\x59\xe6\x90\x29\xc3\xfd\x3a\xa0\x5d\xb0\x0a\x4a\xce\xa6\xbb\x88\x2c\xdb\x18\x73\xdb\xdd\x8a\xdf\xbe\x7d\xa5\xa5\xcb\x59\x5b\xd3\x78\x49\x22\xe0\xca\xef\xa2\xdb\x56\x18\x0d\x7e\x12\x63\x05\x5a\xe8\x7e\xd6\x6f\x87\x8d\xe0\xf6\x07\x58\x53\x6b\x81\x62\x9d\x4a\x9b\x2a\x84\xd2\xae\xdf\x46\x9b\xf9\x20\x36\x21\xf6\x79\x91\x90\xfb\x93\xa4\x12\x23\x0e\x23\x31\x17\x72\x9a\xd8\x14\xc8\x05\xda\x4b\x8c\x79\x6b\x74\x3e\xa8\x56\x56\x96\xdc\x1e\xd6\x99\xd1\x2d\x59\x21\xfb\xe1\x1e\x2f\x43\x8c\xad\xf6\x85\x04\x84\xdc\xcf\x4b\xaa\xf7\x64\x89\xa5\xf9\x2e\x8d\x46\xe6\xcd\xe7\x92\xe7\xd8\x12\xdb\x17\x33\x0e\xd0\xd5\xbf\x45\xdd\x28\xf9\x9e\xdc\x2c\xa6\x92\x71\xff\xdc\x5f\xbd\xe1\x89\x95\xe8\x0b\x23\xdc\xb4\xcb\x39\xd4\x93\x27\x07\xf4\xe5\xbd\x96\x19\xd0\xf6\x65\x74\xbd\x0e\x5e\xed\x80\x90\x97\x4a\x30\x79\x3e\x01\x90\x9e\x49\xc3\x6a\x1f\x86\xf3\x56\xea\x5c\x66\x07\x7a\x0e\x31\xac\xbe\x7f\xb5\x88\xed\x7e\x75\x51\xf5\xd8\xa8\x5a\xed\xe1\xef\x65\x83\x9e\xeb\x2d\x5c\xd1\xb6\x22\x31\x2b\x7c\xa9\xa6\x29\xc6\xcc\xe3\x4f\x4f\xe3\xd0\xce\x1e\xbd\x95\x15\x98\xda\x43\xd0\x9a\x05\x7d\x3e\x3b\x8d\x5f\x3c\xf8\x91\x9d\x9e\xc2\xc0\xe0\x4d\x40\xff\x86\x0e\x7d\x3f\x44\x16\xf3\x98\x43\xe5\xa0\x6f\xb6\x7f\x57\xaa\x57\x47\xd2\x2e\x16\xc9\x15\xa6\xb9\x55\x68\x3d\xbd\xa2\xd1\xeb\x56\x77\xcd\xc3\xff\x68\x2e\x72\x8b\xbe\xb6\xfa\x42\x78\xa7\x6e\x07\xf6\xff\xa1\xe7\x41\xfd\xbf\x38\x3f\xa6\x0b\xbb\xb9\x92\x76\x4c\xf4\x15\xbd\x39\x37\x1d\xa7\xd3\xac\xf5\x98\xd0\x65\x80\x6e\x02\x86\xa9\x64\x6b\xbc\x99\xad\xdb\xa4\x8f\x4e\x86\xdb\x6c\x8d\xba\x7c\x81\x14\xde\x57\xdb\xcd\xa6\x69\x9a\xb1\x0b\xb3\xd4\x94\x9b\x6e\xdb\x4e\xb5\x0f\xf6\x50\x79\xf3\xe5\x91\xdd\x7f\x7c\xd7\x83\xcd\xbb\x1f\x49\x9f\x38\xf7\xdc\xe6\xe8\xe3\x53\xe7\xdd\x21\xcb\xb8\xc0\x27\x43\xdf\x5d\xaf\xe0\x7f\x31\x13\xa8\xf8\x81\xbe\x79\x73\x7d\x1d\x1f\x25\xd7\x2b\xb8\x9e\x7a\xeb\x9a\xcc\x22\xb4\x9b\xc5\xe2\x8a\x0a\x93\xd6\x25\x6a\x1f\x33\x8b\x5c\x1c\x80\x1e\x27\x7a\xe7\xf1\xec\xde\xb1\x42\x0a\x1c\xf8\x9d\xbb\x1b\x83\x42\x90\x77\x7e\x82\x56\x3b\x06\x86\xcf\xa2\x28\x66\x89\xd4\x82\x46\x1f\x1f\xee\x1f\xeb\xc4\x5b\xc4\x7b\x23\x64\x26\x51\x44\xab\xe9\xb3\xe9\x62\x08\x03\x1b\x15\x86\x38\x61\x07\x43\xc8\x2c\x47\x7f\xd7\x9d\xfe\x72\xf8\x2c\x28\x39\x71\x31\x7c\x2e\xf4\x76\xcc\xa5\xd6\x28\xf5\x64\x2a\xd8\xbd\x38\xfb\x84\x32\x2f\xc2\x57\xc6\x15\x25\xec\x74\x66\x93\x98\x29\xee\x3c\x8d\x59\x65\x4d\x45\xc9\xd4\x80\xc4\x6d\x63\xb8\xad\xbd\x81\xee\x7c\xf8\x6c\x4c\x8c\xf7\xa6\x0c\xa9\xff\x15\x00\x00\xff\xff\x8c\xa5\xd7\x5c\xc7\x0a\x00\x00")

func publicAppJsBytes() ([]byte, error) {
	return bindataRead(
		_publicAppJs,
		"public/app.js",
	)
}

func publicAppJs() (*asset, error) {
	bytes, err := publicAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/app.js", size: 2759, mode: os.FileMode(438), modTime: time.Unix(1526832865, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\x6b\x53\xdb\xba\x16\xfd\x9c\xfc\x0a\x55\x33\x77\xe6\xde\x61\x6c\xe5\x55\x1e\x1d\x3b\xb7\x90\x96\x16\x4a\x28\xcf\x52\xf8\x26\x5b\x8a\x2d\x5b\x2f\x24\x99\xe0\xf3\xeb\xcf\xc8\x26\x29\xa4\x0d\xa5\x3d\x9d\xf3\x05\x2b\x4b\xd2\xd6\x5a\x7b\x6d\x5b\x9b\xe8\xd5\xbb\xcf\x93\x8b\xeb\x93\xf7\x20\x77\x82\x8f\xbb\x91\x7f\x00\x8e\x65\x16\x43\x2a\xa1\x07\x28\x26\xe3\x6e\x27\x12\xd4\x61\x90\xe6\xd8\x58\xea\x62\x78\x79\xb1\x1f\x6c\x43\x8f\x3b\xe6\x38\x1d\x4f\x72\xec\x22\xd4\x8e\xbb\x9d\x6e\x27\xe2\x4c\x96\xc0\x50\x1e\x43\xeb\x6a\x4e\x6d\x4e\xa9\x83\x20\x37\x74\x16\xc3\xdc\x39\x6d\xdf\x20\x34\x53\xd2\xd9\x30\x53\x2a\xe3\x14\x6b\x66\xc3\x54\x09\xc4\x52\x25\xff\x3f\xc3\x82\xf1\x3a\x9e\x62\x47\x0d\xc3\x7c\xe3\x20\x55\xd2\x36\xc7\xbd\x24\x6e\x4a\x64\x58\x58\x42\x39\xbb\x33\xa1\xa4\x0e\x51\xa1\x0a\xa6\x24\x45\x83\x70\x10\x6e\x22\x6c\x2d\x75\x16\xa5\xd6\x2e\x67\x42\xc1\x64\x98\x5a\x0b\xd1\xcf\x0e\x41\x0d\xd4\xac\x7d\x29\x1f\xeb\x70\x5a\x6a\xec\xf2\x30\x51\xca\x59\x67\xb0\xf6\x14\xbd\xda\x25\x80\x46\x61\x3f\xec\x37\x9c\x96\xd8\x92\x14\x60\xd2\xd1\xcc\x30\x57\xc7\xd0\xe6\x78\xb8\x3d\x0a\xae\x6c\x99\xe3\xf3\x0f\xfb\xd9\xc7\xeb\xab\x77\x69\x32\x3f\xde\xea\x21\x32\xbb\xde\x2b\x46\x5b\xc5\x5f\x3b\xb7\x89\x9d\x1e\x10\xc4\xce\x8e\x87\x74\xfe\x21\xff\x7a\xba\x7f\x33\x39\x9f\x39\xd2\x3f\xba\x99\xcc\x44\x5e\xba\x3d\x08\x52\xa3\xac\x55\x86\x65\x4c\xc6\x10\x4b\x25\x6b\xa1\xaa\x5f\x4a\x72\x61\xc3\x94\xab\x8a\xcc\x38\x36\xb4\x91\x83\x0b\x7c\x8f\x38\x4b\x6c\xe3\x6d\x80\xe7\xd4\x2a\x41\xd1\x28\xdc\x0a\x7b\x8d\xb6\xc7\xf0\x52\xde\xb8\xdb\x8d\x50\x5b\x67\x51\xa2\x48\xfd\x50\x75\xd4\x80\x94\x63\x6b\x63\x28\xf1\x5d\x82\x0d\x68\x1f\x01\xc1\xa6\x04\x49\xd6\x3e\x1b\xbe\x78\xe1\x0d\x7c\xba\x23\x48\x0c\x96\x04\x3e\x54\x27\x1e\x3f\x1c\x43\xcd\xb8\x1b\x09\xcc\x24\x60\x24\x86\x58\xeb\x26\x08\x61\x77\x8b\xdd\xa9\x92\x0e\x33\x49\x8d\x9f\x78\x32\x63\xd4\x1c\x88\x3a\x18\x36\x13\x2b\x7b\x38\x1c\x77\x3a\x9d\x8e\x9f\x78\x3a\x83\x0d\x01\x3a\x18\x82\x3c\xe8\xf7\x7a\xed\xce\x4e\x27\xaa\x78\x73\x7c\x9a\x63\x17\x08\x6a\x2d\xce\xa8\x5d\xf2\xf7\x28\xf0\x1b\x03\x9f\x10\x08\xee\x02\xff\x62\xb6\xf8\x44\x49\x47\xa5\xf3\x69\x6b\x23\xa1\x8a\xb7\x41\x23\x44\xd8\x5d\xcb\x6c\x31\x5a\x0e\xbe\x17\xe2\xa3\xb2\x59\x0c\x0b\xc5\x24\x25\x6b\x14\x7d\xa7\x86\x49\x5d\xb9\x20\x33\xaa\xd2\x4b\x29\x0d\x06\x5c\xad\x69\x0c\x1d\xbd\x77\x4b\x19\x33\x65\x44\xe0\xb3\x69\x14\x87\x40\x73\x9c\xd2\x5c\x71\x42\x4d\x0c\xa7\xad\x64\x08\xb0\x61\x38\xe0\x38\xf1\xe5\xf6\x14\x24\xd4\xa6\x86\x25\x94\x24\x75\x0c\x13\x6c\x59\x1a\x60\x42\x94\x1c\x78\xe6\x42\x11\xbf\x43\xd2\xf9\xd4\x66\x10\xbc\x2d\x69\x5d\xe9\x90\x4a\xe7\x83\x5b\x2a\xc9\x92\xdd\x8f\xc9\x07\x58\xeb\x47\xab\x3a\x51\x52\x39\xa7\xe4\x62\x65\xe2\x24\x48\x9c\x0c\xb4\x61\x02\x9b\x1a\x3e\xa8\x6b\x17\x41\xf0\x36\xe5\x2c\x2d\x17\x07\x45\x6c\x29\x18\x83\x19\x0e\x34\xd6\xd4\x04\x9a\x63\xb9\x90\x92\x33\x42\xa8\x8c\xa1\x33\x15\x85\xe3\x08\xb1\x31\x38\xa7\x92\x44\xa8\x0d\xb8\xa0\xba\xf4\xef\x19\x2b\xd7\x1b\xf9\xea\xb1\x93\x6b\xac\xfc\x53\x5e\x2e\x1c\x08\x9d\x61\x22\x86\x54\x60\xb6\xea\xf0\xfb\x16\x7b\xec\xef\x63\x68\xbd\xbb\xab\x19\x78\xae\x06\x81\x48\x16\x6f\xe3\x6f\x93\xaf\x2c\x35\x12\x0b\xba\xc2\xff\x72\x09\x3f\x96\xb0\x82\x3e\x53\xa3\x4f\x4b\xd2\x7b\xf3\xaf\x94\xa4\x3f\xe8\xbf\xff\xfb\xbe\x28\x2d\xcb\x64\xc0\xe4\xfa\x82\x3c\x54\x4c\xfe\x4e\x41\x7e\x7b\x3e\x92\x85\x39\x35\x0e\x34\x7f\x03\x82\x65\x46\x0d\x98\xb1\x7b\xea\xbf\x67\xce\x29\x01\x44\x30\x80\xcd\xf7\xaf\xa9\x9c\xa0\x59\x08\x41\x73\xdd\xc4\x70\xce\x88\xcb\xdf\x80\x61\xef\x3f\x0f\x5f\x25\xeb\x8c\x92\xd9\x78\x1f\x33\xfe\x0a\x44\xe8\xe1\xa7\x9f\xba\x56\x15\x10\x95\x75\xa0\x49\x33\xc0\x12\x34\x01\xff\x21\xa9\x45\x45\xfc\x01\x5e\x69\xae\x94\xa5\x00\x83\x45\xcc\x25\xb5\x6e\x84\xfc\x25\x34\xee\x46\xbe\x84\xb4\x03\xd6\xa4\xdf\xee\xd7\x4a\xea\x32\x6b\x2e\xd5\xbb\x8a\xbe\x1d\x84\xfd\x70\x88\x08\xb3\xce\xff\x6c\x6e\xce\xc2\x7a\xdf\xda\xad\x6b\x62\xfc\xac\x11\xe2\x2c\x41\xc5\x4a\x13\xf4\x92\xb0\x8a\xd0\xb0\xb8\xad\xa8\xa9\x1b\x82\xed\x30\x18\x86\xc3\xb0\xbf\x88\xb1\xd2\xb2\x0c\x5e\x6f\x06\xfb\x99\x9e\x24\xe8\xd3\xe1\x29\x3f\x3a\x9e\x7d\xae\x76\xfa\x0e\x0f\x07\x0a\x1d\x4f\x6f\xee\xb9\x9b\x9f\xa9\xed\x53\x27\xca\xe9\x19\xd9\xad\xb6\xe3\xf5\xed\xc9\xcf\x25\x3f\xd7\x96\xa4\xa6\xd6\x4e\x05\x85\x45\x9e\xeb\x00\x19\xc5\x79\xa5\x2d\x12\xe4\xf5\x0b\x13\xfa\x5c\x74\xad\xb4\xa6\x26\x2c\x2c\xea\x87\xfd\x51\x38\x44\x95\x20\x0b\xf0\xc7\x79\xf1\xad\xdc\xcd\xf4\x64\xcb\x7c\x51\x43\x71\x50\x97\x5f\x36\x06\x1b\x3b\x87\xc3\xcb\xc3\xd1\x66\xb1\x57\xf6\xae\x8e\xf0\xe5\x2e\x91\x9b\xdb\x3b\x78\x32\x57\xb7\xc9\xde\x21\x3b\x97\xc5\xee\x27\xc4\xb7\xaf\xee\x26\x57\x27\x07\x27\x62\xb4\xf3\xdb\xb9\xfa\x85\xbe\xb4\x58\x6d\x4b\x7f\x2c\xc5\x8a\x8f\xd7\x9f\xc8\xd1\xee\xbb\x79\xf9\xf5\xb3\xec\xbf\x17\xc7\xfd\xdb\x12\x7d\x9c\xc9\xcb\x34\xf9\x72\x76\x53\x5f\x8b\x9b\xd1\xad\x3e\xa1\x78\xd3\x16\x7b\x48\x5f\x1c\xf6\x68\x55\x9f\xea\xde\xb4\xdc\x4e\xcb\x8d\xd7\x17\xbf\x2c\x05\x61\xad\x57\x8d\x43\x0f\x3d\x24\x6a\xfe\xa5\xf9\x3b\x00\x00\xff\xff\xce\x5c\x1a\x98\xe2\x0c\x00\x00")

func publicIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_publicIndexHtml,
		"public/index.html",
	)
}

func publicIndexHtml() (*asset, error) {
	bytes, err := publicIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/index.html", size: 3298, mode: os.FileMode(438), modTime: time.Unix(1526833136, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicStyleCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x52\xd1\x6e\xdc\x20\x10\x7c\xf7\x57\x6c\x75\xaa\xd4\x4a\xc1\x72\x72\x72\xaa\x70\x4f\xe9\x43\xff\x03\x1b\x62\xd0\x61\x16\xc1\x5e\x6c\xab\xca\xbf\x57\x06\xec\xc6\x52\x5b\xa9\x31\x2f\x66\x76\x76\x66\x80\xed\x50\x2e\xf0\xb3\x02\x00\x90\x26\x7a\x2b\x16\x0e\x2f\x56\xcd\x97\x04\x8d\xc6\x31\xad\xcc\xa0\x89\xc3\x7d\xd3\xbc\xea\x0c\xaf\x04\x26\x4d\x50\x3d\x19\x74\x1c\x7a\xb4\xb7\xd1\x5d\xaa\xb7\xaa\x1a\x85\x71\x45\x6f\x65\x71\xb8\x87\x06\xc4\x8d\x30\x55\x4f\xbd\x16\xc4\x46\x15\xa3\x18\x54\x2c\xbc\xa3\xc9\xe6\xb1\x21\x8f\x3b\x32\x19\x49\x3a\xe5\xf8\x9c\x01\x7c\x55\xe1\xc5\xe2\xc4\x16\x0e\xb1\x0f\x68\x6d\x32\xa9\x57\x93\x2a\x6b\x5b\x13\x89\x45\x5a\xac\xe2\xe0\xd0\xa9\x72\x2c\x11\x06\xe3\x38\x34\x79\xeb\x85\x94\xc6\x0d\x69\xbf\xf5\x83\x35\x45\x22\x93\x59\x87\x44\x38\xae\xf6\x7e\x3e\xb4\xed\x95\x76\x2b\x74\x18\xa4\x0a\xbf\x3b\xfc\x0c\x12\x89\x94\x84\xd3\xf7\xf3\xf3\xd3\xf3\xd3\xc1\x06\xd2\x0f\x4b\x0f\xe1\x0f\x9e\x7b\xc0\x1e\x2d\x06\x0e\xa7\x6f\xe9\xcb\xdd\x5e\x38\x65\xa1\x8e\xd6\x48\x25\x71\x72\x50\x0f\x76\xf1\xda\xf4\xe8\xee\xb2\xe4\x3b\xe4\x78\x94\x90\x6f\x36\xe5\xdd\xa5\x52\x80\xc2\xfb\xd3\xc5\xbe\x7f\x93\x87\xb6\x29\xbd\x9c\xb3\x49\x75\x57\x43\x2c\xf3\x3a\x11\x18\x05\xd1\x5f\x8b\xd2\x56\xed\x70\x66\x51\x0b\x89\x13\x07\xe3\xa2\x22\x68\xa0\x81\x47\x3f\x43\x18\x3a\xf1\xa5\xb9\x4b\xab\x3e\x7f\x2d\x37\x28\xfa\xeb\x10\xf0\xe6\x24\xdb\xce\xfe\xa3\x5d\xd7\xdf\x4c\x8b\xdd\x36\x22\x0f\xfb\x53\xfc\xaf\x10\x23\x7d\x1b\xbb\x8f\xa4\xff\x57\xf8\xb6\xcd\x86\x75\x19\x8d\x73\x19\xfd\xb2\xcd\xb1\xcf\x7e\x86\x4f\x66\xf4\x18\x48\x38\xba\x54\x6f\xbf\x02\x00\x00\xff\xff\x55\x08\x7e\x14\x9c\x03\x00\x00")

func publicStyleCssBytes() ([]byte, error) {
	return bindataRead(
		_publicStyleCss,
		"public/style.css",
	)
}

func publicStyleCss() (*asset, error) {
	bytes, err := publicStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/style.css", size: 924, mode: os.FileMode(438), modTime: time.Unix(1526832898, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"public/app.js": publicAppJs,
	"public/index.html": publicIndexHtml,
	"public/style.css": publicStyleCss,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"public": &bintree{nil, map[string]*bintree{
		"app.js": &bintree{publicAppJs, map[string]*bintree{}},
		"index.html": &bintree{publicIndexHtml, map[string]*bintree{}},
		"style.css": &bintree{publicStyleCss, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}