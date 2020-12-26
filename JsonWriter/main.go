package main

import (
	"config"
	"entities"
	"encoding/json"
	"strings"
	"io/ioutil"
	"os"
	"fmt"
	"log"
	"path/filepath"
	"regexp"
)
var (
	// Create Connection Pool
	dbPool, _ = config.GetDBConnection()
)

type DB_Info struct {
    Language            string  `json:"Language"`
    Version             string  `json:"Version"`
    ApplicationCategory string  `json:"ApplicationCategory"`
    NumberOfSpeaker     string  `json:"NumberOfSpeaker"`
    NumberOfUtterance   string  `json:"NumberOfUtterance"`
    DataCategory        string  `json:"DataCategory"`
    RecordingDate       string  `json:"RecordingDate"`
    FillingDate         string  `json:"FillingDate"`
    RevisionHistory     string  `json:"RevisionHistory"`
    Distributor         string  `json:"Distributor"`
}

type Wave_Info struct {
    SamplingRate        string  `json:"SamplingRate"`
    ByteOrder           string  `json:"ByteOrder"`
    EncodingLaw         string  `json:"EncodingLaw"`
    NumberOfBit         string  `json:"NumberOfBit"`
    NumberOfChannel     string  `json:"NumberOfChannel"`
    SignalToNoiseRatio  string  `json:"SignalToNoiseRatio"`
}

type Label_Info struct {
    LabelText   string  `json:"LabelText"`
}

type Speaker_Info struct {
    SpeakerName string  `json:"SpeakerName"`
    Gender      string  `json:"Gender"`
    Age         string  `json:"Age"`
    Region      string  `json:"Region"`
    Dialect     string  `json:"Dialect"`
}

type Environment_Info struct {
    RecordingEnviron string `json:"RecordingEnviron"`
    NoiseEnviron     string `json:"NoiseEnviron"`
    RecordingDevice  string `json:"RecordingDevice"`
}

type File_Info struct {
    FileCategory    string  `json:"FileCategory"`
    FileName        string  `json:"FileName"`
    DirectoryPath   string  `json:"DirectoryPath"`
    HeaderSize      string  `json:"HeaderSize"`
    FileLength      string  `json:"FileLength"`
    FileFormat      string  `json:"FileFormat"`
    NumberOfRepeat  string  `json:"NumberOfRepeat"`
    TimeInterval    string  `json:"TimeInterval"`
    Distance        string  `json:"Distance"`
}

type Miscellaneous_Info struct {
    QualityStatus   string  `json:"QualityStatus"`
}

type Result struct {
    DB_Info                 DB_Info             `json:"기본정보"`
    Wave_Info               Wave_Info           `json:"음성정보"`
    Label_Info              Label_Info          `json:"전사정보"`
    Speaker_Info            Speaker_Info        `json:"화자정보"`
    Environment_Info        Environment_Info    `json:"환경정보"`
    File_Info               File_Info           `json:"파일정보"`
    Miscellaneous_Info      Miscellaneous_Info  `json:"기타정보"`
}

func getContents(gender *string, age *string, region *string, location *string, tool *string, dialect *string){

    switch *gender {
        case "M":
            *gender = "Male"
        case "F":
            *gender = "Female"
        case "Z":
            *gender = "NotProvided"
    }

    switch *age {
        case "01":
            *age = "3~6"
        case "02":
            *age = "7~10"
        case "03":
            *age = "11~19"
        case "04":
            *age = "20~29"
        case "05":
            *age = "30~39"
        case "06":
            *age = "40~49"
        case "07":
            *age = "50~59"
        case "08":
            *age = "60~69"
        case "09":
            *age = "over70"
        case "10":
            *age = "NotProvided"
    }

    switch *region {
        case "A":
            *region = "서울/인천/경기"
        case "B":
            *region = "대전/세종/충청/강원"
        case "C":
            *region = "광주/전라/제주"
        case "D":
            *region = "부산/대구/울산/경상"
        case "E":
            *region = "기타(외국)"
        case "Z":
            *region = "NotProvided"
    }

    switch *location {
        case "01":
            *location = "가정"
        case "02":
            *location = "사무실"
        case "03":
            *location = "마트/(키즈)카페"
        case "04":
            *location = "공원"
        case "05":
            *location = "지하철 역사"
        case "06":
            *location = "차량"
        case "07":
            *location = "기타"
        case "99":
            *location = "NotProvided"
    }

    switch *tool {
        case "01":
            *tool = "휴대폰"
        case "02":
            *tool = "휴대폰"
        case "03":
            *tool = "노트북"
        case "04":
            *tool = "스마트패드"
        case "05":
            *tool = "녹음 전용 기기"
        case "06":
            *tool = "기타"
        case "07":
            *tool = "NotProvided"
    }

    switch *dialect {
        case "A":
            *dialect = "강원"
        case "B":
            *dialect = "경기/서울"
        case "C":
            *dialect = "경상"
        case "D":
            *dialect = "전라"
        case "E":
            *dialect = "제주"
        case "F":
            *dialect = "충청"
        case "G":
            *dialect = "기타(외국)"
        case "Z":
            *dialect = "NotProvided"
    }
}



func main() {

	var userInfo entities.Worker
    var command string
    var recordingDate string

	// choose Directory
	var Directory = "/home/mz/data/script1_aa_0011/"

    // regualrExpression for file Name
    libRegEx, e := regexp.Compile("^[\\S]+(\\.(?i)txt$)")
    if e != nil {
        log.Fatal(e)
    }

    // search all txt files under the directories
    e = filepath.Walk(Directory, func(path string, info os.FileInfo, err error) error {

        if err == nil && libRegEx.MatchString(info.Name()){

            fmt.Println("Path : "+path)
            fmt.Println("Info : "+info.Name())
            // parse file names
            var dirName = filepath.Dir(path)
            fmt.Println("directory : "+dirName)
            var baseName = filepath.Base(path)
            fmt.Println("baseName : "+baseName)
            var fileName = baseName[:len(baseName)-4]
            fmt.Println("fileName : "+fileName)
            var items = make([]string, 20)
            items = strings.Split(fileName, "-")

            // connect to DB & get location, tool, dialect using worker_id
            dbPool.Debug().Raw("select worker_loc, worker_tool, worker_dialect FROM worker WHERE worker_id = ?", items[0]).Scan(&userInfo)

            if len(userInfo.Worker_loc)==0 {
                userInfo.Worker_loc = "99"
            }
            if len(userInfo.Worker_tool)==0 {
                userInfo.Worker_tool = "99"
            }
            if len(userInfo.Worker_dialect)==0 {
                userInfo.Worker_dialect = "Z"
            }
            var worker_name = items[4]
            fmt.Println("name : "+worker_name)
            if worker_name=="ZZZ" {
                worker_name = "NotProvided"
            }
            var worker_gender = items[5]
            fmt.Println("gender : "+worker_gender)
            var worker_age = items[6]
            fmt.Println("age : "+worker_age)
            var worker_region = items[7]
            fmt.Println("region : "+worker_region)
            var worker_loc = userInfo.Worker_loc
            fmt.Println("location : "+worker_loc)
            var worker_tool = userInfo.Worker_tool
            fmt.Println("tool : "+worker_tool)
            var worker_dialect = userInfo.Worker_dialect
            fmt.Println("dialect : "+worker_dialect)

            getContents(&worker_gender, &worker_age, &worker_region, &worker_loc, &worker_tool, &worker_dialect)

            // read command

            dbPool.Debug().Raw("select comm_sent FROM command WHERE comm_id = ?", items[1]).Scan(&command)

            // @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
            if command[len(command)-1:len(command)] == "\r" {
                command = command[:len(command)-1]
            }
            fmt.Println("command : "+command)
            // @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

            // read Record
            dbPool.Debug().Raw("select create_date FROM record WHERE worker_id = ? and comm_id = ?", items[0], items[1]).Scan(&recordingDate)
            //var recordingDate = recInfo.Create_date
            if len(recordingDate) == 0 {
                recordingDate = "NotProvided"
            }

            // insert variable
            data := make([]Result, 1)

            data[0].DB_Info.Language = "ENG"
            data[0].DB_Info.Version = "N/A"
            data[0].DB_Info.ApplicationCategory = "N/A"
            data[0].DB_Info.NumberOfSpeaker = "N/A"
            data[0].DB_Info.NumberOfUtterance = "N/A"
            data[0].DB_Info.DataCategory = "mariaDB"
            data[0].DB_Info.RecordingDate = recordingDate
            data[0].DB_Info.FillingDate = "N/A"
            data[0].DB_Info.RevisionHistory = "N/A"
            data[0].DB_Info.Distributor = "Mediazen"

            data[0].Wave_Info.SamplingRate = "48000"
            data[0].Wave_Info.ByteOrder = "N/A"
            data[0].Wave_Info.EncodingLaw = "SignedIntegerPCM"
            data[0].Wave_Info.NumberOfBit = "16"
            data[0].Wave_Info.NumberOfChannel = "1"
            data[0].Wave_Info.SignalToNoiseRatio = "N/A"

            data[0].Label_Info.LabelText = command

            data[0].Speaker_Info.SpeakerName = worker_name
            data[0].Speaker_Info.Gender = worker_gender
            data[0].Speaker_Info.Age = worker_age
            data[0].Speaker_Info.Region = worker_region
            data[0].Speaker_Info.Dialect = worker_dialect

            data[0].Environment_Info.RecordingEnviron = worker_loc
            data[0].Environment_Info.NoiseEnviron = worker_loc
            data[0].Environment_Info.RecordingDevice = worker_tool

            data[0].File_Info.FileCategory = "Audio"
            data[0].File_Info.FileName = fileName+".wav"
            data[0].File_Info.DirectoryPath = dirName
            data[0].File_Info.HeaderSize = "44"
            data[0].File_Info.FileLength = "N/A"
            data[0].File_Info.FileFormat = "PCM"
            data[0].File_Info.NumberOfRepeat = "1"
            data[0].File_Info.TimeInterval = "N/A"
            data[0].File_Info.Distance = "30"

            data[0].Miscellaneous_Info.QualityStatus = "Good"

            // json marshal
            doc, _ := json.Marshal(data)

            // save json file
            err := ioutil.WriteFile(dirName+"/"+fileName+".json", doc, os.FileMode(0644))
            if err != nil {
                fmt.Println(err)
            }
        }
        return nil
    })
    if e != nil {
        log.Fatal(e)
    }
}

