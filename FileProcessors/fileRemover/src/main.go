package main

import (
	"fmt"
	"io/ioutil"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"regexp"
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

func checkBothExists(path string, fileName string) {

	var wavPath = path + ".wav"
	var jsonPath = path + ".json"

	var isWavExist = true
	var isJsonExist = true

	if _, err := os.Stat(wavPath); os.IsNotExist(err) {
		isWavExist = false
	}
	if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
		isJsonExist = false
	}	
	
	// remove wav if json does not exist
	if isWavExist==true && isJsonExist==false {

                wavErr := os.Rename(wavPath,"/mnt/data2/mz/data/AI_noPair/"+fileName+".wav")
                if wavErr != nil {
                       log.Fatal(wavErr)
                }	
	}
	
	// remove json if wav does not exist
	if isWavExist==false && isJsonExist==true {

                jsonErr := os.Rename(jsonPath,"/mnt/data2/mz/data/AI_noPair/"+fileName+".json")
                if jsonErr != nil {
                       log.Fatal(jsonErr)
                }
		
	}

}

func main() {

	// choose Directory
	text, err := ioutil.ReadFile("/mnt/data2/mz/goTools/sortJson/searchDirectory.txt")
	if err != nil {
		log.Fatal(err)
	}
	Directory := string(text[:])

	// regualrExpression for file Name
	libRegEx, e := regexp.Compile("^[\\S]+(\\.(?i)wav$)")
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
			fileName = baseName[:len(baseName)-4]
			fmt.Println("fileName : " + fileName)

			var items = make([]string, 20)
			items = strings.Split(fileName, "-")		
			
			var fullID = items[0]
			fmt.Println("full Id : " + fullID)

			
			// check if its  exist
			checkBothExists("/mnt/data2/mz/AI/com/robot/"+fullID+"/"+baseName,fileName)

		}

	return nil
	})
	if e != nil {
		log.Fatal(e)
	}

}
