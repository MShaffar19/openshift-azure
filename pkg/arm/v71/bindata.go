// Package arm Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// data/master-startup.sh
// data/node-startup.sh
package arm

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x59\x5b\x73\xdb\xb6\xf2\x7f\x2e\x3f\xc5\x96\x72\x73\x6b\x40\xda\x69\x3b\xe9\xa8\x55\x66\x1c\x47\xc9\x3f\xff\xb8\xb1\x8e\x9c\x4c\xe7\x4c\x9a\xc9\x40\xc4\x92\x42\x04\x02\x2c\x00\xca\x56\x1d\x7d\xf7\x33\x0b\x52\x77\xd9\x4e\x26\xed\x4b\xf3\xe0\x98\xc0\x62\x77\xb1\x97\xdf\xee\xc2\x9d\x6f\xd3\x91\xd4\xe9\x88\xbb\x31\x30\xbc\x8c\xa2\x0e\x3c\x37\x16\x3c\x3a\x2f\x75\xd1\x05\x65\x0a\xe0\x5a\x80\xb0\xa6\x02\xae\x14\x78\xcb\xf3\x5c\x66\xe0\xc7\xdc\xc3\x85\xa9\x95\x00\x6b\x6a\x8f\x30\x95\x1c\xfc\x18\xa1\xe4\xce\xa3\x85\xfe\xe9\xd3\xa8\x03\xc3\xfe\xf9\xd9\xdb\xe1\x49\xff\xc5\xf0\xec\xed\xa0\x17\xcf\x4c\x6d\x99\x45\x67\x6a\x9b\x21\x2b\xac\xa9\xab\x38\xea\xc0\xd9\xf9\x87\xe7\xff\x79\xf6\xba\x17\x9b\x0a\xb5\x1b\xcb\xdc\x27\x07\x1b\x27\x13\xe3\xb8\xc0\x69\x92\x29\x53\x8b\x38\xea\x44\x1d\x90\x95\xe7\x23\x85\x0e\xd8\x4b\x78\xf9\x7a\xf0\xf6\x0d\x30\x07\x07\xf7\x84\x2c\xe0\x7b\x37\x36\xd6\x43\x7c\xd0\xf2\x8d\xe1\x13\x78\x2e\x15\xb0\xa3\xfb\xc0\x3e\xc2\xe9\xd9\x0b\x60\x4c\x99\x82\x55\x16\x73\x79\x09\xf1\xab\xb7\x4f\xfb\x40\xa4\xf0\x6c\x78\x36\xe8\xc6\x5f\xc7\x9f\x78\x44\xd1\xd5\x15\xc8\x1c\x92\x13\xa3\x73\x59\x24\xe7\x98\xd5\x56\xfa\xd9\x80\xfb\x6c\x3c\xe0\xd9\x84\x17\xe8\x60\x3e\x8f\x94\x29\x0a\xb4\xc0\x7c\x6b\x38\xe6\x3c\xb7\xbe\xae\x12\x37\x86\x58\x6a\xe7\xb9\x52\x52\x17\x60\x51\x00\x99\x3c\x13\x1a\xb2\xc0\xb3\xb6\xdc\x4b\xa3\xc1\x68\x38\xb8\x37\x36\xce\x6b\x5e\xe2\xfd\x38\xca\xb8\x87\x27\xe9\x94\xdb\x54\xc9\x51\x3a\xab\xcb\x34\x53\x12\xb5\x67\x19\x5a\x9f\x54\x58\xc2\xaf\xbf\xde\xed\x9f\x3d\xbf\x4b\x2a\x9e\xa0\xf5\xc7\xee\xe9\xcc\xa3\x5b\xea\x4a\x6b\x32\x97\x19\xf7\xe8\x92\x56\xd7\x21\x56\xc6\x49\x6f\xec\x2c\x6c\xc3\x27\x38\xf7\x96\xf4\x9a\xcf\xa3\xfe\xd9\xf3\xeb\x85\x4e\x70\xb6\x2d\x73\x60\xe5\x94\x7b\x7c\x85\xb3\x2f\x94\xfc\x0a\x67\x3b\x82\x3b\xf0\xe6\xec\xd9\x59\x17\x04\x2a\xf4\x18\x22\x30\x37\x4a\x99\x0b\xa2\x71\x98\x05\x13\xf1\x9c\x42\x92\xc2\x37\x53\x35\x59\xd9\x01\xb7\x08\xb6\xd6\x70\x21\xfd\x18\x38\x4c\x4b\x90\x25\x2f\xb0\xf9\x9e\xc8\x6c\x12\xfc\x90\x58\xac\x0c\x8c\xf8\x04\x05\x48\xdd\xdc\x12\x52\xf4\x19\x5d\x31\x6c\xba\x44\xa4\x5b\xe4\x8b\xab\xbe\xb3\x63\x54\xec\x31\x73\x68\xa7\x68\x99\xad\x4a\xf7\x3e\x22\x27\xf5\x86\x28\xe0\xff\xb8\x87\xbe\xf6\x68\x2b\x2b\x1d\xc2\xa9\xd4\xf5\x25\x3c\x86\xf3\x40\x0c\xf7\x86\x83\xdf\xdc\xfd\x68\xc4\x1d\xd6\x56\xf5\xc6\xde\x57\xae\x9b\xa6\x99\xd0\x89\x45\x31\xe6\x3e\xc9\x4c\x99\x66\x46\x7b\xd4\x3e\x15\xd2\xf9\x94\xa4\xa5\x8d\xac\xf4\x71\xfa\xb8\x61\x94\x1e\x10\x0b\x6e\xb3\x71\x6a\x5c\x54\x54\xc5\x04\x67\xbd\x5c\x2a\xec\xa6\x69\xb8\x47\x35\x91\xa9\xad\x4a\x56\x54\x45\x3a\x1c\xfc\xc6\x5e\x0c\x5e\xb0\x57\xfd\xff\xb2\x46\x0a\xb3\xa8\x90\x3b\x8c\x9c\x53\x19\xa7\x00\xea\x85\x53\x76\xec\xca\x34\xe3\x69\x4b\x55\x63\x45\x4e\x0e\x54\xc1\xed\x0d\xe5\x0d\x31\xb8\x22\x25\x85\x6e\x08\x9c\x08\x35\xe5\xa1\xe8\xcd\xd0\x45\xdb\x26\xc5\x4b\x6f\xb9\xfb\x32\xcb\x32\xe8\x87\x53\xff\x84\x89\x1b\x7d\xfe\x8d\x96\x36\x0e\xd9\x0f\xc9\xd1\xd1\x3e\x5b\x9f\x55\xa8\xcf\x09\xb5\xe1\xc4\x68\xcf\xa5\x46\x0b\x03\xc5\x7d\x6e\x6c\x09\x74\xe8\x9f\x89\x66\x4c\x89\xf7\xbf\xc3\xd8\xfc\xaf\xda\x62\x66\x2c\xb6\xa6\x5d\x7e\xef\xd8\xac\x6a\x6b\x47\x52\xca\xcc\x1a\x67\xf2\xc6\x76\xb3\xba\x0c\x68\x94\xae\x4e\xae\x4b\x28\xaa\x22\x1b\x63\x36\xe9\x69\x13\x50\xf3\xb3\xcb\xce\xf1\xf0\x8c\x50\x34\xd4\x2e\xa8\x2b\x41\xc8\x0c\xef\xae\xae\x5a\x14\x76\xff\x6f\xa4\xbe\xa5\xc8\xc5\x0f\x21\x86\xf9\xfc\xfd\x4e\xa1\xca\x8d\x05\xee\x3d\x96\x95\x07\xa9\xe1\xea\x28\x49\x7e\x9a\xff\x02\xc2\x44\x00\xb3\xba\x84\x56\x0d\x60\x33\x60\x7f\xc2\x97\xc9\x0c\x22\xe1\xce\x1d\x18\x59\xe4\x93\x08\xe0\xc6\x0b\xbf\x5b\xa8\x71\x70\xd5\xfe\x36\x7f\xbf\xff\xea\xad\x4e\x4d\xe5\xcd\xb9\x54\x28\xe2\x08\xa8\xd2\xbf\x7b\xb7\x76\x1a\x98\xf2\xf0\x13\xbc\x7f\xff\x0b\x55\x24\x0d\x4e\x21\x56\x70\xf4\x0b\xa0\x72\x08\x78\x29\x3d\x7d\xe4\x32\x12\x46\xe3\x2d\xde\xb0\x58\x9a\xe9\x97\xb5\x00\x64\xbd\x4c\x21\xd7\x54\xf3\x22\x5b\x02\xb3\x39\xdc\xd8\x12\xdc\x14\xaa\x57\x57\xa8\xc5\x7c\x1e\x45\x32\x87\x6f\xa1\xb0\x58\xad\xc8\x85\xc9\x26\x68\x9b\x92\x98\x3b\xcf\x47\xcd\x85\x23\x00\x37\x73\x1e\xcb\xcc\x2b\x70\xde\x54\xd0\x10\xb2\xa0\x55\x5d\x25\x5e\x96\x68\x6f\xa5\xa2\xfc\x97\x19\x5e\x47\xb7\xb6\x5f\x4e\x72\x97\x5c\xe6\x0e\x58\x0e\xa9\xc0\x29\x81\xc8\xa4\x49\x86\x74\xd9\x70\x56\xdc\xfa\xa3\x08\x00\xb3\xb1\x81\xbb\x37\x93\xc1\xce\x1d\x81\xd8\x43\x61\xab\x3f\x6b\xe3\x39\xc0\x21\x1c\xde\x85\x27\x4f\x56\x57\x27\x35\x4c\xad\xfd\xf6\xc9\x08\xc0\xa2\xf3\x86\xf2\x52\x03\x1b\xee\xd9\x6f\xda\x27\xe2\xd4\x2c\xa5\x82\x63\x69\x74\xf2\xd1\x19\xbd\x6a\x9d\x22\x80\x98\xda\x56\x61\xe5\x14\x6d\xdc\x85\xf8\xa3\xa9\xad\xe6\x4a\xc4\x0f\x69\x4f\x48\x47\x29\xcf\x14\x16\x3c\x9b\x31\x8b\x85\x74\xde\xce\xe2\x2e\x78\x5b\x63\xd4\x34\x4c\x9b\xb6\xe4\xd6\xef\x1a\x73\x3f\xc1\x96\xef\x72\x19\x45\xad\x65\xaa\x5a\x29\xca\xcf\x45\x4e\xbe\x2c\x03\x42\xbd\x36\x02\x43\x16\x3e\x09\xa6\xd6\x44\x75\x67\x6f\x14\xa1\xcf\xc4\xbe\x18\x5a\x7a\x75\xdb\x57\x2e\x73\xf2\x28\x55\xb5\x3e\x84\x4f\x9f\x9a\xdb\x5d\xe7\xd6\x35\xd2\x2d\x81\x8d\x43\x05\xe6\xbc\x56\xde\x7d\x96\x43\xe9\xdc\xf5\xee\x0c\xbb\x64\x17\x42\x36\xe1\x02\xaa\xf9\xac\x7a\xf8\xf3\x8f\x3f\xfe\x18\x70\xed\x9b\xca\x1a\x6f\x7a\x07\x57\xc2\xf9\xef\xbe\x7b\xf8\x60\x1e\x7d\x53\x19\xeb\x9b\x85\x4e\xe7\xc1\xc3\x79\xf4\xcd\x6a\xf0\x38\x0e\x83\xd1\xcb\x61\xff\xf7\xe3\xd3\xd3\x0f\xc7\xa7\xa7\x67\xbf\x03\xab\xe0\x20\x30\x01\x56\x92\x77\x3c\x02\x63\xcd\xff\xaf\xfb\xbf\xd3\xe2\x62\x9b\x09\x62\x0d\x07\xe1\x27\xfb\x08\xc7\x27\x27\xfd\xc1\x1b\x60\x17\x2d\xe6\x2c\xe4\x30\xc7\xa7\xd8\x06\x9f\x9b\xb9\x06\x58\xd2\xc5\x2e\x0d\x84\x17\x01\xc1\x28\x12\x08\x85\x34\x79\xf5\x82\xf3\x02\xb5\x0f\xa3\xa1\x46\x7f\x61\xec\x04\x6a\x2f\x95\xf4\x12\x1d\x14\x26\x20\xa5\x37\x60\x79\x86\x84\x55\x42\x12\x4e\x25\x34\x57\xe5\xcb\xc3\xb6\xd6\x0e\x46\x98\x1b\x8b\x20\xb4\x03\xe9\x60\xa2\xcd\x85\x06\x6f\x42\x1f\xdf\x48\x42\x40\x2d\xa0\xae\x9a\xce\x9c\xd0\x75\x06\x2e\x14\x82\xe8\x62\x2c\x15\x06\xe0\x5d\x82\x1f\x30\x71\x1f\x7a\x3d\x88\xe3\x00\xbe\xc2\xac\xa0\xb7\xb9\x76\x73\xe6\x5b\xb8\x39\x74\xcf\x1b\xfc\x85\xf9\xa2\x1e\xb5\x5c\x1a\xdb\x39\xf4\xf0\xfd\x65\x84\x97\xc1\xb6\xe7\xc7\xe7\x6f\x87\x2f\x7b\x77\xd7\xb8\xfc\x16\x50\xbc\x65\xd2\xec\xc3\x7c\x7e\x37\x1c\x64\x97\x8b\xb4\xa1\xf1\x83\xb1\xca\xca\xa9\x54\x58\xa0\x00\xc6\x08\xaa\xd9\xc2\xa0\x74\x27\x60\x53\x48\xbb\x29\xfd\xda\xfd\x0b\x18\xb6\xd2\x6e\x54\x19\xda\xea\x11\xd5\x9a\x04\x36\x27\xa2\xa8\x29\x61\x2c\xe3\xcc\xdb\xda\x79\xf2\xec\x40\x6a\x98\xd4\x23\x6c\x9c\xee\xc8\xf0\xb5\x43\x50\x26\xe3\x0a\x78\x25\xdb\xee\x2f\x72\xa4\x9c\x04\x66\x11\x62\xd7\xb9\x07\x0f\x9a\xf5\x2e\xdc\x4f\x1e\x74\xfe\x38\x5a\xf4\x26\x6b\x35\xa8\x13\x37\xf9\x6c\xac\x2c\xa4\x4e\x9b\xb2\x96\x2e\xe7\x7b\xd6\x2c\x24\x2b\xe1\x5f\x2f\x83\xc2\x25\xfc\xf8\xfb\xb9\x3a\xa1\xbf\x9e\xa9\x35\xc6\xa7\x81\x4d\xda\xf2\x89\xae\xae\x18\x25\x84\x46\x38\x48\x9e\xf2\x6c\x52\x57\x4f\x95\x19\xbd\xa6\x38\x8e\xe3\x5b\x5f\x07\x96\x29\x49\x48\x34\x45\x3b\xdb\xe9\x03\xa2\x0e\x38\x4f\x71\x0b\x05\xfa\x90\x53\xa3\x20\x25\xb4\x04\xc3\x7c\x13\xb9\xd2\x07\x11\x35\x2f\xa4\xc7\x33\x69\x7b\x9b\x7b\xed\xb9\x72\x22\xa4\x85\x83\x35\xba\x5b\x9a\x17\x61\x2e\xb4\x32\x5c\x90\x9a\x0d\x8f\xf8\x33\x53\xb0\xef\x33\xd1\xd8\xe4\x9a\x2c\xdc\xc8\xa2\xdd\xc4\xf9\x23\x82\x90\x3c\x3b\x51\xd8\xdd\x5d\xda\x47\x1c\xde\x9a\x2a\x6b\xa6\x52\xa0\x4d\xbb\xe9\x07\xc1\x3d\x4f\x3f\x98\x7a\xc9\x7a\xdd\x0c\xdd\xd4\xd4\x94\xa1\xb4\x75\xcb\x5d\x80\x0c\xda\xd8\xa2\xe1\xc4\x46\xad\xd3\x7b\x74\x72\x2b\x0e\xe6\xf3\x96\x48\x84\x27\xb9\xd0\xef\xf5\x48\x58\xeb\x92\x44\x8c\x5a\x02\x1e\xde\x3a\x7a\x0b\x8b\xdf\xec\x97\x56\xfe\x82\x98\xba\xd8\x45\xb0\x3c\x5a\x94\xb7\xdb\xda\x52\x22\x0a\xaf\x2c\x9a\x57\x6e\x6c\xfc\xe7\x7a\xb6\x41\x47\xb2\xc9\xd7\x7b\x96\x6c\xd9\x5d\xfe\xb6\xdc\x5a\x8f\xdd\xee\xe6\x57\xe3\x23\x86\xd0\x7f\x73\xf2\xec\xe4\xcd\xe9\x87\xe3\xc1\xcb\x5e\xfc\x43\x7c\x8d\xeb\x36\x94\x0d\x34\xc4\x25\xf4\x46\xed\xb5\x17\xe6\xda\x88\x87\x1d\xef\x50\xf4\x30\x4a\x9e\xcd\xbc\xd2\x78\xd1\x12\x84\xea\xb5\x96\xbd\xed\xb2\xd4\xd2\x4b\xae\x58\xfb\x70\x05\x71\xeb\x8a\xc3\xf0\x6f\x39\x16\x6e\xac\x76\x1f\xfd\xf0\xf3\xe1\xc3\xf5\xa5\xa3\xbd\x84\x47\xbb\x84\x8f\xf6\x12\x3e\x0a\x84\xf1\x7e\x95\x98\x37\x13\xd4\xc1\x2c\x2c\x37\x96\x85\xb6\x6b\x8b\x94\x8b\x29\x5a\x2f\x1d\xb2\x0a\xd1\xb2\xda\x2a\x07\x7b\x60\x32\x88\x89\xa2\x72\xba\x6b\xa5\xf4\xc1\xd6\xda\xce\x58\xb3\xb4\xe7\x06\x3c\x6d\xb4\x6a\x5b\x7c\x3f\x27\xc0\x31\xf4\x0d\x71\x80\x6a\x6a\x44\x68\x0e\xf2\xb5\x46\xc1\xb8\x28\xa1\xb2\x26\xa7\x90\x5f\x15\xb6\xcc\x68\x6f\x8d\x62\x95\xe2\xd4\x6e\x74\xa8\x81\xe1\xca\x19\xd0\x88\x62\x45\x97\x84\x1a\x9b\x4c\x8d\xaa\x4b\x74\x40\x81\x91\x59\xe4\x1e\xc5\xa2\x23\xa2\xe6\xb4\x99\x34\x32\xea\x83\xa8\x59\x6a\xe1\x97\x95\x70\xf8\xf8\xa7\x43\xea\x05\x97\xd7\x69\x41\xeb\x1a\xfe\xa4\x47\x33\x1c\x84\x72\xe1\x66\x4e\x99\x02\x9c\xd4\x59\x68\xb0\x4a\xae\x79\x81\x80\x54\x43\xfc\x98\x48\xfc\xd8\x9a\xba\x18\xc3\x62\xbe\x88\x56\x23\x41\x3b\x64\x2c\xb8\x2c\xc7\x86\xad\x09\x6d\x7b\x9b\xd0\x05\x7d\x28\x57\x75\x05\x05\x6a\x9c\xf2\x30\x94\x07\x04\xf1\x3c\x9b\xac\x71\xa8\x75\xc9\xdd\x04\x4a\xe1\xc4\x82\x01\xf0\xbf\x1c\x66\x5b\x9f\xa5\xd1\xab\x95\x5c\xd5\xa8\xbd\xd8\xa3\x50\xf3\x12\xf2\xb7\xb1\x6b\x86\xa2\xaf\xe3\x46\xf6\xd0\xc6\x63\x17\xb8\x37\xa5\xcc\xd8\x2a\x82\x42\xdb\x9b\x59\xee\xc6\xa0\x8c\xa9\x1c\xd4\xda\x4b\xb5\xf8\xbb\x8a\x74\x50\x57\xbb\x77\xdb\xcb\xe5\x5a\xdd\x6f\xa4\x86\x3b\x9f\xfb\xb7\x0b\x0a\x69\xc7\x28\x4f\xda\x46\x84\x31\x8b\x23\x63\x28\x86\xfc\xc6\xd8\xf7\xe9\x13\x5c\xdd\xf6\x08\xd3\x1c\x25\x36\xeb\x28\xe8\x0d\x64\xa6\xac\xc2\xe3\xfe\xbe\x17\x99\x98\x8a\xc7\xb8\xf6\x54\xc9\x96\xf2\x41\x9b\x8b\x68\xbe\x7a\xb8\xf8\x5f\x00\x00\x00\xff\xff\x05\xf1\x6c\x4a\xe6\x1a\x00\x00")

func masterStartupShBytes() ([]byte, error) {
	return bindataRead(
		_masterStartupSh,
		"master-startup.sh",
	)
}

func masterStartupSh() (*asset, error) {
	bytes, err := masterStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\xdf\x6f\x1c\xb7\x11\x7e\xe7\x5f\xf1\xf9\x64\x58\x31\x1a\xee\x5a\x28\x0c\x01\x72\x64\xc0\x75\x95\x34\x75\x53\x09\x52\x8b\xa2\x10\xf4\xc0\x23\x67\x77\x99\xe3\x92\x6b\x72\xf6\xa4\xf3\xf9\xfe\xf7\x82\xbb\x77\xbe\x8b\x24\x9f\x63\x34\x79\xc9\x9b\xf6\x38\xbf\x39\xdf\x37\x43\x1d\x3c\x29\xa7\xd6\x97\x53\x95\x1a\x48\xba\x13\x62\xb9\x84\xad\x50\xbc\x0d\xbe\xb2\x75\x71\x45\xba\x8f\x96\x17\x17\x8a\x75\x73\xa1\xf4\x4c\xd5\x94\xb0\x5a\x09\x17\xea\x9a\x22\x24\xc3\x07\x43\x32\xb1\x8a\xdc\x77\x45\x6a\x30\xb1\x3e\xb1\x72\xce\xfa\x1a\x91\x0c\x1a\xc5\xd0\xc6\x43\x0f\x16\xfb\xa8\xd8\x06\x8f\xe0\xf1\xf4\x9b\x26\x24\xf6\xaa\xa5\xe7\x13\xa1\x15\xe3\x75\x39\x57\xb1\x74\x76\x5a\x2e\xfa\xb6\xd4\xce\x92\x67\xa9\x29\x72\xd1\x51\x8b\xef\xbe\x3b\x3c\x3b\xff\xfe\x30\x07\xf8\x96\x22\xbf\x49\x7f\x59\x30\xa5\x4f\x91\xe6\xdf\x6c\x65\xb5\x62\x4a\xc5\x3a\xd2\x4b\xea\x42\xb2\x1c\xe2\x62\x38\xc6\x47\x5c\x71\xcc\x71\xad\x56\xe2\xec\xfc\xfb\xcf\x3b\x9d\xd1\xe2\xbe\xcf\x8b\x68\xe7\x8a\xe9\x1d\x2d\xbe\xd2\xf3\x3b\x5a\x3c\x70\x7c\x80\x7f\x9d\xff\xf5\xfc\x04\x86\x1c\x31\x81\x1b\x42\x15\x9c\x0b\xb7\x59\x26\x91\x1e\x4a\xa4\x2a\xa6\x08\xe5\x1c\xb4\xeb\x13\x53\x4c\x50\x91\x10\x7b\x8f\x5b\xcb\x0d\x14\xe6\x2d\x6c\xab\x6a\x1a\xbf\x67\x56\xcf\x86\x7b\x28\x22\x75\x01\x53\x35\x23\x03\xeb\xc7\x2c\x51\x12\xeb\x9c\xe2\x70\x98\x0a\x53\xde\x13\xdf\xa4\x7a\x1d\x1b\x72\xf2\x58\x26\x8a\x73\x8a\x32\x76\x6d\xba\x11\xf9\x92\x4e\x2f\xc9\xe0\x6f\x8a\x71\xe6\x99\x62\x17\x6d\x22\xfc\xc3\xfa\xfe\x0e\xc7\xb8\x1a\x84\xf1\xcd\xe5\xc5\x4f\xe9\xb9\x98\xaa\x44\x7d\x74\xa7\x0d\x73\x97\x4e\xca\x52\x1b\x5f\x44\x32\x8d\xe2\x42\x87\xb6\xd4\xc1\x33\x79\x2e\x8d\x4d\x5c\x66\x6f\xe5\xe8\xab\x3c\x2e\x8f\x47\x43\xe5\xd3\x6c\x42\x45\xdd\x94\x21\x89\xba\xab\x67\xb4\x38\xad\xac\xa3\x93\xb2\x1c\xf2\xe8\x66\xb6\x8c\x5d\x2b\xeb\xae\x2e\x2f\x2f\x7e\x92\x3f\x5c\xfc\x20\xdf\x9d\xfd\x57\x8e\x5e\x64\x24\x47\x2a\x91\x48\xc9\x69\x95\x1b\xe8\x74\xd0\x8a\x4d\x6a\x4b\xad\xca\xb5\x54\x4f\x5d\xbe\xe4\x41\x6a\xb8\xf6\x51\x72\x4f\x0f\x6e\x45\x73\x40\x7b\x1a\x47\x90\x57\x53\x47\xe6\x74\x41\x49\xdc\x2f\x29\xdd\x71\x54\xe9\xeb\x2a\x2b\x71\x36\x68\xfd\x1e\x25\x1e\xe3\xf9\x23\x56\x3a\x24\x92\x7f\x2e\x8e\x8e\x1e\xab\xf5\x79\x47\xfe\xaa\xb1\x15\xe3\x6d\xf0\xac\xac\xa7\x88\x0b\xa7\xb8\x0a\xb1\x45\x56\xfa\x7d\xba\x99\xca\x6c\xfb\x8f\x51\x6c\xf5\xa1\x8f\xa4\x43\xa4\x75\x69\x3f\x7d\x3f\xa8\x59\xb7\x9e\x1c\x45\x6b\x75\x0c\x29\x54\x63\xed\x16\x7d\x3b\xb0\x51\xb9\xd5\xdc\xf5\x50\x77\xb5\x6e\x48\xcf\x4e\x7d\x18\x58\xf3\x57\x0e\x9d\x37\x97\xe7\x99\x43\x87\xb9\x85\xbe\x33\x99\x97\x71\xbd\x5c\xae\x39\x38\xfd\x3d\x58\xff\x85\x01\x37\xf9\x16\x13\xac\x56\x37\x0f\xc6\x54\x15\x22\x14\x33\xb5\x1d\xc3\x7a\x2c\x8f\x8a\xe2\xe5\xea\x15\x4c\x10\xc0\xa2\x6f\xb1\x0e\x03\x72\x01\xf9\x1e\x5f\xe7\x73\x70\x89\x67\xcf\x30\x8d\xa4\x66\x02\xd8\x93\xee\xf5\x26\x88\xa7\xcb\xf5\x5f\xab\x9b\xc7\x13\x5f\x47\x34\x4e\xdd\x4a\x59\x47\x66\x22\x90\x67\xfc\xf5\xf5\x8e\x36\xa4\x63\xbc\xc4\xcd\xcd\xab\x3c\x8d\x3c\x92\x23\xea\x70\xf4\x0a\xe4\x12\x81\xee\x2c\xe7\x8f\xca\x0a\x13\x3c\xed\xbd\x89\x48\x6d\x98\x7f\xdd\xf0\xcf\x95\xd3\x8e\x94\xcf\xd3\x4e\xc4\x16\x32\x56\xd8\xbb\x0c\xec\x6b\xd2\xe5\x92\xbc\x59\xad\x84\xb0\x15\x9e\xa0\x8e\xd4\x6d\xc5\x4d\xd0\x33\x8a\xe3\x30\xac\x12\xab\xe9\x98\xae\x00\xd2\x22\x31\xb5\x9a\x1d\x12\x87\x0e\xa3\xa0\x1c\xa2\xea\xbb\x82\x6d\x4b\xf1\x8b\x52\x19\xf9\x56\xd3\xe7\xe4\x76\xce\xdb\x59\x95\x8a\xbb\x2a\x41\x56\x28\x0d\xcd\x33\x7d\xcc\x46\x18\x94\x91\x52\xe8\xa3\x26\xd9\xa9\xc8\x47\x02\x20\xdd\x04\x1c\xee\x17\xc3\x83\x1c\x91\xcd\xa3\x8e\xdd\xfb\x3e\xb0\x02\x5e\xe0\xc5\x21\x5e\xbf\xde\xa6\x9e\xc3\x08\xbd\xe7\xfb\x9a\x02\x88\x94\x38\x64\x44\x7a\xc8\xcb\x07\xe7\xcb\xa5\xcc\xfd\x43\xef\x51\x5c\x06\x47\x19\x7c\x55\x54\xb9\x7b\x05\x30\xee\x54\xd9\xc9\x28\x5d\x1a\x45\x6d\xf0\xc5\xcf\x29\xf8\xed\x3e\x25\x80\x89\x0b\xb5\x34\xd1\xce\x29\x4e\x4e\x30\xf9\x39\xf4\xd1\x2b\x67\x26\xdf\xe6\x33\x63\x53\xe6\x01\xe9\xa8\x56\x7a\x21\x23\xd5\x36\x71\x5c\x4c\x4e\xc0\xb1\x27\x31\x6e\x51\x39\x0e\xf2\x66\xf4\xbb\x5b\x71\x15\xf9\x61\xc9\x1f\x17\xb8\x77\xc3\x95\x15\x62\x5d\xbf\xae\x77\x2e\x23\x78\x83\xda\x1f\xdb\x81\xc1\xfe\x19\x0c\x0d\x38\x7d\x3d\x5c\x88\xcf\x52\xcf\x84\x38\xc0\xed\x00\x9b\x6c\x3a\x37\x7f\x86\x05\x6e\x95\xaa\xc9\x33\x94\x37\xf0\xc4\xb7\x21\xce\xd0\xb3\x75\x96\x2d\x25\xd4\x61\x80\x27\x07\x44\xa5\x29\x43\xc4\xd8\x0c\x8f\x42\x1c\xe4\xf2\x6e\x94\x63\xef\x13\xa6\x54\x85\x48\x30\x3e\xc1\x26\xcc\x7c\xb8\xf5\xe0\x30\x2c\x8e\xa3\x27\x1a\x2a\xd1\x77\xe3\x2a\x98\x21\xbd\x40\x1a\xb8\x47\xdc\x36\xd6\xd1\x80\xf6\x4f\x98\x83\x34\xcf\x71\x7a\x8a\xc9\x64\x40\xbc\x09\x5b\xbc\x8f\xf8\x1e\x75\x9e\x60\x7f\x2d\xae\x46\xd8\x63\xb5\xa1\xc0\xb5\x95\x91\x24\x12\x31\xfe\x74\x27\xe8\xae\x0b\x91\x71\xf5\xe6\xea\xdf\x97\x3f\x9e\x1e\xee\x58\xf9\x4f\x88\x33\x8a\x6b\x23\xe3\x39\x56\xab\xc3\x41\x51\xde\x6d\xee\x21\xef\xbb\x52\x76\xd1\xce\xad\xa3\x9a\x0c\xa4\xcc\x0c\x21\x37\x05\xcd\x39\x41\xce\x51\x9e\x94\xf9\xcf\x93\x0f\x90\xb4\xf6\xb6\x37\x64\xac\x49\x4b\xf4\x3e\x3b\x1c\x35\x84\x18\x79\x53\x6a\x25\x39\xf6\x89\xc5\xbe\x6e\xe7\xde\x93\x91\xca\xb4\xe8\x62\xc8\x83\x1c\xa1\x23\x9f\xf2\x56\x21\xf3\x5e\x10\x83\x93\x9d\x53\x9e\xc6\x56\xcd\x4c\xfa\x05\xad\x7c\x99\xbb\x7d\x9d\xfb\x8a\xa0\x5c\x0a\xf0\x44\x66\x2b\x59\xb8\xa0\x95\x2b\xe6\xc1\xf5\x2d\x25\x18\x1b\xa1\x23\x29\x26\xb3\xe9\x95\x0c\xff\x11\xfa\x3a\x77\x48\x6e\x23\xd1\xce\xb2\xa0\x6c\xf1\xe2\xf8\xe5\x0b\xc8\x1d\x62\x0c\xd1\xd6\xd6\x97\x9f\xb1\x9f\xe3\x18\x21\x39\x10\x7b\x5a\x24\x17\x6a\x24\xeb\xf5\xd0\x7a\xad\xf2\xf9\x15\x42\x73\x8a\x0b\x6e\xb2\x08\x37\x31\xf4\x75\x83\x0d\xaa\xc5\x16\x7d\x6b\x68\x6f\xac\x7c\x42\xe8\x3d\xca\xbc\x7f\x2c\x0e\x90\x88\x07\x6c\xf5\x1d\x6a\xf2\x34\x57\xc3\x84\x1c\x9e\x4c\xac\xf4\x6c\xc7\x42\xef\x5b\x95\x66\x68\x4d\x32\x1b\x03\x50\x1f\x12\xe9\x7b\x9f\x6d\xf0\xdb\x5f\x2a\xd7\x93\x67\xf3\x48\x40\xe3\x52\xf2\x9b\x99\x1b\xf9\xe7\xff\xb3\x26\x0e\xe0\x03\xd3\x09\x14\x87\xd6\x6a\xf9\xcb\x16\x82\x8e\xf9\x41\xef\x42\xe8\x12\x7a\xcf\xd6\xa1\x55\xf9\x09\x99\xd9\xa3\xef\x1e\xa6\xf6\xa8\x95\xcf\x86\xbe\x57\x1a\xcf\x7e\xed\xff\x10\x72\x47\x27\x99\x67\xcd\x9a\x34\xa5\x8c\x34\x0d\x21\xb7\x10\xff\x82\x60\x3f\x7e\xc4\x72\xff\x42\x34\x2a\x66\x23\x3b\x9b\x45\xa6\x48\x1d\xda\x6e\x78\x64\x3f\xb6\x1d\xe5\x4d\x28\x35\x3d\x9b\xcc\xa7\x1b\xef\xf0\xe1\x56\xac\xb6\x6b\xc4\xff\x02\x00\x00\xff\xff\x14\x88\xe8\x2d\x23\x11\x00\x00")

func nodeStartupShBytes() ([]byte, error) {
	return bindataRead(
		_nodeStartupSh,
		"node-startup.sh",
	)
}

func nodeStartupSh() (*asset, error) {
	bytes, err := nodeStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"master-startup.sh": masterStartupSh,
	"node-startup.sh":   nodeStartupSh,
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
	"master-startup.sh": {masterStartupSh, map[string]*bintree{}},
	"node-startup.sh":   {nodeStartupSh, map[string]*bintree{}},
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