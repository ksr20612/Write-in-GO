package main

import (
	"config"
	"fmt"
	"io/ioutil"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"regexp"
)

var (
	dbPool, _ = config.GetDBConnection()
)

func createDir(dir string){
	
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("NO DIR : "+dir)
		err := os.Mkdir(dir,0777)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func copy(preDir string, newDir string){

	original, err6 := os.Open(preDir) 
	
	if err6 != nil{
		panic(err6)
	} 
	defer original.Close() 

	copy, err4 := os.Create(newDir) 
	if err4 != nil{ 
		panic(err4) 
	}
	defer copy.Close() 

	_, err5 := io.Copy(copy , original) 
	if err5 != nil {
		panic(err5) 
	}
	fmt.Println("COPIED : "+newDir)

}

func parseId(fullID string) string {
	
	var idContents = make([]string, 15)
	idContents = strings.Split(fullID,"_")

	if idContents[0] == "script1" {

		return idContents[1]

	}else {

		return idContents[0]		
	
	}	
}

func main() {

	// choose Directory
	text, err := ioutil.ReadFile("/home/mz/goTools/temp/moveJson/searchDirectory.txt")
	if err != nil {
		log.Fatal(err)
	}
	Directory := string(text[:])

	// regualrExpression for file Name
	libRegEx, e := regexp.Compile("^zz[\\S]+(\\.(?i)[\\S]+$)")
	if e != nil {
		log.Fatal(e)
	}

	// search all txt files under the directories
	Err := filepath.Walk(Directory, func(path string, info os.FileInfo, err error) error {

		if err == nil && libRegEx.MatchString(info.Name()) {

			fmt.Println("Path : " + path)
			fmt.Println("Info : " + info.Name())

			// parse file names
			var dirName = filepath.Dir(path)
			fmt.Println("directory : " + dirName)
			var baseName = filepath.Base(path)
			fmt.Println("baseName : " + baseName)
			var fileName = baseName[:len(baseName)-4]
			fmt.Println("fileName : " + fileName)

			var items = make([]string, 20)
			items = strings.Split(fileName, "-")		
			
			var fullID = items[0]
			fmt.Println("full ID : " + fullID)

			var domain = items[2]
			var topic = items[3]
			fmt.Println("domain : "+domain+", topic : "+topic)
			
			if domain == "02" {

				createDir("/mnt/data1/mz/jsons/old/offline_old/"+fullID)

				renameErr := os.Rename(path,"/mnt/data1/mz/jsons/old/offline_old/"+fullID+"/"+baseName)
                		if renameErr != nil {
                       			log.Fatal(renameErr)
                		}				
			
			} else if domain == "03" {

				createDir("/mnt/data1/mz/jsons/kid/offline_kid/"+fullID)

				renameErr2 := os.Rename(path,"/mnt/data1/mz/jsons/kid/offline_kid/"+fullID+"/"+baseName)
                		if renameErr2 != nil {
                       			log.Fatal(renameErr2)
                		}

			
			} else {}
			
		}	

		return nil
	})
	if Err != nil {
		log.Fatal(e)
	}

}
