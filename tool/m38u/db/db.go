package db

import (
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

type Db struct {
	file string
	lock sync.Mutex
}

func (i *Db) Add(record string) error {
	i.lock.Lock()
	defer i.lock.Unlock()
	byes, err := ioutil.ReadFile(i.file)
	if err != nil {
		return err
	}
	str := string(byes)
	if strings.Index(str, record) > -1 {
		return nil
	}
	AppendFile(i.file, record)
	return nil
}
func (i *Db) Remove(record string) error {
	i.lock.Lock()
	defer i.lock.Unlock()
	byes, err := ioutil.ReadFile(i.file)
	if err != nil {
		return err
	}
	str := string(byes)
	strings.Replace(str, record+"\n", "", -1)
	return ioutil.WriteFile(i.file, []byte(str), os.ModePerm)
}
func (i *Db) Read() (string, error) {
	bys, err := ioutil.ReadFile(i.file)
	if err != nil {
		return "", err
	}
	return string(bys), nil
}

func AppendFile(filename string, data string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = f.WriteString(data + "\n")
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
