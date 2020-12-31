// Code generated by vfsgen; DO NOT EDIT.

package resources

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

// FileSystem statically implements the virtual filesystem provided to vfsgen.
var FileSystem = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2021, 1, 1, 11, 14, 7, 431629619, time.UTC),
		},
		"/db_migrations": &vfsgen۰DirInfo{
			name:    "db_migrations",
			modTime: time.Date(2021, 1, 1, 10, 54, 22, 893038569, time.UTC),
		},
		"/db_migrations/000001_create_project_table.down.sql": &vfsgen۰FileInfo{
			name:    "000001_create_project_table.down.sql",
			modTime: time.Date(2020, 12, 18, 6, 45, 40, 917000000, time.UTC),
			content: []byte("\x44\x52\x4f\x50\x20\x54\x41\x42\x4c\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x70\x72\x6f\x6a\x65\x63\x74\x3b"),
		},
		"/db_migrations/000001_create_project_table.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000001_create_project_table.up.sql",
			modTime:          time.Date(2021, 1, 1, 9, 30, 39, 157449214, time.UTC),
			uncompressedSize: 326,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x8c\xc1\x4a\xc3\x40\x14\x45\xf7\xf9\x8a\x4b\x57\x09\x28\x54\x70\xd7\xd5\xb4\x7d\xa5\xa3\xc9\xa4\xce\xbc\xd1\xd6\x4d\x08\x9d\x67\x89\x68\x12\xd2\xc4\xef\x97\x8e\xe2\xc2\x8d\xb8\xbc\x97\x73\xce\xca\x92\x62\x02\xed\x99\x8c\xd3\xa5\x81\xde\xc0\x94\x0c\xda\x6b\xc7\x0e\xb3\x69\x6a\xc2\x75\x77\x3e\xf7\xb3\x45\xf2\xcd\xb2\x5a\xe6\xf4\x8b\xeb\x87\xee\x55\x8e\x23\xd2\x04\x40\x13\xe0\xbd\x5e\x63\x67\x75\xa1\xec\x01\xf7\x74\xc0\x9a\x36\xca\xe7\x8c\x4b\xaf\x3a\x49\x2b\x43\x3d\x4a\xf5\x71\x9b\x66\x57\x17\xa5\xad\xdf\x05\x8f\xca\xae\xb6\xca\xa6\x37\xf3\x79\x16\xe3\xc6\xe7\x39\xbc\xd1\x0f\x9e\x22\x75\xec\xda\x97\xe6\x84\x3b\x57\x9a\xe5\xd7\x31\x48\x3d\x4a\xa8\xea\x11\xac\x0b\x72\xac\x8a\x1d\x9e\x34\x6f\xe3\xc4\x73\x69\xe8\x27\x14\x85\xa9\x0f\xff\x13\x82\xbc\xc9\x1f\x42\x92\x2d\x3e\x03\x00\x00\xff\xff\x0f\x1c\xea\x55\x46\x01\x00\x00"),
		},
		"/db_migrations/000002_create_job_table.down.sql": &vfsgen۰FileInfo{
			name:    "000002_create_job_table.down.sql",
			modTime: time.Date(2020, 12, 19, 15, 17, 27, 530370581, time.UTC),
			content: []byte("\x44\x52\x4f\x50\x20\x54\x41\x42\x4c\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x6a\x6f\x62\x3b"),
		},
		"/db_migrations/000002_create_job_table.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000002_create_job_table.up.sql",
			modTime:          time.Date(2021, 1, 1, 10, 57, 24, 779310759, time.UTC),
			uncompressedSize: 822,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x92\x4f\x73\xd3\x30\x10\xc5\xef\xf9\x14\x7b\xb4\x67\x72\x08\x1d\x38\xf5\xa4\x24\x4a\x2b\x70\xe4\x22\xcb\xd0\x70\xd1\x08\x6b\x53\x54\x8a\xe4\xb1\xe4\x84\xe1\xd3\x33\x91\x9b\x3f\x06\x66\x32\x1c\xfd\x9e\x7e\x4f\xda\xf5\x5b\x08\x4a\x24\x05\x49\xe6\x05\x05\xb6\x02\x5e\x4a\xa0\x8f\xac\x92\x15\x3c\xfb\xaf\x90\x4d\x00\xc0\x1a\xa8\x6b\xb6\x84\x07\xc1\xd6\x44\x6c\xe0\x03\xdd\xc0\x92\xae\x48\x5d\x48\xe8\x7b\x6b\xd4\x13\x3a\xec\x74\x44\xb5\x7b\x9b\xe5\xd3\x03\xd2\x76\xfe\x19\x9b\xa8\x8e\xe8\x21\x96\xd7\x45\x01\x82\xae\xa8\xa0\x7c\x41\xab\xe3\x19\xc8\xac\x19\xa0\x1d\x76\xc1\x7a\x07\x8c\x4b\x7a\x47\x45\xd2\x9c\xfe\x81\xf0\x89\x88\xc5\x3d\x11\xd9\xcd\xcd\x2c\x3f\x45\x25\xdb\xef\x1d\x76\x27\xff\xcd\x6c\x36\x24\x85\xa8\xbb\xa8\x8c\x8e\x08\x92\xad\x69\x25\xc9\xfa\x61\x0c\xa2\x33\x7f\xf8\x49\xb6\x2e\x62\xb7\xd3\x2f\xa7\xc8\x77\xaf\x89\x06\x5b\x74\x26\x28\xef\x54\xab\x43\x84\x79\x59\x16\x94\xf0\xe4\x35\x3a\x36\xdf\x54\xdf\x8e\xc4\x01\x40\xd7\x58\x0c\xe3\x01\x88\x10\x64\x33\x9d\x1c\x0e\x45\x1d\xbe\xab\xf1\x88\xc7\x11\x92\xd5\x78\xb7\xb5\x4f\xf0\xbe\x2a\xf9\x3c\xa9\x7b\xeb\x8c\xdf\xab\x60\x7f\x21\xcc\xd9\x1d\xe3\xf2\x52\xf6\xdb\x6d\xc0\xf8\x0f\x23\x76\xbd\x6b\x0e\x3f\x28\xfa\x8b\x65\xe5\xc3\x23\x74\x08\x18\xc3\xf1\x92\x34\x50\x87\x3a\xa2\x51\x3a\x5e\xac\xef\x33\x93\xf7\xe9\x13\xbe\x94\x9c\x8e\xb7\xd9\xb7\xe6\xff\x00\x83\x2f\x78\x05\x18\x9e\x52\x73\xf6\xb1\xa6\x90\x9d\x0b\x35\x4d\x9d\xc8\x27\xf9\xed\xe4\xb5\xbc\x8c\x2f\xe9\xe3\xdf\xe5\x4d\x8b\x55\xd6\xfc\x84\x92\x0f\x65\x4e\xe0\x55\xec\x7c\xd5\x08\x3e\xcb\xf9\xed\xef\x00\x00\x00\xff\xff\xde\xb0\xe3\x06\x36\x03\x00\x00"),
		},
		"/gen.go": &vfsgen۰CompressedFileInfo{
			name:             "gen.go",
			modTime:          time.Date(2020, 12, 19, 15, 41, 7, 28702208, time.UTC),
			uncompressedSize: 93,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xd2\xd7\x4f\xcf\xb7\x4a\x4f\xcd\x4b\x2d\x4a\x2c\x49\x55\x28\xca\x55\xd0\x4d\x53\x28\x4a\x2d\xce\x2f\x2d\x4a\x4e\x8d\x4f\x2b\xd6\x4b\xcf\xe7\x42\x55\x92\x9e\xaf\x50\x54\x9a\x87\xac\x26\x3e\x3d\x35\x0f\xa4\xae\x20\x31\x39\x3b\x31\x3d\x15\x2e\x55\xcc\x05\x08\x00\x00\xff\xff\x19\x1b\x3c\x9d\x5d\x00\x00\x00"),
		},
		"/resource_fs_gen.go": &vfsgen۰CompressedFileInfo{
			name:             "resource_fs_gen.go",
			modTime:          time.Date(2020, 12, 18, 6, 35, 24, 131140660, time.UTC),
			uncompressedSize: 330,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x4c\x8e\x41\x4b\xc3\x40\x10\x85\xcf\x3b\xbf\x62\x9c\x53\x82\x65\x73\xaf\xf4\x26\xf5\x22\x2a\x14\xbc\xca\x36\x4e\x36\x83\x9b\xdd\x30\xbb\x29\x48\xc9\x7f\x97\x24\x16\x3d\x0e\xef\x7d\x6f\xbe\xa6\xc1\xfb\xf3\x24\xe1\x13\xc5\xc7\xa4\x0c\x30\xba\xf6\xcb\x79\xc6\xc1\x49\x04\x90\x61\x4c\x5a\xb0\x02\x43\x21\x79\x02\x43\x91\x4b\xd3\x97\x32\x12\x80\x21\x2f\xa5\x9f\xce\xb6\x4d\x43\x93\xfb\x49\xdb\x94\x9e\x9b\x4b\x97\x3d\x47\x82\x1a\xe0\xe2\x14\x95\x73\x9a\xb4\xe5\xe3\x09\x0f\xb8\x80\xf6\x51\xb4\x22\x4b\x35\x40\x37\xc5\x76\x7d\x54\xd5\x78\x05\x23\x1d\xb2\x2a\xee\x0f\xb8\x6d\xd8\x27\x8e\xac\xae\x70\xf5\x37\xb2\xbb\x65\xaf\x63\x91\x14\xf3\x15\x8c\x79\xdb\x94\x5f\xdc\xc0\x7b\x44\xba\x95\x33\xed\xc0\x98\x77\xa7\xe2\xce\xe1\x37\xa5\xa3\x04\x3e\x7d\xe7\xc2\xc3\x9a\x2e\x67\xdc\x38\xfc\x87\x7e\x74\xd9\xfa\xb4\x34\xe6\xfa\x61\x95\xba\x3b\x60\x94\xb0\x58\x9a\x90\xbc\x3d\xba\xe2\x42\x88\x15\xab\xd6\x60\x66\x98\xe1\x27\x00\x00\xff\xff\x70\x53\x6a\x0d\x4a\x01\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/db_migrations"].(os.FileInfo),
		fs["/gen.go"].(os.FileInfo),
		fs["/resource_fs_gen.go"].(os.FileInfo),
	}
	fs["/db_migrations"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/db_migrations/000001_create_project_table.down.sql"].(os.FileInfo),
		fs["/db_migrations/000001_create_project_table.up.sql"].(os.FileInfo),
		fs["/db_migrations/000002_create_job_table.down.sql"].(os.FileInfo),
		fs["/db_migrations/000002_create_job_table.up.sql"].(os.FileInfo),
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
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
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

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
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
