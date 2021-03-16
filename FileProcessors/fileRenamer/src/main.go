package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {

	// choose Directory
	text, err := ioutil.ReadFile("/home/mz/wavRenamer/searchDirectory.txt")
	if err != nil {
		log.Fatal(err)
	}
	Directory := string(text[:])

	// regualrExpression for file Name
	libRegEx, e := regexp.Compile("^[\\S]+(\\.(?i)json$)")
	if e != nil {
		log.Fatal(e)
	}

	// regualrExpression for file Name
	libWav, e := regexp.Compile("^[\\S]+(\\.(?i)wav$)")
	if e != nil {
		log.Fatal(e)
	}

	// regualrExpression for file Name
	libTxt, e := regexp.Compile("^[\\S]+(\\.(?i)txt$)")
	if e != nil {
		log.Fatal(e)
	}


	// search all txt files under the directories
	e = filepath.Walk(Directory, func(path string, info os.FileInfo, err error) error {

		if err == nil && libRegEx.MatchString(info.Name()) {

			fmt.Println("Path : " + path)
			fmt.Println("Info : " + info.Name())

			// parse file names
			var dirName = filepath.Dir(path)
			fmt.Println("directory : " + dirName)
			var baseName = filepath.Base(path)
			fmt.Println("baseName : " + baseName)
			var fileName = baseName[:len(baseName)-5]
			fmt.Println("searching : "+fileName+".json")
			var items = make([]string, 20)
			items = strings.Split(fileName, "-")

			// set new name
			var newName = items[0] +"-"+items[1] +"-"+ items[2] + "-03-" + items[4] +"-"+ items[5] +"-"+ items[6]+".json"
			fmt.Println("new Name : "+newName)

 			// modify "fileName"	
			
			// write file
			err := os.Rename(path, dirName+"/"+newName)
                    	if err != nil {
                        	log.Fatal(err)
                    	}


		}else if err == nil && libWav.MatchString(info.Name()) {

			fmt.Println("Path : " + path)
			fmt.Println("Info : " + info.Name())

			// parse file names
			var dirName = filepath.Dir(path)
			fmt.Println("directory : " + dirName)
			var baseName = filepath.Base(path)
			fmt.Println("baseName : " + baseName)
			var fileName = baseName[:len(baseName)-4]
			fmt.Println("searching : "+fileName+".wav")
			var items = make([]string, 20)
			items = strings.Split(fileName, "-")

			// set new name
			var newName = items[0] +"-"+items[1] +"-"+ items[2] + "-03-" + items[4] +"-"+ items[5] +"-"+ items[6]+".wav"
			fmt.Println("new Name : "+newName)

 			// modify "fileName"	
			
			// write file
			err := os.Rename(path, dirName+"/"+newName)
                    	if err != nil {
                        	log.Fatal(err)
                    	}


		}else if err == nil && libTxt.MatchString(info.Name()) {

			fmt.Println("Path : " + path)
			fmt.Println("Info : " + info.Name())

			// parse file names
			var dirName = filepath.Dir(path)
			fmt.Println("directory : " + dirName)
			var baseName = filepath.Base(path)
			fmt.Println("baseName : " + baseName)
			var fileName = baseName[:len(baseName)-4]
			fmt.Println("searching : "+fileName+".txt")
			var items = make([]string, 20)
			items = strings.Split(fileName, "-")

			// set new name
			var newName = items[0] +"-"+items[1] +"-"+ items[2] + "-03-" + items[4] +"-"+ items[5] +"-"+ items[6]+".txt"
			fmt.Println("new Name : "+newName)

 			// modify "fileName"	
			
			// write file
			err := os.Rename(path, dirName+"/"+newName)
                    	if err != nil {
                        	log.Fatal(err)
                    	}

		}else {}

            return nil
	})
	if e != nil {
		log.Fatal(e)
	}
}
