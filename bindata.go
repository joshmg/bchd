// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// sample-bchd.conf
package main

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _sampleBchdConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x7c\x6b\x73\x1c\xb7\xb1\xf6\xf7\xfd\x15\x5d\xa9\xa4\x48\xb9\x96\xcb\x5d\x4a\x94\x15\xd2\xeb\x2a\x4a\x72\x1c\xbe\xaf\x2e\x2c\x53\xca\xa5\x5c\x29\x17\x76\xa6\x67\x07\xe1\x0c\x30\x06\x30\xbb\x5c\x9f\x4a\x7e\xfb\xa9\xee\x06\xe6\xb2\x14\x4d\xd9\x47\xd4\x87\x84\x3b\x33\x68\x34\xfa\xfa\x74\x03\xf0\x8f\x17\x4d\x53\xe9\x4c\x05\x6d\x0d\xbc\x6f\xe8\xff\xfc\xbf\x26\x93\x73\x38\xfa\xa2\xff\x26\xe7\xf0\x5a\x05\x05\x1e\x43\xd0\x66\xed\xbf\xfc\x04\x93\x73\xf8\x50\x22\xe4\xda\x61\x16\xac\xdb\x41\xb0\xe0\x83\x75\x08\x39\x4f\xdc\x66\x25\x28\x0f\xa1\x44\x58\x55\x36\xbb\x81\xac\x54\xda\x80\x32\x39\x34\x88\x0e\x54\x9e\x3b\xf4\x1e\xfd\x0c\x88\xd0\xe4\x7c\xf4\x59\x50\x37\xe8\xc1\xe3\x06\x9d\xaa\xe0\xfb\x97\x53\xf0\x16\x42\xa9\x3d\x54\x36\x0a\xaf\x6e\x7d\x80\x52\x6d\x10\x14\x54\x36\x80\x2d\xa0\x70\x88\xe0\x1b\x95\xe1\x2c\xb1\x87\x85\x6a\xab\x00\xda\xc3\x7f\x8f\x67\xab\xac\xcc\x8f\x99\x3d\x6b\xe0\xea\xfd\xf5\xe5\x3f\xe0\xfd\x35\xfa\x29\xfc\xf1\xcd\xfb\x57\x17\x6f\x2e\xae\xae\x5e\x5f\x7c\xb8\x38\x7e\x39\xfc\xec\xef\xda\xe4\x76\xeb\xa7\x93\x73\xf8\xef\xf1\x1b\xbd\x72\xca\xed\x8e\x87\x4a\xbc\x6e\x9b\xc6\xba\x30\x1e\xf5\x56\x65\xf0\xfe\x7a\xca\xcb\xfd\x63\x69\x6b\x3c\x1e\xce\x3d\x39\x87\xab\x4a\x99\x3f\xcf\x00\xbe\x33\x1b\xed\xac\xa9\xd1\x04\xd8\x28\xa7\xd5\xaa\x42\x0f\xca\x21\xe0\x6d\xa3\x4c\x8e\xb9\xac\x1c\x77\x50\xab\x1d\xac\x10\x5a\x8f\xf9\x0c\xe0\xdd\xfb\x0f\xdf\x9d\x25\xee\x26\xe7\x80\xf7\x12\x0a\xbb\x46\x67\xaa\xaa\x76\xf0\xa7\xbf\x5d\xfc\x70\x79\xf1\xf2\xcd\x77\x7f\x9a\xc2\xaa\x0d\x91\x2c\xc9\x71\x85\xa0\xb2\x8c\xf4\x91\xc3\x56\x87\x72\x72\x0e\x7f\x4c\x1f\x43\x89\x0e\x67\x00\x17\x95\xb7\x53\xf8\x2f\xc9\xb2\xe3\x2d\xd8\xb1\xec\x06\x12\x23\x15\xd0\x7a\x73\xed\x96\x43\xd9\x4f\x1e\xc5\xda\xdf\x61\xd8\x5a\x77\xf3\xb8\x06\xff\xd1\x23\x04\xf4\xc1\x60\xa0\xd5\xc5\x3f\x97\x8b\xee\x5d\x89\xe0\x70\x4d\x76\x4d\x96\x41\xef\xc1\x08\x63\xf4\xbd\xc3\x35\x3d\x92\xef\x2f\xaa\xca\x6e\x21\xb3\xc6\x60\x46\x1c\x93\x28\xc9\x31\x3c\x14\xce\xd6\xa0\xcc\x0e\x4a\xeb\x03\x6c\x4b\x34\xd0\x7a\xfa\x62\x9f\x74\x6d\x73\x9c\xc1\xcb\x1d\x09\x5a\xec\x7c\x9a\xe6\x00\x63\x73\xf4\xb0\xd5\x55\x05\xd6\x54\xbb\x34\x11\xcd\x62\x43\x89\x2e\x7e\x40\x53\x60\x4e\x5a\x43\x4d\x8f\x27\xe7\xec\x60\x15\x4f\x6d\x1d\x2c\x4e\xbe\x9e\xcd\x67\xf3\xd9\x62\x06\x1f\xc8\xfb\x2c\x47\x2c\x32\x81\xd6\x63\xd1\x56\x43\xf6\x6a\x72\xfe\x50\x2a\x03\xd6\x20\x10\x53\x36\xbb\x41\x47\x53\x07\xa5\x0d\x2d\x2d\x58\x70\xad\xd9\x5f\x88\x1f\x08\x47\x99\x1d\xcd\x2d\x32\x7a\x6d\xcd\x41\x00\x87\x1e\x43\x1f\x48\x24\x40\x90\x25\xad\x94\x47\xd0\x77\xc8\x75\x72\xe9\xa4\x42\xba\xda\x1b\xbe\x12\xd9\xac\x30\x92\x57\x01\x7c\x50\x2e\xb4\xcd\x80\x19\x63\xf9\xe5\x58\xc1\x5e\xd7\x6d\x25\xae\x3f\x54\xf0\xe4\x9c\xde\x74\xe6\xf0\x2a\xca\x7b\xa3\x15\x28\xb8\x7e\xff\xea\xff\x5f\x9f\x42\xe3\xec\xed\xae\xf3\xdd\xeb\x06\x33\x5d\xec\x48\x74\x4a\x5e\x09\x4f\xb9\xf6\xe4\xbc\x50\x69\x1f\xd0\x68\xb3\x9e\x9c\x43\x61\x1d\x68\x93\xd9\x9a\xbe\x4e\x46\x63\x8d\x87\xd6\x54\xe8\x7d\xfc\xb6\x0f\xaa\xec\xf8\x8d\xb3\x1b\x4d\x5e\x4a\x4c\x10\xeb\x07\xf2\xd9\xc1\xe4\x3c\x2a\x92\xd6\xca\x33\x2f\x3b\x45\x9f\xfd\x79\x7e\x3a\x4f\x8f\x5b\x8f\x6e\x99\x7e\x34\xca\xfb\x65\x8a\xfb\xc3\x15\x81\x5a\xd9\x0d\x92\x51\x28\xef\xdb\x5a\xc2\xc2\x0a\xe1\x83\x75\x70\x58\x86\xd0\xf8\xb3\xe3\xe3\xed\x76\x3b\x0b\xd6\x35\xce\xfe\x1b\xb3\x30\xb3\x6e\xfd\x84\x66\xbf\x2c\x98\x33\xa1\xa3\x3d\x18\x4b\x36\xea\xf8\x61\x61\xc9\x47\xd8\xb4\xfa\xd0\xc7\x7e\xe2\x70\x43\x71\x4e\xec\x2e\x58\x32\x5b\xcf\xd2\xd4\x99\x70\x06\x3f\xb7\xe8\x34\xb2\xc5\x55\xd6\xde\xb4\xcd\x40\x36\x87\x9c\x48\xb4\xc9\x1c\x2a\x96\x95\xb1\x66\x57\xeb\xb0\x13\x6b\x16\x7a\x62\xe2\x39\xac\x76\x69\x3a\x9a\x6b\x67\x5b\x07\x97\x57\xb0\x42\xfa\x55\xa1\xba\x89\xe2\x7d\xfd\xee\x9a\xd7\x63\xac\x35\xda\x9a\xde\x64\x94\x01\x55\x05\x74\x46\x05\xbd\x49\x0b\x0d\x76\xe8\x90\x33\x1e\x32\xcc\x88\x1f\x86\x22\x89\x42\x25\x23\x66\xb1\x2a\x16\xac\x61\x23\x7f\x67\xcd\x9d\xe1\x9d\x65\xb3\xe3\x65\x21\x86\x74\x16\x69\x4d\xc6\xcf\x94\xc9\x06\x1c\xbf\xb0\x6d\xe8\x0c\x50\x17\x60\xc8\x7b\x35\x25\x5f\x0e\x72\x71\x39\x43\xf3\x58\xa4\xc7\xc9\x3c\xf8\x47\x67\x1e\xdf\x19\x36\x5f\x62\xd2\x07\x87\xaa\x06\xed\x6d\xf4\x98\xd5\x0e\x9c\x32\xb9\xad\xf5\x2f\x24\x40\x99\x94\xc8\x40\xe6\x30\x27\x21\xab\xca\x93\x4b\xb6\x15\xcb\x5b\x53\xae\x24\x4a\xa4\xab\x20\xae\x62\x70\x0b\x99\x76\x59\xab\x03\xfb\x05\xaa\xac\x1c\xf8\x04\xe3\x09\xed\xa1\x66\x08\xa1\x83\xc4\xa5\x5c\x17\x85\xce\xda\x2a\x88\x18\x33\xeb\x1c\x56\x2a\xe0\xd0\x99\x38\xa6\x5b\xd7\x71\xdb\x2b\xf1\xa3\xd1\x1b\x74\x5e\x55\x70\x55\xb5\x6b\xce\xea\x57\x95\xda\xc1\xe1\xc7\x2b\x73\xf5\x84\x08\xaa\x36\xd8\x5a\x85\x98\x68\x6d\x43\x76\x54\x26\x1f\x06\x02\x08\x93\x73\x1e\x67\x57\x81\x01\x4e\x49\x19\x9e\xed\xa2\x22\x83\x8a\xda\x93\xe0\xef\x05\x52\x60\x0e\x39\x6e\x74\xc6\x18\x49\x62\xc6\x20\x08\x4f\xce\x45\xd1\x0c\x81\x8c\x05\x2c\x0a\xb2\x27\x5d\x00\xde\xee\x93\x8d\x01\x21\xfa\x08\xe6\xb4\xd2\xb6\x31\x8d\xac\x30\x06\xa2\xfb\x78\x42\x2f\x66\x4f\x16\x47\x86\xd1\xc5\x25\x60\x61\xbf\x37\xd8\x71\xdf\x48\x06\xd1\x86\xf0\x02\x65\x7c\x61\x91\x9c\x3a\x1a\x23\x3c\x75\xf9\x51\xa3\x5c\xd8\x81\xd7\x41\x1c\x34\x8a\xa4\x9b\x5a\x0f\x9c\x55\x80\x1c\xa9\x13\x95\xf1\xb4\xba\x9d\x6d\x79\x31\x2b\x2c\xb5\xc9\xe1\xdd\xc5\x87\xe9\x80\xbf\x6e\x3e\x72\x14\x32\x43\xd2\x4d\xbe\x41\x17\x34\x39\x23\xc7\x76\x95\x95\xfc\x2a\x71\x1d\x63\x28\x11\xf6\x51\x14\x3a\x30\xea\x21\x17\x41\x31\x67\x8e\x9f\x24\xb3\x83\x94\x04\x0f\x95\x21\xa7\x8c\x10\x74\x5f\x67\x1c\x0d\xd2\x92\x74\xb3\x5c\xcc\x4e\x66\x4f\x67\xcf\xc6\x0f\x4f\xe6\xf3\x93\xb3\xb3\xc5\xc9\xd3\x67\xa4\x87\xaf\xbe\xe8\x3f\x52\x6c\x5b\xd7\xca\xed\x08\x1a\x1f\xa8\x3c\x27\x78\x71\x00\x64\xc8\xad\x87\x83\x68\xf8\x07\xb3\xc9\xf9\xe4\x1c\xde\x13\x48\x20\xcf\xb7\xc5\x5e\xec\x0d\x5b\x1b\x57\xec\xa7\x03\x32\x64\xcb\x1d\x8d\x69\x8c\xd0\x7d\x1c\x66\x8c\xfa\xd2\x86\x12\x14\x03\x1d\x12\x2e\x55\x08\x51\xbe\x82\x74\x42\xa9\x02\xbf\xd9\x2a\x13\xa4\x80\x50\x1d\x54\xe1\x88\x47\xf1\xa9\xc3\xaa\xa4\x72\xb5\xd1\x94\x0a\x3d\xf8\x4a\xaf\xcb\x50\xed\xd8\xb3\xd1\xa1\x09\x34\x61\x9f\xf3\x07\xe6\x47\xb9\x60\x47\x11\x5a\x7c\xbf\xd0\xb1\x06\xf1\xe3\x38\xcc\x35\xc6\xc0\x16\x92\x62\x53\x62\xa1\x6c\x61\x4d\xaa\x6c\x18\x53\x58\x4f\x25\x86\xcf\x9c\x5e\x51\xaa\xc0\xca\x6e\xd9\x18\x29\xb0\xad\xd4\xaa\xda\xc1\x96\x21\x8c\x41\x49\x5c\xb5\xcd\x69\xf5\xca\xec\x42\x49\xb2\x65\x64\xcd\xf2\xef\x05\x9b\x5b\x94\x34\x98\xb2\xdc\x30\xdd\x77\xf0\x4d\xf8\xcf\xb5\xcf\xec\x06\x1d\xe6\x1c\x38\x22\xce\x91\x77\xc9\x4f\x3a\x71\xb2\x2b\x98\x1c\x54\xe5\x2d\x54\x18\x7c\xc4\xab\xb5\x0d\x69\xcc\x8d\x89\xaa\x52\x8e\x74\xa9\x36\x4a\x57\x6c\xfd\xa9\x06\xc9\x94\x21\xde\x68\x11\x43\x3e\xba\x77\xe3\xc4\xb6\xb3\x6d\x8c\xc6\x1d\xe2\x80\x9a\xd4\x16\x93\x39\x01\xc8\x81\x47\x93\x72\x25\x29\xac\x2a\xac\x3d\x2b\x2a\x86\x7c\x72\x6d\x8a\xf5\xde\xd6\x28\x3e\x4c\xaa\x38\x6c\xd0\x95\xaa\xf1\x90\xb7\xe2\xe8\x50\x68\x87\x5b\x55\x55\x4f\xa2\x54\x7b\x03\xb5\x12\x73\x85\xeb\x52\x99\x7c\x2a\xc6\xf1\xfe\xdd\x9b\x7f\x0e\x79\x66\x88\x97\x6c\x38\x2e\x4f\x1c\xdd\x44\xd9\x53\x34\xbe\x0c\x22\xc6\x88\xd5\x86\x41\xf1\x70\x60\x42\x78\x4b\x75\xa2\x26\x33\x25\x90\x29\x1f\x91\x60\xbb\xc8\xba\x0f\xcd\xa2\x98\x9e\xb0\xa6\x5e\xbf\xbb\x06\x8f\x48\x42\x60\xe3\x64\x57\xe9\x03\x1c\x13\x8a\xa1\x2d\xa7\x62\x9b\xb0\x46\xa7\x32\xae\xb2\xe3\x82\x7a\x8b\x18\xac\x94\x66\x10\xf3\xa4\xd2\xaf\xa1\x04\x35\x34\xb5\xae\x0c\x19\x28\x7a\x06\x70\x6d\xa7\xc2\x70\x12\x6d\x52\xac\xe4\x1f\xbd\xc1\x6a\x27\x3e\x4f\x4a\x8f\x6e\xbf\x5f\x82\xfc\x21\xb8\x96\x0a\x8f\x3f\x44\xb2\x5f\x3e\xf8\x51\x85\x95\x93\xfe\x9c\x67\xc1\x86\x4f\x79\x3c\xc9\x2c\x47\xaf\x1d\x47\x2b\x4a\x64\x2c\xb4\x06\x9d\xe4\xb0\xc9\x39\xfc\xd3\xb6\x1c\xdb\x52\xe0\x62\xb0\x31\xc8\xd7\x8c\xac\xc6\x40\xca\xba\x20\x3d\x8d\xae\xfb\x40\x8f\x58\x71\x93\x73\xce\x4b\x39\x81\xf2\x31\x62\xd0\x05\x44\xdc\x45\xba\xed\x0d\x30\x46\x08\x88\xe1\x61\xb9\xf8\xf3\xc9\x6c\xf1\xfc\xc5\x6c\x31\x5b\x0c\x9f\xce\x19\x9d\x9d\x9c\xbd\x78\xfa\xf4\xe9\xe0\x79\x81\x2f\xe6\x67\x67\xc3\x2f\x7f\x94\x47\x27\xff\x92\x4f\xef\x15\x53\x8a\xcc\xec\x1e\x29\x3c\x3f\x24\x39\xaa\x14\x92\xec\xe0\xff\x24\xba\xbe\xa4\xed\x85\xf7\x7b\x45\x77\xa7\xda\x0a\x83\x4a\xb6\x54\x3e\x1a\xb8\xd7\x39\x46\x23\xf6\x71\x79\x31\xae\xc7\xf2\xc6\xc4\xf0\x7a\x7f\x2a\x05\x1f\x13\xae\x8f\x50\xb4\x77\xa9\x3d\xc5\x75\x4f\xf7\x14\x97\x9e\xf7\x8a\x4b\x4f\xee\x2a\xee\xad\xba\xd5\x75\x5b\x83\x69\xeb\x15\x3a\x4a\xdc\xda\xac\x6c\x4b\x11\x9e\x70\x66\x1b\xe4\x47\xe7\x61\xb5\xba\xe5\xbf\x97\x8b\x93\xd3\x38\xfe\xb3\xc6\xb2\x4e\x2f\xaf\x86\x24\x1a\x74\xba\x59\x32\x95\xd7\x94\x82\xc4\x2c\xfc\xce\x64\x71\x88\xaf\xec\x96\xc2\x4f\x49\x39\x81\xc4\x1d\x4a\x87\xbe\xb4\x55\x4e\x05\xfb\x6a\x17\xd0\x1f\x7b\xcc\x98\xa6\x36\x34\x90\xc6\xc5\x52\xda\x37\x88\xf9\xf2\x74\x71\x32\x9f\xd3\x0c\xef\x3a\x1e\x3b\xbe\xf6\x52\x22\x55\x35\x04\x21\x19\xc0\x2b\xb7\xc6\x90\xbe\x94\x05\xbf\x88\x8c\x72\x5c\x5c\x29\x23\xe0\xb5\x80\x5a\x7b\x81\x14\x94\x75\x92\x98\x8c\x8d\x5f\x08\x38\x4e\x52\xe6\x90\x46\x49\x5e\x19\xf0\x99\x65\x00\x5a\x48\x65\xe1\x07\x5d\x1c\x12\x5f\x9a\xe1\x93\xe4\x57\xca\x74\xb2\x58\x2e\x64\x85\x7f\xb5\x5b\xa8\xac\xf4\x80\x88\xfe\xdd\x81\xf0\x37\x55\xe9\x1c\x82\xae\x11\x5a\xa3\x83\xe0\xf9\xff\xf1\x53\xa8\xa7\x50\xfe\x87\x08\xbf\xd5\x86\x19\x5d\xa4\x69\xf2\xd6\x49\x19\x73\xf2\xac\xdc\x7b\xb2\x58\x94\x4f\xe7\xf5\xe2\xd4\xa7\x00\xb0\x2d\x75\x40\x4e\x41\x39\xb9\x68\x52\x04\x2f\xe7\xf2\xca\xcf\x52\x3f\xa3\x4b\x89\x5b\xc6\x3e\x97\x57\x50\xab\x90\x95\x54\x5f\x10\x5a\x4b\x54\xfa\x2c\xc5\x20\x2a\x94\xa8\xdd\x40\x72\xa9\xf4\xe6\x3a\xa4\x1b\xd4\x17\x99\xa3\xa7\xe2\x06\x83\xaf\xa2\x23\xcd\x67\xf3\xe3\x93\x67\xa3\x57\x45\x3e\x9f\x9f\x9d\x1d\x2f\x9e\x0f\xf5\x3d\x48\xa2\x0c\x21\x52\x22\x1b\x62\x45\xae\xfb\x19\x30\x72\x13\xc8\x4f\xa9\x00\xe0\x35\xb4\x9e\x62\x11\xd1\x08\x36\x02\x49\x22\x32\x4e\xb3\xa3\xb4\x42\xe1\x4c\xec\x28\x37\x9e\x26\xbe\x5b\x64\x69\x13\xd0\x15\x2a\x8b\xfd\x09\xa9\x11\xbb\x62\x6a\xdc\xcb\x19\x65\xa3\x54\x04\xee\xa5\x16\x2a\x8f\x08\x59\x6a\xe9\x57\x10\x48\x4a\x90\xbd\x6b\xc4\x1f\xc4\x6e\xe5\x01\x23\x09\x4d\x83\x18\x48\x65\xb6\xae\x31\xf5\x72\xfb\x00\xba\x8b\xe1\x38\x22\x46\x82\xf0\xdc\x23\x20\x6e\xd2\xdc\xd2\x8e\xc8\xc8\x12\x28\x36\x3e\x0c\x9d\x83\x85\x3c\x82\xa8\xad\xf6\xbc\xa2\x8b\xaa\x1a\x8a\xc3\x9a\xf1\xca\x62\xab\x46\xa0\x6a\x7c\xf3\xe4\x6c\x72\x0e\x51\x4a\xcb\x44\xa2\xd9\x3c\xfb\x15\x3a\xc3\x11\x12\x71\xe7\xfd\xc0\xe7\x0f\x0d\x4c\x23\xcf\xce\x3e\xc9\x30\x33\x4a\x41\x79\xfc\x71\x8c\xe8\xf7\x70\xf7\xe9\x41\x91\xb7\xbd\xb1\xfb\x0c\x7e\x7a\xec\x8f\x67\x67\xff\x4a\x03\xb9\x9a\xe3\x59\x07\xdd\xdc\xfb\x06\xf6\xbd\x9d\xbd\xd1\xcf\x3f\x67\xf4\x8f\x67\x67\x8b\x87\xe6\x35\xd6\x1c\xf9\xa0\x4c\xae\x5c\xde\x91\x79\x7e\x3f\x13\xcf\x3f\x29\xe7\xcf\xa0\x32\x1a\x7c\x57\xe8\x9f\x41\x61\xa0\x81\xe7\xf7\x6b\xe0\x33\x08\x25\x75\x8c\x62\x51\x5f\x25\xdc\xd7\xd0\x4d\xdd\x2b\xe9\xdc\x6f\xd0\x39\x82\x25\xaa\xaa\xe2\xd8\x2e\x49\x25\xb9\x0d\xc9\x33\x7a\x5d\x55\xd6\xd6\x50\xe8\x2a\xa0\xd3\x66\x4d\x90\x1d\x11\x5e\x5e\x5e\xcd\x17\x8b\x85\x8c\xa5\xef\xf8\x33\xf9\xca\xc7\x1d\x89\x3c\xa7\xa0\xa3\x89\x0f\x55\x41\x56\x62\x76\xd3\x58\x6d\x82\x9f\xc1\x5f\xac\xab\x55\x38\x83\x83\x6f\x4a\xa4\x02\xee\xdb\xb3\x6f\x4a\xe5\xcb\x6f\x0f\x04\x5a\xf6\xdf\x2e\xf7\x3e\x18\x65\xde\x56\x57\xe1\x48\x9b\x31\xe9\xd8\xe5\xcf\xe3\xfe\xde\x20\x8a\x70\x35\xba\x8d\x48\xf4\x80\x52\xad\xa5\x05\xf1\x12\x06\x24\x7a\xee\x25\x96\x05\x9f\xca\x1e\x6e\x2c\xaa\x35\xc1\x5a\x06\xb4\xda\x0f\x0b\xa6\xb4\xe1\xc2\xc9\xb3\xe5\x3d\x93\x40\x3a\xa9\xda\x9c\xa2\x9a\x72\x2a\x23\xe1\xc0\xc1\xf1\xc1\x14\x0e\xce\xe8\x7f\x0e\x63\xdf\xe3\xc9\x01\x77\xd0\x54\x9c\x70\x39\x5c\x25\x3d\xd3\x21\x65\xca\x5e\x11\x70\xf8\xea\x2f\xb1\x45\x9c\x0d\xe4\xfe\x18\x9b\x61\x3f\x5c\xbd\x02\x8f\x6e\x43\xa8\x29\xa6\x81\x23\xce\x1a\x7d\x57\x27\x3d\xcf\xac\x09\xce\x56\xd2\x5e\x49\xfa\xe9\xc7\x4b\x7a\xcd\xca\xae\x1d\x2e\x89\x8e\x87\x90\x24\x24\x23\x6a\x53\xb0\x7d\x10\xa0\x96\xb2\x11\x5c\x2b\x18\x88\x93\x6a\xe3\x6c\x86\xde\x4b\x4d\xde\xe7\xb0\x01\x9b\xda\xa7\x52\x9a\x33\x58\xb7\x6d\x5b\x80\x6b\x32\x56\xe3\xc5\xbb\xd7\xf4\x77\xa3\xbc\x9f\x02\x77\xe8\x5d\x93\x55\xba\xd6\x61\xf8\x9a\x1f\xc8\x37\x04\x90\x46\x05\xc1\xec\x51\xf6\x04\xaf\x31\x6b\x9d\xb4\x07\x69\x3d\x17\x57\x97\x9c\x82\x87\xd5\x86\x18\xa2\x51\x35\xca\xf6\xb7\xf2\x7e\x6b\x5d\x1e\x4b\xa4\x8c\x77\x08\xbc\xed\x7a\x67\x94\x7a\x79\x1d\x98\xff\xea\x40\xde\xb8\xed\x86\x04\xa8\x50\x71\xb8\x25\xc0\x52\xb4\x55\xc5\x2d\x07\x5b\x8c\x3a\xeb\x47\x1d\x65\x02\x31\x79\xad\x0d\x1c\x41\xdc\x6e\x19\xa8\xa3\xaf\x55\x93\x56\x66\x22\x70\xee\xf8\x93\x4b\xe2\x06\xdd\x4f\x4c\xe0\xa7\xc4\xe3\x4f\x3b\xdb\xfe\x44\xa5\xa2\x7c\x2a\xfb\x01\x63\x35\xf5\x43\x23\x1b\xf7\x0d\xee\xf4\xb8\xfc\x15\xec\x54\xdc\x65\xfc\x61\x2c\xd5\xf7\xa7\xbf\x08\x98\x22\xad\x45\x38\xf5\x05\xc0\x14\xd5\x7b\x0c\xa7\x7e\x07\x98\x1a\x23\xda\xc0\x35\xf4\x9e\x4a\xa5\x0f\xd3\xc9\x68\x90\xa4\x49\x94\x97\x57\x9b\x67\x11\xf0\x6f\x9e\x3f\x8c\xcd\x24\xdb\xb1\xae\x7e\x2b\x12\x1b\x8c\xfa\x1d\x68\xac\x1f\xfc\x00\x20\x7b\x76\xe7\x7b\x7a\xf8\x30\x26\xbb\x33\x6e\x00\x0a\x9e\x3d\x0c\xcb\xee\x0c\x4f\x50\xe0\xd9\xc3\xc8\xec\xce\xd8\x11\x2e\x7a\xf6\x30\x38\xfb\xd4\xe4\x8b\x87\x66\xff\x24\x9c\xf9\xfa\x57\x59\xf9\xfa\xf3\x21\xda\x1d\x42\xa3\xf1\x9f\x89\xd2\xee\x10\x19\xe8\xe4\xeb\xdf\x08\xd4\xee\xd0\x4a\x0a\xfa\x9a\x82\xcd\x5f\x74\x85\xe9\xcc\x42\x8a\xdf\x19\x61\x86\x42\x67\x2a\x20\x25\x74\x94\x20\x45\x4f\xbb\xa3\x2d\xae\xc9\x66\xf4\xe0\x73\x48\xdc\xe0\x4e\x28\xdc\xe0\x6e\x44\x80\x5e\xec\xc5\xbb\xfa\x4e\x97\x27\xb3\x26\x6b\x9d\x23\x54\x43\xfe\x9d\x55\x9a\x31\x0f\xb7\xc7\xd3\x52\xf7\x76\x35\x5d\x93\xd5\xea\x36\x7e\xb9\x5c\xcc\x7f\xf3\x24\x5b\x5c\x79\x9b\xdd\x60\x48\xd3\xf5\x54\xbb\x57\x7e\xf9\xa9\xbe\xd2\x1e\x21\x87\x3f\xb7\xe8\x43\xec\x30\xc6\x7d\xfd\x88\x0f\x30\x1f\x7c\x5d\xed\x06\x8c\x77\x4f\x1d\xfe\xec\x97\x27\xcc\xff\x5b\xed\x5c\xdc\x11\x80\xff\x77\xfd\xfe\xdd\x11\x91\xff\xb9\xd5\xee\xc6\xd3\xbc\x2f\x75\xc8\xac\x36\xf0\xca\x3a\x84\xa3\xa3\x18\xed\xb9\x5b\xd5\x3a\xb5\xa6\xdc\xca\x21\x76\x72\x2e\x26\x43\xc1\x58\xad\x74\xa5\xc3\x0e\xb4\xf7\x2d\xfa\x6e\xd7\x66\x85\xb0\xb5\xee\x06\x73\x50\xce\xb6\x26\xe5\x42\x99\x6b\x7c\x38\xa4\x07\x58\xf1\x20\x12\x87\xe9\x88\x0c\xf7\x72\x15\x6e\xd0\x10\xc2\xe1\x5d\x98\x88\x72\x64\xa7\x21\x66\xcf\xf1\x9e\xad\xb4\x21\x53\x7d\x20\x1d\x76\x6e\x7c\x72\x63\x42\x67\x37\xbc\x17\x36\x9a\x89\x92\x53\x8a\xfe\xd2\x63\x8d\x4d\xa5\x60\x79\xdf\x67\x83\x23\x70\xc0\xd0\x8d\x6d\xd5\x9a\x42\xaf\xd9\xd2\x05\xb0\xba\x26\xfb\x0d\xeb\xfc\xf0\xe6\xfa\x13\xb9\x79\xb4\x87\xdd\xef\x07\x71\x4a\x92\x66\x53\x94\xc5\x18\x16\x4a\xc3\x8f\x8f\x71\xa4\x88\x35\x70\xf1\xc3\x84\x4e\x63\x6f\x36\xb5\x08\x84\xed\x50\x3d\x1a\xc4\x5e\xff\x3e\x8c\x3d\x39\x87\xaf\xf0\xb6\x41\xa7\xa9\x6e\x50\xd5\x57\x23\x42\x0f\x43\x6d\xb6\xd6\xdf\x05\xb6\x87\xf3\xf0\xd6\x55\xeb\xd1\x27\xdb\xa3\xe0\xc4\x93\x48\x4c\x4a\x9c\x4b\x17\x5e\x1b\xe4\xae\x6d\xaf\x1b\x0e\x2c\xd1\x1e\x1f\x05\x54\x73\x29\x6a\x7a\x45\x1f\x73\xe8\xee\x7b\x65\x7c\xec\x63\x20\x46\x59\xdd\x20\xe8\x4d\xce\xe1\x70\x84\x1c\x28\xee\x9f\x4e\xd3\x91\xbd\x33\x58\xd0\x6f\x36\x93\x75\x9f\x07\xe8\x19\x4f\x6f\x40\xb5\xa1\x24\xbf\x88\x87\x3c\x83\xbd\x89\xd3\x86\x24\x4b\x42\xfa\x72\xf0\x23\x7d\x88\xc3\xe0\x48\x64\xe9\x1d\x8f\x5c\x7e\x63\xe9\xef\x93\x23\xfe\xf5\xed\xe3\x98\xe4\xeb\x74\x24\xee\x3a\x9e\x81\xfc\xac\xa2\xaf\x3b\x48\xc7\x01\xb9\xb4\x55\xee\x87\xc7\xe5\xe4\x3c\xee\xe3\x68\xb9\x63\x78\xa5\xb2\x1b\x14\x2f\x6f\x3d\x76\x62\x7e\xc9\x0c\xbc\x4a\x0c\x48\x0b\x3c\x77\x7c\x14\xe7\x0c\x8a\xa2\xca\x57\x64\xa1\xab\xb0\x6b\x70\x29\x3f\x89\x2a\x56\x18\x10\x4a\xed\x83\x75\x3a\x53\x95\x2c\x64\x18\xde\x98\x22\x5c\xc0\xaa\x2d\x0a\xc9\x52\xf1\x93\xb8\xcd\xc4\x8d\x51\x3e\xb3\xc8\x86\x9f\x11\x8f\xb6\x20\x47\x43\xeb\xd6\x72\x68\xae\x35\x28\x21\x91\x44\xdc\xa7\xbb\x48\x88\x03\x2c\x9f\x6a\xe1\xd6\x72\x72\x50\x3e\x7c\xd5\x12\x59\x3e\xa3\x38\x39\x87\x57\xca\xc4\xd3\x2a\x8c\xf7\x79\xc3\xe4\xe4\xc5\x8b\x6e\x8e\x1c\x9b\x50\x2e\x9f\x3d\x95\x9c\xf7\x03\x52\x8d\x9e\xf3\x2a\x3e\x7e\xf8\xc7\xfb\x5e\x7b\xbc\xb8\x2e\x75\x82\x36\x39\xde\x52\x01\x23\xec\x10\x8e\xd6\x3e\x1e\x42\xe5\x77\x2c\x03\x1f\x54\xc0\xe5\x3c\xad\x22\xa1\x00\xaf\x7f\xe1\x73\x95\x6f\xf5\xcb\x74\x1c\xa4\x9b\x27\x53\x59\xc9\x8c\xe7\x2b\xfe\x93\xbe\x5d\x9e\xce\xe7\x77\x25\xe1\x31\xb3\x26\xf7\xb0\xc2\xb0\x45\x1c\x9c\xd8\x2c\xaa\xd6\x97\x72\xb2\x27\x5f\xf1\x0f\xf6\xf3\x8d\xaa\x96\x8b\x17\x44\xe9\x31\x9c\xe3\x7a\x67\xb2\xd2\x59\xa3\x7f\x89\xa7\xb6\x3f\xd7\x47\x4a\xbb\x65\x09\x74\xa7\x8b\x28\xa9\x76\xc4\x10\x74\xa0\x6f\x9b\x5d\x92\xd4\xa3\x7b\x0d\xad\x44\xea\xfb\x7d\xbb\xae\xa8\xfe\xef\x1b\x63\xa9\x0b\x16\x74\x03\x4e\xf1\x1e\x27\x9b\x17\x7f\xbf\x46\x83\x5e\xb3\x12\x0a\xe5\x03\xad\xe8\xb1\x52\xe5\x5b\xac\x1b\x6b\xab\x07\x45\xfe\x68\x77\x14\x46\x76\x2d\xfa\x39\x4c\xbb\x90\x4f\xa4\xeb\xd8\x9f\x1d\x93\xb3\x10\xf7\xb9\xe6\xd3\x93\x39\xff\x93\xc3\x5d\x94\x67\xf5\x06\x45\x0f\xe4\x08\xe9\xb5\x34\x86\xe4\xd0\x52\x1d\xf7\xe5\x82\x53\xc6\x2b\x49\x9d\x05\x62\xda\x3d\xb1\xc6\xeb\x9c\x4f\xf0\x28\xae\x56\x7e\x41\x67\xe9\xfd\x54\xb6\x46\x1d\x56\x6a\x17\x6e\x0b\x44\xaa\x78\xe6\xf3\x39\xc7\x9c\x1f\x54\xc0\x23\xee\x91\xc8\x9d\x87\x01\xed\xae\xf9\xb9\x51\x55\x8b\xb0\x38\x85\xaf\x60\x31\x9f\xcf\x65\xb9\xb1\x03\x52\x6b\xd3\x06\x76\x63\x26\x42\x34\x78\xa2\xe5\xe2\x54\xc2\x0c\x61\x5b\x8a\xa1\xeb\x12\x1a\xa7\xad\x23\x54\x4c\x61\x99\xbf\xe2\xfe\x35\x4d\x6b\x1d\x54\x76\x7b\x54\xec\x71\x10\x31\x23\x7d\x9a\x06\x2f\x23\x64\x8f\xa2\xd0\x35\x76\x71\x41\x85\x80\x75\x23\x4d\x5b\x4f\x59\xc0\xe0\x16\xb4\xd9\xa0\x49\xd7\x4a\xd4\xe8\x48\x58\x83\x8c\x25\x3f\x8c\xb7\x42\x45\x6d\x98\x9f\x81\xf1\x70\x68\x94\xb1\x31\xfe\x3c\x99\x42\xeb\xe1\xb0\xd6\x99\xeb\x1f\x91\x08\xf8\x61\x55\xe9\xfe\x3b\x0f\x87\xfd\x8f\x9a\x5e\x93\x94\xe8\x47\x09\x87\xa5\x6d\x9d\x67\xc0\x10\x1c\x81\x6d\xec\x82\xd6\xe9\xbc\xe6\x7d\xd4\x37\xac\x0f\xeb\x1a\xde\xf4\x1e\x68\x9b\xad\x3f\x58\x52\xc3\x1d\x39\xd5\xea\x56\x46\x84\xdb\xb4\x1b\xfc\x5a\x5a\x47\xb2\xa2\xb1\x6a\xd9\x75\x87\x27\xb6\x66\xe9\xde\x8c\x27\x1c\x2d\x09\xe9\x07\x12\xfc\xb8\xf2\x1d\x11\x71\xb8\x56\x2e\x67\x9b\xb6\x45\xd7\x5f\x31\x7b\x17\x28\x24\x57\x54\x6a\x67\xac\xf1\x21\x4f\xa4\xff\x8d\x59\xf8\x42\xb4\x89\xd4\x90\xf8\x03\x29\x88\xf3\x5d\x97\x7e\xda\x70\x6b\xf9\x47\xad\x6e\xd9\xf5\x9e\x9d\x3e\x52\xe6\x90\x4b\x5a\xaa\x82\x4b\xce\xa9\x8f\x13\xa6\x5e\x72\x5a\x27\x04\x9e\xce\x32\x80\x92\x20\x5f\x2a\x5f\x1e\x51\xde\x1c\x49\x5a\x12\x7c\x2c\x18\xe4\xa4\xb3\xe2\xd4\x34\xd2\x46\xbf\x4b\x9d\x8e\x99\xad\x31\x38\xb5\x1d\x12\xfa\xe1\xea\x15\x1b\xf5\x2d\x53\x14\x4d\xdc\xcf\x4d\x6c\xa4\x7e\x16\x43\x52\xee\x78\x54\x2e\x2b\xc7\x93\x7a\xae\x23\x3a\xee\xe2\xf9\x26\xf7\x00\x07\x69\x0e\x5b\x50\x6c\xd3\x39\x5c\xeb\xba\xa9\x10\xde\x60\xbe\x46\x07\x57\xce\x06\x9b\xd9\x0a\x0e\xaf\xdf\xf0\x99\xec\x1b\xde\x95\x1e\x39\xdc\xf0\x64\xb8\x1a\x60\x15\x06\xf4\x35\x86\xd2\xe6\xd2\x48\x91\x23\xc9\x72\xf7\x42\xb0\x7f\x8d\x41\xf1\xf5\xad\x11\xdb\xbe\x6a\x06\x5c\xbf\xb1\x6a\x8f\x69\x5f\x35\x71\xfc\xda\xa9\xa6\xf4\xa0\xcd\x51\x8d\xb5\x75\xbb\xc8\x8b\xe6\x53\x7b\xa3\x5a\xa6\x40\x15\x5a\x3e\x7f\xc9\x90\xdc\x49\x30\xf6\xdd\x5c\x3c\x43\xd4\x97\xdc\x30\x23\x8c\x99\x36\xca\xe5\xb8\x72\xce\x07\xdf\x07\x6a\xf8\x1e\xc3\x75\xd5\x7c\x4f\x4c\x5c\xb3\x46\x86\x6b\xbe\xb3\x26\x61\x96\xbf\x93\x95\x7d\x61\x8b\x3f\x87\x6b\xbd\x36\xbc\x4c\xf8\x1b\x3a\xe9\x8b\x91\x15\xbd\x22\x97\x7e\x14\x0f\x93\xe8\x2c\x17\x78\xd2\xd4\x1c\x40\x24\xc3\xd4\x8a\x0d\xeb\x94\x52\x2c\xa0\x09\x4e\x0b\x42\xf5\x7a\x3d\x0a\x33\xa7\x92\xde\x1f\x43\x24\xaf\xac\x36\xf0\x3d\x1a\x94\x43\x3a\x70\xf8\x96\x2d\xf0\xc9\x6f\xaa\xe7\xa8\xf2\xec\x49\xa4\x82\x04\x28\xc5\x56\x8a\x00\x40\xba\xd4\xd2\x9d\xb6\xaf\xa5\x47\xa9\xfa\x2b\x8d\x7c\x5a\xcb\xb6\xeb\x32\xf6\x90\x32\x55\x55\x7c\x24\x6e\x8b\x14\x02\xfc\x78\xb3\xf0\xd5\xd5\x47\xa2\x81\x0e\x0e\x75\x91\xcc\xef\xc9\xe3\x60\xdf\x78\xb5\x64\x7f\x6e\xd9\x15\x1e\xb4\x3b\xe2\x99\x91\xee\xe2\x23\x37\x99\xe2\x81\x63\xae\x2e\xd1\x73\xd3\xaa\x69\x5d\x63\x3d\xf6\x1b\x3b\xb1\x3f\x20\x1b\x8c\x72\x9f\x0b\xbc\x36\x99\x74\xbd\xba\x4b\x24\x44\x91\x9d\x9c\xde\x6b\x0f\x85\x72\x10\xac\x15\xc0\x44\x13\xf4\x8c\x75\x1b\x3b\x5b\xeb\x42\x29\x37\x15\xb6\x65\xec\xac\x45\x55\xe1\xb2\x50\x95\xc7\xc1\x26\x7c\x3c\x1d\x1c\x2c\x34\x6a\xc7\xe2\xcd\x07\xa5\xe5\xde\x0c\xbc\x37\x68\x83\xb4\xf1\xb8\xf8\x8b\x96\xb0\xaf\xfb\x34\x5d\xde\x77\x32\x30\xf0\x47\xe9\x1b\x4e\x07\x77\xee\x74\x74\xe7\x8f\x64\x42\x7a\xb3\x5c\xd0\x4a\x56\xd2\x64\x8d\x9f\x3e\xf8\xc1\xc9\x83\x5f\x3c\xbd\xd3\x99\x8e\x90\x51\x96\x92\x60\x81\x60\xd9\x60\xe5\x5a\xd0\xfe\xde\x17\x69\x7b\x3f\x15\x4a\x62\xe2\xad\x34\x34\x6c\xdb\x05\xa2\xe7\x5d\x50\xd1\x5a\x7c\xda\xa1\xdd\x74\xe2\x34\x1e\x08\xc8\xb9\xf3\x36\x90\xe0\x9e\x6c\x67\x30\xbe\xcf\xf7\x09\xbe\x99\x22\x15\xdd\xa8\x88\x5a\xc4\xbd\x72\x23\xa1\xaa\xee\x27\x0d\x6d\x13\x2f\x70\x0d\x17\xd4\x9a\xa0\xab\xb8\xf1\xa8\x02\x21\x79\xde\xfe\x1d\x1f\x5f\x1f\xf4\x25\x48\x62\x1d\x5a\xac\xb5\xe1\x68\x76\xef\x46\xc0\x43\xe2\x66\x17\x93\x0a\x23\x09\x2a\xee\xe8\xca\x1d\x8b\xd8\x45\xf6\x68\x7c\xeb\xe5\x15\xe8\x22\xb2\x5b\x29\xb7\xee\x0f\x87\xaa\x20\x84\x7a\xe6\x62\xa8\xfd\xfa\xb4\x2b\xa5\x06\x1c\x8e\x79\x8a\xd8\x90\x34\x78\x94\x54\x77\x9c\x6a\x12\xe5\x50\x49\xff\xa5\xbf\x3d\x96\x26\xa1\x35\xdc\xb5\x0f\xae\xc6\x3c\xef\x9a\x33\x93\xa0\x6a\xdb\x9a\xe0\xa7\x20\x27\x57\x9b\x96\xfe\x26\x7f\xf3\xb5\xe0\x2d\x62\xc7\x77\xa7\x1c\x99\x11\xf4\xa1\x33\xa3\xe8\x4b\x2c\x8a\x9f\x5b\xa4\x40\xc2\x2c\xa7\x43\x6d\x2a\x50\xd4\x89\x0c\xcb\x95\xbc\xbb\x05\x9d\x5a\xe3\x54\xae\x58\x38\xa5\xe5\xca\x6b\x67\xa6\x54\xd6\xe3\x6a\x47\xb9\x3e\xee\x34\xd5\x7c\xe1\x8d\x4b\xbb\x1b\xac\x76\xb1\xce\x4c\x66\x2c\xc1\x81\xaf\xf7\x65\x29\x45\xc8\x95\xba\x4e\x2c\xbd\x6a\xb5\x1f\x68\x75\xb5\x8b\xfd\x0d\xd5\xa9\x28\xf5\xf0\x19\x83\x8c\xcd\x80\x96\x65\x10\x73\x39\x39\xc0\xb4\x13\xcf\xc3\x3c\x7a\x3e\x44\x41\x36\x21\x6d\x99\x9d\xf2\x72\xd3\xc4\xbe\xa0\x18\x94\x36\xf1\xce\x6f\xd3\xc6\x9b\xd1\x29\xd9\xb1\xd3\xf1\x34\x7c\x31\x5b\x5c\x67\x06\x97\x21\x85\x05\xb6\x5f\xb9\x85\xcf\x95\x26\x81\x80\x4a\x34\xd0\xb7\x6d\xb6\xca\xc7\x60\xcb\x0e\x47\x5f\xcf\xe0\xb2\x88\x17\x19\x72\x29\xcc\xe4\x34\x3c\xb1\x5d\xb4\x26\x13\x96\x49\x15\xbb\x78\x98\xc3\x53\x4e\xe8\x6e\x58\x90\x8f\xef\xc0\x07\x17\x8f\x2e\x65\xab\xa2\x52\x6b\xbf\x14\x56\x1e\xa7\x11\x8c\xab\x76\xfd\x38\x1d\x5b\xa2\x0c\x95\x5d\xaf\xe5\xfa\xeb\x06\xab\xbe\x33\xcb\x3f\xe3\xc1\xe4\xe0\x54\x86\x53\xc8\xe9\xfb\x29\x6f\x59\x4c\x61\xab\x9c\x99\x02\x3a\x67\xdd\x14\x32\xa7\xf9\xfc\xfe\x7f\x06\x77\x2c\x78\x5b\x22\x9d\x98\xf8\xc6\xb7\x2b\xbf\xf3\x01\xeb\x6f\x97\xdf\x30\xe9\x6f\xa7\xfd\xb3\x93\xfe\xe1\x6c\x36\x93\xd6\x01\x07\x41\x1b\xd9\x8a\x07\xec\x72\xbd\xd1\x79\xab\x2a\xe8\x46\x52\xa4\xfe\xe8\x51\xf6\x4a\x8e\x8e\x98\x43\x1e\xb1\xf4\xdc\xea\x93\x3d\x86\xf1\xe5\xa7\x7e\x2c\xef\x88\x74\x23\x68\x5d\xa9\x72\xe5\x1d\x8a\xb4\x6f\x33\xd8\xa6\xf8\xeb\x87\x0f\x57\xd0\x38\x5b\xe8\x0a\xbb\xdd\xcf\x78\x3f\x24\x3d\xbe\xf7\xd4\x8d\x6c\x89\x8d\xee\xdd\x8f\x8f\x34\x8d\xe9\x0c\xb7\x86\xc8\x12\xe5\x8a\x50\xfa\xef\x3a\xa8\x00\x65\x08\xcd\xd9\xf1\x71\xb7\x95\x76\xf6\x4d\x1c\x4a\xdc\x7f\x7b\xcc\x4b\x3b\x6e\xe8\x19\x58\x8a\x55\xb1\x6f\x1d\xaf\x89\xd3\x87\xcb\xe7\xf3\xe7\x5c\x11\xfc\xdd\xe9\x80\x8c\x42\xd2\xec\x77\x2e\x4f\xa5\xfd\xc3\xac\x69\xd3\xe8\xe3\x50\x37\x6c\xf9\x33\x7a\x32\xf9\xdf\x00\x00\x00\xff\xff\x44\x24\x08\x2d\x0e\x45\x00\x00")

func sampleBchdConfBytes() ([]byte, error) {
	return bindataRead(
		_sampleBchdConf,
		"sample-bchd.conf",
	)
}

func sampleBchdConf() (*asset, error) {
	bytes, err := sampleBchdConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sample-bchd.conf", size: 17678, mode: os.FileMode(436), modTime: time.Unix(1620840435, 0)}
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
	"sample-bchd.conf": sampleBchdConf,
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
	"sample-bchd.conf": {sampleBchdConf, map[string]*bintree{}},
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
