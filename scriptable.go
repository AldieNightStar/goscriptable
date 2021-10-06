package goscriptable

import (
	"archive/zip"
	"bytes"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/otiai10/copy"
)

func SaveData(fileName string, data map[string]string) {
	sb := strings.Builder{}
	sb.Grow(32)

	for k, v := range data {
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(v)
		sb.WriteString("\n")
	}

	WriteToFile(fileName, []byte(sb.String()))
}

func LoadData(fileName string) map[string]string {
	if !IsFileExist(fileName) {
		return make(map[string]string)
	}
	f := Open(fileName)
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil
	}

	source := string(data)

	configMap := make(map[string]string)

	for _, line := range strings.Split(source, "\n") {
		if len(line) < 1 {
			continue
		}
		if !strings.Contains(line, "=") {
			continue
		}
		arr := strings.SplitN(line, "=", 2)
		if len(arr) < 2 {
			continue
		}
		configMap[arr[0]] = arr[1]
	}

	return configMap
}

func GetOsArgs() []string {
	return os.Args[1:]
}

func In(elem string, list []string) int {
	for i := 0; i < len(list); i++ {
		if list[i] == elem {
			return i
		}
	}
	return -1
}

func IsIn(elem string, list []string) bool {
	return In(elem, list) > -1
}

func CreateFolder(dir string) {
	os.MkdirAll(dir, 0750)
}

func CreateFile(name string) *os.File {
	f, err := os.Create(name)
	if err != nil {
		return nil
	}
	return f
}

func Run(workDir, cmd string, args ...string) {
	c := exec.Command(cmd, args...)
	c.Dir = workDir
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Run()
}

func RunAndReturn(workDir, cmd string, args ...string) string {
	data := make([]byte, 0, 128)
	c := exec.Command(cmd, args...)
	out := bytes.NewBuffer(data)
	c.Dir = workDir
	c.Stdout = out
	c.Stderr = out
	c.Stdin = os.Stdin
	c.Run()
	return strings.Trim(out.String(), " \n")
}

func MultiplyString(n int, s string) string {
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		sb.WriteString(s)
	}
	return sb.String()
}

func Tabulated(n int, text string) string {
	arr := strings.Split(text, "\n")
	for i := 0; i < len(arr); i++ {
		arr[i] = MultiplyString(n, " ") + arr[i]
	}
	return strings.Join(arr, "\n")
}

func Block(title string, code ...string) string {
	s := Tabulated(4, strings.Join(code, "\n"))
	return title + " {\n" + s + "\n}\n"
}

func HttpGet(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	return data
}

func WriteToFile(filename string, data []byte) {
	f := CreateFile(filename)
	defer f.Close()
	if f != nil {
		f.Write(data)
	}
}

func cloneZipItem(f *zip.File, dest string) bool {
	// Create full directory path
	path := filepath.Join(dest, f.Name)
	err := os.MkdirAll(filepath.Dir(path), os.ModeDir|os.ModePerm)
	if err != nil {
		return false
	}

	// Clone if item is a file
	rc, err := f.Open()
	if err != nil {
		return false
	}
	if !f.FileInfo().IsDir() {
		fileCopy, err := os.Create(path)
		if err != nil {
			return false
		}
		_, err = io.Copy(fileCopy, rc)
		fileCopy.Close()
		if err != nil {
			return false
		}
	}
	rc.Close()
	return true
}

func UnpackZip(filename string, dest string) bool {
	r, err := zip.OpenReader(filename)
	if err != nil {
		return false
	}
	defer r.Close()
	for _, f := range r.File {
		cloneZipItem(f, dest)
	}
	return true
}

func Delete(dir string) error {
	return os.RemoveAll(dir)
}

func Open(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	return f
}

func CopyFile(filename string, dest string) bool {
	f := Open(filename)
	f2 := CreateFile(dest)

	defer f.Close()
	defer f2.Close()

	if f == nil || f2 == nil {
		return false
	}

	w, err := io.Copy(f2, f)
	if err != nil {
		return false
	}
	if w < 1 {
		return false
	}
	return true
}

func CopyDirectory(source, dest string) error {
	return copy.Copy(source, dest, copy.Options{})
}

func TimeString() string {
	ms := time.Now().UnixNano()
	return strconv.FormatInt(ms, 16)
}

func RemoveSpaces(s []string) []string {
	arr := make([]string, 0, 32)
	for i := 0; i < len(s); i++ {
		arr = append(arr, strings.Trim(s[i], " \n"))
	}
	return arr
}

func IsFileExist(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
