# Write-in-GO
Write in GO

1. GO server
- A server for audio data collection; people can record and save
- using MediaRecorder APIs for recording 
- using fasthttp ( https://github.com/valyala/fasthttp ), fasthttp/routing ( https://github.com/jackwhelpton/fasthttp-routing )
- using GORM.io ( https://gorm.io/index.html )

2. GO json file Writer
- search all .wav files under certain directories, search DB, and write json file containing meta-data about the wav files
- using Standard library of GO
- using GORM.io ( https://gorm.io/index.html )
