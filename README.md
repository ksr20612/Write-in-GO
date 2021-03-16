# Write-in-GO
Write in GO

1. GO server
- A server for audio data collection; people can record and save
- using MediaRecorder APIs for recording 
- using fasthttp ( https://github.com/valyala/fasthttp ), fasthttp/routing ( https://github.com/jackwhelpton/fasthttp-routing )
- using GORM.io ( https://gorm.io/index.html )

2. GO json file Writer & Reader

	2.1. jsonWriter
	- search all .wav files under certain directories, search DB, and write json file containing meta-data about the wav files
	- using Standard library of GO
	- using GORM.io ( https://gorm.io/index.html )

	2.2. jsonWriterV2
	- read certain txt file line by line, get information from DB, and write json

	2.3. jsonModifer
	- read json, modify its content, and overwrite

	2.4. jsonReader
	- read json

3. GO file Processing Programs
- scripts are too slow for processing tons of files.... here's some alternatives

	3.1. fileCopier
	- search all files under certain directories, and copy those into different directories
	- find /path/to/directories/ -name '@@@' -exec cp (-r) {} /path/to/directories/ \;
	
	3.2. fileMover
	- search ceratin files, and move those
	- find /path/to/directories/ -name '@@@' -exec mv {} /path/to/directories/ \;
	
	3.3. fileRemover
	- remove files
	- find /path/to/directories/ -name '@@@' -exec rm (-r) {} \;
    
	3.4. fileRenamer
	- rename files
	- find /path/to/directories/ -name '@@@' -exec mv (-r) {} /path/to/directories/ \;
	- ren @@@ ### (cmd)
    
	3.5. fileSorter
	- sort files based on their names
