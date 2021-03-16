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
	text, err := ioutil.ReadFile("/home/mz/jsonModifier/searchDirectory.txt")
	if err != nil {
		log.Fatal(err)
	}
	Directory := string(text[:])

	// regualrExpression for file Name
	libRegEx, e := regexp.Compile("^[\\S]+(\\.(?i)json$)")
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
			fmt.Println("searching : "+fileName+".wav")
			var items = make([]string, 20)
			items = strings.Split(fileName, "-")

			// set new name
			var newName = items[0] +"-"+items[1] +"-"+ items[2] + "-03-" + items[4] +"-"+ items[5] +"-"+ items[6]+".wav"
			fmt.Println("new Name : "+newName)

 			// modify "fileName"
			
			// open file
			oldText, err1 := ioutil.ReadFile(path)
			if err1 != nil {
				log.Fatal(err1)
			}
			
			// search fileName
			newText := strings.Replace(string(oldText),fileName+".wav",newName,1)
			fmt.Println("@@new json : "+newText)		
			
			// write file
			err := ioutil.WriteFile(path, []byte(newText), os.FileMode(0644))
                    	if err != nil {
                        	log.Fatal(err)
                    	}
			
		}
            return nil
	})
	if e != nil {
		log.Fatal(e)
	}
}
