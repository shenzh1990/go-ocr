package main

import (
	"BitCoin/pkg/settings"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"os"
	"path/filepath"
	"time"
)

func main() {

	fmt.Println(settings.FConfig.Seconds)
	//  doAdd(settings.FConfig.FilePath,"a1.txt")
	// doDelete()
}
func exec() {
	gocron.Every(settings.FConfig.Seconds).Seconds().Do(doDelete)
	<-gocron.Start()
}
func doAdd(dirpath string, filename string) error {
	file, err := os.Create(dirpath + "/" + filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	fmt.Println("add file " + dirpath + "/" + filename)
	return nil
}
func doDelete() {
	list, err := doList(settings.FConfig.FilePath + "/")
	if err != nil {
		fmt.Println(err)
		return
	}
	d := time.Duration(settings.FConfig.TimeSetting) * time.Hour
	for _, v := range list {
		if v.ModTime().Add(d).Before(time.Now()) {
			err = os.Remove(settings.FConfig.FilePath + "/" + v.Name())
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("delete file " + settings.FConfig.FilePath + "/" + v.Name())
		}

	}
}
func doList(dirpath string) ([]os.FileInfo, error) {
	var dir_list []os.FileInfo
	dir_err := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if !f.IsDir() {
				dir_list = append(dir_list, f)
				return nil
			}

			return nil
		})
	return dir_list, dir_err
}
