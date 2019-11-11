// Code generated by go-bindata.
// sources:
// data/data.go
// data/substitutions.txt
// data/weights.txt
// DO NOT EDIT!

package data

import (
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

var _dataDataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func dataDataGoBytes() ([]byte, error) {
	return bindataRead(
		_dataDataGo,
		"data/data.go",
	)
}

func dataDataGo() (*asset, error) {
	bytes, err := dataDataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/data.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1550004291, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSubstitutionsTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcf\xd1\x0d\x03\x31\x08\x03\xd0\xff\x93\x6e\x87\x8e\x80\x81\x10\xd8\x7f\xb1\x2a\x0e\x95\xca\xe7\x93\xef\xb0\x53\x96\xd8\xf6\x29\x4b\xc1\x7e\x9f\xb2\xd4\x2c\x32\xf2\xb2\x36\xb9\x83\x0c\x91\xc3\x00\x9a\x7a\x68\x6e\x4d\x67\x2a\x3f\xe6\xa1\x4b\x36\x79\xd9\xd5\x2f\x31\x7a\x03\xf1\xdf\x1b\xe0\xbf\xb1\x3a\x55\x19\x45\xca\x5e\x98\x34\xf9\x31\xb2\x57\x39\x2f\xab\xf7\x66\xdf\x63\xb3\xe7\x4c\xb9\xca\xaa\x57\x2d\x90\xb6\x9a\x36\x36\xaf\xfb\x40\xc5\xfb\x7c\x03\x00\x00\xff\xff\xa9\xbf\x54\x83\x3b\x01\x00\x00")

func dataSubstitutionsTxtBytes() ([]byte, error) {
	return bindataRead(
		_dataSubstitutionsTxt,
		"data/substitutions.txt",
	)
}

func dataSubstitutionsTxt() (*asset, error) {
	bytes, err := dataSubstitutionsTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/substitutions.txt", size: 315, mode: os.FileMode(420), modTime: time.Unix(1549922049, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataWeightsTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x5d\xcb\xae\x6b\x29\x8f\x9e\x9f\x67\xa1\x25\x0c\x18\xf0\xb0\x5b\xff\xb0\x4a\xfd\xfe\x6f\xf3\x0b\x9b\xb3\x43\x12\xe3\xf5\x6d\x95\x6a\xfb\x0c\xbe\x8b\x61\x71\x67\x25\xc9\x94\x73\x6e\x29\x53\x1f\xc4\xe9\xdf\xff\xff\x0f\x51\xca\xc7\x7f\x33\x8d\xd4\x13\xa7\x96\x6a\x2a\x89\xfe\xe4\x66\x78\x0b\xff\xf9\xbf\x7f\xfe\xf7\x9f\x47\x3c\x65\xc5\x53\x73\xf4\x73\x88\x37\xfd\x92\xe8\xf6\xbf\xe1\x39\xed\x80\xe9\x17\xd3\x2f\xf5\x07\x5f\x14\x29\xaa\x39\x15\x3b\x12\xa7\xfe\x17\x6f\xe5\xad\x82\xea\xff\xe0\xb1\xfc\x9b\xe5\xc3\x02\xe6\x33\x8a\xe2\x47\x05\x9f\xd7\xb0\x7c\x06\xfa\x7c\xa7\xe9\xcf\xa2\xf8\xfc\x96\x73\x4e\xbd\xa5\x5a\x12\x75\x4d\x6c\x53\x5a\x5e\x45\xa0\x2a\x7f\x8b\x4c\xa9\xaa\xa4\xa4\x9a\x48\x93\x5f\x16\x9c\x4a\x6a\x7f\x72\xa3\x8d\x6f\x0d\x2a\x32\xe7\xbc\xf0\xac\x4f\xee\xbb\x08\x65\x3b\xcc\xa5\xbe\xf1\xa5\xa4\xcc\x53\xc4\x7b\x64\x5f\xf8\x91\x73\xcd\x69\x87\x17\xde\x2a\x66\x26\xd1\xcc\x7b\x1a\xfa\xef\xff\x31\x42\x33\x42\x83\x09\xcc\x69\x07\x90\xd0\x7b\xda\x01\x23\x10\x29\x61\x85\x67\x42\xa2\xe2\x70\x16\xa3\x6a\xad\xf8\x9c\xba\x38\xa5\x29\x67\x05\x2c\xb1\x56\x95\xb0\x02\x46\x98\x59\x09\x2b\x60\x04\x19\x4a\x58\x01\x22\x90\x3d\x3e\x82\x1f\x1f\x65\xe9\x69\x07\x8c\x40\x45\x1d\xa8\xa0\x0e\xa5\xa8\xc3\x0a\x18\xa1\x6a\x2d\x69\x00\x09\xa4\x29\xad\x00\x12\xd8\x08\x8c\x12\x9a\x28\x61\x05\x8c\xc0\x56\x4b\x0c\xd7\x52\xef\x4a\x58\x01\x23\xc8\xd4\x5a\x5a\x01\x22\xb4\xc6\x8b\xa0\x01\xed\x46\x9f\x1c\xa0\x1b\xcd\x5c\x96\xb4\x85\xbf\x9c\x73\x4e\x21\x1d\x00\xf7\x5f\xc5\x97\x99\x76\x40\xf0\x3d\xe7\xa5\xaf\x01\xc3\xaf\x47\x61\x01\xc3\xaf\x8a\xb5\xf0\x39\xc4\x7a\x15\x3b\x7b\x16\x33\x90\xd3\xa0\x7e\x4a\xef\xbf\xf3\x4f\x9e\xa2\xa3\xfe\x1a\xc4\x05\x49\x49\x72\xa6\x9a\x76\x78\x74\x58\x78\xca\x9c\x76\xc0\xf0\x34\xd3\x0e\xaf\x22\xbf\x26\xcf\x8f\x89\x54\x32\xad\x2e\xbd\x82\x80\xfa\x95\x14\x5f\xbf\xab\x54\xbe\x27\x6a\xc9\xb4\xfa\xe7\x0a\x28\x7e\x1a\x7e\xa2\xe5\x15\xc3\x4b\xc7\xf0\x65\x2d\x0c\x2d\x60\xf8\x35\x4b\x5b\x80\xf0\x95\xb5\x3e\x2b\x5a\x5e\xce\x9a\x3f\x0b\x58\xff\xc3\xf2\x1f\x68\xfe\xc3\xda\xcf\x40\xdb\xcf\x20\xcd\x67\x10\xd0\x03\x0c\x6f\xfa\x84\xea\x17\xd3\x2f\x60\xfd\x8c\xaa\xf5\x39\x2a\x5a\x3f\xd6\x1e\x06\xd2\x83\xff\x64\xa1\xae\xf5\xaf\xe1\xb3\xff\x92\x8f\xa7\xb4\x03\xa4\x3f\x9a\xea\x8f\x56\x31\xfd\x69\xf9\xcc\xfc\x3d\x9e\xf8\x78\xd2\x7c\x66\x67\x00\xbf\xf6\x51\x39\x27\xca\x94\xdd\x55\xe7\x67\xfd\x53\x26\xca\xb4\xf0\x4d\x26\x86\x67\xd3\xf7\x57\xb5\xdf\xf8\xd2\x14\x3f\x50\xfc\xb4\xfc\x67\x1e\xdf\xf8\xef\xf1\x6d\x01\xa7\xe1\xbf\xf4\xbd\xf1\x90\xf2\x24\xd5\x17\x30\x1f\xb2\xfa\x24\x29\x33\x1f\x1b\xcd\xdb\x66\x6a\x31\xa4\xcc\x62\x8c\x8a\x33\x78\x31\xe4\x67\xe7\xf2\xc0\x28\x96\x55\xc9\xd2\xe9\x28\x05\xe9\x6e\xa2\x6b\x5b\x10\x9d\xf6\x6d\x6f\xa4\xc0\x9a\xa8\x94\x9c\x05\xc2\x97\x4c\x6b\x59\x51\x28\x63\xfa\x85\xb2\xea\xff\x9d\x65\x9e\xf1\xd5\xf4\x2b\x63\xf8\x42\xaa\x5f\x04\xd4\xd7\x56\xb7\xe0\x18\xbe\x59\x7d\x36\x10\x5f\xab\xe2\x6b\xf5\x5a\x11\x65\x5d\x71\x34\x25\xf2\x26\xb4\x4c\x25\xad\xf0\xb1\x79\x5c\x5b\x40\x79\x43\x6b\xd3\x5d\x0c\x2a\x6b\x95\xa6\xe1\x64\xd0\x6e\xa8\x5f\x0c\xdb\xa0\x12\x6b\x0f\x32\xc6\x6a\xc6\x7f\xd3\xfa\xea\x0a\x2c\x53\xf1\x1a\xbe\xbb\xc2\x1b\x3a\xe5\x85\x17\xc3\x0b\x8c\x27\xc3\x13\x84\xef\x96\x7f\xa7\x5c\x06\xb0\xb4\x51\xa0\xa4\x1d\xbc\x0d\xf6\x37\xbe\x9a\x7e\x23\x50\x9f\x0d\xff\xb6\xf5\x08\xf1\x9c\x76\xc0\xf0\xdd\xf4\x3b\xaa\xdf\x4d\xbf\xa3\xfa\xc3\xf4\x07\xaa\x3f\x4c\x7f\xa0\xfa\xd3\xf4\x27\xaa\x3f\x4d\x7f\xa2\xfa\x62\xfa\x02\xea\xeb\xd0\xde\x4b\x2e\xc8\xd2\x78\x01\xb5\x3d\xb4\x9a\x31\x7d\x96\x35\x55\xf6\x9f\xdd\xcc\x13\xbe\x5b\xfb\x19\x1d\xd4\x5f\xfb\xa4\xb4\x03\xd0\x9e\x57\x3f\x2f\x69\x87\xc7\xa9\x2c\x51\x33\x0a\xa7\x1d\x70\x8a\xa4\x1d\x60\x4a\xed\x69\x07\x9c\x32\xd3\x0e\x30\x45\x48\x29\x82\x17\x9f\x72\x4b\x3b\xe0\x14\x2d\x3e\xb9\x83\xd8\x85\xc2\x5a\x7c\x6f\x4b\x74\xa1\xf0\x6a\x53\x3b\xe0\x94\x6a\x14\xe4\x44\x78\x01\xc9\x2c\x0a\x72\x42\xad\xc0\xa1\x78\x77\xa1\xf4\x89\x2f\xb6\xf0\xd4\xe0\x3d\x8e\xcf\xf1\xbe\x7f\x30\x1e\xcf\xa8\xff\x32\x5a\xda\x01\xf6\xf8\x61\xa0\x1e\xa5\xa7\x1d\x50\x8f\x17\x03\xf5\x60\x52\xc6\x18\xb0\xc7\x8b\x81\x7a\x0c\x51\x86\xe0\x1e\x2f\x06\xea\x21\xca\xf0\x5b\xba\xef\xf1\x62\x80\x1e\xc4\x73\x31\xea\x84\xcb\x71\x30\x40\x8f\x96\x6b\xd2\x00\xf4\x8d\x2f\x06\xec\x31\xd2\x0e\xb0\xc7\x0f\x03\xf5\x20\x6d\x25\x8d\xe0\x3e\x78\x30\x60\x0f\xed\x51\xad\x78\x37\x3f\x17\x8f\x1f\x06\xea\x51\x78\x31\x26\x32\xf6\x7c\x31\x40\x8f\x35\x9b\xaf\x7d\x0e\xc3\x1e\x07\x03\xf3\x20\xea\xc6\x80\xfb\xf9\xc9\x40\x3d\xb4\xd7\x52\x65\x6f\xea\xf0\x3d\x5e\x0c\xd0\xa3\xf2\x6a\xed\xa5\xbb\x87\x04\xae\xc7\xc1\xc0\x3c\xca\xd0\xd9\xa0\x16\xe4\x86\xf3\x8b\x81\x79\xd4\xd2\x48\x19\x8c\xf6\xf3\x93\x81\x7a\xf0\xaa\xdd\xca\x84\x96\xe3\x64\x80\x1e\x5c\x56\x56\xad\x0d\x74\xdc\x3d\x19\x98\x47\x6b\x63\x8d\xa2\xdc\xdc\x1b\x5b\xcf\xe3\x64\x60\x1e\xdc\xc6\x1a\xe1\x58\x18\x9d\xcf\x4f\x06\xe8\x21\xbc\x66\xe7\x4e\x05\xf6\x38\x18\x98\x47\xa7\x62\x0c\x46\xfb\xf9\xc9\x40\x3d\xb4\x95\xf4\x2a\xde\x42\xd4\xf7\x78\x31\x40\x8f\x2a\x65\x31\x06\x3c\xee\x9e\x0c\xd0\x63\x66\x65\xcc\x81\x3e\x8f\x93\x81\x7a\x8c\x55\xbb\x83\x06\x5c\x57\x07\x03\xf3\x18\x34\xd6\x08\x37\xfd\xed\x87\xe7\x71\x32\x30\x8f\x99\x65\xcd\x9c\x73\x14\x74\x6c\x3f\x19\xa0\xc7\x28\x2b\x2b\xc9\x15\x7d\x1e\x27\x03\xf3\x90\xb5\x13\x2c\x59\xa8\xa0\xf3\xc7\xc9\x00\x3d\xa8\xae\xd9\xe0\x72\xb0\xeb\x7a\x1c\x0c\xc8\xa3\xe6\x3c\x67\xda\xe1\xdf\xaf\xb7\x52\xd6\x0e\xdd\x6e\x1c\x6b\x9a\x6b\x4f\x54\x73\xad\x8a\xaf\xee\xae\xf6\x6b\x0f\xf5\x86\x7f\x7e\x6b\xa7\xd4\xcc\x33\xa7\x1d\x1e\x4f\x0c\x12\x95\x47\x0a\x7f\x52\xea\xa2\xf4\xb5\x7e\xb3\x00\x95\xe2\xc0\x43\xa5\x18\x59\xd2\x0e\x90\xfe\x81\x87\xf4\xe7\x28\x69\x07\x48\xff\xc0\x43\xfa\x52\x6b\xda\x01\xd2\x3f\xf0\x88\x3e\x65\x5a\xad\x42\x03\xa2\x7f\xe2\x21\xfd\x42\x2b\x1f\x0d\x90\xfe\x81\xc7\xf4\x8b\xe6\x53\x90\xde\xff\x81\x87\xf4\x2b\xaf\xf6\xa6\x01\xd2\x3f\xf0\x90\x7e\xeb\xab\xbd\x69\x80\xf4\x0f\x3c\xa4\xcf\xac\xe5\x5d\x01\xd2\x3f\xf0\x90\x7e\xb7\xf6\xd0\xd1\xf6\x73\xe0\x31\xfd\xa1\xe5\xed\xce\x95\x9d\xaf\xff\xc2\x43\xfa\xb3\x69\x7b\x9b\x0d\x6c\x9f\x07\x1e\xd2\x97\xc9\x69\x07\x48\xff\xc0\x23\xfa\x45\xe7\x19\x0b\x88\xfe\x89\x87\xf4\x8b\xd6\xa7\x06\x48\xff\xc0\x63\xfa\xb3\xa6\x1d\x30\xfd\x17\x1e\xd3\x17\x2d\x6f\x41\x56\xb6\x0b\xdf\xda\xaa\x7f\x0d\xbf\xc5\x43\xf9\xf0\xa0\xb4\x03\xa4\x7f\xe0\x21\xfd\xbe\xd6\x36\x16\x20\xfd\x03\x0f\xe9\x8f\xbc\xc6\xb7\x32\x90\x53\xd2\x0f\x3c\xa6\x5f\xb4\x3e\x07\x74\x92\xfc\x8e\x87\xf4\xa7\xce\x2f\x1a\x20\xfd\x03\x0f\xe9\x4b\xd5\xf6\x2f\xd0\xbb\xd4\xef\x78\x44\xbf\x66\x5d\x62\xd5\xec\x2e\xb1\xbe\xf5\x4f\x3c\xa4\x4f\x75\xd5\x67\x7d\xbf\x3c\x0f\xf4\x0f\x3c\xa6\xdf\x4a\xda\xe1\x7b\x95\x5b\x73\x5a\xcd\x71\xed\x32\x12\xf5\x44\xfa\xa2\xb4\xb2\xb4\x17\x68\xc0\xb2\x7a\xe1\xb1\xac\x74\xed\xad\x01\xd3\x7f\xe1\x21\xfd\x52\x35\x9f\x15\x20\xfd\x03\x0f\xe9\xd7\xb6\x7a\x99\x06\x48\xff\xc0\x43\xfa\x4d\xf7\x0e\x1a\x20\xfd\x03\x8f\xe9\x73\x4f\x3b\x7c\xb7\x0a\xe7\x8d\xfc\x85\xd5\x89\x40\xc3\xf3\x85\x66\xa9\x95\xcd\x82\x91\xeb\x84\x0f\x3c\x54\x84\xae\x0b\x27\x0d\x90\xfe\x81\xc7\xf4\x45\xcb\xbb\x02\xa6\xff\xc2\x43\xfa\x83\xb5\x63\xae\x00\xe9\x1f\x78\x44\xbf\xe5\xb9\x9a\xb4\x06\x44\xff\xc4\x43\xfa\xa4\xe5\xd5\x00\xe9\x1f\x78\x48\xbf\x68\x79\x35\x40\xfa\x07\x1e\xd2\xaf\x45\xcb\x5b\x91\xe3\x98\x0f\x3c\xa6\x3f\x46\xda\x01\xd3\x7f\xe1\x21\x7d\x1e\x6b\xa2\xd1\x00\xe9\x1f\x78\x48\xbf\xf7\xd5\x1f\x35\x40\xfa\x07\x1e\xd2\x1f\xba\x70\xd5\x00\xe9\x1f\x78\x48\x7f\xea\xc6\x53\x03\xa4\x7f\xe0\x21\x7d\xdb\x38\x34\x74\xa3\x71\xe2\x11\x7d\xce\xba\xb1\xd5\x80\xe8\x9f\x78\x48\x9f\xba\xe2\x57\x80\xf4\x0f\x3c\xa4\x5f\xfa\xea\x8f\x1a\x20\xfd\x03\x0f\xe9\xd7\xb2\xc6\x13\x0d\x90\xfe\x81\x87\xf4\x1b\xaf\xfe\xae\x01\x99\xef\xb8\xb1\xa4\x1d\xa0\x7c\x0e\x3c\x94\x0f\x93\xe2\xa1\xab\xa8\x0f\x3c\xa4\xdf\xc7\xea\xbf\x1a\x20\xfd\x03\x0f\xe9\x0f\xd2\xfa\x1c\x84\x8d\xb7\x27\x1e\xd3\xd7\xf5\x83\x06\x4c\xff\x85\x87\xf4\xe7\xd4\xfa\x5c\x01\x5c\x32\xb1\xe8\x5e\x55\x03\x94\xd2\x81\x47\x52\xea\x39\xaf\x22\x68\x40\xf4\x4f\x3c\xa4\x4f\xda\x84\x34\x40\xfa\x07\x1e\xd2\x2f\x7a\x76\xaa\x01\xd2\x3f\xf0\x98\x7e\x1b\x69\x87\x97\xfe\xed\x45\xe0\x05\x14\xd3\x07\x97\x34\x27\x1e\xca\xa7\x95\x35\xc4\x69\x80\xf4\x0f\x3c\xa4\xcf\x45\xcb\xcb\x05\x5b\x72\x9c\x78\x4c\x5f\x77\x19\x1a\x30\xfd\x17\x1e\xd2\xef\x7a\xd6\xa7\x01\xd2\x3f\xf0\x90\xfe\xe8\x6b\x48\xd1\x00\xe9\x1f\x78\x48\x5f\xb2\xd6\xa7\x20\x6f\x05\x7d\xe0\x31\x7d\xbd\xcb\xd0\x80\xe0\x87\x9d\x3d\x0e\xf4\xac\xf2\xc4\x43\xfa\xa5\xaf\xf1\x4a\x03\xa4\x7f\xe0\x21\xfd\xca\x8a\xaf\xc8\x1b\x1e\x1f\x78\x48\xbf\xe9\x92\x60\x40\x6f\x3f\x7d\xe0\x41\xfd\x91\x76\x00\xf5\x7f\xf0\x90\x3e\x77\xcd\x67\x05\x48\xff\xc0\x43\xfa\x5d\xc7\x1f\x0d\x90\xfe\x81\x87\xf4\x87\x2e\xb9\x35\x40\xfa\x07\x1e\xd2\x9f\x43\xcb\x3b\xdd\xef\x35\x70\xf4\x0f\x3c\xa2\x3f\x73\x59\xfd\x65\xfa\x5f\x22\xf0\xad\x7f\xe2\x31\x7d\x1d\xdf\x34\x60\xfa\x2f\x3c\xa6\xaf\x4b\x5c\x0d\x98\xfe\x0b\x0f\xe9\xd3\x58\xfd\x51\x03\xa4\x7f\xe0\x21\xfd\xc2\xab\xbf\x68\x80\xf4\x0f\x3c\xa4\xaf\xef\xc3\xd4\xd9\x7e\x4e\x11\x4f\xc5\xf7\xcb\xf5\xfc\x81\x87\xf4\x5b\xdd\x78\xac\x7f\x9d\x78\x50\x5f\xdb\x03\xff\x1c\xd9\x3d\xe4\x7f\xe0\x21\x7d\xd6\x2d\xed\xe4\x9f\xbb\x9e\x07\xfd\x03\x8f\xe9\x8b\xa4\x1d\xa0\xfa\x39\xf0\x90\x7e\xd7\xbb\x06\x0d\x90\xfe\x81\xc7\xf4\x75\xbe\xd6\x80\xe9\xbf\xf0\x90\xfe\x20\x6d\xcf\x2b\x40\xfa\x07\x1e\xd2\x9f\x7a\x97\x37\xbf\xdf\xdc\x75\x9e\xf0\xc2\x4b\x5e\xe3\xb3\x06\x28\x9f\x03\x8f\xe4\x23\x59\x6f\x19\x34\x20\xfa\x27\x1e\xd2\xd7\xf7\xf8\x2c\x40\xfa\x07\x1e\xd3\xaf\x86\x6f\x58\x7f\x39\xf1\x98\xbe\x8e\xff\x42\x3f\xb7\x12\x4f\xfa\x2f\x3c\xa4\x5f\xe6\x6a\x9f\x52\xc0\xf1\xe4\xc4\x83\xfa\xac\x78\x01\xf3\x3f\xf0\x98\xbe\x6c\x3c\x36\x1f\x9d\x78\x50\x5f\x9f\x57\xfd\x79\x37\xe3\x29\xff\x17\x1e\xd2\xaf\x3d\xa7\x1d\xa0\xfc\x0f\x3c\xa6\x3f\x0d\x0f\xde\x75\x9e\x78\x48\xbf\xe9\xdd\xa5\x06\x48\xff\xc0\x43\xfa\xac\x77\xd3\x1a\x20\xfd\x03\x0f\xe9\x77\xbd\x4b\xd5\x00\xe9\x1f\x78\x48\x7f\xe8\xfa\x4a\x03\xa4\x7f\xe0\x9f\xf5\xab\x7d\x46\x4b\xc3\x39\xbe\x2d\x35\x49\xf3\xeb\x75\xc1\xf2\xc8\xa0\x0f\x86\x18\x63\x1a\x43\x70\x8f\x80\xe1\x7b\xb0\x65\xe5\xd6\x13\x95\xf9\x7d\x18\x57\x73\xa6\xda\x92\x86\x89\x26\x16\x32\xdc\xc4\xa8\x53\xda\xe1\xd1\x43\xf1\xa3\xa7\x1d\x10\x3c\xe5\xbc\xf4\xe9\xfd\x93\x85\x51\x19\x62\x86\x57\x06\xca\xb9\xa5\x1d\x60\x8f\x80\x71\xf1\x18\x69\x07\xd8\x23\x60\xf8\x1e\x54\xd2\x0e\xa8\x47\xc4\xf0\x3d\x8a\x31\xdc\x0d\xf2\xa5\x21\xea\x47\xe2\x77\xf8\x05\xa9\x19\xa9\xe1\x35\x16\x31\xfc\xd2\x34\x49\x3b\xc0\x1e\x01\xc3\xf7\x60\xab\x31\xc6\x9f\x4a\xc4\xf0\x3d\x06\x2b\x03\xe9\x55\xdb\x23\x62\xb8\x1e\xa4\x83\x96\x06\xd0\x23\x64\xf8\x1e\xd6\x1e\x09\x6f\xc1\x21\xc3\xf7\xb0\xc6\x48\x05\x6e\x57\x21\xc3\xf7\xa8\xc6\xa8\xbf\x69\xf0\xd4\xac\x28\xee\x5c\x7b\x27\x59\x1d\x37\xfc\xa9\x44\x0c\xbf\x34\xac\xa3\x1d\xb9\x87\x92\xd7\xc4\x74\x62\x20\x68\x62\xd8\x89\x45\x0c\x3f\xb1\xde\xd3\x0e\xbf\x48\x6c\xe4\xb4\xc3\x6f\x48\x56\x05\x6f\xe7\x3a\x71\x69\x22\x86\x5f\x1a\xb1\xf2\x0b\x5e\x63\x11\xe3\xe2\x61\x59\x09\x5e\x8e\x88\xe1\x7a\x14\x9b\x1c\x4b\x86\x3d\x42\xc6\xc5\x43\x87\xe1\xf2\xf6\xb5\x06\x0f\x1e\x01\xc3\xf7\x20\x4e\x3b\xa0\x1e\x11\xe3\xe2\x31\xd3\x0e\xb0\x47\xc0\xf0\x3d\x8a\x36\xf8\x52\xe0\x76\x15\x32\x7c\x8f\x6a\x4f\xb0\xc2\x4b\xa8\x90\xe1\x7b\xb0\xb6\xf6\xc2\x78\x39\x22\xc6\xc5\xc3\x5a\x09\xc3\x0b\x82\x90\xe1\x7b\xd8\xc8\x50\x06\x5e\x57\x11\xc3\xf7\x98\xf6\x04\x27\xde\x3f\x22\xc6\xc5\xa3\xa7\x1d\x60\x8f\x80\xe1\x7b\x88\xf5\x28\xc1\x3d\x22\xc6\xc5\xc3\x9e\xa0\xe0\xcf\x3c\x62\xb8\x1e\xd5\x36\x25\x15\xdf\xc6\x84\x8c\x8b\x07\xa7\x1d\x60\x8f\x80\xe1\x7b\xd8\xe8\x53\xf1\xf1\x2a\x64\xf8\x1e\x55\x5b\x62\xad\x70\xdb\x0d\x19\x17\x8f\x92\x76\x80\x3d\x02\xc6\xc5\xc3\x6a\xb7\xe2\xcf\x23\x62\xf8\x1e\xb6\x66\xac\x0d\x2f\x47\xc4\xf0\x3d\xd8\x6a\x97\xf1\x72\x44\x0c\xdf\xa3\xeb\x08\x57\x3b\x3c\x26\x86\x8c\x8b\x87\xb5\xc4\x8e\xb7\xdd\x88\xe1\x7b\xe8\xf1\x86\x06\xd4\x23\x62\xf8\x1e\x36\x52\x57\x7c\x6c\x0f\x19\x17\x0f\x2b\xf9\xc4\xeb\x2a\x62\xf8\x1e\x62\x59\x09\x5e\x8e\x88\x71\xf1\xb0\x1e\x85\xaf\x77\x43\x86\xeb\xd1\xf4\xc4\x50\x03\xe8\x11\x32\x7c\x8f\xa2\xad\xbd\x15\xb8\x7f\x84\x0c\xdf\xc3\x46\xb8\x86\x8f\x89\x21\xe3\xe2\x51\xd3\x0e\xf8\xe6\xab\xd9\xb0\xe8\x1f\xa2\xdf\x49\x23\xed\x00\x97\x26\x60\x5c\x4a\x23\x69\x87\x5f\x24\xd6\xec\xe1\x37\xbc\xb9\x44\x0c\x3f\xb1\x66\xd5\xec\xbe\x46\x71\x4f\xcc\x5a\x4c\xc3\xdb\x58\xc4\xb8\x24\x66\x75\x8c\x1f\xea\x85\x0c\xdf\xc3\x56\xfd\x0d\xdf\x27\x84\x8c\x8b\x47\x4f\x3b\xc0\x1e\x01\xe3\xe2\x31\xd3\x0e\xbf\x78\x88\xdd\xda\x4a\xc7\x5b\x57\xc4\xf0\x13\xeb\xd6\xba\xdc\x97\xb0\xee\x89\x59\x5b\xc1\x67\xf8\x90\x71\x49\xcc\xea\xf8\x57\xe7\x41\xad\x5b\x27\xee\xf0\xf2\x3f\x64\xf8\x89\x0d\x6b\x60\x03\x6f\x92\x11\xe3\xe2\x61\xd5\xe5\x7e\x31\xcd\xb5\xf0\xc3\x7a\xd7\xc0\xfb\x63\xc4\xf0\x13\x9b\x36\x4b\xcc\xdf\x9c\x6b\xb6\x69\x6d\xec\xed\xb3\x77\x71\x62\x11\xe3\x92\x98\xcd\x2b\xee\xe7\x59\xee\x89\x59\xf9\xdd\xef\xf3\xbb\x91\xd8\x0e\xb8\x18\xbf\x61\x0a\x19\x6e\x69\x98\xb4\x1f\xb3\xfb\x4d\x9f\xd7\xc4\xc8\x6c\xdc\xef\x46\xb9\x93\x46\xda\xe1\x17\xa4\xa2\x95\xed\xdf\x54\x5f\x49\x36\xc1\x72\x85\xfb\x66\xc8\xf0\xeb\xcd\x36\x44\x8c\x6f\xa1\x42\x86\xef\x61\x37\x3a\xcc\x70\x6b\x0e\x19\x17\x0f\x7b\x2a\x0c\x77\xe5\x90\xe1\x7b\x0c\xed\x63\xec\xbe\x5e\x7a\x7d\x88\x62\x24\xc1\x0b\x1f\x31\xfc\xc4\x6c\xd5\xce\xf8\x3a\x3f\x64\x5c\x3c\xac\xba\x04\xaf\xe0\x88\x71\xf1\xb0\xe6\x8b\x9f\x45\x85\x0c\xd7\xa3\xdb\xc5\x72\x77\xdf\xd9\xbf\x3d\xc4\x6e\xc7\xe7\x1d\x7e\xc9\x21\x66\xf8\x89\x91\xce\x79\x9d\xe0\x59\x32\x64\xf8\x1e\x76\xb4\xdd\x0b\xbc\x3e\x0a\x19\x17\x8f\x99\x76\x80\x3d\x02\x86\xef\x61\xc7\xe7\x1d\x3f\x70\x0f\x19\xbe\x47\xb3\xda\xc5\x47\xc6\x90\xe1\x7b\xb0\x76\xf5\x8e\x8f\x8c\x21\xe3\xe2\x61\x0d\xde\x7d\x69\xfa\xda\xe0\xbb\x0e\xc0\xbd\xe3\x85\x8f\x18\x7e\x62\xb6\xc2\xed\xf8\x9a\x38\x64\xf8\x1e\x76\xdd\xd9\x07\xde\xe0\x23\xc6\xc5\xc3\xb2\xc2\x6f\x34\x42\x86\xef\x31\x8d\x31\x71\x8f\x88\xe1\x7b\xd8\x9c\xd0\xf1\xdb\x86\x90\xe1\x7a\x0c\x3b\xfb\x19\x19\x6e\x57\x21\xe3\xe2\xc1\x69\x07\xbc\xc1\x0f\x7b\x93\x62\xe0\xef\x5e\x84\x0c\x3f\x31\xea\x69\x07\xd8\x23\x60\xf8\x1e\xf6\x7e\xc3\xc0\xdf\x88\x08\x19\xbe\x87\xad\x83\xc6\x80\x47\xad\x90\x71\xf1\x18\x69\x07\xd8\x23\x60\xf8\x1e\x76\x30\x3c\xf0\xa3\xe4\x90\x71\xf1\xb0\xc6\x38\xe1\xa5\x59\xc8\x70\x3d\x66\xd6\xda\x5d\x01\xf4\x08\x19\x17\x0f\x4e\x3b\xc0\x1e\x01\xe3\xe2\xd1\xd3\x0e\xb0\x47\xc0\xf0\x3d\xac\x47\x4d\xbc\x0f\x86\x0c\xdf\xc3\x36\x7b\xb3\xe0\x75\x15\x31\x7c\x0f\xbb\xf4\x9b\xf8\x35\x61\xc8\xb8\x78\xb4\xb4\x03\xec\x11\x30\x7c\x0f\x3b\xab\x7d\xff\x7a\xc0\xd8\x23\x62\x5c\x3c\xac\x76\x1b\xfe\x3c\x22\x86\xef\xc1\xc6\xc0\x8f\x44\x43\x86\xef\x61\xcb\x9f\x89\x2f\x98\x42\xc6\xc5\x63\x1a\x03\xde\xe9\x84\x0c\xdf\xc3\xce\xaa\x26\x7e\xba\x15\x32\x2e\x1e\xd6\x6b\xf1\x2b\xbc\x90\xe1\x7b\x88\x95\x5c\x70\x8f\x88\xe1\x7a\x88\x1d\x51\x09\xbe\x28\x0b\x19\x9e\x47\x61\x5d\x62\x69\xc0\x3c\x62\x86\xef\xa1\x17\x00\x85\xe1\xfe\x11\x33\x7c\x8f\x21\xca\x80\xe7\xf3\x98\xe1\x7a\x4c\x5d\x01\x68\xf8\x5e\x5c\x7e\xfe\x02\x61\xcd\xd5\xee\x17\x6a\x83\x5f\x04\x8d\x19\x5e\x4e\x2d\xf7\x35\x92\xb4\xdc\x91\xb5\xd2\x62\xb0\xe8\x21\x08\x0b\x7c\xb2\x18\x33\xbc\xac\x58\x74\x37\xcc\x02\xdf\xc5\xc4\x0c\xdf\x63\x18\x03\xbe\x59\x88\x19\x17\x8f\x96\x76\x80\x3d\x02\xc6\xc5\x63\xa6\x1d\x60\x8f\x80\xe1\x7b\xe8\x45\x84\x06\xd4\x23\x62\x5c\x3c\xac\xe4\xf0\xa8\x1b\x33\x7c\x0f\x7d\x0d\x82\x05\x3e\x50\x8d\x19\x9e\x47\xcf\x7a\xc6\xdf\x33\xfc\xa2\x57\xcc\xf0\x3d\xf4\x64\xad\x67\xf8\x2c\x2e\x66\x5c\x3c\x66\xda\x01\xf6\x08\x18\xbe\x47\x35\x06\xfc\x39\xab\x98\xe1\x7a\x90\xbe\x52\xd5\x09\xbe\xd4\x8e\x19\x17\x8f\x66\x0c\x74\x95\x18\x33\x5c\x8f\x62\xb5\x5b\xf0\xe7\x11\x32\x2e\x1e\x92\x76\x80\x3d\x02\x86\xef\xa1\xaf\xbe\x68\x40\x3d\x22\x86\xef\xd1\x8c\x01\x9f\xbf\xc6\x8c\x8b\x07\xa7\x1d\x60\x8f\x80\x71\xf1\xb0\xda\x85\x3f\x65\x15\x33\x7c\x0f\xb6\xac\xf0\xb6\x1b\x32\x2e\x1e\x96\x55\x47\xc7\xdd\x98\xe1\x7b\x0c\x7b\x82\x03\xaf\xab\x88\xe1\x7b\x4c\x4a\x3b\xa0\x1e\x11\xe3\xe2\x61\x25\x9f\x78\x39\x22\x86\xef\x21\x3d\xed\x80\x7a\x44\x8c\x8b\x87\x65\x05\x5f\xc8\xc5\x0c\xd7\xa3\xea\x5e\xa2\x57\x78\xf7\x11\x33\x7c\x0f\x7d\x8d\x4a\x03\xea\x11\x31\x7c\x0f\x9b\x0d\xf0\x17\x85\x63\xc6\xc5\xc3\xb2\x82\x6f\xb9\x63\xc6\xc5\x43\xd2\x0e\xb0\x47\xc0\xf0\x3d\xba\x95\x1c\x3e\xc9\x88\x19\xbe\x87\x18\x43\x70\x8f\x88\x71\xf1\xb0\xda\x85\x2f\xc5\x63\x86\xeb\xd1\xf4\xcd\x16\x0d\xa0\x47\xc8\xb8\x78\x94\xb4\x03\xec\x11\x30\x2e\x1e\xcd\x18\x70\xff\x08\x19\x17\x8f\x69\x0c\xb8\xed\x86\x0c\xdf\xa3\x58\xc9\x0b\x5e\x57\x11\xc3\xf7\xa8\x56\x72\xf8\xe4\x35\x66\x5c\x3c\x86\x31\xe0\xb5\x68\xc8\xf0\x3d\x6c\x45\x86\xbf\x54\x1b\x33\x7c\x0f\x1b\xe1\x1a\x3e\x26\x86\x0c\xdf\xc3\x56\x19\x6d\xe0\xcf\x3c\x62\x5c\x3c\x24\xed\x00\x7b\x04\x0c\xdf\xc3\x46\x9f\x86\x8f\x57\x21\xc3\xf5\xe0\xa2\xab\x25\xc6\xfb\x47\xc8\xf0\x3d\xac\x25\xe2\x6f\xc6\xc5\x0c\xdf\xa3\x59\x56\x0d\x5e\x27\x86\x8c\x8b\x87\xa4\x1d\x60\x8f\x80\xe1\x7b\xe8\x8b\x19\x9d\x3b\x3c\xee\x86\x0c\xdf\xc3\x5a\x3b\xe3\xfd\x23\x64\xf8\x1e\x7a\x2f\xda\x19\xfe\xc0\x65\xcc\xf0\x3d\xc4\xb2\x82\xcf\x96\x62\x86\xeb\xd1\xf5\xbb\x27\x7a\x87\x6f\xb6\x63\x86\xef\x61\xbb\x89\x8e\xef\x3f\x42\x86\xef\x21\x96\x15\xbe\xff\x08\x19\xae\xc7\xd0\x0f\xcc\xf4\x01\x7f\x88\x30\x66\xf8\x1e\xb6\xd2\x1f\x78\x1f\x0c\x19\xbe\x07\x5b\x56\xf0\x7b\x5e\x31\xe3\xe2\xd1\xd3\x0e\xb0\x47\xc0\xb8\x78\x48\xda\x01\xf6\x08\x18\xbe\x87\x8d\x3e\x03\x7e\xf5\x2c\x66\x5c\x3c\x5a\xda\x01\xf6\x08\x18\x17\x0f\xab\x5d\xe8\x56\xa6\x3c\x32\x2e\x1e\x56\xbb\xf0\x6d\x6d\xcc\xf0\x3d\x6c\xa4\x1e\xf8\xd8\x1e\x32\x2e\x1e\x6c\x0c\xbc\xed\x46\x8c\x8b\x87\x95\x1c\x5f\x5f\x85\x0c\xdf\x63\x5a\xc9\xe1\x3b\x96\x98\x71\xf1\x98\xc6\xc0\xcb\x11\x31\x7c\x0f\x3b\xc5\x19\xf8\xb9\x4f\xc8\xf0\x3c\x06\xe9\xcd\x8f\x06\xcc\x23\x66\x5c\x3c\xa6\x32\xe0\x0f\xd9\xc6\x0c\xdf\x43\x36\x03\xdd\x47\xc5\x0c\xd7\xa3\xe8\xad\xa5\x06\xd0\x23\x64\x5c\x3c\x5a\xda\x01\xf6\x08\x18\xae\x47\xd7\x13\x2f\x0d\xa0\x47\xc8\x70\x3d\xac\xd7\x0e\xbc\x9f\xc7\x0c\xdf\xc3\x5a\x09\xde\x07\x63\x86\xeb\x31\xf5\x13\x04\x1a\x40\x8f\x90\xe1\x79\x4c\xfb\x06\xbc\x89\x7f\x67\x5e\xcc\xb8\x78\xf4\xb4\x03\xec\x11\x30\x7c\x8f\x62\x0c\xf8\x5b\xcd\x62\x86\xef\xa1\x2b\xcb\x99\xe1\x33\x99\x98\x71\xf1\x18\x69\x07\xd8\x23\x60\xf8\x1e\xcd\x9e\x20\x7c\xaf\x16\x33\x2e\x1e\x9c\x76\x80\x3d\x02\xc6\xc5\x63\x1a\x03\xed\x1f\x31\xc3\xf7\xd0\xd1\x67\x66\xf8\x4c\x3f\x66\xf8\x1e\xdd\x5a\x09\x7c\xaf\x16\x33\x7c\x8f\x61\xad\x1d\xde\x73\xc6\x0c\xdf\x63\x5a\x2b\x81\xdf\xea\x8b\x19\x17\x0f\x6b\x25\xf0\xdb\xd4\x31\xe3\xe2\x61\x3d\x0a\x1e\xdb\x63\x86\xef\x21\x96\x15\x7c\x0e\x17\x33\x5c\x8f\xa6\x9f\xa1\xd5\x00\x7a\x84\x0c\xdf\xc3\x66\x03\xfc\xde\x20\x66\xf8\x1e\xfa\xd6\xcb\x6c\xf8\xd8\x1e\x32\x7c\x0f\x1b\xa9\xf1\xf3\xf6\x98\xe1\x7b\xe8\xe9\xe0\x6c\x0d\x9e\x07\x43\xc6\xc5\x63\xa6\x1d\x60\x8f\x80\xe1\x7b\xb0\x65\x05\xaf\xe1\x62\xc6\xc5\x63\x18\x03\xee\x83\x21\xc3\xf7\xe8\xd6\x12\xf1\x71\x37\x64\x5c\x3c\x2c\x2b\x78\x7f\x1e\x33\x7c\x8f\x61\x59\xc1\x9f\xc6\x89\x19\x17\x0f\x36\x06\xde\x07\x23\xc6\xc5\x43\xd2\x0e\xb0\x47\xc0\xf0\x3d\x6c\xc6\x69\xf0\xfe\x3c\x66\x5c\x3c\x6c\x14\x9d\x78\x5d\x45\x0c\xd7\x63\x18\x63\xc0\x7b\xe7\x98\xe1\x7a\x4c\x9b\x0d\x26\x3e\x7f\x84\x0c\xdf\x43\x6f\x46\x35\xa0\x1e\x11\xc3\xf5\x10\xfd\x2e\x01\x0d\xa0\x47\xc8\xf0\x3c\xc4\xbe\x75\x5d\x7a\x45\xfb\x60\xcc\x70\x3d\xec\x13\xcc\x02\x7d\x82\xd9\xf0\xab\x07\xca\x80\xe7\x9b\x98\xe1\xe6\x64\x9f\x1d\x94\x09\x9f\x4b\xc4\x0c\xd7\x63\xea\x67\x07\x45\xe0\xf7\x06\x62\x86\xeb\x21\xfa\xf6\xae\x08\x7c\xde\x15\x33\xbe\x3d\x9a\xfd\xb4\x42\xcb\x19\xfb\xfd\xec\x77\xfc\xf3\x8f\x3d\x28\xb0\x2f\xbc\xff\x65\x26\x8e\xfe\x81\x87\xf4\x99\x38\xed\xf0\xad\x6f\xb5\x54\x53\x4b\xf2\xc2\x2f\x7d\xc2\x7e\xcf\xf8\x1d\x8f\xe4\x43\x6b\x8b\xb5\x03\x92\x0f\x65\x56\xfd\xf7\x53\xb7\x7b\x3e\x27\x1e\xca\x87\x44\xd2\x0e\x50\x3e\x45\x9f\x2f\x61\x3f\xa6\xf8\x8e\x87\xf2\x29\xbd\xa7\x1d\xb0\x7c\xfa\x50\x3c\xf4\xe3\x6d\xef\x78\x2c\x9f\x61\xf9\xb8\xbf\xff\xed\xe5\x33\xb5\xbc\xdf\x3f\x46\x75\xcb\xe7\x85\x87\xf2\x11\xab\x4f\xc9\xfe\x27\x79\xbe\xf2\x91\x4c\x8a\x6f\x60\x3e\x07\x1e\xcb\x87\x2d\x1f\x46\xf3\xe1\xa5\xdf\xaa\xfb\x8d\x57\x5e\x3e\x2f\x3c\x92\x4f\xab\x83\x15\x3f\x3d\xfd\xef\x7c\x5a\x9d\x8a\x1f\xe0\xf3\x3a\xf1\x48\x3e\x7d\xed\xea\x76\xf8\x57\x7f\x8f\xf9\x3d\x9f\x77\x7d\xb6\xf1\x93\x33\xe5\xe2\xb7\xb7\x4f\x7c\xa1\x4c\x89\x7b\xce\xee\x37\xe3\x7f\xe2\xbb\xe9\xf7\x9c\xc9\xfd\xf1\x33\x07\x4f\x64\xf8\xe2\xd7\xe7\x37\xbe\x0c\xc5\xbb\x3f\x88\xe0\xe1\x5b\x33\x3c\x52\xde\x05\x9c\x9a\x4f\xcd\x60\x3e\x35\x6b\x3e\x95\xc0\x7c\x2a\x69\x3e\xfe\x8f\xb9\x7a\x78\x56\xfd\xe9\xfe\xc8\x8f\x83\x5f\xfd\xab\x67\xf2\x7f\x1c\xeb\x1b\x4f\x2d\xd7\xd4\xf3\xcc\xee\x8f\x8f\x7e\xe3\x67\x55\xfd\xf9\xf7\x2b\xa8\x3f\x5b\x33\xe9\x0f\x4a\xef\xbf\x8a\xa7\x9e\x76\x78\x6e\x9f\x0b\x38\x72\xda\x01\xcb\x67\x58\x3e\xee\x8f\xe1\x15\xfd\x2e\x8b\x95\x0b\xa7\xae\x78\x61\x59\xe5\x95\xbf\x6f\xa6\x3c\xe5\x2f\x2c\x92\x76\xf8\xc6\x67\xed\xef\xfc\xd3\x1f\xfb\x5a\xa0\x50\xda\xe1\x39\xff\x51\xb4\xbf\x8c\x92\x8b\xd3\x3e\xc5\xc3\x17\xa6\x34\x4a\xf3\xbe\xc5\xd3\xc3\xb3\xe9\x73\x71\xda\x8f\x8b\x2f\x5c\x17\xbe\xd3\xf7\xfc\xe8\xe1\xbb\xe9\x77\x10\x3f\x32\x2d\xfc\xa0\xf7\xf1\x70\xd5\x1e\xa7\x9a\x7a\x6a\xca\xa1\x45\xae\x6b\xc1\xf8\x67\x0c\x9a\x63\xa4\x1d\x50\x8a\xa8\x4b\xf9\x85\x4b\x99\x54\x93\x06\xd8\xa5\xac\xbe\x36\x46\x15\xdc\xa5\x59\xf1\x1b\xff\x82\x32\x8c\xf2\xfe\x0d\x05\x31\x65\xd6\x92\xc6\xf8\xb8\x56\x0b\x29\x63\xc8\x2a\xcb\xf8\x45\x59\x66\xd6\xe2\xcf\xf7\xa6\x1e\x53\xaa\x96\x65\xfe\xe2\xb9\xcc\xc9\x9c\x76\x40\x29\xf6\xf4\x85\xde\xd6\x00\x21\x45\x1a\xb5\xa4\xa5\x07\x13\x9b\x36\xd3\xcd\x5c\x72\x7e\x4f\xcc\xff\x59\x7d\x05\x0e\xc5\xbf\xcd\x5c\x11\xbe\x35\xc5\x77\x54\xbf\xab\x3e\x65\xc1\xf0\x44\xa4\x78\x6a\x28\xbe\x2b\xfe\xed\x67\xc2\x23\x3c\x9b\x3e\xa3\xfa\x6c\xfa\x03\xcd\x7f\x2e\xfd\xfa\xd9\x95\x6e\xf8\xda\xb3\xa4\x49\x5f\x8f\xf8\x82\xb7\x91\x79\x96\x2e\x97\x9f\x95\xcd\x6f\x2b\xbd\x77\xfc\xe3\xca\x2d\xd5\x3f\x0a\x95\x34\xcb\x47\x9f\xbb\x3b\x9c\x78\xcc\x41\x2c\x27\x41\x1d\x4e\x3c\xe4\x50\xad\xd4\x95\xc7\xdb\x7e\xc1\xc6\xfb\xf1\xf5\xe3\xb8\xb3\x76\x56\x7c\xe7\xb7\x41\x23\xc4\xb3\xe1\x3b\x88\x27\xd3\x27\x34\x1f\x32\xfd\xb7\xfd\x7b\x84\xaf\xa6\x5f\x51\xfd\x6a\xfa\x15\xd5\x6f\xa6\xdf\x50\xfd\x6e\xf8\x0e\xe3\x2d\x9f\x8e\xe6\x33\x4c\x7f\xa0\xfa\xd3\xf4\x27\xaa\x2f\xa6\x2f\xa8\xbe\x98\xbe\x80\xfa\xbd\xa8\xfe\xfb\x37\x35\x86\x78\x4e\x3b\x60\x78\x6b\x0f\x1d\x6c\x0f\x83\x79\xe1\x35\x60\xf8\x6e\x78\xf0\xf9\x0e\x7b\x5e\x03\x7d\x5e\x83\xa7\xe1\x27\x86\x17\xd2\xfe\x28\xf4\xd3\x1f\x4f\xc4\xfb\xf9\x5e\xfe\xc0\x3f\xef\x67\x17\x50\xeb\x53\x08\xfa\xb1\xf9\x77\x3c\xa6\xdf\x46\xda\xe1\x79\x3d\x3f\x87\x8d\x6f\x6b\x25\x09\x2d\x72\x28\x5f\x39\x6c\xa8\xf9\xb9\xe9\x4f\xb4\x7c\xca\x90\x9a\xe6\xe8\x13\x5b\x80\xa8\x8f\xcf\x89\x7d\x6c\x0e\xb1\x80\xfa\xf8\x9c\x27\x9f\x42\xca\xf9\x58\x2d\xc4\x3e\x1e\xe7\xd1\x87\x95\x53\xcb\x6f\x7c\x1c\xce\x93\x4f\xdd\x1c\x68\x19\xba\x7d\x3c\xce\x93\x0f\x5b\x1d\x30\xde\xde\x7c\xce\xa3\x4f\x35\xce\x6f\xca\xe3\x71\x1e\x7d\x46\xda\xe1\x17\x3e\x0e\xe7\xc9\xa7\x5b\x1d\xf4\xdf\x94\xc7\xe3\xc4\x3e\xc3\xfa\xf6\x78\x3f\x0f\x8b\x7d\x7c\xce\x93\xcf\xa0\xb4\x03\xee\xe3\x71\x1e\x7d\x66\xda\xe1\x17\x3e\x0e\xe7\xc9\x67\x1a\x67\xfe\xc6\xc7\xe3\x3c\xf9\x88\xd5\xb5\xfc\xe6\xf9\x78\x9c\x47\x1f\xcb\x4d\x7e\x53\x1e\x8f\xf3\xe8\xb3\xc6\x5e\x74\x63\xfa\xd7\xc7\xe1\x84\x3e\x62\x73\x96\xe4\xee\xec\x14\xbe\xcf\x71\xa6\x90\xe1\x09\x3b\xf7\x99\x52\x0c\x5f\x30\xbc\xd8\xee\x5a\x72\xa9\x02\xff\x6c\x87\x47\x7a\x43\x7c\x6e\x7e\xfe\x28\x74\x26\xf9\xbc\xb0\x7c\xb0\xf9\x26\x3d\xd8\xd4\xa9\x89\xd5\xf9\xfa\x8c\xd7\x6d\x55\xa3\x4f\xea\x83\xf1\xb4\xae\xd9\x8c\xbe\xb2\xaa\xf3\x95\xd5\xa3\xc7\xc1\x80\x3c\xd8\x2a\x78\x05\xe4\xa7\xff\xdf\xf1\xcf\xab\x33\x05\xb6\x85\x6f\xee\x6d\x98\xab\xff\x83\x87\xf4\x59\xf3\xb9\xee\x87\xbf\xf4\x0f\x3c\xa0\x3f\xad\xbc\x33\x67\xf7\xf6\xe0\x73\xf5\xfa\x8e\x07\xf5\x7b\x12\xdd\x75\xa3\xfa\x3f\x78\x44\xbf\xe6\x5c\x93\xe8\x0a\x16\xd2\x3f\xf1\x88\xfe\xd0\xe7\x35\x2f\xf5\xff\xad\x7f\xe2\x9f\xf5\xff\x1b\x00\x00\xff\xff\xd8\xbd\x37\xa8\x73\xcf\x00\x00")

func dataWeightsTxtBytes() ([]byte, error) {
	return bindataRead(
		_dataWeightsTxt,
		"data/weights.txt",
	)
}

func dataWeightsTxt() (*asset, error) {
	bytes, err := dataWeightsTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/weights.txt", size: 53107, mode: os.FileMode(420), modTime: time.Unix(1549922049, 0)}
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
	"data/data.go": dataDataGo,
	"data/substitutions.txt": dataSubstitutionsTxt,
	"data/weights.txt": dataWeightsTxt,
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
	"data": &bintree{nil, map[string]*bintree{
		"data.go": &bintree{dataDataGo, map[string]*bintree{}},
		"substitutions.txt": &bintree{dataSubstitutionsTxt, map[string]*bintree{}},
		"weights.txt": &bintree{dataWeightsTxt, map[string]*bintree{}},
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

