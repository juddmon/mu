// Code generated by go-bindata.
// sources:
// assets/cluster.yml
// assets/vpc.yml
// DO NOT EDIT!

package templates

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

var _assetsClusterYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x59\x7b\x6f\xdb\x38\x12\xff\x3f\x9f\x62\xea\x2b\xe0\xed\xa2\xf2\x43\x4e\xda\x54\x40\x76\xe1\x38\x6e\x6b\xac\xdd\x1a\x51\xda\x02\x4d\x8b\x82\xa2\x28\x9b\xa8\x44\xea\x48\x2a\x8f\xed\xf5\xbb\x1f\x48\x4a\xb2\x5e\xce\x26\xbb\x3d\xec\x1d\x70\x2e\x82\xda\xe4\x70\xde\xf3\x9b\xa1\xe4\x38\xce\xc1\xf4\x83\x7f\x41\x92\x34\x46\x8a\xbc\xe4\x22\x41\xea\x3d\x11\x92\x72\xe6\x41\xdf\x1d\x8d\x47\xce\xe8\x85\x33\x7a\xd1\x3f\x58\x23\x81\x12\xa2\x88\x90\xde\x01\xc0\x82\x49\x85\x18\x26\x17\xb7\x29\xd1\xbf\x01\xcc\x37\xf0\x95\xa0\x6c\x63\x16\xce\x88\xc4\x82\xa6\xca\xb0\x2a\xe8\x41\xdd\xa6\x04\x14\x87\x4c\x92\x41\x4e\x16\xa1\x2c\x56\x1e\x28\x77\x90\x50\x2c\xf8\x81\x39\x4a\x05\x09\x67\x28\x45\x98\xaa\xdb\xaa\x80\x37\x59\x12\x10\x51\x3f\xd9\x1f\xf7\xdb\x12\x2d\x21\xf0\x08\x68\x2e\x5b\x6a\xb9\x31\xca\x18\xde\x02\x65\x70\xcb\x33\x01\xf3\x99\x0f\x38\xce\xa4\x32\x3c\x57\xe8\xc6\xa7\xbf\x93\x3f\x94\xe7\x76\xc8\x5b\xa1\x1b\x9a\x64\x09\xb0\x2e\xb9\x5b\xa4\x00\x23\x06\x01\xc9\x15\x20\xe1\x1e\x15\x7e\x23\xb7\x6f\x50\x72\x2f\x9f\xe6\xa4\xda\x2a\x24\x25\xc7\x14\x29\x02\xd7\x54\x6d\xe1\x9a\x8b\xaf\x44\xec\x14\x18\x00\x2c\x09\xba\x22\x10\xc4\x88\x7d\xd5\x07\x42\x2a\x51\x10\x13\xf0\xfd\xd7\x80\x30\x26\x52\x36\xa2\xd1\xd7\x26\xfa\x72\x3b\x8d\x63\x7e\xed\xb5\x85\xfb\x59\xc0\x88\x82\x48\xf0\x04\xae\xb7\x14\x6f\x8d\x1a\x9a\xb8\xc5\xb3\x65\xc5\x8a\xb2\x25\x61\x1b\xb5\xf5\xa0\xff\xc2\xba\x72\x85\x6e\xca\xa5\xf1\x71\xbf\xae\xcb\x68\x60\xfe\x0d\x47\x66\xd9\x68\x44\xc2\x35\x52\x8a\x08\xe6\x41\xef\xa7\x4f\x9f\xc2\x6f\xe3\xa7\x93\xef\x4f\x3e\x7d\x1a\xdc\xe7\xc7\x30\xff\xea\x7e\x7f\xd2\x33\x2c\x67\x9c\x49\x25\x10\x65\xaa\x66\x63\x3f\xc9\xa4\xd2\x31\x43\x70\x85\x62\x1a\xc2\x6c\x71\x76\x0e\x41\xcc\xf1\x57\x0f\x6e\x06\xe6\xdf\xf0\x66\xa0\xb5\x7d\x9f\xe2\x45\x78\x9f\xa0\x99\x88\xf1\x08\xd4\x96\x68\xa6\x99\x09\x1f\x4d\x52\x2e\x14\x44\x5c\x98\x75\xc3\xec\x00\x60\x9d\x05\x31\xc5\xd6\xd3\xd3\x8f\xe3\x1f\x27\x60\xfa\x71\x0c\xd2\x06\x90\xb6\x05\xb9\x3f\x52\x90\x5b\x13\xd4\x4c\xb0\xba\xe0\xc9\x8f\x14\x3c\xb9\x43\xf0\x8a\x28\x14\x22\x85\xb4\xb4\xe9\x07\xdf\xf3\x66\x31\xcf\x42\x8b\x7e\x5a\x84\xb7\x60\x8a\x88\x08\xe1\xbc\x0e\x4b\xec\x7b\x25\x78\x96\x4a\xbb\x08\xe0\xc0\x12\x05\x24\x2e\x7e\xea\x4f\x58\x48\xe9\x95\x88\x37\xe3\x2c\xa2\x9b\x4c\x18\xd6\xbd\x92\xb6\x8e\xa7\xc5\xc7\xa9\x21\x6b\x6d\x23\x2f\xf7\xda\x5a\x51\xa0\xf7\x51\x68\x9a\x29\x0e\x3e\x46\x31\x65\x9b\x87\x2a\xd5\x00\xe4\xda\x5e\x0e\x9a\x75\x47\x19\x3d\x4a\x26\xed\x6e\xb1\xc7\x57\x45\x77\xb0\x20\xf9\x6b\xa1\x58\x0d\x14\xeb\x47\x7f\x23\xb7\xfa\xc0\x46\x20\xa6\x2a\xc8\x03\x3f\x59\xa8\xd3\xf9\xc0\x38\x23\x4f\x4a\x5e\x75\x4c\xab\x33\xdb\xd5\x77\x17\xcf\x92\x45\x67\x7b\xaa\x73\xca\x49\xaa\xe0\x5e\xc2\x31\x60\x9e\x31\x55\x72\xab\x35\x9d\x3a\x97\xa2\xa7\xdc\xc9\x65\xc6\x59\x48\x75\x18\x8d\xbb\x5f\x23\x59\xf3\x56\xef\x25\xf3\xbc\x37\x5c\xf5\x76\x49\x6b\x96\xe6\xff\xcc\x50\x2c\x7b\x1e\x5c\x3e\x3a\x27\x51\xe1\xe1\xa7\xd0\xef\x7f\xb6\x5c\x1a\xe0\xf3\x20\x6e\x2d\xe0\xda\xcb\xd7\xfd\x0b\x7c\xdd\x3b\xf8\x4e\xfe\x02\xdf\x49\xc1\x77\x85\xd2\x94\xb2\x8d\xcc\x61\xe2\x9c\x6c\x28\x67\x17\x7c\xba\x5a\x58\x76\x99\x74\x08\x92\xca\x19\x17\xdc\xa7\xab\xc5\xe2\xcc\x03\x94\x50\xc7\x0d\x26\xc1\xb3\xd1\xe1\xb8\x20\xbc\x26\x52\x39\x6e\x07\x21\xc2\xcf\x8e\x9f\xbb\xd8\x82\x14\xc9\x2c\x61\x17\xc7\xd1\xc4\x9d\x1c\x07\xcf\x6d\x13\x44\xa9\xc3\xb8\x50\xdb\xbd\xf2\xa3\xc0\x8d\xc6\xee\x8b\xa3\x82\x5a\xf2\x2c\xa7\xee\x52\xe2\x70\x72\x74\xf8\x7c\xec\x8e\x6a\xda\x76\xb1\x0d\x22\x32\x7a\x71\x14\x46\x6d\xb6\x5d\xd4\xf8\xf9\x71\x74\x38\x41\x87\x85\x6d\x98\x30\x25\x50\xdc\x49\x4b\xc6\xe4\x59\x74\x7c\x1c\x1e\x9c\x13\xc9\x33\x81\x89\x71\xfb\x1c\xcb\x99\x4d\xfc\x6a\x67\x30\x98\x3d\x9f\x19\xe0\x2e\x06\xa7\xf9\xcc\xd7\x08\x97\x03\x9c\x01\xea\xd6\x91\x0a\x41\xed\x87\xa1\xce\xbb\x44\x4a\x58\x28\xdf\x32\x0f\x2e\x3f\x5b\x48\x13\x3c\x25\x42\x51\x52\xa2\xd9\xfb\xf5\xec\x23\x67\x64\x11\x12\xa6\x68\x44\x0b\xd5\x74\x72\xe9\xdc\x5a\x44\xbb\x52\x76\x3a\x2a\xa9\xb2\x69\xc8\x4d\xe3\x7a\xaf\xfb\x98\x07\x8f\xfc\x2c\x80\xc7\xdf\x5a\xf5\xf3\xbd\x72\xc8\x64\xac\x31\xe7\x0d\x37\xc7\x1e\x22\xdd\x7d\xb0\x74\xf7\x07\x4a\x9f\x3c\x58\xfa\xe4\x7e\xd2\x97\xa6\x5f\xd4\x9a\x9a\x81\x40\x7b\x60\xc6\x99\x42\x94\x11\x51\xf4\x19\x59\x40\x2f\x65\x06\x7a\xcb\x1b\xc4\x0e\x8d\xed\xc9\x6a\x6f\x6b\xe3\xbe\xa5\xe9\xea\x8d\x33\x41\x8c\x12\x6b\x1e\x53\x5c\x36\x88\x22\xb3\x7d\xba\x61\xa8\xd2\xa5\x2f\x68\x42\x78\xa6\x3c\x58\x5f\x8c\x8f\x56\x66\xf9\x5d\x1a\x22\x45\xea\xc7\x2b\x09\x7b\xce\x63\xfd\x9f\xa5\xda\x31\x5a\x51\x56\x9a\xb8\x60\x3e\x11\x57\x14\xd7\xac\x33\xf6\x9d\x22\x85\xb7\x4d\xbb\x75\xef\xce\x24\xd1\xaa\x54\xf5\xd0\x9f\x0f\x88\xaa\xb7\xac\xae\xbc\xf4\xa0\xaf\x44\x46\xf4\xf1\xb6\x7b\xef\x2e\xbc\x8e\x60\xd9\x3b\x40\x65\x2a\x33\xf6\x76\x4f\x66\x54\xed\x0c\xc6\x86\x49\x75\x56\xc1\x3c\x49\x10\x0b\x6b\xf3\x0b\xc0\x68\xfc\x05\x85\xe1\x97\xa2\x77\x7e\x51\xfc\x0b\xae\xc2\x4a\xeb\x7c\x9e\x8e\xff\x6a\xec\x02\xfc\xe3\xd1\x30\xa0\x6c\x18\x20\xb9\x6d\xed\x11\xbc\xe5\x1a\x87\xbe\xcc\x96\xef\xfc\x8b\xf9\xf9\xc9\xe3\x6f\x3b\xfc\xfa\x0e\xf0\xcb\x2f\x30\x24\x0a\x0f\x09\x96\xfa\x6f\x60\xb5\xaf\xb0\x89\x68\x4c\x1a\x9a\xf7\xcc\x09\x1c\x31\xfd\xe7\x6c\xb3\xd4\x9c\xea\xb5\xd5\x66\x8a\x30\xb5\x57\xed\xcb\x04\x51\xf6\xb9\xb5\x2c\x15\xc2\x5f\x4f\x1e\x7f\x33\xae\xf6\xf5\x8f\x6a\xbd\x15\x1f\x61\x1a\x5f\x41\x66\xdb\x60\x93\x2a\xe1\xa1\xce\xa7\xd1\x68\x74\x38\x1a\xf5\x1b\x9b\xfc\x9a\x11\xe1\x81\xe0\x5c\x35\x76\x36\x06\xa7\xdb\x3b\x3b\xb3\xb7\x9c\x7f\x95\x83\xd0\x98\x8f\x32\xc5\x1d\x41\x62\x8e\x42\x22\xfe\xa4\x23\x5a\x7c\x1c\x2d\xa1\xed\x1a\x25\xe8\x66\x43\x84\x3c\x49\xb9\x54\x83\xcc\x54\x5a\x8b\x28\x45\x6a\x7b\x52\x36\xac\x41\xbb\x12\x06\x45\x52\x0f\xf6\x66\x73\x8b\x29\xc2\x7a\xf3\x64\xc8\x53\x35\x44\xd7\xd2\xe4\x9b\xd6\x9a\x32\xaa\xc0\xb9\x02\xc7\x31\x61\x83\x6a\xd8\x34\xda\x7d\x07\xc7\x11\xb9\x2e\x1d\x45\x69\x76\x75\xe8\xe0\xce\x40\x02\x88\x8c\x21\x79\xd2\x08\x89\xb4\x60\xd2\xc8\x4e\x79\x2b\xaf\x68\xad\x22\xf3\x28\xd8\x5c\x6d\x2e\x03\x10\x86\x82\x98\x84\x15\xf4\x68\xee\xcb\x4c\x90\xf3\x8c\x31\x0d\x15\xfb\xa8\x3a\xea\x04\xec\x70\xd7\x5d\x2d\x77\x52\xfe\x41\x82\xed\x19\x02\x16\x09\xda\x90\x85\xc6\x89\x97\x94\x85\x0b\xb6\x42\x29\x5c\x36\xa6\xc4\xa7\xb6\x41\xf4\x2a\xde\xee\x3d\xb5\x33\x0f\x14\x09\xe7\x13\x9c\x09\xaa\x6e\xf3\x9b\x25\x5c\xda\x33\xaf\xb9\x54\xfe\xab\x92\xaa\x76\x81\xb2\x14\x1d\xf7\xc4\x05\x4a\x8a\xd5\xb5\xe0\xda\x49\x39\xed\x7c\xe6\x36\x36\x1a\x17\x2b\x78\xb4\x88\xe0\xb2\x72\x79\xc8\x55\xaf\xff\xea\x55\x3b\x6f\xaf\xd0\xed\x9d\x24\xe2\xac\x02\xdb\x60\x5a\xfb\x29\x92\xe4\xd9\x61\x35\x46\x1d\x05\x59\x01\x53\x70\x6e\xea\xe5\x75\x9b\x25\xf6\xae\x13\xc7\xe0\xdc\x02\xba\x96\x8e\x8e\x50\xc0\xb9\x92\x4a\xa0\xb4\x46\xfc\xb7\xd4\x4a\x4b\xa8\x34\xad\x11\x1c\x02\x8f\x7f\xbd\x9f\xe4\x8e\xa1\xf5\x0e\xd1\xed\x30\xb6\x1a\xed\x62\xba\xd2\xa8\xd2\x8e\x75\x3b\x83\xd7\x48\x6d\x3d\xe8\x0d\x8b\xea\x38\xe7\x95\xa2\x72\xca\xc4\xd1\xcb\x56\xb6\xfe\xd6\x2d\x30\xa7\xe9\x92\x32\x95\x32\x4b\x88\x26\xb0\xc3\xcc\x19\xc7\x59\xa2\x01\xba\x74\xa5\xaf\x90\x22\xf5\x25\x07\xe6\x51\x44\xb0\xf2\xa0\xfa\x74\xc3\x0a\xa0\x0c\xd3\x14\xc5\xf5\xea\x2f\x46\x9d\x83\x7a\x91\x13\xec\x0e\x50\x82\x7e\xe7\x0c\x5d\xeb\x76\x9b\x54\xf6\xa7\x06\x65\xeb\x8f\x39\xa4\x92\xde\x4e\xe1\x3d\x7e\x32\x76\xd0\xaa\xab\xac\x65\xb6\x90\x08\x96\x4e\x8e\x95\xbb\xc9\x6a\x8f\xe5\x9d\xb6\xdf\x65\x7d\x97\xd6\xd6\x4e\xe9\x99\x91\x93\xec\xee\x43\xcd\x7d\x9d\x46\x7a\xab\x95\xec\x1d\xb4\x67\x44\x3c\x84\x9a\x4a\xcc\xaf\x88\x58\xf3\x38\x9e\xb3\x30\xe5\x94\xa9\x0e\x32\x3f\x0b\x12\xaa\x7e\x6e\xed\x08\xaf\xbd\x26\x3d\xcd\xac\xb6\x5c\x74\x59\x0f\x7a\x3f\xeb\x50\x58\x84\xec\xb8\x15\xba\x9e\x57\x03\xd5\x7d\xb7\xb8\xdd\x13\x5b\xc8\x31\xab\xeb\x3a\x62\xc8\x8a\xba\x37\xfc\x6a\xcf\x22\xe7\x33\xdf\x68\x52\xe2\x38\xec\x64\x36\xc0\x7d\xc1\x36\x82\xc8\x4a\xda\x2c\xd2\xb5\xe0\x8a\x63\x1e\x7b\xa0\xf0\x0e\xd0\x5e\x0a\x9e\xac\xb9\x30\x2f\x1a\xdc\x5d\xf3\xbb\xe0\x1d\x8b\x33\x1a\x8a\x45\x9a\xc3\x7c\xe5\x69\xe0\x3c\x0e\xfe\x0b\x9c\xb3\x3c\xfd\x0f\xf9\xe5\x78\xd4\xe1\x97\xea\x62\xe1\x97\xea\xab\x83\xf9\xf2\xd4\xd5\xb1\x3a\xcf\x3a\x70\xac\xed\x9a\x5c\xaf\x7d\xfd\xbf\x53\xc9\x8a\x8a\xa5\x32\xa5\x7e\xcf\x8e\x8e\x26\x47\xc5\xaa\x6f\x2f\x53\x35\x81\x7a\x9a\x78\x45\xd4\x54\x29\x1b\xbf\x41\xbe\x5c\x75\x70\x95\xc8\x96\x40\x85\x4a\x2f\xb8\xf3\xe5\xe9\xff\x82\x85\x2d\xe5\x3b\x4d\x6c\xfa\x61\x8e\xe5\x3c\x0e\xda\xb6\xc5\x48\x2a\x8a\x97\x1c\x85\xa7\x28\x46\x0c\x53\xb6\x79\xef\x7a\xde\x6e\x21\xc7\xc4\xb6\x99\xf6\x49\x83\xfc\xff\x33\x9c\xbf\xfb\x19\x4e\x63\x12\x6e\x0c\x22\x3a\x0f\xca\xf8\x2f\x75\x6f\x62\x5d\xcf\x04\xf7\xe5\x41\x7e\x60\x4f\x0e\x54\xd3\x64\x2a\x58\x31\x35\x1b\x59\x39\x49\xfe\xfa\xc8\xf6\xdf\x8a\x76\x56\x78\xc4\xc5\x35\x12\xe1\x0e\x93\x90\xd8\x10\x65\x2c\x69\xf2\xcb\x19\x55\x28\xca\xb9\xa2\x81\x62\xbb\xf2\x7b\x7d\x71\xb1\x2e\x8d\x6f\x33\xb8\xb7\x1b\x9a\x42\x3b\x86\xc2\x42\x89\x3b\xd4\x80\x07\x37\x88\xb7\x99\x4a\x33\x5b\x63\xfa\x5e\xf0\x4e\xe4\xe3\x5b\x95\x78\xab\x54\xea\x0d\x87\xe6\x91\xc9\x3c\x0e\x06\x67\x6f\x7c\x33\x2e\x0f\x0f\xa0\xf9\x02\x50\xf7\x95\x77\xe7\xcb\x56\x3a\x68\x57\xd7\xf8\xee\xbc\x5e\x4b\x80\x1a\xb3\xa9\x60\xc5\xcb\x44\xcd\xb7\x20\xb4\xaf\xae\xe7\x37\xda\xa6\xc2\xce\xfc\xaa\x64\x4d\x6b\x4c\xf5\x4e\x4b\x95\x8e\xa7\xd7\x0d\xbd\xaa\xc3\xda\xde\x37\x9c\x95\x77\x3f\x7f\x46\xa7\x42\xc6\xc1\xbf\x03\x00\x00\xff\xff\x75\x16\x3e\x1e\xeb\x21\x00\x00")

func assetsClusterYmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsClusterYml,
		"assets/cluster.yml",
	)
}

func assetsClusterYml() (*asset, error) {
	bytes, err := assetsClusterYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/cluster.yml", size: 8683, mode: os.FileMode(420), modTime: time.Unix(1484338456, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsVpcYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x99\x4b\x6f\xdb\x38\x10\xc7\xef\xf9\x14\xd3\x60\x01\x6d\x81\x3a\xb5\xa5\xa4\x68\x78\xf3\x3a\x4e\x63\x6c\x9b\x18\x96\xe1\x02\x69\xf6\x40\x4b\xe3\x58\x88\x44\x6a\x49\x2a\x0f\x14\xf9\xee\x0b\x51\x92\x1f\xb4\xe4\xfa\x15\xaf\x8b\x5c\x6a\x72\xc8\x99\xff\x4f\x43\xce\x48\xad\xd5\x6a\x47\xcd\xef\x6e\x1f\xa3\x38\xa4\x0a\x2f\xb9\x88\xa8\x1a\xa0\x90\x01\x67\x04\x2c\xbb\xde\xa8\xd7\xea\xe7\xb5\xfa\xb9\x75\xd4\xa5\x82\x46\xa8\x50\x48\x72\x04\xd0\x61\x52\x51\xe6\x61\x1f\x19\x65\xde\x4b\x3a\x04\x70\x81\xd2\x13\x41\xac\xf4\xe2\xc2\x02\x54\x66\x02\x8a\x43\x22\x11\x46\x5c\xc0\xa0\xdb\xd2\x0b\xfa\x2f\x31\x12\x70\x95\x08\xd8\xbd\x1e\x68\x86\x21\x7f\x42\x7f\x40\xc3\x04\x65\xb6\x69\x0d\x7c\x1c\xd1\x24\x54\x93\x5f\x7e\xe0\x51\x85\x7e\xee\x52\xcf\x91\x19\x23\x57\x8e\xf5\x36\x25\x31\xb9\xc9\x90\xa1\x82\x91\xe0\x11\x3c\x8d\x03\x6f\x9c\x06\x45\x53\x63\x70\xdd\x2b\xa0\x9e\x87\x52\x9e\x94\x87\xf6\x2d\x60\x5f\x91\xdd\xab\x31\x01\xeb\xdc\xca\x86\xe8\xf3\x64\xa8\xf1\xd9\x9a\x0f\xa8\x7e\xa2\xff\x3e\xd6\x67\x85\x75\xa9\x52\x28\x18\x81\xe3\x3f\xef\xee\xfc\x9f\x8d\x0f\xce\xeb\xfb\xbb\xbb\x93\x55\x7e\x7c\xcc\xff\x69\xbf\xbe\x3f\xd6\x5b\xb6\x38\x93\x4a\xd0\x80\xa9\x39\x8d\x56\x94\x48\x05\x43\x04\x0a\x8f\x34\x0c\x7c\x68\x75\x2e\x7a\x30\x0c\xb9\xf7\x40\xe0\xf9\x44\xff\x7d\x7c\x3e\x49\xa3\x1d\xc4\x5e\x2b\xf0\xc5\x5f\x7a\xae\x92\x96\x5e\xba\xfc\xb1\xad\xcb\xa6\x51\xc0\x69\x7c\x3a\x5c\x3a\xdd\x64\x18\x06\x5e\x06\xa1\x79\xdb\x58\x8b\x54\xf3\xb6\xb1\x63\x52\xf6\xe9\xef\x42\xca\x5e\x93\x94\xbd\x43\x52\x8d\xdf\x8a\x94\xb3\x26\x29\x67\x87\xa4\xec\x83\x26\xd5\xe2\xcc\x0f\xd2\x35\xba\x08\x5c\x51\x69\x1c\xc6\x8c\xd7\xf1\x25\x23\xe4\x9a\xab\xe3\xec\x67\x5a\x1d\xf4\x50\xfb\xdf\x84\x86\xf2\x98\xc0\x8f\x77\x3d\x1c\x55\x1e\xe4\x0f\x60\x59\xff\x94\x6d\x6f\x6f\xb1\xbd\xfd\xeb\xed\x9d\x2d\xb6\x77\x8c\xed\x7b\x28\x79\x22\xbc\xac\x58\x0e\xba\x2d\x32\x93\x22\xcd\xef\x2e\x21\xed\x96\x4d\x48\x71\x71\x77\x05\x8f\x51\xa8\xa0\xa8\xad\x00\xd3\x0c\x04\xed\x6d\xb6\x24\xe4\x26\x6d\x46\x87\x21\x5e\x30\xe9\x26\x71\xcc\x85\x22\x60\x29\x91\xa0\x65\x4e\x5f\x71\xa9\x18\x8d\x50\x1a\x06\x66\xab\x90\x39\x32\x46\x73\xdb\x3e\xbd\x97\x53\x1c\x7f\xe3\x0b\x81\x6b\x1a\x61\x3e\x02\xa0\x1b\x03\x02\xef\xdc\x64\x08\x7f\xfc\xd4\x02\x5d\x45\xbd\x87\xd4\xe8\x75\xf1\xce\x2e\xa7\x91\x4d\x17\x79\x9a\xe5\x19\x29\x49\xb2\x0a\x64\x83\xd8\xeb\xf8\x05\xae\x1c\xec\x22\xc8\xaa\xa4\xcb\xcd\xbf\xd1\x38\xb3\xe8\xc4\x37\xec\x2b\x4d\x98\x37\x26\x90\x52\xcb\xe7\x9b\x8f\x34\x08\xe9\x30\x08\x03\xf5\x72\xcb\x99\xd6\x8c\x21\x7a\x0a\x7e\x40\xfd\x03\xbc\xfb\x92\xee\x2a\xf3\x0c\xab\x22\x87\xea\x89\x8b\x07\x13\x5e\xe6\x77\x53\xc8\xb5\x58\x2f\xaf\x35\x16\xef\xfd\xad\x68\xdb\x3b\xa4\x6d\xef\x92\x76\xe3\x10\x68\xdb\x8b\xb5\x63\x2b\xda\xce\x0e\x69\x3b\xbb\xa4\x6d\x1f\x02\x6d\x47\xbf\xe0\xa4\xc5\x10\xd5\x17\xaa\xf0\x89\xbe\x94\xd3\x36\x8c\x2a\xa0\xee\x2b\xfc\xac\x02\xac\x14\xf8\xa0\xdb\xca\xe7\x9b\x4a\x51\x6f\x1c\x21\x53\x6b\xa5\x84\xe1\x65\x62\xb1\x48\x24\xd3\xd4\xe3\x89\xc2\x7e\x5a\x29\xca\x03\x9a\xce\xaf\x15\xc6\x9e\x33\x63\x5e\xce\x12\x25\x79\xc3\x15\x23\xf3\xe5\x0d\x23\x25\xcf\xa5\x42\xe6\x94\xc3\x44\xad\x09\x30\xb7\xbc\x40\xa9\x02\x46\xd3\x03\x3e\x73\x3e\xe7\xdf\x3a\x01\x56\x7d\x3e\x93\x42\x35\xf5\xd3\x94\x92\x7b\x81\x76\xb0\xec\xae\x29\x5d\xb0\x71\x71\xcd\xe6\x0d\xed\xf3\x8b\x56\x85\x64\xd4\x84\x3d\x09\xab\xaa\x63\x4b\x85\xd9\x5b\x08\x73\xf6\x24\xac\xaa\x64\x2c\x15\xe6\x6c\x20\x2c\x3f\xc0\x4d\x2f\x2c\x17\x31\x9d\x3f\xf0\xab\xa2\xc3\x86\x3c\x61\x7e\x3b\x1e\x63\x84\x82\x86\x5d\x2e\x94\x29\xb1\xcd\x94\xa8\xb8\xa3\x0d\xa3\x0a\xb1\x53\x2b\x83\xac\x81\x09\xa0\x97\x84\x78\x9d\x44\x43\x14\xe9\x4b\x61\xdd\x29\xfa\xf3\xae\xe0\x8a\x7b\x3c\x24\x60\x7d\xb2\x66\x6c\x9b\x5e\x96\x09\xfa\xfb\x58\xd1\xec\xdf\x0b\x94\x69\x83\x3f\xa2\xa1\x9c\x74\xf8\x4b\xee\x9f\x54\x73\x8f\xb2\x7b\x24\x13\x7a\x97\x82\x47\x3a\x02\xfb\xd4\x9a\x0c\xf6\x79\xea\xfe\xec\xcc\x39\xb3\xa6\xe4\x5c\xf7\xea\x70\x78\x9d\xbe\x09\x2f\x1d\x40\xf1\xc5\xf2\x97\xcc\x6c\xdb\x20\x96\x0d\xe4\xb8\xae\x94\x8a\x0f\x87\xd7\xd9\xff\x9c\x5f\x9f\xeb\x06\xab\x6c\xe0\x26\x51\x1a\xd6\xe1\x80\xaa\x6f\x05\x6a\xf6\x4d\x7b\x23\x4e\x26\xa6\xc9\x21\x34\x6a\xaf\x29\x66\xc5\x62\x53\xba\xe0\x6d\xdb\x83\xd5\x9e\x84\x51\x81\xf7\x2a\x6f\xab\x26\x61\x13\x79\xce\x5e\xe5\x6d\xd5\x2a\xac\x22\xef\x26\x51\x71\xa2\xb2\x6f\x5e\xba\xd8\xe7\xfd\xf6\xcc\xa7\xc6\xfe\x18\x21\xf0\x81\x8f\x40\x8d\x11\x1e\xe3\xac\x9e\x17\x95\x7b\xb6\x35\x68\x3f\xeb\xaf\x5a\x85\x7b\x1a\x55\x97\x76\xed\x6c\xf1\x68\x54\x06\x90\xb5\x02\x20\xb5\x61\xe0\xcf\xfd\xd7\x40\x16\x4a\xee\xf5\x92\x11\xd2\x19\x4d\xdb\x93\x8a\x03\x91\x4e\x2d\x49\xfc\x7c\x52\x47\x7d\xcd\xb5\x83\x75\x15\x2e\x28\x5b\x3c\x29\x6b\xaa\xb5\x37\x50\x6b\x2f\x53\x6b\xbf\x95\x5a\xbb\x44\xad\xb3\xa6\x5a\x67\x03\xb5\xce\x32\xb5\xce\x5b\xa9\x75\x3a\xfe\xd1\x7f\x01\x00\x00\xff\xff\x7b\x24\x36\xe9\xff\x1d\x00\x00")

func assetsVpcYmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsVpcYml,
		"assets/vpc.yml",
	)
}

func assetsVpcYml() (*asset, error) {
	bytes, err := assetsVpcYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/vpc.yml", size: 7679, mode: os.FileMode(420), modTime: time.Unix(1484338624, 0)}
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
	"assets/cluster.yml": assetsClusterYml,
	"assets/vpc.yml":     assetsVpcYml,
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
		"cluster.yml": &bintree{assetsClusterYml, map[string]*bintree{}},
		"vpc.yml":     &bintree{assetsVpcYml, map[string]*bintree{}},
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
