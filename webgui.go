package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _static_gui_index_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x7d\x73\xdb\x36\xd2\xff\xdf\x9f\x02\x61\x9e\x96\xd4\x48\x24\xe5\xb7\x99\x56\x96\x94\x27\x75\xdd\x36\xd7\x5e\x93\xc6\x4e\xe7\x3a\x3e\x5d\x06\x22\x21\x11\x36\x45\xb0\x20\xa8\x97\x3a\xfe\xee\xb7\x0b\x82\x14\x49\xc9\xb1\x73\xd3\xeb\x74\x6e\x9a\x99\x58\xe4\x62\x77\xf1\xc3\xee\x62\xb1\x58\x69\xf8\xec\xeb\xd7\xe7\x57\xbf\xbc\xb9\x20\x91\x5a\xc4\xe3\x83\x21\x7e\x90\x98\x26\xf3\x91\xc5\x12\x0b\x09\x8c\x86\xe3\x03\x42\x86\x8a\xab\x98\x8d\x7f\xca\x69\x70\xcb\xe4\x0d\xfc\x1d\x90\x5f\x44\x7e\x95\x4f\x19\x39\x17\x8b\x05\x4b\x14\x79\x99\xd0\x78\xf3\x1b\x93\x43\xbf\x60\x3e\x40\xb9\x05\x53\x94\x04\x11\x95\x19\x53\x23\x2b\x57\x33\xf7\x0b\x6b\x5c\x0d\x44\x4a\xa5\x2e\xfb\x35\xe7\xcb\x91\xf5\x0f\xf7\xdd\x4b\x17\x74\xa5\x54\xf1\x69\xcc\x2c\x12\x88\x44\x81\xde\x91\xf5\xea\x62\xc4\xc2\x39\xab\xc9\x25\x74\xc1\x46\xd6\x92\xb3\x55\x2a\xa4\xaa\xb1\xae\x78\xa8\xa2\x51\xc8\x96\x3c\x60\xae\x7e\xe9\x11\x9e\x70\xc5\x69\xec\x66\x01\x8d\xd9\xe8\xd0\x2a\x80\xc5\x3c\xb9\x25\x92\xc5\x23\x2b\x53\x9b\x98\x65\x11\x63\xa0\x28\x92\x6c\x36\xb2\x10\x56\x36\xf0\xfd\x05\x5d\x07\x61\xe2\x4d\x85\x50\x99\x92\x34\xc5\x97\x40\x2c\xfc\x8a\xe0\x1f\x7b\xc7\xde\x91\x1f\x64\xd9\x96\xe6\x2d\x38\x70\x65\x59\x01\x57\x2b\x27\x6a\x93\x02\x5e\xc5\xd6\xca\x2f\x47\x08\xf1\xa6\x7c\x9e\xe4\x8b\x8c\xdc\x91\x05\x95\x73\x9e\xb8\x53\xa1\x94\x58\x0c\xc8\x51\x3f\x5d\x9f\x91\x19\x2c\xca\xcd\xf8\x6f\x6c\x40\x4e\x34\x01\xe5\x5d\x1a\x83\xd4\x80\x04\xb0\x5c\x26\xcf\xc8\x7d\x53\x55\x96\xd2\xc4\x0b\x59\x16\x80\xd2\x9a\xfc\xe1\x09\xca\x87\x3c\x4b\x63\xba\x19\x90\x69\x2c\x82\xdb\x52\xf6\x79\x2c\x68\xc8\x93\x39\x48\x54\x0c\x89\x48\xd8\x19\xd1\xf6\x1b\x90\x2f\x4e\x3f\x3b\x33\x08\x07\xe4\x18\x90\x10\x9a\x2b\xd1\xc0\x77\xfc\x20\xbe\x40\xc4\x42\x0e\xc8\xf3\x2f\xbf\xfc\xb2\x9a\x50\x32\xf4\xda\xee\x7c\xcd\x61\x2f\x95\x62\x2e\x59\x56\xb3\x8f\x12\x69\x39\x97\xe1\xc5\xf8\x64\x12\x38\xa6\x10\x92\x73\x29\xf2\x24\x2c\x97\xdc\xef\x03\xec\x1a\x39\x15\x19\xc4\x81\xa8\x61\x4b\x69\x88\x0b\x47\x5e\x54\x39\x15\x12\x74\x19\x27\xb8\x92\xcf\x23\xe5\x4a\x30\x4d\x9e\xed\xe7\x88\xd9\xac\xcd\x60\xbc\x01\x71\x20\x6e\x19\xa0\x2a\x57\x3f\x9b\xcd\x8c\x79\xb2\x88\x86\x62\x05\xa4\x7e\xbf\x4f\xd0\x96\xf8\xff\x14\x65\xdd\x15\x9b\xde\x72\xe5\x16\x66\x5d\x40\x38\x45\x1a\x1c\x4d\x30\x7a\x39\xcd\x58\x58\x4d\x10\x98\x2d\x77\x67\x20\xc1\xfc\xa0\x26\x13\x31\x0f\xc9\xf3\x20\x08\x2a\xa8\x4d\x78\xad\x05\x97\x3e\x2d\x61\x1c\x9e\x16\x0f\x3b\xd3\x44\x27\x95\x0f\x6a\xdc\xfd\x92\xbb\x52\xbb\x3b\xd4\xb0\xd8\x2e\xca\xd6\x3c\x69\x19\xb5\x33\xba\xe0\x31\xc4\xc5\x42\x24\x02\x42\x3a\x30\xc1\x31\xf4\xf5\x76\x82\xbc\xe4\x17\x89\x69\x38\x15\xe1\x46\xef\xb3\x90\x2f\x49\x10\xd3\x2c\x1b\x59\x98\x0c\x28\x4f\x98\x34\xfb\x6c\xf8\xcc\x75\xc9\xa5\x82\xb4\x12\x40\xe2\x58\x4e\xa9\x24\xae\x6b\x86\xe0\xbd\x14\x33\x43\xc5\x87\x1b\xb2\x19\xcd\x63\x65\x54\x3c\x30\x81\x3b\x8b\x73\x1e\x56\x3c\x4d\x2e\xa3\xa8\x88\xd0\x1a\x0f\x70\x4d\x73\x30\x47\x62\xb2\x42\xf1\x62\xb5\xc4\x94\x98\xcf\x21\x71\x40\x00\xc5\x34\x05\xd7\x5b\x24\xa4\x8a\x1a\x32\x42\x28\xe8\x25\x19\x7c\x83\xf9\xf5\x79\x21\x6d\x11\x2a\x39\x75\xd9\x1a\xb2\x41\xc8\xc2\x91\x35\xa3\x31\xf2\x6a\x2a\xa2\x97\x22\xae\xa6\x6a\x40\xc3\x8c\x05\x42\x25\x98\x4c\xba\x22\x89\x37\xd6\xf8\xaa\x80\x03\x12\x7c\x4e\x71\x17\x81\x2b\x80\xef\x23\xa2\x1c\xe6\x71\xb5\xfa\x3f\x8a\x75\xe8\x17\xa6\x6c\xd0\x68\xcb\xae\x53\x09\x26\x29\xf3\xfc\x73\xab\x7e\xa2\x55\x07\xd9\x5b\x9d\x80\x86\x3e\xad\x39\xd6\x07\xcf\xb6\xfc\xcc\xc3\xca\x84\xad\x49\x4a\xef\x54\xee\x6b\xba\x7f\x26\xe4\xa2\x25\xa1\x49\xe6\x19\xf3\x8a\x45\xc0\x45\xe0\xe7\x8c\x51\x19\x44\x16\x81\x53\x2f\x12\x30\x1f\x78\xb9\xed\xaf\x5a\xc8\xa1\x16\x17\xd3\x5d\xda\x62\x02\x36\x9e\xa4\xb9\xaa\x1d\x44\x56\x43\xc8\x04\x85\x45\x20\x1f\x07\x2c\x12\x31\xc4\xec\xc8\x7a\x23\x32\x45\xde\xbd\xfd\xc1\xaa\x0e\xdc\xb0\x3d\x7b\xd3\x2e\xbb\xd1\x9d\xe5\xd3\x05\xdf\x4e\x36\x55\x09\x81\xff\xdb\xed\xf5\x36\x4f\xf6\xfa\xcd\x47\x58\x0d\x4a\x1e\xd7\x6c\x56\xda\x0a\x3e\xda\x88\x60\xb7\x37\x09\x31\x2f\x05\x69\xa0\xf8\x12\x9c\x01\x51\x51\x05\xc0\x77\x62\xc1\xd0\xd5\x43\x3f\xe6\xe3\xb6\x60\x83\xf3\x4a\xc2\xa6\x87\x3c\xb7\x9f\xbb\x4c\x29\x25\xfe\x3c\x7e\x12\xfa\xf2\x51\x1f\x37\xed\xa5\xd4\x01\x60\x4d\x02\x25\xc9\x9c\xab\x28\x9f\xea\x2a\x64\xc1\x6f\xd9\x2c\xde\x24\x89\xff\xeb\x36\x88\xad\xf1\xa5\xc8\x65\xc0\x08\x78\xe0\x5b\xae\xbe\xcb\xa7\xfb\xe0\x36\xe1\x15\x4e\x44\xc3\xf9\x1e\xc0\xd9\x86\xef\x76\x4d\x75\x96\x56\xfa\xdb\x26\x53\x1f\x84\x75\x71\x55\xdb\x20\xa6\xb8\xa8\x56\xf6\x96\x15\xc5\x86\x49\xf9\x99\xe7\x79\xc4\xf9\x05\x56\x92\xf5\xe0\x64\x58\x00\xee\x19\x51\x11\x83\xc9\xa9\x84\xfd\x23\xe9\x6f\x9b\x67\x9d\x3d\x39\xb8\x2c\x0e\x9a\x3b\x6b\x0f\x03\x66\x0b\x52\x7f\x71\xe1\x74\xe6\x29\x0b\x89\x89\x06\xb3\xd3\x4a\x96\x6d\xfa\x5c\xd2\x38\x67\x89\x58\x8d\x2c\xa8\x25\xea\x34\x28\xf0\x46\x56\x93\x42\xd7\x86\x4b\x1f\x50\xa6\x0e\x2d\x8a\x90\x26\xc2\x56\x1e\xd9\xbe\x9a\xc7\xa6\xf5\x8a\x52\xa8\x79\x08\x21\xdd\x1c\x2a\x55\x9a\xd6\xf5\x46\xfd\x1c\x8a\x0e\x35\x1f\x6c\x58\x26\xde\xeb\x72\x1c\x33\x66\x74\x58\x67\x39\x69\x00\xd3\x59\x16\x65\xa0\x58\x4f\x12\x16\x6f\xa5\x74\x9e\xc5\x78\xda\xf2\x24\x4c\xad\x84\xbc\x6d\xf1\xd4\x96\x85\xca\x6b\x53\x49\xe2\xd7\xdf\x6b\x6e\x92\x62\x45\x4c\xf1\xfa\xa0\x2b\x21\x1e\xdd\x75\xe6\x9e\x10\xf3\x20\x66\x33\xb8\x4d\xb8\x27\x7b\x4f\x2e\x84\xa7\x84\xa2\xf1\xfb\x32\xc4\xac\x71\xff\xd1\x73\x05\x4b\x66\x3c\xe5\x40\xae\x3c\x06\xb2\x7d\x27\x4c\x2b\xdb\xed\x43\xf9\x20\xaa\x12\xcf\xfb\x94\xc9\xf7\x21\xdd\x3c\x1d\x57\x89\x88\xbc\x81\x5a\xf7\x6b\xba\x79\x14\x59\xe3\xb5\x1e\x59\x8f\x6e\xa2\x27\x6e\xa1\x3c\x08\x50\xee\x09\x7b\xe7\x3f\xdc\x39\x50\xbc\x37\x0c\x59\x99\xc0\xdc\x32\x43\x64\xf9\xd4\xf5\x42\xb0\x3d\xb0\x54\x74\xde\x22\x74\x8f\x9b\x31\x18\x9d\x42\x48\xa4\xe4\x7b\xb6\x81\x68\x0f\x21\x20\x80\x50\x1f\x57\x14\x2e\xaa\xda\xb7\xb7\x86\xa5\xda\x92\xc5\x90\xfe\x5b\x26\x9c\x56\x22\xd0\x63\x3b\x2e\xfb\x18\xba\xd3\x1d\x74\x65\xc5\x72\x19\x08\xb0\xff\x0e\xbe\x80\x26\x4b\x9a\x69\x80\x8b\xcd\x39\x5c\xc2\xe1\x20\x2e\x6e\xc8\x96\xce\x4c\x24\x62\x78\xec\x8c\xac\x63\x70\x00\xec\xe2\x82\xff\x93\x30\x9d\xb4\x31\xe9\xd9\x32\xba\x48\xe1\x82\x3d\x7f\x5f\x56\xbf\x17\x6b\xa4\xb0\xda\xde\x6a\x41\x2d\x33\x5b\x29\x89\x70\x1e\x0d\xe9\xc6\xe3\x16\xec\x30\x0b\xc0\xe0\x8a\x64\x32\xd8\xde\xe5\xe9\x0d\x5d\x7b\x73\x21\xa0\x8c\xa5\x29\xcf\xf4\x09\x8a\x34\x38\x19\xa7\x99\x7f\xf3\x6b\xce\xe4\xc6\x3f\xf2\x0e\xbd\x63\xf3\xa2\xef\xf1\x37\x99\xce\x6e\x5a\xdf\xf8\x21\xd5\x4f\x6d\x13\xdc\xb4\xbb\x04\x4f\xd2\x0e\xda\x6e\x00\x6e\x2c\xf2\x70\x16\xc3\xa1\xd8\x42\xae\xdd\x0a\x9a\xfc\x43\xaf\xef\x1d\x9a\xd7\x5d\xe5\x35\xed\xb5\x66\xc4\x0d\x05\x7f\x6b\xaa\x76\xe3\x2c\x4f\x02\x2c\xee\x09\x94\x99\x50\xf6\xbd\xa1\x12\xca\x3e\xb8\x2a\x3b\x58\xfe\x75\xc8\x9d\x36\xb8\x64\x2a\x97\x09\x09\x59\x20\x42\xf6\xee\xed\x2b\x6c\xd9\xc0\xf5\x3d\x51\x8e\x93\xb0\x15\x1c\xf0\xf3\x8b\x75\xea\xd8\xd7\x2f\x3e\x7c\x3e\xb1\x49\x57\x97\x8e\xf0\x61\x8f\xf0\xc5\x76\xae\xff\xf5\xf9\xd9\xa4\xfb\xa2\xe3\x7c\xfe\xe1\xf9\x87\xb3\x0f\xff\xd7\xb1\x3b\x1e\x5b\xb3\xc0\x89\x45\xa0\x6f\x16\x5e\x51\xf3\x76\x3e\x7c\xb8\xee\x59\xd6\xa4\x73\x7d\x38\xf1\xe0\x1c\xc4\xba\xd4\xf1\xff\xd9\xf5\xe7\x3d\x62\x7f\x76\xd4\xb7\x3b\xc0\x91\xe4\x71\x0c\xa0\xee\x0f\xea\xe0\xc3\xa5\x83\x97\xa2\x1e\x31\x45\xe6\xd7\xf0\xd2\x02\xaf\x19\xc8\x0b\x7d\x79\x22\x83\x06\xe3\xd9\x8e\x3e\x25\xe1\x89\x2a\xe6\x80\xe3\x7a\x70\x0d\x5e\x97\xca\xf8\x0c\x49\x5e\xcc\x92\xb9\x8a\xc8\xb8\x3e\x44\x20\x9d\x49\x32\xc2\xbf\x1e\x94\xc0\x98\x01\x92\xb9\xd3\x37\xe2\x5d\x62\x41\xdd\x63\x9d\x69\xd6\xfb\x83\x3a\x32\xe0\xdc\x45\x90\x67\x4c\xfe\xc0\x93\x5b\x07\x1f\xd0\x9e\x3d\xed\xc4\x1a\x0e\x7c\x25\xa3\x11\xb1\x4d\x53\xee\x67\x3c\xf8\xed\x2d\x1a\xa3\xbd\xac\x22\x37\x22\x57\xc0\xa5\x43\x09\x75\xfa\x16\x60\x2a\x95\x1b\x58\x84\xc1\x9d\xb1\xa1\xfb\x55\x92\x29\x3a\x87\x98\x78\xc3\x83\x87\x75\xf3\x92\x4b\x6b\xdf\xa3\xb8\xb1\x5e\x28\xac\x77\xd7\x1b\x4a\xba\x2a\x2e\x61\xce\x4d\x26\x92\x72\xaa\x9b\x9f\x70\x5f\x3a\xba\x82\xd6\xd7\x19\xab\xe3\xc1\x21\xa2\x79\x3c\x88\xd7\xce\x59\x83\xcd\x34\x92\x80\x29\x8b\xc4\xca\x31\xa3\x7c\x46\x0a\x81\xab\x47\x4c\x56\xa9\xa9\x17\x51\x1d\x0f\xf7\x8d\x53\x85\xc4\x9e\x5d\xa0\x95\xff\x9d\x29\x8a\xc1\xe5\x5d\xa1\x58\xa7\x47\x4e\xfb\x1d\x03\xa0\xa6\xb9\x59\x6a\x19\xdd\x8f\xaa\x3c\x2f\xa4\x0a\xcd\xbb\x4a\x9b\xb5\x99\x51\x6a\x99\x45\x5a\xbb\xfc\xad\x62\xe9\xa9\x28\x74\xad\x54\xa6\xf3\x3d\x30\x4c\xe6\xef\x60\x2b\xd4\xb9\xab\xd2\xb8\x55\xeb\xc8\xf1\x05\x9d\x33\x6b\x40\xac\x5c\xc6\x8e\x6d\x75\x1f\x9f\x33\xca\x17\xd3\x84\xf2\xb8\xd3\xb5\xec\x8e\x65\x94\xde\x77\x5a\x11\xdb\x76\xf0\x03\x71\xfb\xfb\x38\xf8\x9c\xa6\x18\xb3\xbf\xaf\x8b\xdf\xc1\x8e\xf9\x11\x73\xee\x53\xdd\x5b\x2d\xf1\x8f\x73\x70\x61\x35\xec\xc1\x7f\xa2\xce\x1f\xf0\xb6\xf7\xe7\x8e\x98\x6f\xe0\xb0\x81\xa3\xfa\x16\x1b\x20\xff\xa3\x11\x53\x2e\xf1\xaf\x80\xf9\x3d\x02\xe6\x67\x9e\xfc\x97\x0e\x90\x3f\x47\xb4\xe0\xfa\xfe\x8a\x94\x4f\x89\x94\x83\x66\x2d\xb2\x73\xf9\x37\xeb\xd2\x3a\x8d\x29\x5e\x2e\xe7\x70\xc1\x87\xfb\xbd\xa7\xc4\x37\x7c\xcd\x42\xe7\xa8\xd3\xaa\x69\xbc\xfa\x35\x1c\x74\x50\xa5\xa4\x63\x37\xae\xdc\x76\x8f\xd4\x95\x9e\x8b\x25\x93\xb0\x34\xd0\x8c\xdf\x7b\x75\x5a\x7d\xe1\xc6\xbf\xc2\x3c\xb6\xbe\x22\xda\x83\x8f\xa9\xe9\xda\x9f\xd9\xf7\x1f\xd7\xa5\x97\x67\xef\xb9\xb7\xdb\xdd\x47\x14\xc3\x9a\xb5\x66\xdf\x27\x97\xf8\x4d\xa0\x8a\x18\x31\x17\x6c\x12\xf3\x4c\xe9\xc1\x0c\x46\xf4\x1d\x7b\x44\xae\x27\x9a\x32\x13\x92\x38\x4b\x2a\xc9\x1a\xaa\xc3\x02\x7c\x79\x73\xaf\x95\xe5\x46\xcc\x4b\xf3\x2c\x72\xae\xd7\xbd\x26\xe3\xf5\x7a\x32\x69\x7a\xb0\x12\xc0\x07\xa7\x2c\x51\x1d\xb8\x5e\x4c\x41\xab\xa9\x62\xa7\x70\x49\x21\x2e\xa1\xf0\x01\x56\xd1\x72\x1b\x00\xd6\xdf\x83\xab\xd4\xb7\x85\xb4\xe9\x76\x77\x02\xb6\xea\x27\x80\x8b\xd3\x94\x25\xa1\x63\x0f\x95\x1c\x0f\x55\x38\xb6\xbb\x9b\xae\xed\x0d\x7d\x78\x2c\x5e\x4b\x8d\x80\xfd\xba\x3f\xe9\xda\xc4\x69\xd2\x0e\x81\xd6\x29\xf8\x7d\xd0\x61\x37\xd7\xb7\xcd\x65\x22\xad\xb6\xe4\x6e\x0e\x6b\xdf\xe8\xcb\xc4\x80\xfd\x91\xf3\x6a\xff\x97\x52\x2d\x7d\x1e\x98\xe0\x82\x06\xd1\xd6\x7c\xcb\x1e\xe1\x3d\x42\xb7\x13\xed\x99\xaa\xbe\xf6\x46\xe3\x41\x2b\x85\xcb\x6d\x74\xb2\x6d\x8f\xdb\xdd\xea\x92\xb4\xf4\x5e\xe6\x2a\x12\x3a\xc1\x19\xf7\x62\x9e\xee\x74\x6d\x6b\xfc\xff\x76\xb7\x3e\xdc\xb5\xb1\x3b\x8e\x06\x5b\x7a\x3a\x77\x80\xf5\x62\x9d\x43\x74\x0b\x73\x98\x8e\x71\xe4\xbc\xf8\x75\x01\x32\xa7\xa6\x2d\x61\x77\x1e\x39\x19\x2e\x75\xbf\x63\x8f\x41\x77\x07\xff\xe4\xd6\xf9\x64\x4b\x14\x71\x85\x01\x1f\xa6\xb0\x09\xee\xee\xcf\x1e\xda\x9e\x97\xa0\x8b\x23\xdc\xed\x4a\x91\x21\x01\xa9\xe6\x38\x86\xb1\x85\x90\xac\xc9\x59\x8d\x31\xdd\xcf\x68\x52\x09\xf2\x1a\xe6\x30\xbd\x4e\x26\xc0\x9c\xee\x40\x0c\xd4\x1a\xe8\xa1\x08\x72\x14\xf7\xe6\x4c\x5d\xc4\x0c\x1f\xbf\xda\xbc\x0a\x9d\xaa\x6d\xd6\xc1\x11\xbd\x7a\x0c\xfa\xa3\xd0\xaa\xe5\xa8\xaf\x40\x8d\xe6\xda\xae\x1b\xdb\x0a\xa3\x6a\x4d\x31\x9d\x42\x70\x0c\xc8\xb5\x75\xc5\xa4\xc4\x1f\xb7\x3c\xb3\x7a\xc4\xba\xcc\x83\xdb\x0c\x1f\xbe\xa2\x21\x7e\xfc\x28\x14\xf9\x56\x08\xfd\x7c\x11\x69\x0a\xcb\x95\xa4\x31\x3e\xbe\xfe\x1e\xff\x96\xc3\x18\xab\xe4\x95\xd2\x8f\x90\x3f\x43\xf3\xfc\x72\xc5\xf0\x3b\x94\x67\xd6\xa4\x57\x2e\x1c\x90\x64\x4c\xe1\xe4\x55\x28\xdd\xd5\xd2\xb6\x86\x06\x87\x5e\x69\xb2\x5e\x6d\x0c\x65\x41\x0e\xdb\x29\xe9\xb5\xe5\x9e\x82\x56\xd2\x87\x3a\xa4\x24\x9c\xb4\x09\xc7\x6d\xc2\x51\x9b\x70\xd8\x22\xf4\x5b\xef\xed\xf1\xb6\x82\xf6\x0c\x6d\x08\x06\xe3\xa4\x5a\xc5\xbd\x79\x9a\x1c\x6c\xdf\xd0\x43\x8b\x0d\x78\x4d\x3b\x0d\xfc\x84\x6d\x2b\xfd\xec\x40\x34\x74\x3c\x18\x31\x0d\xa4\x3b\xfd\x93\xa1\xcb\x48\xac\xbe\x95\x3c\x84\xed\xc3\xc0\x90\xfa\xbb\xf4\x1e\x81\x43\x18\xca\x82\x8c\x2f\xd9\x00\x3b\x44\x0c\x5b\x3b\x5c\x7f\x31\xf6\x32\x4b\x59\xa0\xde\x62\x33\xcb\x70\xdf\x6f\xbb\x4a\x2b\x9e\x84\x62\xe5\x89\x04\xbf\x18\x83\xb9\xab\x3d\x5f\x6e\x01\x44\x07\xf5\x0f\x0c\xb5\x9b\x70\x36\x90\xcb\xe4\x8d\x5c\x18\xf9\x97\xba\xcb\xb4\xc3\x89\x43\x76\xd5\xf0\x70\x40\x70\x37\xa5\xeb\xaf\x9c\xeb\x6d\x14\xe4\x2a\x77\x57\x5b\xa8\x96\x80\xca\x6f\xf4\x3a\xde\x0c\x4e\x82\x57\x89\x73\xda\xef\x57\x72\x25\x1f\x6e\x97\xbf\x5d\xbe\xfe\xd1\xb1\x7c\x9a\xf2\x17\x4b\x6c\xed\x76\xe1\x6f\x6f\xbb\x60\x34\x60\xa7\x11\x8c\x30\xe9\x33\xa4\x7a\xb0\x4b\x84\x6c\x8e\x3d\x08\xe0\x75\xae\x10\x41\x6f\xd7\x92\xdb\x7f\xc6\xe8\x68\x15\xb0\x16\x4e\x71\xd6\xe2\xa8\x75\x9c\x34\xae\xe6\xf8\x7d\xe3\xfd\xfe\x60\x97\x7e\xdf\xea\x99\x19\xdf\x34\xf3\x1a\x12\x5f\x4f\x6f\x00\x02\x9a\xc6\x4b\xf1\x47\x71\x15\x67\xcd\xf2\x86\xaf\xbe\x8e\x56\x47\x0c\x47\xdb\x26\xd7\x7d\x31\x30\x4a\x04\xc5\xb3\xd3\x06\x76\x40\xca\x1f\xd0\x98\x7e\xf0\xd0\x2f\x7e\x3a\x03\xd9\x1d\x7f\xfa\xf7\xef\x00\x00\x00\xff\xff\x19\x6c\xf3\xb4\x0a\x28\x00\x00")

func static_gui_index_html_bytes() ([]byte, error) {
	return bindata_read(
		_static_gui_index_html,
		"static/gui/index.html",
	)
}

func static_gui_index_html() (*asset, error) {
	bytes, err := static_gui_index_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "static/gui/index.html", size: 10250, mode: os.FileMode(420), modTime: time.Unix(1452898227, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"static/gui/index.html": static_gui_index_html,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"static": &_bintree_t{nil, map[string]*_bintree_t{
		"gui": &_bintree_t{nil, map[string]*_bintree_t{
			"index.html": &_bintree_t{static_gui_index_html, map[string]*_bintree_t{
			}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

