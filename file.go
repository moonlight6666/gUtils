package gUtils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func EnsureDir(dir string) error {
	_, err := os.Stat(dir)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0655)
		return err
	} else {
		return err
	}
}

func EnsureDirPerm(dir string, perm os.FileMode) error {
	_, err := os.Stat(dir)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(dir, perm)
		return err
	} else {
		return err
	}
}

func IsDir(path string) bool {
	path = filepath.ToSlash(path)
	length := len(path)
	s := []byte(path)
	return string(s[length-1]) == "/"
}

// 文件是否存在
func IsFileExist(filename string) (bool) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

//读取文件内容
func ReadFileContext(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	context, err := ioutil.ReadAll(f)
	return context, err
}

//写文件
func WriteFileContext(filename string, context []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm) //创建文件
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(context)
	return err
}

// 获取该路径下所有文件列表
func GetAllFile(path string, suffix string) ([]string, error) {
	files := make([]string, 0, 30)
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if suffix == "" || strings.HasSuffix(path, suffix) {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

// 一次性写文件
func FilePutContext(filename string, context string) error {
	if IsFileExist(filename) {
		//如果文件存在
		err := os.Remove(filename)
		if err != nil {
			return err
		}
	}
	f, err := os.Create(filename) //创建文件
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.WriteString(f, context)
	return err
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func CopyFile(src, dst string) (int64, error) {
	//fmt.Println("拷贝文件", src, "->", dst)
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func Copy(src, dst string, ignoreList [] string) error {
	fInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if fInfo.IsDir() {
		return CopyDir(src, dst, ignoreList)
	} else {
		if IsDir(dst) {
			dst = dst + filepath.Base(src)
		}
		_, err = CopyFile(src, dst)
		return err
	}
}

func CopyDir(src, dst string, ignoreList [] string) error {
	err := filepath.Walk(src, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil {
			return err
		}
		if HasStrings(filename, ignoreList) { //过滤后缀
			return nil
		}
		relPath, err := filepath.Rel(src, filename)
		CheckErrorExit(err)

		dstPath := path.Join(dst, relPath)

		if fi.IsDir() { // 目录
			//fmt.Println("创建目录", dstPath)
			return os.MkdirAll(dstPath, fi.Mode())
		} else { // 文件
			_, err := CopyFile(filename, dstPath)
			return err
		}
	})
	return err
}
