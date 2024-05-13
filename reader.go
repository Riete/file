package file

import (
	"bufio"
	"io"
	"os"

	"github.com/riete/convert/str"
)

type FileReader struct {
	path string
	file *os.File
}

func (f *FileReader) open() error {
	var err error
	f.file, err = os.OpenFile(f.path, os.O_RDONLY, 0666)
	return err
}

func (f *FileReader) File() *os.File {
	return f.file
}

func (f *FileReader) Close() error {
	return f.file.Close()
}

func (f *FileReader) Read(size int64) ([]byte, error) {
	b := make([]byte, size)
	n, err := f.file.Read(b)
	if size > int64(n) {
		return b[0:n], err
	}
	return b, err
}

func (f *FileReader) ReadString(size int64) (string, error) {
	b, err := f.Read(size)
	return str.FromBytes(b), err
}

func (f *FileReader) ReadAt(size, offset int64) ([]byte, error) {
	b := make([]byte, size)
	n, err := f.file.ReadAt(b, offset)
	if err == io.EOF {
		err = nil
	}
	if size > int64(n) {
		return b[0:n], err
	}
	return b, err
}

func (f *FileReader) ReadStringAt(size, offset int64) (string, error) {
	b, err := f.ReadAt(size, offset)
	return str.FromBytes(b), err
}

func (f *FileReader) ReadAll() ([]byte, error) {
	b, err := io.ReadAll(f.file)
	return b, err
}

func (f *FileReader) ReadStringAll() (string, error) {
	b, err := f.ReadAll()
	return str.FromBytes(b), err
}

func (f *FileReader) Seek(offset int64, whence int) (int64, error) {
	n, err := f.file.Seek(offset, whence)
	return n, err
}

func (f *FileReader) BufferReader() *bufio.Reader {
	return bufio.NewReader(f.file)
}

func NewReader(path string) (*FileReader, error) {
	f := &FileReader{path: path}
	return f, f.open()
}
