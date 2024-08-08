package file

import (
	"io"
	"log"
	"os"

)

type File interface {
	Save(path string,content io.Reader)error
	CreateDir(path string)error
	close()
}

func NewFileMngmnt()File{
	return &fileMngmnt{}
}

type fileMngmnt struct {
	file *os.File
}

func (f *fileMngmnt)Save(path string,content io.Reader)error{
	out, osCreateError := os.Create(path)
	if osCreateError != nil {
		return osCreateError
	}

	_, ioCopyError := io.Copy(out, content)
	if ioCopyError != nil {
		return ioCopyError
	}

	f.file = out

	defer f.close()

	return nil
}

func (f *fileMngmnt)CreateDir(path string)error{
	osMkdirError := os.MkdirAll(path, 0o700)
	
	if osMkdirError != nil {
		return osMkdirError
	}

	return nil
}


func (f *fileMngmnt)close(){
	if f.file != nil{
		err := f.file.Close()
		if err != nil{
			log.Fatalln(err.Error())
		}
		return 
	}

	log.Fatalln("missing file")
}