package gdk

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// SelfPath gets compiled executable file absolute path
func SelfPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// RealPath get absolute filepath, based on built executable file
func RealPath(fp string) (string, error) {
	if path.IsAbs(fp) {
		return fp, nil
	}
	wd, err := os.Getwd()
	return path.Join(wd, fp), err
}

// SelfDir get compiled executable file directory
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// Name get filepath base name
func Name(fp string) string {
	return path.Base(fp)
}

// Dir get filepath dir name
func Dir(fp string) string {
	return path.Dir(fp)
}

// InsureDir mkdir if not exist
func InsureDir(fp string) error {
	if IsExist(fp) {
		return nil
	}
	return os.MkdirAll(fp, os.ModePerm)
}

// EnsureDirRW ensure the datadir and make sure it's rw-able
func EnsureDirRW(dataDir string) error {
	err := InsureDir(dataDir)
	if err != nil {
		return err
	}

	checkFile := fmt.Sprintf("%s/rw.%d", dataDir, time.Now().UnixNano())
	fd, err := Create(checkFile)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("open %s: rw permission denied", dataDir)
		}
		return err
	}
	fd.Close()
	Remove(checkFile)
	return nil
}

// Create create one file
func Create(name string) (*os.File, error) {
	return os.Create(name)
}

// Remove remove one file
func Remove(name string) error {
	return os.Remove(name)
}

// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final slash-separated element of path;
// it is empty if there is no dot.
func Ext(fp string) string {
	return path.Ext(fp)
}

// Rename rename file name
func Rename(src string, target string) error {
	return os.Rename(src, target)
}

// IsFile checks whether the path is a file,
// it returns false when it's a directory or does not exist.
func IsFile(fp string) bool {
	f, e := os.Stat(fp)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// IsExist checks whether a file or directory exists
// It returns false when the file or directory does not exist.
func IsExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

// Search a file in the give paths.
// this is often used in search config file in /etc ~/
func SearchFile(filename string, paths ...string) (fullPath string, err error) {
	for _, path := range paths {
		if fullPath = filepath.Join(path, filename); IsExist(fullPath) {
			return
		}
	}
	err = fmt.Errorf("%s not found in paths %s", filename, paths)
	return
}

// FileMTime get file modified time
func FileMTime(fp string) (int64, error) {
	f, e := os.Stat(fp)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// FileSize get file size as how many bytes
func FileSize(fp string) (int64, error) {
	f, e := os.Stat(fp)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

// DirsUnder list dirs under dirPath
func DirsUnder(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return []string{}, nil
	}
	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}
	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	ret := make([]string, 0, sz)
	for i := 0; i < sz; i++ {
		if fs[i].IsDir() {
			name := fs[i].Name()
			if name != "." && name != ".." {
				ret = append(ret, name)
			}
		}
	}
	return ret, nil
}

// FilesUnder list files under dirPath
func FilesUnder(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return []string{}, nil
	}
	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}
	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}
	ret := make([]string, 0, sz)
	for i := 0; i < sz; i++ {
		if !fs[i].IsDir() {
			ret = append(ret, fs[i].Name())
		}
	}
	return ret, nil
}

// SearchFileWithAffix search file under dirPath and meet the followinng conditions
// match prefix and suffix
// prefix and suffix must been set and not be empty

func SearchFileWithAffix(dirPath, prefix, suffix string) (fullPath string, exist bool) {
	if !IsExist(dirPath) {
		return "", false
	}
	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return "", false
	}
	sz := len(fs)
	if sz == 0 {
		return "", false
	}
	if len(prefix) == 0 || len(suffix) == 0 {
		return "", false
	}
	for i := 0; i < sz; i++ {
		if !fs[i].IsDir() {
			if strings.HasPrefix(fs[i].Name(), prefix) && strings.HasSuffix(fs[i].Name(), suffix) {
				return fs[i].Name(), true
			}
		}
	}
	return "", false
}

func OpenFile(filename string) (*os.File, error) {
	if !FileExists(filename) {
		err := os.MkdirAll(filepath.Dir(filename), 0755)
		if err != nil {
			return nil, err
		}
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	if f == nil {
		return nil, errors.New("open file must success")
	}

	return f, nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func FileEmpty(name string) bool {
	stat, err := os.Stat(name)
	return os.IsNotExist(err) || stat.Size() <= 0
}

func ReadFile(path string) string {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}

	return string(buff)
}
