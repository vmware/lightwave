// Code generated by vfsgen; DO NOT EDIT.

package static

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// SetupAssets statically implements the virtual filesystem provided to vfsgen.
var SetupAssets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 12, 6, 17, 37, 51, 135495215, time.UTC),
		},
		"/configs": &vfsgen۰DirInfo{
			name:    "configs",
			modTime: time.Date(2019, 10, 31, 16, 46, 8, 895154054, time.UTC),
		},
		"/configs/schema.ldif": &vfsgen۰CompressedFileInfo{
			name:             "schema.ldif",
			modTime:          time.Date(2019, 10, 31, 16, 46, 8, 895154054, time.UTC),
			uncompressedSize: 8840,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x4d\x73\xda\x48\x13\xbe\xf3\x2b\xba\x8a\x54\xc5\x3e\x58\x65\xd9\x8e\xf3\xbe\x5b\xb5\x07\x82\xa9\x84\x0a\x60\x16\x89\xec\xe6\xb4\x35\x91\x1a\x98\x8d\x98\x61\x67\x06\x08\xfb\xeb\xb7\x24\xa1\x2f\x42\x4b\xe2\xc3\x78\x7d\x32\xe8\xe9\xa1\x9f\x99\x9e\x7e\xba\x5b\x4d\xd0\xde\x0c\xe7\xec\xcf\x15\x2a\xcd\xa5\xf8\xf5\xd6\xba\xb5\x6c\xeb\xb6\xd1\x68\xde\xd4\xf9\x6b\x34\xa1\xc7\xa7\x33\xb3\x66\x2b\x04\xc7\x75\xa0\x65\x8c\xe2\xdf\x96\x06\x75\xdd\x15\x1a\x2c\x31\x31\x9b\x05\xea\x5f\xe0\xaa\x01\x00\xd9\xb2\x96\xe3\x3a\x56\x8a\xb1\xec\xe8\xe9\xa0\xd5\xef\xc0\xdb\x20\xc1\x38\xae\xf3\x24\xe7\x8c\x8b\x01\x9b\xe3\xdb\x08\xf1\xd4\x71\xda\xf0\xb6\xe8\xdc\x0d\xec\xa2\x9c\xaf\x03\xb7\xf5\x07\xd8\xd6\xbd\xf5\x68\xd9\xd6\x83\x65\x5b\xf6\xc3\xe3\xa3\x65\xdb\xef\x2c\xfb\xce\x0e\x3f\xbe\x8b\x81\xdd\xc1\xc7\x5e\xe7\xe6\x4b\xab\x37\xee\x44\x5f\x5c\x1f\xea\xf8\x1d\xe1\xb8\x8b\x82\x09\xf3\x19\x37\x65\x7e\xc7\x20\xf8\x9e\xa0\xaa\xfd\x7e\xb8\x3d\x93\xdf\xf7\x84\xdf\xed\x40\x7a\xdf\x5d\x19\xa0\x62\xc2\x2b\xdd\xf4\x08\x09\x3b\xd0\x6a\x06\x77\xef\xcf\xc4\xe0\x81\x0a\x19\x0c\x70\xca\x0c\x97\xa2\x2d\x97\xc2\x94\xc6\x4d\x0a\x85\x1c\xf6\x82\x1c\xde\x11\x1c\x46\x28\x70\x5d\xe9\x7e\x84\x7a\x25\xcf\x1f\x09\xcf\xfb\xec\xc7\x07\x64\x0a\x95\x2b\xbf\xa3\xe8\xf1\x09\x1a\x5e\x7e\x79\xfb\xec\x07\xc4\x26\x10\xd9\x40\xd1\xe8\x82\x9c\xde\x57\x71\x1a\xe1\x44\xa1\x9e\x1d\x43\x6d\x6b\xfa\xda\x14\xff\x47\x53\xfc\xf4\xfc\xf9\x20\x62\x9f\x64\xe0\xa3\x02\x39\x81\xcf\xb8\x79\x6d\x5e\xff\x2f\xe5\x75\xcc\xb9\x15\xe9\xfd\x47\x8e\xcf\xbe\x25\x78\xa6\x02\xed\x6e\x16\xa5\xd4\x52\x20\x64\xc8\x0b\xca\xa5\x4d\x09\x7d\xea\xd7\x17\x16\x2c\x6b\x32\xc8\x41\x6b\x53\x38\xdc\x63\x52\xe1\xf9\x1c\xe5\xb2\x34\x41\x17\x20\x97\x0c\x13\x4a\xdd\xc7\x1a\xd5\x07\xa6\xf1\x69\x50\xe6\x76\x88\x02\x07\x99\xf2\x66\x10\xa2\x21\x81\x5f\x32\x50\x28\x79\xff\xa8\xe4\x72\x51\xcd\x21\x82\xbd\x3a\x09\x4a\xdf\xdb\x52\x08\xf4\xc2\xc2\xc3\x31\x8a\x8b\xa9\x2e\x2d\xb4\x52\x30\x14\xd0\x2f\x18\xf3\x94\xba\xb7\x96\x66\x86\xc2\x70\x2f\x2a\x9a\x2a\x73\x4d\x01\xfd\x4a\x09\x87\x12\xf5\xd0\xb9\x88\x41\xf9\xde\x07\x1c\x85\x89\x98\x08\x88\x7e\x10\x58\x10\xc8\x35\xfa\x27\x9c\x81\x5b\xcb\x71\x4a\xaa\x1d\x3e\x15\xa8\x4a\xbd\x8e\x55\x4a\xf3\xa9\xe0\x62\x0a\x1e\x2a\xc3\x27\xe1\x29\xa0\x86\xab\x80\x69\x03\xe2\xfa\xc0\x86\xe3\xf0\x7d\xa7\x14\x79\xa8\xe4\x8a\xfb\xa8\xaa\x82\xe7\x67\xdc\x25\xbb\x3a\x52\x67\x03\xce\x4a\xb7\x3e\x07\xb8\xa4\xbf\x94\xac\x86\xa9\xbc\xaa\x7b\x2e\x62\x2e\xe9\x35\x25\xad\x43\xa6\xf5\x5a\x2a\xbf\x34\x3e\xd6\xa7\x5c\xc1\x7a\xfe\x51\x32\xea\xa0\xd6\x5c\xd6\x2a\x25\x13\x4c\x58\x48\x32\xd0\xb1\xe1\x39\x0b\x82\x5a\xc9\xe4\x8e\x52\xd3\x11\xfa\x5c\xa1\x67\xc6\xa3\x6e\x21\xac\x93\xef\xc3\xbc\x9d\x3d\x7b\xb9\x6c\x77\x47\x29\xe5\x33\xf7\xbd\x38\x09\x87\x39\xb8\x8f\x66\x26\x4b\xc3\x22\xc4\x83\x17\x67\x6d\x56\xd4\x9f\x79\xce\xf8\x3c\x51\x5e\x8f\x1a\x25\xa6\x43\xa9\x4d\x4f\x4e\xe5\xd2\xe4\x4e\xa1\x34\xe4\xa5\x36\x10\x44\x16\xa0\xb6\x26\x90\xda\xbc\xe0\xe1\x50\x1a\x1a\x7b\x5f\xe1\x75\x0c\x3a\xc2\xcf\x33\x6c\x3d\xa5\xa1\xdb\x88\xd2\x3a\x94\xc6\xec\x22\xf7\x9d\x03\xca\x99\x11\xfe\xbd\x44\x6d\xd2\x12\x21\x59\x2c\xbb\xf2\xc9\x72\x17\xac\xfc\xef\x28\xd9\x6d\x2b\xa9\x75\x3c\x6f\x2c\xe3\xf8\xfb\x0c\xcd\x0c\x15\x30\xad\xa5\xc7\x99\x41\x1f\x42\xca\x6a\x03\x72\x81\x2a\xaa\x20\x98\x17\x2e\x05\x26\x5a\x4b\xc3\x95\x5e\x7a\x33\x60\x1a\x98\x00\xc9\x7d\xef\x26\xbe\x7c\xb5\x0b\x8c\x73\xa5\xb8\x7b\x4a\xb2\x87\xcb\x6f\x01\xf7\x3a\xc2\x5f\x48\x5e\x4e\x3e\x46\x02\x16\xa0\x17\x0c\xd7\x7b\x4a\xc5\x9d\x68\x9a\xff\x25\x1e\xe6\x97\x31\x88\xc7\xfe\xb0\xca\x23\x4f\x22\x70\xe4\xfb\x82\xe7\x6f\x7f\x85\xb9\xc9\x0b\x98\xd6\x07\xbc\x33\x90\x91\x59\x3b\xb6\xda\xbf\x5b\x31\x24\x5a\x98\x7c\x69\x10\x87\xf9\x08\x27\xa8\xb0\x62\x88\x1d\x95\x6b\x60\x24\xe4\xaf\x86\x33\x1e\x02\x0b\x1f\xc4\x9f\xdc\xd1\xb8\xed\x8e\x47\xad\x5e\xf4\xb1\x3f\x76\x5c\xb8\x02\x4f\x44\x9f\xa2\xbf\x37\xb0\x77\xf2\x9f\x01\xae\x93\xed\x3c\x90\x60\xf9\xcb\x85\xea\x37\x0b\x19\x1d\x23\x17\xf5\xc8\x14\xb9\x64\x2f\x56\x72\x80\x85\xe2\x2b\x66\x30\xcf\xf0\x0d\x2c\x35\xaa\x76\xd6\x68\x10\xeb\x6d\x7b\x97\xc6\x76\x47\xfa\xad\xaf\x70\x05\xf4\xeb\x07\xca\xab\xe2\x84\x9f\x40\x65\x33\x74\x02\xb0\x7f\x54\x4d\x83\x77\x07\xa4\x04\x32\x6b\x2c\xab\x7e\x77\xdf\x58\xb2\xf4\xe7\x0f\x30\xd8\x29\x51\x1b\xa7\x86\x22\x55\x0a\x77\xfd\x50\x1f\xcc\xc6\x31\x52\x95\xde\xb4\x3d\xc0\x83\x02\xb3\x3a\x32\xf7\x9c\x43\x71\x56\x41\x21\xf3\xad\xe6\xee\xa5\x0d\x43\x74\x27\x46\x5b\x69\x6a\xd8\xb3\x56\xd2\x4c\x91\xbf\xb5\x6d\x6d\xca\xec\xe3\xe9\x16\x85\xc8\x0d\xc0\x28\xc8\x4f\xe3\xa5\x3c\x90\xba\xa9\xd5\x61\x7c\x74\xf0\x50\xdd\x47\xf7\x29\x9b\xfb\xd6\x8e\x1d\xd8\x31\x39\x25\x8a\x0a\x63\xf3\x3a\x47\x5f\x98\x52\x1f\xbd\x1f\xd5\xcd\x4e\x69\x8b\xb3\x40\xd1\x7d\x4a\xc6\x82\x90\x37\x38\x94\x7f\xbe\xfd\xdb\x47\xbf\xc8\x7e\x5f\x2f\x46\xa4\x9f\xbd\xbd\x0d\x81\x4d\xbb\x08\xe2\x39\x59\xac\x53\xf8\xac\xd4\x3d\x39\xed\x51\xbd\xdb\x13\x2e\x02\xb9\x99\x57\x9c\x54\x86\x02\x2e\x26\x52\xcd\xa3\x54\x74\xdc\x51\x15\xcb\x58\x0a\x55\xa8\x14\xf7\x5c\xdc\x66\xa3\x19\xfe\x97\xd6\x9c\x5d\xe1\x73\x2f\xde\x88\x66\x02\x8e\xd9\xae\xe6\xeb\xb4\xd2\xe0\xff\xa0\x1f\x26\xa6\xa1\xe2\xc2\xe3\x0b\x16\xc4\xe3\xa2\xd4\xe2\x63\xef\xf9\x43\xab\xd7\xfb\x7a\x33\x1e\x74\x7f\x1b\x77\xe2\x07\xd7\x8d\xe6\xbf\x01\x00\x00\xff\xff\x67\xa5\x62\x0e\x88\x22\x00\x00"),
		},
		"/init": &vfsgen۰DirInfo{
			name:    "init",
			modTime: time.Date(2019, 10, 31, 16, 46, 8, 899153917, time.UTC),
		},
		"/init/lightwave-stsd.service.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "lightwave-stsd.service.tmpl",
			modTime:          time.Date(2019, 10, 31, 16, 46, 8, 899153917, time.UTC),
			uncompressedSize: 276,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x41\x6a\xc3\x30\x10\x45\xf7\x73\x0a\x5f\xa0\xa6\x17\xd0\xa2\x6d\x42\x09\x74\x51\xaa\x84\x2e\x8c\x17\x42\xfe\x71\x85\x6d\xc9\x9d\x19\xc5\x31\xc1\x77\x2f\x0e\x2e\xb4\xcb\x37\x6f\xfe\xe7\x57\xa7\x18\xb4\xa6\x1d\xc4\x73\x18\x35\xa4\x68\xde\x42\xfb\xa5\x93\xbb\xa0\xb0\xf0\x99\x51\x1c\x53\x87\x58\x58\xf0\x25\x78\xd0\xd3\x59\xc1\x46\x66\xe9\x53\x5b\xaa\xe3\x16\x5a\x44\xe8\x94\xb8\xfb\xc5\x7e\x92\xa1\x29\x65\x0b\x7c\xe0\x3b\x07\x86\x98\xff\x67\xaa\xb6\xc6\x9a\x8e\xf3\x08\x73\x4e\xdc\x85\xd8\xd2\xfe\x0a\x6f\xd5\xb1\x9a\xdb\xad\xb4\x2a\x2f\x43\xb3\x2c\x64\xb3\xf7\x10\xd9\x5f\x83\x5a\x75\x9a\xc5\x3c\xd2\x49\xc0\xeb\xd3\x3b\xa7\xd5\xad\x78\xd8\x2d\x0b\xbd\x72\xca\xe3\x1f\x71\xe7\xbb\xa1\xea\x10\x45\x5d\xdf\xd7\xf4\xe9\xa2\xa2\x79\x9e\xcd\x90\x7b\x0d\x0f\x59\xc0\xdb\x7c\xfa\x09\x00\x00\xff\xff\x14\x59\x6a\xe5\x14\x01\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/configs"].(os.FileInfo),
		fs["/init"].(os.FileInfo),
	}
	fs["/configs"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/configs/schema.ldif"].(os.FileInfo),
	}
	fs["/init"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/init/lightwave-stsd.service.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
