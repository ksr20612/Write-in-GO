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
			var fileName = baseName[:len(baseName)-4]
			fmt.Println("fileName : " + fileName)
			var items = make([]string, 20)
			items = strings.Split(fileName, "-")		
			
			var fullID = items[0]

			fmt.Println("full Id : " + fullID)
			
			var keyId = parseId(fullID)

			fmt.Println("key id : " + keyId)

			// 나누기
			
			// 일반, ai로봇
			if keyId == "a" || keyId == "b" || keyId == "c" || keyId == "d" || keyId == "e" || keyId == "f" {
				
				// check if dir exists, create if not
				createDir("/mnt/data2/mz/AI/com/robot/"+fullID)
					
				// copy
				copy(path,"/mnt/data2/mz/AI/com/robot/"+fullID+"/"+baseName)
				
			// 일반, ai비서
			} else if keyId== "g" || keyId== "h" || keyId== "i" {
				
				// check if dir exists, create if not
				createDir("/mnt/data2/mz/AI/com/secr/"+fullID)
					
				// copy
				copy(path,"/mnt/data2/mz/AI/com/secr/"+fullID+"/"+baseName)

			// 일반, 키오스크
			} else if keyId== "j" || keyId== "k" || keyId== "m" {

				// check if dir exists, create if not
				createDir("/mnt/data2/mz/AI/com/kiosk/"+fullID)
					
				// copy
				copy(path,"/mnt/data2/mz/AI/com/kiosk/"+fullID+"/"+baseName)

			// 노인, AI로봇
			} else if keyId== "q" || keyId== "r" || keyId== "s" {

				// check if dir exists, create if not
				createDir("/mnt/data2/mz/AI/old/robot/"+fullID)
					
				// copy
				copy(path,"/mnt/data2/mz/AI/old/robot/"+fullID+"/"+baseName)	
				
			// 노인, AI비서		
			} else if keyId== "n" || keyId== "o" || keyId== "p" {

				// check if dir exists, create if not
				createDir("/mnt/data2/mz/AI/old/secr/"+fullID)
					
				// copy
				copy(path,"/mnt/data2/mz/AI/old/secr/"+fullID+"/"+baseName)	


			// 노인, 키오스크
			} else if keyId== "w" || keyId== "x" || keyId== "y" {

				// check if dir exists, create if not
				createDir("/mnt/data2/mz/AI/old/kiosk/"+fullID)
					
				// copy
				copy(path,"/mnt/data2/mz/AI/old/kiosk/"+fullID+"/"+baseName)	

			// 유소아, AI로봇
			} else if keyId== "t" || keyId== "u" || keyId== "v" || keyId== "ad" || keyId== "af" {

				// check if dir exists, create if not
				createDir("/mnt/data2/mz/AI/kid/robot/"+fullID)
					
				// copy
				copy(path,"/mnt/data2/mz/AI/kid/robot/"+fullID+"/"+baseName)	

			// 유소아, AI비서
			} else if keyId== "z" || keyId== "aa" || keyId== "ab" || keyId== "ac" || keyId== "ae" {

				// check if dir exists, create if not
				createDir("/mnt/data2/mz/AI/kid/secr/"+fullID)
					
				// copy
				copy(path,"/mnt/data2/mz/AI/kid/secr/"+fullID+"/"+baseName)	

			}

		}

	return nil
	})
	if e != nil {
		log.Fatal(e)
	}

}
