// Code generated by go-bindata.
// sources:
// client/bootstrap/daemon-down.sh
// client/bootstrap/daemon-up.sh
// client/bootstrap/docker.sh
// client/bootstrap/keygen.sh
// client/bootstrap/token.sh
// DO NOT EDIT!

package client

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

var _clientBootstrapDaemonDownSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xce\xc1\x4a\xc3\x40\x10\xc6\xf1\xfb\x3c\xc5\x47\xdb\xeb\xe6\x0d\x72\x88\x36\x88\x60\x57\x28\x7a\xf0\xd4\xae\xd9\xd9\x74\x48\x33\x5b\x67\xb7\xa8\x6f\x2f\x22\x42\xe8\xfd\xe3\xff\xfb\xd6\xb8\x0b\x45\x06\x94\xc1\xe4\x52\x91\xb2\xe1\xdd\x44\x47\xd1\x11\x31\x7f\x2a\xea\x89\x11\x03\xcf\x59\x1b\xa2\xc2\x15\x8e\x89\xb6\x5d\xbf\x7b\xf6\x07\xdf\xed\xfa\x56\x94\xad\x4a\x70\x7f\x23\xa2\x35\xee\x4f\x3c\x4c\x90\x84\x70\x36\x0e\xf1\x1b\x76\x55\x15\x1d\x1b\xea\x9e\xf6\x7d\xb7\x7d\x3b\xec\x5f\xbd\x7f\xf4\x0f\xed\xb1\x5c\x63\x46\xcc\xc3\xc4\x86\x4b\x81\xfb\x80\x73\x49\xce\x95\x0d\x2b\x0d\x33\xb7\x9b\x05\xb5\x3a\xfe\xd6\x5f\xc2\xc4\xe0\x2f\x29\xf5\xff\x63\x43\xcb\x8c\xcd\x70\x09\x9b\x1b\x8a\x7e\x02\x00\x00\xff\xff\x8a\x49\xa7\x95\xe9\x00\x00\x00")

func clientBootstrapDaemonDownShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapDaemonDownSh,
		"client/bootstrap/daemon-down.sh",
	)
}

func clientBootstrapDaemonDownSh() (*asset, error) {
	bytes, err := clientBootstrapDaemonDownShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/daemon-down.sh", size: 233, mode: os.FileMode(420), modTime: time.Unix(1514021421, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientBootstrapDaemonUpSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x54\x51\x6f\xdb\x38\x0c\x7e\xd7\xaf\xe0\x9c\x14\x7d\x99\xec\x76\xc3\x01\xbb\x0e\x7e\xc8\xad\xc6\x1a\xac\x4d\x8a\xa4\x87\xc3\xa1\x57\x64\x8a\xcd\xd4\x5a\x6c\x49\x27\xd2\xc9\xd6\x5f\x3f\xc8\x4e\xe2\xa4\xe8\x53\x62\x89\xfc\xf8\xf1\xe3\x27\x0e\xe0\x2f\x45\x3a\x07\xca\xbd\x76\x0c\x2b\xeb\x61\xe9\xb5\x79\xd6\xe6\x19\x1a\x07\x5c\x22\x14\x0a\x6b\x6b\x62\x21\x08\x19\x24\x0a\x71\x3d\xca\xee\xa6\x93\xc5\x2c\xbb\xcd\x46\xf3\x2c\x3d\x7b\xbc\x7c\xa2\xfd\xe1\xfd\x74\xf6\x90\x9e\x3d\x7e\x78\x22\x71\x33\x9d\x3f\x2c\x46\xd7\xd7\xb3\x6c\x3e\x4f\xcf\x1e\x3f\x3e\xd1\x21\x75\x32\xba\xcb\x52\x6d\xd0\xb3\x56\xb2\xc3\x17\xe3\xbb\xd1\xd7\x2c\x6d\x96\x79\xa5\x1a\x93\x97\x4e\x15\xc9\x2e\xe2\x6a\x78\x5a\x51\x7c\x99\x4e\x1e\x46\xe3\x49\x36\xeb\xca\x7d\xba\xf8\x74\x29\xc4\x00\xe6\xc8\x81\x73\xa1\x3d\xe6\x6c\xbd\x46\x12\xf5\xba\xd0\x1e\xa4\x83\xe1\xcd\xf4\x2e\x4b\x9c\xb7\x3f\x30\xe7\xd7\xc7\x44\x55\xc8\xff\x52\x62\xbe\x06\xbd\x02\x55\x79\x54\xc5\x2f\xf0\x8d\x31\x41\x08\x65\x0a\x60\xb5\x46\x28\xec\xd6\x00\xfe\xd4\xc4\xe1\x78\x2f\xcc\xe8\x76\x96\x8d\xae\xff\x5d\xcc\xfe\x9e\x4c\xc6\x93\xaf\xe9\x77\x6a\x0a\x0b\x85\xcd\xd7\xe8\xc1\x11\xc8\xff\x41\xca\x95\xae\x18\x3d\x44\x46\xd5\x98\x0e\x8f\x74\x88\xbe\x0b\xbd\x82\x47\x78\x07\xf2\x05\xa2\xe1\x2b\xb0\x08\x9e\x3e\x87\x29\x18\x01\x00\x80\x79\x69\x21\xfa\xa6\xab\x2a\xd4\x3f\x10\xc9\xad\x61\x15\xc4\x8a\xe3\x38\x6a\x03\x8f\x19\xf8\x1a\xe4\x0a\x5e\x03\x8b\x95\xfe\x1c\x9a\xbe\xf7\xe8\x94\x47\x50\xce\x79\xeb\xbc\x56\xbc\x9f\x38\xe8\x5a\x3d\x63\xdc\xb1\x8b\x5e\xcd\x20\x82\x77\x29\x44\x8c\xc4\xa7\x0c\x07\x70\xdf\x54\x55\xeb\x9b\xdd\xf4\x0e\x32\xf5\x0d\x84\x90\xc0\x7b\x7c\x1a\xf1\x16\x79\x17\xd0\x86\xad\x37\x04\x56\x84\x47\x28\xb7\xad\x4f\x4e\x84\x38\x05\xdc\x35\xf0\x16\x6c\x65\x55\x01\x52\x43\xd2\x45\xca\x36\x52\xac\x74\x50\x64\xd6\x98\x5e\x51\xd8\x6a\x2e\x41\xe5\x39\x12\x01\xdb\xb6\xb1\xd2\x12\xef\x81\x28\xfc\x70\xeb\x10\x8f\x15\x6e\x94\xe1\x63\x03\x82\x14\x03\xe0\x52\x13\x68\x02\x83\x01\x45\xf9\x5f\xb0\xc4\x5c\x35\x84\xb0\x45\xd8\x86\x8c\xfe\x99\x85\x1a\x4b\x04\xb5\xac\x10\x88\x95\x67\x31\x68\xc1\x89\xad\xeb\x59\x11\x84\xc8\x1d\x95\x18\xc6\x7c\x4e\xa0\x2a\xb2\x6d\x84\xb7\x1b\xf4\xa4\x55\xf5\x5e\x0c\xa0\x64\x76\x74\x95\x24\xdb\xed\x36\xae\x36\x65\xac\x6d\xe2\x2c\x31\x25\x85\x35\x2c\xf1\xa7\xb3\x84\x92\x4b\x94\x5d\x3f\xb2\xeb\x47\x1a\xcb\x12\x37\x68\x24\x5b\xa9\x64\xef\xaf\x92\xeb\x4a\x0c\x8e\x0a\x7a\xcc\x6d\x5d\xa3\x29\xb0\x38\x2e\xf7\xc3\x21\xab\x97\x17\x1b\x3f\x6b\x2e\x9b\x65\x28\xfb\xe1\xe2\xf2\x8f\xe4\xe2\xcf\xe4\xe2\x63\x52\xd8\xb6\x40\x43\x87\xb2\xda\xec\xff\xad\xac\x97\xb9\x4e\xc4\x00\x46\x04\x0a\x3c\x52\x53\xf1\xfb\x4e\xc3\x7e\x2a\xa5\x22\xf0\xd6\xf2\x7e\x32\x3b\x39\x3c\xd6\x96\x11\x36\x8e\x62\x71\xf2\x04\x1a\x03\xb2\x00\x29\x7d\x0d\xff\xb5\x66\x90\xae\xb7\x74\x58\x22\xd1\x55\x34\x3c\x5d\x2b\xd1\x3e\x72\x03\xc9\x46\xf9\xc4\x37\x26\xe9\xe0\xe2\x20\xd2\xd5\x5b\x87\x7d\x4a\xd4\x6e\x96\xe8\x2a\x51\xce\x25\xad\x61\x76\x57\x08\xe1\x22\xdd\xdd\xf7\xa7\xf3\xf9\xcd\xe2\xdb\x64\xfa\xcf\x64\x11\xf6\xe6\x3c\x3d\x3f\x64\x26\x31\x51\x99\xac\x8d\xdd\x9a\x45\xf8\xa6\xf3\x7d\x96\x0c\xcb\xa4\xef\xa3\xdd\x26\xbb\xbb\xa8\x7b\x33\x51\x4b\xa4\xdf\xc3\x91\xf8\x1d\x00\x00\xff\xff\x74\x47\xad\x86\xf1\x05\x00\x00")

func clientBootstrapDaemonUpShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapDaemonUpSh,
		"client/bootstrap/daemon-up.sh",
	)
}

func clientBootstrapDaemonUpSh() (*asset, error) {
	bytes, err := clientBootstrapDaemonUpShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/daemon-up.sh", size: 1521, mode: os.FileMode(493), modTime: time.Unix(1521347519, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientBootstrapDockerSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\x5f\x4f\xdc\x3e\x10\x7c\xf7\xa7\x18\xb8\x13\x7f\x24\x92\x00\xfa\xe9\xf7\x50\x44\x25\x7a\x9c\xaa\xaa\x48\x48\x1c\x3c\x57\x26\xd9\x5c\x2c\x12\x3b\x78\x37\xd0\xb4\xea\x77\xaf\xe2\x8b\x53\xae\xa5\x2a\x79\x72\xbc\xb3\xeb\x99\xd9\x99\xed\x64\xf7\xc6\x66\x5c\x29\x35\xc3\x07\xe7\x84\xc5\xeb\x96\xa1\xd1\xe8\xbc\x32\x96\x50\x3a\x8f\xc2\xe5\x0f\xe4\xa1\x6d\x31\x1e\x93\xdc\x35\xad\x63\x4a\xd5\x0c\xb7\x95\x61\x18\x46\xeb\x49\xa4\x87\xae\xdb\x4a\xa7\xf8\x64\x59\x74\x5d\x33\xf2\xce\xd7\x70\xb6\xee\x61\x4a\x18\x41\xe1\x88\xed\xbe\x80\xbe\x1a\x96\xd0\x4f\x2c\x54\xc0\x59\xdc\xdd\x77\x56\x3a\x9c\xfc\x9f\x1e\xff\x77\x04\x4f\x8f\x9d\xf1\xc4\xe0\xae\x70\x81\x86\x46\x49\xcf\x60\xa1\x96\x53\xa5\x98\x04\x09\x29\x75\x79\xbd\xf8\xbc\xbc\xf9\xb2\xba\xbe\xbb\x59\x2c\xcf\xd7\x24\xe9\x86\x63\x9a\xbb\x26\x16\x2f\x97\xab\xdb\xf3\xfd\x4c\x9a\x36\x5b\x93\x24\x23\x80\xab\xfd\x41\xf6\xa2\xa2\xfc\x61\x60\x17\x65\x7a\xd2\x45\x0f\xb3\x51\x40\x45\xaa\x4c\x89\x4a\x73\x15\x01\xa7\xef\xb3\x82\x9e\x32\xdb\xd5\xf5\x19\xa4\x22\xab\x00\x0c\x82\x04\xc7\xaa\x34\x67\x4a\x95\x24\x79\x55\x9a\x9a\x0e\x0e\xf1\x3d\x54\x67\xb8\xf0\x6b\x7e\x37\x9e\x81\xf9\x09\xd8\x75\x3e\x27\xdc\xdd\x5c\xfd\xba\x3d\x45\x41\x2c\xc6\x6a\x31\xce\x62\x98\x90\x86\x62\x64\x10\xdc\x7c\xfd\xfd\xe1\x0b\x56\x05\x4c\x52\xf2\xea\x6a\x78\x25\x71\xd8\x9d\x9f\xee\x6e\x28\xd6\x71\xce\xf3\x9a\xe4\x5f\x73\x02\x26\xb9\x0e\xed\x98\x9f\x8c\x13\x98\x26\x94\x27\xe9\xbc\xc5\xa6\x32\xe8\xfe\x31\xd8\xf9\x91\x24\x1a\xf5\xea\xc6\x4d\x89\x9d\x83\xbf\xd8\x79\xf8\x82\xc7\x0c\x17\x8d\xfe\xe6\x2c\x96\x8b\xd5\x66\x19\x36\x27\x8e\xb1\x40\xde\xb1\xb8\x26\x6e\x29\x9a\xb4\xf6\xd4\x22\x79\x8c\xad\x19\x49\x9e\x71\xcf\x42\x4d\xe2\xa9\x26\xcd\xf4\x9a\xd2\xbe\x9b\x06\x21\xe9\x47\x5e\x7f\xea\x9d\xe1\xd6\xf7\x10\x87\xc2\x3d\xdb\xda\xe9\x02\x1d\x1b\xbb\x1e\x23\xee\x83\x61\x47\x2f\xd0\xf7\x54\x3a\x4f\xf0\xc4\xce\xcb\x00\x14\x17\x9f\x89\x6d\xe9\x04\x37\x25\xa6\xd4\x60\xbe\x15\xea\xe9\x77\x88\xf1\x6f\xfc\x83\x86\x6a\x0b\x31\xd5\xb6\xc8\x0f\x9f\x6e\x25\x19\x96\xda\xb5\x85\x16\xc2\xde\xde\x74\x93\x4c\x79\x0f\xb4\xb6\xba\xde\x44\xeb\x4d\x84\x4a\x33\x46\x45\x95\x46\xa9\x60\x3d\x93\x7f\x32\x39\xc5\x2c\xb0\x68\x2f\xea\x67\x00\x00\x00\xff\xff\x99\xac\x13\xfb\x96\x04\x00\x00")

func clientBootstrapDockerShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapDockerSh,
		"client/bootstrap/docker.sh",
	)
}

func clientBootstrapDockerSh() (*asset, error) {
	bytes, err := clientBootstrapDockerShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/docker.sh", size: 1174, mode: os.FileMode(493), modTime: time.Unix(1521099304, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientBootstrapKeygenSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x41\x6b\xdc\x30\x10\x46\xef\xf3\x2b\xbe\xb0\x81\x5c\x2a\xef\xbd\xa1\x0b\x6d\x13\xda\x3d\x74\x13\x68\x7a\x2a\xc5\x68\xad\x71\x34\xd8\x91\x8c\x66\xdc\xad\x2f\xfd\xed\xc5\xeb\xa5\x24\x64\x2f\xd1\x55\x6f\x78\x4f\x9a\xd5\xc5\x7a\x2f\x69\xad\x91\x68\x85\xfb\x92\xc3\xd8\xb0\xc2\x63\x18\xf7\xbd\x34\x6e\x28\xf2\xdb\x1b\xa3\xe3\xc9\x0d\x5e\x0a\x7c\x0a\xc8\xa3\x0d\xa3\x29\x2c\xf2\x89\x9b\xef\x2b\x22\x65\x83\x63\xa2\xed\x4d\x7d\x73\xfb\xfd\x61\xbb\xfb\xf8\xb0\xbd\xdb\x7d\xb8\xfc\x7a\xf7\xed\x76\x5d\xa9\xc6\xb5\x84\xba\xa8\xaf\x25\x71\x31\xf1\x75\xe0\xa1\xcf\x13\xdd\xff\xf8\x54\xbf\x71\xa6\x1a\xc6\xfd\x9c\xfc\x39\x72\xd3\x41\x5a\x04\x56\x93\xe4\x4d\x72\x42\x2b\x3d\xc3\xf7\x85\x7d\x98\xc0\x7f\x44\x4d\xdf\x93\xb4\xf8\x09\xd7\xe2\xf2\xa5\x09\xbf\xae\xe7\x87\x24\x02\x80\x23\x73\x71\xa4\x5e\x37\xbd\x20\xe7\xb3\xc2\xb6\x7d\xf6\x01\x08\x99\x35\xd9\x22\x7c\x87\x27\xdf\x31\xc4\xaa\xff\xb8\x6a\x74\x1d\x4f\x8f\x9c\xe0\xa6\x73\x25\x9b\x73\xd6\xe3\x78\x2b\xd7\xc4\xbd\x32\x2d\xda\x2f\x9c\xb8\x9c\xd6\x82\x83\x58\x44\xca\x18\xbc\xea\x21\x97\xb0\x08\x9f\xcb\x5e\x9b\x9c\xa1\xa8\x87\xdb\xe1\xea\x8a\x5a\x21\x3a\xe1\xda\xf8\x84\x47\xb1\x38\xee\xab\x26\x3f\x61\xb3\xc1\xdf\x65\x09\x5d\xca\x87\x54\xc7\xac\xa6\x44\x8d\xb7\xb3\xa9\xff\x02\x00\x00\xff\xff\x37\x00\x91\x4b\x4d\x02\x00\x00")

func clientBootstrapKeygenShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapKeygenSh,
		"client/bootstrap/keygen.sh",
	)
}

func clientBootstrapKeygenSh() (*asset, error) {
	bytes, err := clientBootstrapKeygenShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/keygen.sh", size: 589, mode: os.FileMode(493), modTime: time.Unix(1515346036, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientBootstrapTokenSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x90\xc1\x4a\xc3\x40\x10\x86\xef\xfb\x14\xbf\xb4\xd2\x53\xb2\xf7\x42\x0e\x45\x82\x29\xd6\x46\x8c\xe0\x45\x08\xdb\x64\x74\x97\xd8\xd9\xb8\xb3\x6b\xf1\xed\x25\x9a\x6a\x8f\x33\xdf\x7c\x3f\x33\xb3\xb8\xd2\x07\xc7\x5a\xac\x52\x42\x11\x19\x29\xf5\x58\xee\xca\x4d\x53\x16\xd7\xa2\xd4\x02\xb7\xc4\x14\x4c\x24\x81\x61\x34\x4d\x85\xe8\x07\x62\xbc\xfa\x80\x24\x84\x93\x8b\x16\x9b\x87\x2d\x02\x7d\x24\x92\x28\xf9\x85\x03\x83\xde\xd0\xd1\xf3\x2c\x25\x71\xfc\x86\x9b\xdd\x36\x57\x92\x7a\x8f\xde\x77\x03\x05\x84\xc4\xc8\xb2\x70\xc4\x8b\x02\x80\xec\x13\xcb\xaa\xbe\x2f\xd7\xda\x8c\xa3\xb6\x5e\xe2\x19\xd0\xb4\x40\x7b\xb7\xaf\x9f\xf7\x6d\x55\x37\x4f\x4d\xb1\xfa\x9b\xd1\xb9\x88\xd5\x03\xfb\x13\xb7\x53\x2d\xab\x7f\x6b\x4a\x2b\x7e\x32\xcf\xbd\x8c\x38\x86\xaf\xd1\x3b\x8e\x85\x63\x0a\xd1\x99\x19\xa5\x43\xf7\x6e\x12\x77\x76\x34\xbd\x9e\xd1\x7a\x39\xff\xe4\xf7\x0e\xf5\x1d\x00\x00\xff\xff\xa4\x4f\x24\xd1\x35\x01\x00\x00")

func clientBootstrapTokenShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapTokenSh,
		"client/bootstrap/token.sh",
	)
}

func clientBootstrapTokenSh() (*asset, error) {
	bytes, err := clientBootstrapTokenShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/token.sh", size: 309, mode: os.FileMode(493), modTime: time.Unix(1521347945, 0)}
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
	"client/bootstrap/daemon-down.sh": clientBootstrapDaemonDownSh,
	"client/bootstrap/daemon-up.sh":   clientBootstrapDaemonUpSh,
	"client/bootstrap/docker.sh":      clientBootstrapDockerSh,
	"client/bootstrap/keygen.sh":      clientBootstrapKeygenSh,
	"client/bootstrap/token.sh":       clientBootstrapTokenSh,
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
	"client": &bintree{nil, map[string]*bintree{
		"bootstrap": &bintree{nil, map[string]*bintree{
			"daemon-down.sh": &bintree{clientBootstrapDaemonDownSh, map[string]*bintree{}},
			"daemon-up.sh":   &bintree{clientBootstrapDaemonUpSh, map[string]*bintree{}},
			"docker.sh":      &bintree{clientBootstrapDockerSh, map[string]*bintree{}},
			"keygen.sh":      &bintree{clientBootstrapKeygenSh, map[string]*bintree{}},
			"token.sh":       &bintree{clientBootstrapTokenSh, map[string]*bintree{}},
		}},
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
