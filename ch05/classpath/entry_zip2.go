package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry2 struct {
	absPath string
	zipRc   *zip.ReadCloser
}

func newZipEntry2(path string) *ZipEntry2 {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry2{absPath, nil}
}
func (self *ZipEntry2) readClass(className string) ([]byte, Entry, error) {
	if self.zipRc == nil {
		err := self.openJar()
		if err != nil {
			return nil, nil, err
		}
	}
	classFile := self.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}
	data, err := readClass(classFile)
	return data, self, err
}
func (self *ZipEntry2) openJar() error {
	r, err := zip.OpenReader(self.absPath)
	if err == nil {
		self.zipRc = r
	}
	return err
}
func (self *ZipEntry2) findClass(className string) *zip.File {
	for _, f := range self.zipRc.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}
func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (self *ZipEntry2) String() string {
	return self.absPath
}
