package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	source := flag.String("src", "", "Source directory")
	destination := flag.String("dst", "", "Destination directory")
	flag.Parse()

	if flag.NFlag() != 2 {
		log.Println("Incorrect number of flags")
		flag.Usage()
		return
	}

	srcfiles, err := getFileList(*source)
	if err != nil {
		log.Println("Error getting source file list: " + err.Error())
	}
	dstfiles, err := getFileList(*destination)
	if err != nil {
		log.Println("Error getting source file list: " + err.Error())
	}

	missing(srcfiles, dstfiles)

}

func getFileList(srcpath string) ([]string, error) {

	var filelist []string

	err := filepath.Walk(srcpath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			filelist = append(filelist, strings.Replace(path, srcpath, "", -1))
			return nil
		})

	return filelist, err

}

// stolen from stack overflow
func missing(a, b []string) {
	mb := make(map[string]bool, len(a))

	for _, kb := range b {
		mb[kb] = true
	}
	for _, ka := range a {
		if !mb[ka] {
			fmt.Println(ka)
		}
	}
	return
}
