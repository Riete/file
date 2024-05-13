package file

import (
	"errors"
	"io"
	"os"

	"github.com/riete/convert/str"
)

type FileWriter struct {
	path string
	file *os.File
}

func (f *FileWriter) open() error {
	var err error
	f.file, err = os.OpenFile(f.path, os.O_WRONLY|os.O_CREATE, 0666)
	return err
}

func (f *FileWriter) Close() error {
	return f.file.Close()
}

func (f *FileWriter) Write(b []byte) (int, error) {
	n, err := f.file.Write(b)
	return n, err
}

func (f *FileWriter) WriteString(s string) (int, error) {
	n, err := f.file.WriteString(s)
	return n, err
}

func (f *FileWriter) Append(b []byte) (int, error) {
	if _, err := f.file.Seek(0, io.SeekEnd); err != nil {
		return 0, err
	}
	return f.Write(b)
}

func (f *FileWriter) AppendString(s string) (int, error) {
	if _, err := f.file.Seek(0, io.SeekEnd); err != nil {
		return 0, err
	}
	return f.WriteString(s)
}

func (f *FileWriter) WriteAt(b []byte, offset int64) (int, error) {
	n, err := f.file.WriteAt(b, offset)
	return n, err
}

func (f *FileWriter) WriteStringAt(s string, offset int64) (int, error) {
	return f.WriteAt(str.ToBytes(s), offset)
}

func (f *FileWriter) WriteWithTruncate(b []byte) (int, error) {
	if err := f.Truncate(0); err != nil {
		return 0, errors.New("truncate file error")
	}
	return f.Write(b)
}

func (f *FileWriter) WriteStringWithTruncate(s string) (int, error) {
	if err := f.Truncate(0); err != nil {
		return 0, errors.New("truncate file error")
	}
	return f.WriteString(s)
}

func (f *FileWriter) Truncate(size int64) error {
	return f.file.Truncate(size)
}

func NewFileWriter(path string) (*FileWriter, error) {
	f := &FileWriter{path: path}
	return f, f.open()
}
