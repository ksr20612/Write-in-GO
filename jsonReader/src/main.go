package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"encoding/json"
	"github.com/bom-master/bom"
)

type DB_Info struct {
	Language            string `json:"Language"`
	Version             string `json:"Version"`
	ApplicationCategory string `json:"ApplicationCategory"`
	NumberOfSpeaker     string `json:"NumberOfSpeaker"`
	NumberOfUtterance   string `json:"NumberOfUtterance"`
	DataCategory        string `json:"DataCategory"`
	RecordingDate       string `json:"RecordingDate"`
	FillingDate         string `json:"FillingDate"`
	RevisionHistory     string `json:"RevisionHistory"`
	Distributor         string `json:"Distributor"`
}

type Wave_Info struct {
	SamplingRate       string `json:"SamplingRate"`
	ByteOrder          string `json:"ByteOrder"`
	EncodingLaw        string `json:"EncodingLaw"`
	NumberOfBit        string `json:"NumberOfBit"`
	NumberOfChannel    string `json:"NumberOfChannel"`
	SignalToNoiseRatio string `json:"SignalToNoiseRatio"`
}

type Label_Info struct {
	LabelText string `json:"LabelText"`
}

type Speaker_Info struct {
	SpeakerName string `json:"SpeakerName"`
	Gender      string `json:"Gender"`
	Age         string `json:"Age"`
	Region      string `json:"Region"`
	Dialect     string `json:"Dialect"`
}

type Environment_Info struct {
	RecordingEnviron string `json:"RecordingEnviron"`
	NoiseEnviron     string `json:"NoiseEnviron"`
	RecordingDevice  string `json:"RecordingDevice"`
}

type File_Info struct {
	FileCategory   string `json:"FileCategory"`
	FileName       string `json:"FileName"`
	DirectoryPath  string `json:"DirectoryPath"`
	HeaderSize     string `json:"HeaderSize"`
	FileLength     string `json:"FileLength"`
	FileFormat     string `json:"FileFormat"`
	NumberOfRepeat string `json:"NumberOfRepeat"`
	TimeInterval   string `json:"TimeInterval"`
	Distance       string `json:"Distance"`
}

type Miscellaneous_Info struct {
	QualityStatus string `json:"QualityStatus"`
}

type Result struct {
	DB_Info            DB_Info            `json:"기본정보"`
	Wave_Info          Wave_Info          `json:"음성정보"`
	Label_Info         Label_Info         `json:"전사정보"`
	Speaker_Info       Speaker_Info       `json:"화자정보"`
	Environment_Info   Environment_Info   `json:"환경정보"`
	File_Info          File_Info          `json:"파일정보"`
	Miscellaneous_Info Miscellaneous_Info `json:"기타정보"`
}

func main() {

	var result Result
	var wavDurSec float64
	var strSec string

	// choose Directory
	text, err := ioutil.ReadFile("/home/mz/goTools/jsonReader/searchDirectory.txt")
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

			fmt.Println("Name : " + info.Name())

			// open file
			jsonFile, openErr := os.Open(path)
			if openErr != nil {
				log.Fatal(openErr)
			}

			defer jsonFile.Close()

			// read json
			byteValue, readErr := ioutil.ReadAll(jsonFile)
			refinedVal := bom.Clean(byteValue)
			if readErr != nil {
				log.Fatal(readErr)
			}
			json.Unmarshal(refinedVal, &result)

			// get FileLength
			length, lenErr := strconv.ParseFloat(result.File_Info.FileLength, 64)
			if lenErr != nil {
				log.Fatal(lenErr)
			}
			wavDurSec += length
			strSec = fmt.Sprintf("%f", wavDurSec)
			fmt.Println("wavDurSec : "+strSec)
		
		}
            return nil
	})
	if e != nil {
		log.Fatal(e)
	}
	
	
	fmt.Println("@@@@@@@@@@@@@@@ TOTAL @@@@@@@@@@@@@@@")
	fmt.Println("TOTAL SECONDS : "+strSec)
	fmt.Printf("T0TAL HOURS : ")
	fmt.Sprintf("%f",wavDurSec/3600)

}
