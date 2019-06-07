// Code generated by go-bindata.
// sources:
// assets/app/app.js
// assets/css/app.css
// assets/index.html
// assets/js/angular-websocket.min.js
// assets/ui.html
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

var _assetsAppAppJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x56\x5f\x8b\xe3\xb6\x17\x7d\xde\x7c\x8a\xfb\x33\x81\x38\x33\xc6\x4e\x7e\x85\x76\x9a\x90\x87\x32\x6d\xd9\x96\xee\x6c\x21\x94\x3e\x2c\x0b\xa3\x71\xae\xff\x74\x65\x5d\x23\xc9\x09\xdb\xc1\xdf\xbd\x48\xf2\x1f\x39\x93\x76\x67\x16\xc6\x2f\x89\xa5\x7b\x8e\xce\x91\x8e\x2e\x4e\x49\x28\xe2\x18\x73\xca\xc3\x80\x13\x3b\xe0\x01\x58\x5d\xc7\x7f\xa9\x60\xb9\x9d\xcd\x8e\x4c\x9a\x57\xd8\x01\x13\x79\xc3\x99\x8c\x2b\x3a\x34\x1c\xc3\x85\x2e\x50\x56\xa4\x34\xd3\x3f\xd4\xf5\x22\x82\x0f\x0b\x91\xff\x89\x0f\x7b\x4a\x3f\xa1\x5e\x7c\x34\xe0\xe4\x6a\x66\xa8\x32\x96\x6a\x92\x9f\x7d\xc8\x4f\x47\x14\x5a\x2d\x22\xc8\x1a\x91\xea\x92\x44\x38\x3f\xe1\x83\xb2\xd8\x25\x3c\xce\xc0\x3e\x49\x02\xef\x6b\x14\xc0\x60\x20\x86\x94\x84\x40\x0b\xe9\x8a\x8c\xc2\x03\xd3\x6c\xaf\x25\xb2\x0a\x76\x30\x32\x85\xc1\x49\x6d\x92\x64\xfd\xfd\xff\xe3\xf5\xb7\x37\xf1\x3a\xfe\xe6\xbb\xcd\xcd\xea\x66\x95\xb0\xba\x4c\x4e\xce\xe0\x48\x92\x12\xe7\x8e\x19\x76\xf0\xe1\xe3\x30\x37\x92\xc7\x24\xde\xa1\x52\x2c\xc7\x70\xd0\x5d\xb9\x81\x51\x34\xc0\x64\x4b\x73\xd2\xd0\xd5\xfc\xcf\x2c\xf8\xc6\x9f\xed\xc1\x5b\x0f\xdb\x8b\x88\xeb\x46\x15\xe1\xaf\xfb\xf7\x77\x71\xcd\xa4\xc2\xbe\x38\x36\x7a\x96\x03\xa4\x9d\x9a\xa8\x50\x17\x74\x50\xb0\x9b\xe8\xe9\x39\x37\xde\xff\x68\x98\xcf\x51\x6f\xc6\x83\xf0\x9d\x4c\xcc\x2b\x14\x07\xa7\x47\x69\x59\x8a\xbc\xcc\x3e\x87\x8f\xc0\x3a\xe2\x45\x8e\x7a\x01\xed\xd2\xf3\xd2\xf6\x12\x07\x85\x12\x75\x23\x45\x2f\x72\x3b\x33\xe2\xaf\x92\xd9\xb3\x53\x02\x17\x63\x32\xf0\xfa\xc2\x01\x94\x66\xd2\x33\x06\x61\x23\x79\x04\x29\xe3\xfc\x81\xa5\x9f\x96\x67\xd5\xfd\x06\xaa\x69\x82\x2c\x46\x34\x9c\x47\xf0\x28\xb1\x0b\xdf\x2f\xd9\x1d\xe9\x3b\x92\x15\xe3\xb7\x9c\x14\x6e\x40\xcb\x06\x5b\xcf\xfa\xa0\xe1\x99\x91\x19\x9f\xff\x0c\xcf\x58\xf6\x66\x5a\xf9\x34\x48\x1e\x63\xe7\xf8\xcb\x59\x1a\x9f\x73\x2f\xed\xd9\xa1\xb6\x33\x7b\xbf\x13\xf8\x11\xb3\x52\x20\xe8\x02\xe1\xfe\xf7\x82\x04\xfe\x56\x2a\x7d\x4b\x42\x4b\x13\x34\x79\x6f\x54\x76\xff\x81\x84\xab\xab\x4d\x5d\x6a\xfb\xc6\x3d\xb8\x76\x62\x13\x30\x96\xfa\x21\x18\xc9\x26\xed\x42\xa5\x54\x63\x04\xf3\x42\xeb\x3a\x82\x79\x29\x34\xca\x23\xe3\x11\xcc\x39\xa5\xcc\x06\x1c\xce\x93\xe4\x76\xdc\x41\x63\x6d\xc6\xcd\x3d\x69\xb7\xe3\xe0\x03\xa7\x13\xca\x6e\x74\xe6\x12\x31\x3f\xa9\x3f\x24\x87\x1d\xb8\x6e\x12\xc0\xf5\xb8\x46\x5c\x90\xd2\xe1\x12\xae\x21\xd8\x4c\x27\x6a\x92\xdd\x44\xdf\x6d\x2c\xe1\xb9\xa4\xd8\xa6\x34\x74\x6b\xf8\x41\xaf\x54\xde\x07\xa4\xcc\xec\x6b\xac\xa8\x91\x29\xc2\x6e\x07\x81\x15\x1f\x78\x09\x3a\x33\x65\xca\xcd\xd9\xba\x53\x6c\x01\xb9\xc2\x4b\x3c\xce\xef\x05\xa2\x61\x23\xce\x98\x66\x43\xc7\x49\x92\xbe\x16\xad\x13\xd8\x3d\x31\x67\xeb\xba\x22\x89\x99\x44\x55\xec\x35\xd3\x08\x3b\xcf\xe8\xb0\xb6\x3d\xca\x38\x37\x6d\xdb\x6e\xd9\xdf\x24\x30\x59\x27\x29\x89\xac\xcc\x83\x65\xac\x0b\x14\xe3\x25\x92\xa8\x6a\x12\xca\xbb\x45\x67\x5b\xd0\x17\xf8\xea\x6d\xaa\x5b\x5f\x96\x42\xfd\x33\x13\xfb\x1a\xf1\xe0\xa9\x0a\x95\x19\xe8\xa9\x9d\xb0\xba\xb9\x2c\x2c\x82\x47\x08\x32\x26\xde\xd1\x01\x83\x0d\x58\x24\xb4\x5f\x96\x3b\xb9\xe8\x0a\x35\x64\x4c\x38\x74\xb0\x84\x89\xde\xa9\x5c\xb3\x8e\x2f\xb5\xa2\x03\xbe\x40\x69\xe5\x64\x9a\x9f\xd7\x54\xf9\x96\xf8\x64\x43\x0b\xe2\x2f\xd9\x4f\x53\x1e\x6c\xc0\xfc\x7c\x95\x4a\x8b\xff\x77\x81\xa5\x48\x6f\x89\xf8\x1e\x75\x4d\xa5\xd0\xbe\xd0\x23\xe3\x3d\xaf\xb9\xfc\x1a\x2b\xf3\xf1\xe3\x67\x2b\x4e\x7d\xe8\x35\x1c\x19\xdf\x3e\xd7\x97\x0f\x0d\x36\x8e\xfd\xb5\x4e\xa1\x14\xe9\x5b\x64\xfa\x2b\x4d\x16\x3e\xf4\x65\x26\x7d\xe8\x2b\x99\x4c\xae\x8c\xcf\xbe\xe9\x87\x97\xba\xc9\xd3\xb6\x13\x3a\x74\x04\xeb\xd5\x6a\xe5\xbe\x3c\x0c\xe1\x3f\x01\x00\x00\xff\xff\x77\xb7\xf8\xcf\x02\x0b\x00\x00")

func assetsAppAppJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsAppAppJs,
		"assets/app/app.js",
	)
}

func assetsAppAppJs() (*asset, error) {
	bytes, err := assetsAppAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/app/app.js", size: 2818, mode: os.FileMode(420), modTime: time.Unix(1481854601, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCssAppCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8e\x41\x0a\x83\x40\x0c\x45\xf7\x73\x8a\xac\x85\xd1\x4d\x57\xe3\x69\xa6\x63\xd4\xd0\x98\x91\x49\x4a\x95\xd2\xbb\x17\x2b\x15\x71\x19\x5e\xfe\xff\xaf\xa9\x20\xa1\x18\x16\xec\x20\x65\x7e\x4e\xa2\xa0\xb6\x32\x2a\x54\x8d\xab\x4b\x7e\xf9\x83\xbf\x1d\x00\x80\xe1\x62\x3e\x32\x0d\x12\x76\xd2\xba\x8f\xab\x53\xe6\xeb\x63\x47\x3a\x73\x5c\x03\x09\x93\xa0\xbf\x73\x4e\x8f\xf6\x47\x7a\xce\xd1\x82\x64\xc1\xfd\x6e\x2a\x28\xa8\x68\x60\x23\x9e\xfa\x37\x83\xcb\x22\x63\x6f\x47\xe6\x5c\x0c\x3a\xc7\x84\xd0\xd3\xf2\x4f\x4d\xb1\x0c\x24\xbe\xd0\x30\x5a\xf0\xb7\x79\xd9\x3c\xbf\x01\x00\x00\xff\xff\xe1\xda\xcb\x43\xef\x00\x00\x00")

func assetsCssAppCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsCssAppCss,
		"assets/css/app.css",
	)
}

func assetsCssAppCss() (*asset, error) {
	bytes, err := assetsCssAppCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/css/app.css", size: 239, mode: os.FileMode(420), modTime: time.Unix(1481854601, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x94\x4f\x53\xe3\x38\x10\xc5\xef\xfe\x14\x42\x17\x0e\xbb\xb6\x37\x24\xbb\x61\xb7\x6c\x57\x65\x03\x43\x12\xa8\x81\x49\x20\x90\xa3\x22\x75\xec\x0e\xfa\x37\x92\xec\x90\xf9\xf4\x53\x8e\x8b\xbf\x07\x6a\x2e\x73\xb2\xfa\xf9\xb5\x9e\x7e\x55\x6a\x65\x47\x67\xd7\xe3\xdb\xd5\xcd\x39\xa9\x82\x92\x45\x14\x65\xed\x97\x48\xa6\xcb\x9c\x82\xa6\x44\x97\x31\xb3\x36\xa7\xa1\x02\xa7\x8c\x0f\x2c\x8c\xac\xa5\x45\x44\x48\x56\x01\x13\xed\x82\x90\x4c\x41\x60\x84\x57\xcc\x79\x08\x39\xad\xc3\x26\x3e\xa5\x6f\x7f\x55\x21\xd8\x18\xbe\xd7\xd8\xe4\xf4\x21\xbe\x1b\xc5\x63\xa3\x2c\x0b\xb8\x96\x40\x09\x37\x3a\x80\x0e\x39\x9d\x9e\xe7\x20\x4a\x78\xd7\xa9\x99\x82\x9c\x36\x08\x3b\x6b\x5c\x78\x63\xde\xa1\x08\x55\x2e\xa0\x41\x0e\xf1\xa1\xf8\x93\xa0\xc6\x80\x4c\xc6\x9e\x33\x09\x79\x8f\x16\x51\xb7\x53\xc0\x20\xa1\x98\xea\xcd\xc1\xd0\x40\x96\x76\x4a\x14\x65\x47\x71\x4c\xae\x58\x00\x1f\x08\x37\xca\xa2\x04\x41\x98\x16\x44\xa1\xc6\x0d\x82\x20\xe3\xc5\x82\xc4\x71\x11\x65\x12\xf5\x23\x71\x20\x73\xea\xc3\x5e\x82\xaf\x00\x02\x25\x95\x83\x4d\x4e\x5b\x3e\xff\x5f\x9a\x2a\xf6\xc4\x85\x4e\xd6\xc6\x04\x1f\x1c\xb3\x6d\xc1\x8d\x4a\x5f\x84\xb4\x9f\xf4\x93\x61\xca\xbd\x7f\xd5\x12\x85\x3a\xe1\xde\x53\x82\x3a\x40\xe9\x30\xec\x73\xea\x2b\xd6\x3f\x1d\xc4\xff\x2f\x57\x88\x8b\xe9\x17\xb8\xec\x89\x0b\x35\x9b\x8f\x1e\xf7\xbc\x9e\x8c\x26\xf3\xb2\x7f\x72\xad\xee\xf8\x6e\x37\x34\xba\x3f\x5f\x89\x72\xb0\x64\x7f\xdc\xa8\xc5\xad\xff\x91\x5e\xfe\x73\xda\xac\xc5\xf9\xb6\x1a\xd4\x94\x70\x67\xbc\x37\x0e\x4b\xd4\x39\x65\xda\xe8\xbd\x32\xb5\xa7\xcf\xec\xd7\x36\xa0\xd1\x4c\x92\x50\x81\x82\xdf\x4c\x1a\x1f\x42\x3e\xe3\x75\x93\xbd\xf9\xda\xc3\xb9\x5f\x3e\x2c\x07\xfa\xec\xaf\x59\x1d\xa4\xbe\x60\x5e\x8e\x67\xf5\x78\x58\xef\xb6\xa2\xbe\xff\x77\xb1\x74\x57\xcd\x7c\x65\xcc\x8d\x3d\x59\xdf\xaf\x4a\x55\xce\xbe\x4d\x1f\x76\x32\x5d\xd8\x4f\x79\x0f\x5c\x1d\x46\x7b\x2c\x66\x6d\x77\x8c\x8f\xa8\xad\xd7\x73\x87\x36\x10\xef\xf8\x2b\x33\xdb\xb2\xa7\xa4\x34\xa6\x94\xc0\x2c\xfa\x03\x6f\xab\xa5\x12\xd7\x3e\x65\xba\xac\x25\x73\x5b\x9f\xf6\x92\xbf\x93\xe1\x73\x7d\xa0\xdd\x7a\x5a\x64\x69\xb7\x67\x77\xbb\xc9\xbb\x84\xed\x4b\x7b\xbc\x83\xb5\x37\xfc\x11\xc2\x2f\x35\x32\x6b\x0f\x1c\x1f\x7d\x59\xda\x4d\x67\xbb\x5c\x1b\xb1\x6f\xe7\xb8\x1d\x1d\x67\xa4\x04\xf7\x76\x9c\xc7\x2f\x2a\x7d\x0e\x10\xd8\xb4\x7e\xd4\x5c\xd6\x02\xba\xa0\xe3\x1a\x93\xf6\x69\x38\xa6\xc5\x04\x1c\x1c\x7b\x62\x6a\x47\x04\x36\x49\x92\x64\xa9\xc0\xa6\x4b\x6d\xb3\x8a\x28\x4b\xbb\xd7\xe4\x67\x00\x00\x00\xff\xff\x39\xf8\xd7\xd6\x5e\x04\x00\x00")

func assetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsIndexHtml,
		"assets/index.html",
	)
}

func assetsIndexHtml() (*asset, error) {
	bytes, err := assetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/index.html", size: 1118, mode: os.FileMode(420), modTime: time.Unix(1481854601, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsJsAngularWebsocketMinJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x59\x51\x73\xdb\x36\xd6\x7d\xef\xaf\xa0\xf8\x65\x38\x40\x8d\xd0\x72\x93\xbe\x48\x85\x35\xa9\xe3\x7e\x9b\x9d\xda\xce\xc6\xe9\xf4\xc1\xf5\x64\x20\xf2\x52\x42\x43\x01\x2a\x00\xda\x51\x45\xfe\xf7\x1d\x10\x20\x45\xd2\x94\xbd\xb3\xb3\x2f\x36\x71\x01\x1c\x5c\x5c\x5c\x1c\x1c\x40\x93\xac\x10\x89\xe1\x52\x20\x20\x06\xef\x79\x86\xc2\xc6\x12\x52\x6a\x76\x5b\x90\x59\x90\x42\xc6\x05\x44\x91\xfb\x1f\xb3\x4d\x8a\xdd\x27\xba\x0b\x37\x32\x2d\x72\x08\x49\x08\xdf\xb6\x52\x19\x1d\x92\x90\x89\x55\x91\x33\x15\x92\xf0\x51\x87\xf7\xc4\xe0\x39\xe4\x1a\x02\x8b\x5d\x08\xd7\x31\x0d\x27\x0d\xb8\xef\x87\x0d\x72\x50\xc4\x1b\x88\x82\xbf\x0a\xae\x00\xb5\x78\xf8\x60\x7a\xd4\x21\x76\xb0\xfb\x07\xa6\x02\x49\xf7\xbe\xd7\x6c\x5f\x55\x73\x83\x24\x91\x71\x83\x03\xb1\x07\x20\x10\x3f\x6a\x7c\x28\xff\x0e\x4b\x2d\x93\xaf\x60\x68\xdb\xb8\xaa\x90\x59\x73\x4d\xba\x51\x21\x92\x08\xbc\x0f\x0b\x0d\x81\x36\x8a\x27\x26\x9c\x37\xd5\x81\x46\x80\xf7\x0a\x4c\xa1\x44\x00\x51\x04\xf1\x97\x2f\xa0\xaf\xea\x79\x2c\x60\xb6\x0f\x53\xc8\x58\x91\x9b\x70\x06\x55\xd5\x76\x52\x07\xd4\x0e\x90\xb7\x88\xb2\x9c\x7c\x45\x12\x97\xe5\x45\xfd\x17\x09\x2a\x89\xa4\x0f\x92\xa7\xc1\x14\x13\xeb\x5d\xbc\x55\xd2\xc8\x44\xe6\x9a\x4a\x67\x28\x54\x4e\x4d\x59\x86\x57\x5c\x6b\x2e\x56\xc1\x6f\x9f\x7e\x0d\x5d\x8d\xd6\x39\x3d\x45\x8f\x5a\xe3\x53\x1e\x1b\xd0\x06\x35\x1d\x3c\x96\x4e\xe4\x16\xa8\x88\x22\xe1\x3e\xcb\x12\x5c\x85\x92\xd2\xdc\x5a\xcb\x2f\x8c\xe7\xf2\x01\x94\x6b\xf4\xc4\x1c\x45\x93\xa9\xf7\x42\xc3\xbb\xed\x36\xdf\xbd\xd3\x3b\x91\xb8\xd6\x3d\x53\x59\x4e\xce\x5c\x4b\x2e\xb8\xe1\x2c\xff\xcc\x37\x20\x0b\xe3\x9a\xf6\x6d\x65\xf9\xe3\xd4\xc3\x6e\xd8\xb7\x5e\xc3\x43\xb9\x2c\xdf\xc0\x8f\xde\x5b\x48\xa4\x10\x90\x98\x0f\xd9\xb5\x34\xd7\x52\x6d\x58\x7e\x91\x4b\xed\xa7\x76\xb4\xfa\xe0\xd3\x92\x0b\xa6\x76\x9f\x77\x4d\x34\x0e\xe5\xb2\x0c\x97\xb9\x5c\xfa\x88\x7e\x69\xb1\xde\x19\x03\x9b\xad\xd1\xd4\x3b\xaa\x41\xa4\xff\x2a\xa0\x00\x7a\x77\xef\x2c\x52\xdc\x6c\x41\x5c\xb0\x3c\x5f\xb2\xe4\xab\xee\xd8\xaf\x40\x6b\xb6\x82\xb1\xaa\x4b\xa5\xa4\x1a\xab\xa8\x1d\xee\x55\x64\xa8\x71\x89\xa5\xbb\x5b\xc3\x0c\x5c\x48\xa1\x0d\x13\x46\x63\x62\x16\xae\xd2\x7b\x8b\xf0\xcc\x95\x35\x98\x0f\xc2\x80\x12\x2c\xaf\xbb\xa0\x29\xae\x7c\x0a\xfb\xd4\xb2\x3b\x73\x14\x94\xee\x2f\x6e\xae\xaf\x2f\x2f\x3e\x7f\xb8\xfe\xff\xd9\x94\xdc\x7c\xbc\xbc\x9e\x9d\x91\x8b\x5f\x6f\x6e\xad\xe1\x87\xfa\xeb\xf2\xfd\xec\x0d\xf9\x74\xe9\x1b\x7e\x79\xf7\xf3\xcd\xa7\xcf\x97\xef\x67\x6f\x2b\xd2\x43\x17\x87\x35\xb8\x90\x29\xd0\x33\x78\x43\x06\xc3\x7b\xc7\xd9\x32\x07\xeb\x46\xa1\x6d\x43\x4d\xef\xde\xc2\x9b\xfb\x5e\x5b\xcd\x32\x78\xcf\x57\xa0\x0d\x3d\x6c\x5c\xbc\x87\x28\x9a\x1c\x72\x3c\x7e\xf5\x6a\xbb\x66\x1a\xa2\xa8\x6b\x4b\xeb\x6e\x08\xf7\xbd\x5b\x72\x91\x7e\x96\x75\x9a\x1f\x10\x0d\xf6\x5c\x63\xfb\xcf\x7d\xc4\x4c\x14\xa1\xce\x3e\x32\x47\xf6\x4e\x7f\x50\x29\x50\xf8\x2a\x05\x6d\x94\xdc\x85\x07\xae\xc1\x7b\xe9\x61\xa0\xc2\x98\xc8\x41\xc4\x7c\x38\x7a\x53\x9c\xb4\xd3\xa9\x89\xac\x57\x88\x0f\x0b\x48\x29\x3d\x9a\x28\xb1\x5d\xc5\xb2\x44\x9d\x9e\x54\xc4\x89\x02\x9b\x1a\x0d\x59\x0c\x78\xa7\xe1\x0e\x37\x8e\x14\x1b\x97\xcd\x34\xb9\x6b\x19\xef\xbe\x0e\xa2\x63\x53\x37\x76\x9b\xf4\xff\x60\x22\xcd\x41\x0d\x41\xe4\x16\xc4\x4b\x08\x76\x3b\x1d\xe9\x0e\x76\xdb\xbc\xd4\xbf\xde\x5b\x47\x00\x92\x9a\x2e\x5e\x00\xa8\x13\x76\x14\xa0\xc3\x1f\x03\x3e\x19\x24\x57\xc6\x15\x38\x9a\xe8\xac\x7c\x26\x15\x9a\xf7\x59\x24\xce\x41\xac\xcc\xfa\xbf\x5e\xd3\xb9\xcb\x57\xef\xce\x01\x56\xaf\x79\x66\x10\x9e\x77\x61\x6d\x2d\x4a\x11\xc4\x7e\x25\x71\xc3\x79\x93\xe1\x64\x16\x6d\x9b\xd9\x3f\x6f\x6f\xae\x63\x7b\x24\x8a\x15\xcf\x76\x9d\xce\xf6\x8c\x4d\x21\x03\xa5\x20\x8d\x15\x68\x99\x3f\x00\xc2\x55\x3f\x0c\x42\x1a\x9e\xed\xfa\xfc\xd8\x4d\x6d\x1b\x11\xeb\xbf\xa1\xd3\xb9\xf9\x69\x8c\x4e\x7d\x80\xe6\xe6\xe4\x04\x8f\xd5\xdf\x99\xfb\x38\x61\x79\xee\x56\x70\xb8\x0c\x6e\xfc\x01\xa7\xbe\xe8\x40\xbf\xfd\x88\x07\xfd\x06\xff\x89\x0b\x03\xbe\x7f\xd1\x85\x7e\xfb\x11\x17\xfa\x0d\x9e\x77\xc1\x85\xac\x37\x68\xc3\x6c\x63\x11\xdf\x16\x7a\x8d\xc0\x25\xfe\x10\xc8\x1d\xb7\xc7\x91\x06\xa1\x7b\x0e\xaa\x9e\xc1\x33\x50\x83\x10\x3c\x07\xe5\x39\x87\x3e\xd1\xb9\x93\x1d\x02\x8c\xcd\x5a\xc9\xc7\x40\xc0\x63\x50\x63\xa2\xb0\x41\x0d\x36\x85\x36\xc1\x12\x02\x16\xb4\x82\x18\xcf\x79\x86\x4c\x14\x2d\x91\x89\x33\x9e\x1b\x50\x38\x8a\x26\x69\xaf\xd4\x16\x02\x5e\x6f\xc6\xc4\x6a\xdc\x4f\xb0\xba\xfc\xb6\x1d\x19\xed\x23\x33\xf6\x24\xee\x0c\xe6\xf6\x53\x20\x55\xa0\xa0\x56\xa9\x56\x20\x2b\xd0\xda\x39\xd0\x8f\xc3\x50\x45\xb8\x48\xec\x33\x31\x03\xb2\x75\xd0\x33\xb3\x68\x3c\x9a\x39\x11\x49\x58\x61\x64\xad\xca\xea\xba\x43\x69\x32\xad\xc6\x62\xd8\xa7\xdd\xde\xaa\xbc\x20\x88\x46\xb6\x78\xb3\x4c\x07\x16\x1c\x9e\xbd\x03\x92\xed\x8d\xe7\xf6\x42\x7d\xfe\x9a\xbe\xb6\x5c\x98\xe6\x70\x65\xad\x0d\x75\xf8\xd5\x8c\xee\x77\x04\xb8\xc2\x33\x74\xb4\x92\x98\x8e\xb6\x40\x93\x29\xc6\x04\x3d\x2f\x37\xed\x1d\x20\x91\x29\x4c\x1a\x76\x1e\xe8\x9c\xb2\x1c\xc4\x6c\xa0\x6e\x62\x2e\x52\xf8\x76\x93\x21\x07\x83\xcf\x5f\x9f\x61\x7f\x02\xb4\x5d\x46\x22\xd6\x3d\xd7\xfe\xc7\x11\xeb\xef\xb5\x41\xc4\x9e\x54\x3e\x8d\xd8\x13\x5f\xfb\x2a\xa0\xcf\x77\xcd\x75\xc8\xb8\x2b\x12\xde\x4b\xba\x71\xec\xc5\xd4\xaa\xd8\x80\x30\x9a\xfc\x80\xc9\xe0\xae\xb1\xd0\xcf\x4f\x05\xe2\xda\x8e\x34\x91\xb5\xf7\xdd\x32\xd1\x5d\x87\x0d\xc6\x55\x43\xba\x92\x08\xa2\xeb\xd0\x11\x45\xa7\x73\xf5\xd3\xe8\x8e\xf3\xf4\xab\x4e\x4e\xb0\xa0\x63\x2d\xee\xd4\x3d\x91\x54\xc4\x7e\x3f\x12\xb9\x48\x91\xc4\x36\x4f\x52\x66\x18\xa5\x54\x2e\x0c\x12\x71\x26\x88\x38\xec\x45\x02\x78\x26\x9f\x32\x48\x14\xd9\x9b\x2a\x24\xc8\x75\xb6\x99\x31\xda\x75\xd4\xda\x5f\x88\xe4\x08\x5b\x0f\x14\x65\xbc\x2c\xb2\x0c\x14\xa4\xef\x36\xb2\x10\xc6\xa7\xaf\xaf\xab\x31\xd0\x18\x67\x58\x3d\x31\xbe\xb2\xb2\x56\xe6\x71\x62\xe7\x95\x53\x3d\x77\x09\x0a\xb1\x59\x83\x68\xd8\xcd\x95\xba\xf2\xc8\x0b\x19\xbf\x6e\xf5\x9a\xb4\x09\xd1\x92\xa2\x85\xae\x08\x54\xdd\x4b\x75\x7b\x72\xf0\xae\x02\xda\xe6\x3c\x01\xd4\x35\xb5\xdb\x0e\x93\x33\x4c\x54\xac\xe0\x4f\xbb\xd3\x0c\x26\xbc\xb2\x83\x2b\x6a\x9c\xa4\x41\x98\x70\x97\x14\x8c\x4a\xa4\xec\x9c\x37\x5c\x43\xeb\x04\xef\x4b\x34\x3e\xae\xcf\x9e\xdc\x8f\x16\xed\x90\xe1\x6d\x1d\xdc\xc0\x6f\x76\x3b\x8f\x35\xd3\xc1\x12\x40\x04\x75\xc0\xd3\x10\xcf\x7a\xbe\x3b\xd2\x6f\x54\x19\x90\x46\x7a\xcd\x54\x85\x09\xef\x32\x2d\x26\x22\xe6\xfa\xca\x0e\x90\xd6\x77\x6e\xff\x8d\xb0\x2b\x5d\xb8\x31\x21\x3d\xbc\x12\xf8\x6c\x78\xb2\x6f\x91\x88\x37\x32\xf9\x7a\x6b\x85\x23\xc6\x84\xf5\xd7\xbf\xe5\xaa\xee\x22\xd6\x40\x3e\x67\xe6\x1d\x69\xfa\x65\x05\xe6\x67\x96\x7c\x95\x59\xf6\x1e\x72\xb6\x43\x27\x27\x47\x0e\x16\x4c\x0c\x85\xd3\x33\x78\xd3\x44\x3b\x91\x42\xcb\x1c\xe2\x5c\xae\x50\xf8\xa9\x69\x6e\x0f\x50\x2e\x82\xf0\xc4\x9c\x84\x81\xb6\xd6\x54\x87\x98\x48\xf4\xac\xb8\xf7\x9d\x31\x19\x55\x12\x43\x27\x47\x58\xf6\x8a\x99\x75\xac\x98\x48\xe5\x06\xe1\x93\x33\xe2\x6e\x8a\x83\x87\x0d\x22\xe8\x0f\x44\x53\x20\x8a\x0e\x5e\x37\x9a\x39\xd5\x30\x59\x2e\xa5\x42\xf5\xe7\x86\x0b\x64\xbe\x97\xdf\xd7\x85\xad\x7c\x44\x82\x68\x4c\xd4\x13\x66\x1d\xde\xeb\x7b\x1e\xf2\x0c\x75\x70\x01\x4f\x28\x85\xb2\x9c\x9e\x43\x59\xc2\xf9\xdb\xa7\xc2\x44\x5b\x84\x83\x2c\x11\x01\x17\x06\x56\xa0\x82\x25\x98\x47\x9b\x8a\xd3\x80\x89\x34\x78\x4b\x82\x95\x34\xb3\x20\x3c\x01\x3c\x5f\x37\xb7\xc8\xce\x0e\x80\x3e\x63\x1c\x6a\xbc\x0c\xf8\xc2\xbd\xcb\x17\x6d\xba\xfb\x8e\x64\x85\xfa\x17\x17\xd2\xbb\xdb\x77\x2f\x18\x6e\xdf\xf8\xfc\x0c\x1c\xb5\xe4\x90\x06\x69\x01\x81\x91\x7e\xd7\x04\x7a\xb8\xaf\x42\x5c\xe1\x8a\xac\xa3\x68\x8d\x3a\x81\x24\xe1\xc1\xc9\x90\xec\x57\x60\x66\x9d\x1c\xee\xea\xaf\x63\xbe\x1f\x9b\x72\x45\x74\x1f\xec\x49\xd4\x3f\xaf\x21\x38\xb4\x0f\xb6\x4a\x6e\x41\x99\x5d\xc0\x75\x6d\x7e\x2d\x45\xbe\x0b\x71\x55\xe1\xc1\x03\x65\xeb\x97\xc5\xd2\xde\xd6\x79\x71\xe4\xad\x58\x73\xf7\xfa\x81\x1a\x76\xcf\x1a\xa7\x8f\x5a\x2f\x66\x7f\x9c\xfe\x71\x7a\xea\x0f\x99\x5a\xef\x4e\xe4\xd3\xec\xf8\x20\x1e\x58\xce\xd3\xa0\x50\xb9\x75\xf2\x81\xa7\x96\x95\x5a\x71\xba\xb0\x4d\x59\x8d\x3d\xf3\x9f\xb8\x22\x9d\xf1\x7f\x87\xa5\x63\x39\xbb\xa3\x7a\x47\x45\x77\x32\x10\x3f\x32\x25\x50\xf8\x1e\xb6\x0a\x12\x66\x20\x9d\x05\x1f\x73\x60\x1a\x82\x42\x43\xd0\xbc\x51\x14\x2a\x27\xc1\xe1\x69\x22\xc4\xdd\x91\x90\x8b\xc4\xcd\xd2\xe6\x48\xec\x1e\x9f\x3f\xfa\xb0\x22\x43\xc2\xc3\x7b\x6d\x48\xf6\x0f\x2c\x2f\xa0\x96\xc1\x35\x3d\x31\x92\x50\x8d\x24\x26\x39\x1d\x79\x14\xbf\xdd\x6d\x96\x32\x8f\xa2\x50\xd7\x1f\xc3\x8a\x98\x1b\x50\xcc\x48\xb5\x18\xbb\xc6\xf8\xa7\xef\x6a\x36\x52\x09\x51\xf4\xcc\x70\x56\x17\x0a\x6d\x54\x91\x18\xa9\x28\xa5\xce\xbe\x68\xbc\x98\xb5\xd0\x76\xed\x42\x59\xcf\x3b\xa4\x94\x76\xdf\xde\x5b\x50\xb3\xe8\x58\x67\xb9\x95\x3f\xa3\x83\xfb\x77\x77\x6c\xd4\x6e\xcf\xa8\x88\x2f\x72\x0e\x56\x0c\x88\x38\x69\xbe\xaa\x84\x99\x64\x8d\x0a\xbc\xaf\x18\x65\x65\xf9\xc8\x45\x2a\x1f\xe3\x76\xa5\x5b\xcb\x95\xfc\xbb\x35\xd6\x61\xde\xf6\xdf\x5c\x84\x94\x5b\x92\x51\xbf\x62\x99\x02\xf8\x1b\x16\xbd\xd2\x6c\x4b\xd6\x74\x74\x45\x49\xda\xc7\xe2\xfa\xb6\xbe\x4e\x91\xdd\xd0\xfe\x8b\x9f\x23\x59\x0e\x6b\xde\xbb\x70\x90\xaf\xc3\x0a\x37\x22\xb9\x18\xda\xdf\x29\xc5\x76\x64\xd5\x37\x67\x52\x5d\xb2\x64\x4d\x36\xb4\xae\xee\x4a\x23\xab\x3d\xe6\x43\xab\xd7\x1f\x65\x89\x8e\xd4\x1c\x51\xf5\x5e\x7f\x9e\x9f\x9f\x4f\x89\xa4\xd7\xc5\x66\x09\xea\xa0\x95\xef\xce\xee\x71\x59\x4e\xe7\x56\xd2\x4a\x3a\x3d\x97\x8b\xfa\x10\x48\x80\xe7\x48\xe2\x59\xe7\x44\x90\x98\x4c\xcf\x65\x14\x21\x79\x42\x0d\x9e\x9b\x73\x39\x97\x27\x27\x98\x67\xc8\xca\xd0\x9a\xef\x9c\x12\xb8\x93\xf7\x94\x52\xc0\x8d\xe6\xf2\x9b\xfe\xf5\x59\x85\x49\x2f\x02\xee\xd7\x1c\x14\x8a\x55\xbb\xdc\x21\xb9\xbb\xc7\x71\xc6\x6c\xea\xee\x50\xf8\xea\xb1\xf9\x01\x26\x24\x77\xe1\xab\xf6\xd1\x34\x24\xe1\xab\xbf\xec\x1f\xe3\x0e\x46\xfb\xd9\x36\xf5\x8c\x11\x12\xd5\x85\xea\x0e\xf1\x3c\xd2\xef\xa3\x40\x1a\xd4\x83\x95\x84\x63\x03\xdd\x85\xaf\x72\xb9\x0a\x09\xef\xb6\x1b\x32\x58\xaf\xd9\x78\x20\xfc\xaf\x4e\xaf\x7b\xb3\xee\x46\xe7\x1e\x13\xd3\xe9\x48\x5f\x0c\x27\x26\xd0\xfc\x70\x45\xbb\x3d\x2b\x3c\xff\xee\xf4\xf4\xff\x02\x2d\x0b\x95\xc0\x15\xdb\x6e\xb9\x58\xfd\xf6\xe9\x57\xfa\xc4\x05\x2b\x2e\xe2\x3f\xad\x0a\xd9\x7e\xf7\xef\x00\x00\x00\xff\xff\x1b\xb6\xdc\x05\x0b\x1c\x00\x00")

func assetsJsAngularWebsocketMinJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsJsAngularWebsocketMinJs,
		"assets/js/angular-websocket.min.js",
	)
}

func assetsJsAngularWebsocketMinJs() (*asset, error) {
	bytes, err := assetsJsAngularWebsocketMinJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/js/angular-websocket.min.js", size: 7179, mode: os.FileMode(420), modTime: time.Unix(1481854601, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsUiHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x57\x5f\x73\xda\x38\x10\x7f\x2e\x9f\x42\xa7\x9b\x2b\xe4\xc1\xa5\xd0\xd2\xbb\x69\xc1\x37\xb9\x3f\x9d\xbc\x30\xe9\x24\xf7\x05\x84\x25\xdb\xba\x93\xb5\x1e\x59\x0e\x64\x32\x7c\xf7\x1b\x09\x9b\xd8\xb2\x4c\xa0\xa1\x69\x79\x41\x62\x97\xdd\xdf\xfe\xf6\x8f\xbd\xf3\x95\x0a\x07\x73\xca\xef\x50\x24\x48\x51\x2c\x70\x04\x52\x13\x2e\x99\xc2\xe1\x00\xa1\xa6\xe4\xdf\x32\x5b\x81\x56\x20\x71\x38\x18\x20\xd4\x16\x2a\x58\xd7\x3f\xb7\x05\x11\x88\x60\x53\x04\x93\x29\x32\xa7\x8c\x06\x1f\xea\x43\x5e\x16\x69\xf0\x0e\xad\x99\x10\x38\x1c\xbc\x72\xac\x21\x05\xeb\x20\x62\x52\x33\xc5\xa8\xc5\x52\x7f\x3c\xd6\x3f\xb4\x15\xd2\x59\x78\x03\x90\xa1\x7f\x58\x96\x33\x45\x74\xa9\xd8\x7c\x9c\xce\x5a\x3a\x63\xca\xef\x4e\xb6\x7a\x5d\x6a\x0a\xa0\xce\x6d\x78\x12\x3e\x3c\x20\x5d\x68\xa2\xdf\x44\xa5\x52\x4c\x6a\xe3\x00\x6d\xb7\xaf\x29\x4b\x3e\xcd\xc7\xe9\xa4\xa9\xef\xea\x5e\x95\x19\xa7\x5c\xdf\xa3\xed\xf6\x17\x54\x5f\xce\x87\x08\x76\x31\x1f\x42\x54\x79\x78\xb5\xff\x3e\x36\x95\x96\xd4\x25\x50\xf6\x11\x35\x7f\x35\x82\x22\x27\xb2\xb6\xa1\xd9\x46\x07\xb9\xe2\x19\x51\xf7\x18\xc9\x24\x28\x52\x58\x2f\xf0\x0e\x5f\xa1\x49\xc2\xd0\x62\x81\x26\x38\x1c\x5d\x31\xa2\xd1\xad\xfd\x65\x72\x31\x1f\x1b\x23\xe1\xf3\x2d\x4f\xdb\x96\xa7\xb5\xe5\xa6\xd9\x27\xbd\x50\x22\x13\xa6\x1a\x4e\x56\x02\xd6\x4c\xbd\x21\xe5\xc6\xd8\xc6\xe1\xe8\xb2\xdc\x20\xeb\xe6\x32\xd2\xfc\x8e\xf9\xbc\xb8\x9f\x16\x97\x6e\x25\x36\xb2\xb0\xd2\x32\x48\x14\x94\x39\x46\x0a\x04\x5b\xe0\xdd\xa5\xcd\xcd\x9c\x54\xc2\x55\xa9\x35\x48\xdc\xf8\x33\x32\x06\x28\x8b\x49\x29\x34\x5a\xe9\x2c\xd8\x14\x36\x12\xca\x0b\xb2\x12\x8c\xd6\x94\x65\x40\x2d\x63\x43\x88\xe3\x21\x6e\x73\x22\x93\x20\x12\x3c\xfa\x6f\x81\x0b\xa6\x4d\xd2\x47\x56\xeb\xa2\xad\xe6\xa6\x0b\x5d\xc7\x71\x1b\xe5\x98\x7c\x3b\xd8\xa4\xd4\xe0\xc5\xbd\xcb\xa3\x47\x19\xfd\x8e\x86\xc6\x4b\x55\x44\x43\xf4\x71\x77\xaf\xbc\x1e\x41\x82\x35\xe3\xb0\xd0\xa5\xe1\xb2\xd4\xf0\x72\x3c\xa4\x8c\xf8\xa1\xfb\x78\xb0\xca\x35\x0f\xbb\x32\xff\x0a\x1a\xac\x95\x27\x69\x30\xfd\xf1\x1c\x1a\x9e\x0e\x3d\x02\x10\x47\x87\x6e\x95\x9f\x5f\x02\xd6\xcc\x13\xb1\xff\x09\x20\x9c\x60\xbb\x93\xdd\x4e\x1d\x67\x8c\xb5\xf2\x94\x97\x59\x3e\xc4\xa1\x1d\x33\xe6\x8c\x40\x8a\xfb\xee\xa0\x3c\x6c\x88\x09\x16\x69\xc5\xa3\x21\x0e\xff\xae\x8e\xc8\x18\xf7\x1a\xab\x40\x3e\xde\x0f\x3f\x1c\x5e\x79\x5f\x2c\x3c\x8f\xac\x5f\x71\x67\xb2\xa7\xb3\xf0\x33\x91\x68\xe4\xa0\xaf\x06\xed\xee\xeb\xe6\xcb\xd2\xc4\xf0\x16\x87\xd7\x71\x5c\x61\xdd\xeb\xa7\x9c\xb2\x5e\xfd\x87\x07\xd4\x91\x6c\xb7\xe8\xe6\xcb\xb2\x32\x73\xe1\x4e\xe0\xaf\x19\xc2\x27\x36\x72\x91\x79\xab\x39\x26\x72\xd9\x9e\x69\xed\xb2\xfb\x4c\xe4\x6d\xce\x18\xdd\x4f\x9f\x0e\x08\xcf\xc4\xf1\x75\xdb\xf9\xe1\x0a\x58\xf7\xa3\x35\xc2\x2e\xd8\xd7\x3f\x6f\xa6\xb3\x3f\x7e\xfb\xf4\x1d\xd0\x66\x8c\xf6\xa3\x35\x42\x2f\xb5\x35\xe0\xef\x08\x3c\xe5\x49\xda\x8f\xdc\x4a\x8f\x82\x7e\x7c\x08\x9d\x69\xb5\x1f\x0d\x9e\xee\x7e\x6f\x9a\x3e\x9d\x85\x57\x20\x28\x72\xfa\x73\x17\x4c\x0a\x82\xfe\x55\x2a\xa2\x39\xc8\x25\x97\xc5\x62\xf1\x16\x87\xa3\xfd\x6b\xab\x2b\x36\xad\x9a\x71\x59\xbf\x54\xfd\x08\xad\xfa\xd3\x23\x52\x27\x11\x26\xea\x51\x4c\x44\xc1\x2e\xaa\x39\xf5\x4d\xea\xa1\xcf\xb3\x56\xa5\x75\x2c\x5d\xbf\x4e\x0a\x1b\xd7\xee\x71\xbf\x0c\x36\xff\xe3\x5b\x3d\xea\x7d\xf0\xdd\xe3\x62\x28\x44\x6b\x23\xf1\xef\x93\xd5\xe2\xe8\x4f\xa0\x77\xed\x70\x9f\x2d\x9d\xdd\x74\x32\xf5\x94\xbc\x29\x43\xf3\xf0\x45\xb7\x4c\xe7\xc0\xa5\xf6\xd5\x4e\x1d\x64\x8f\x8b\xf7\xfd\x1e\x8e\xcf\xa3\xac\xf3\x58\xa5\x8b\xcb\xc8\xe0\xaa\x61\x8d\x26\x9d\x86\x6d\x6d\x21\x89\xb8\xcf\x53\x1e\x81\x44\xfb\x53\xa0\x15\x27\x32\x11\x2c\xd0\x90\xe3\xd0\xb3\x30\xf5\x54\xc0\xa9\x61\xce\xd3\x69\x63\xc7\x6d\x80\x6e\xac\x94\xd3\xb3\x38\x3a\x1b\x9b\x41\x97\xce\x13\xf8\x5c\x81\xd6\x90\x9d\x42\xa9\xf7\xda\x77\x39\xd8\x47\xa7\x75\xce\x91\xcb\xfa\x49\x7d\x63\x27\xb7\x5d\x98\xbd\x2d\xe3\x46\xfc\x52\xf9\x35\x90\x0e\x74\xcb\xb9\xdb\xe5\x39\x61\x36\xfb\x25\x6d\xc0\xf6\xf7\xcb\x0f\x41\xe8\xcb\x34\xcc\x71\xed\x51\x1f\xab\x43\xf5\xf5\x7f\x00\x00\x00\xff\xff\x45\x6e\x05\x31\xed\x14\x00\x00")

func assetsUiHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsUiHtml,
		"assets/ui.html",
	)
}

func assetsUiHtml() (*asset, error) {
	bytes, err := assetsUiHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/ui.html", size: 5357, mode: os.FileMode(420), modTime: time.Unix(1518988751, 0)}
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
	"assets/app/app.js": assetsAppAppJs,
	"assets/css/app.css": assetsCssAppCss,
	"assets/index.html": assetsIndexHtml,
	"assets/js/angular-websocket.min.js": assetsJsAngularWebsocketMinJs,
	"assets/ui.html": assetsUiHtml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"app": &bintree{nil, map[string]*bintree{
			"app.js": &bintree{assetsAppAppJs, map[string]*bintree{}},
		}},
		"css": &bintree{nil, map[string]*bintree{
			"app.css": &bintree{assetsCssAppCss, map[string]*bintree{}},
		}},
		"index.html": &bintree{assetsIndexHtml, map[string]*bintree{}},
		"js": &bintree{nil, map[string]*bintree{
			"angular-websocket.min.js": &bintree{assetsJsAngularWebsocketMinJs, map[string]*bintree{}},
		}},
		"ui.html": &bintree{assetsUiHtml, map[string]*bintree{}},
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
