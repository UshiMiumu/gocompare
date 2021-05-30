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
		log.Println("Error getting destination file list: " + err.Error())
	}

	missing(srcfiles, dstfiles)

}

// getFileList takes a ssrcpath and returns a slice
// of strings of every file & folder within that path
func getFileList(srcpath string) ([]string, error) {

	var filelist []string

	err := filepath.Walk(srcpath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// remove the srcpath to rerun just a relative location
			filelist = append(filelist, strings.Replace(path, srcpath, "", -1))
			return nil
		})

	return filelist, err

}

// find missing strings from a source slice in a
// destination slice and print them out
func missing(src []string, dest []string) {

	// create a map to hold the destination strings
	// string as key (the strings), then bool to
	// be able to use logic operator
	destmap := make(map[string]bool, len(dest))

	for _, kdest := range dest {
		// add the destination strings to the map
		// setting them all as true
		destmap[kdest] = true
	}
	for _, ksrc := range src {
		// check if src strings exist
		// if key not found then bool is false
		if !destmap[ksrc] {
			fmt.Println(ksrc)
		}
	}
	return
}
