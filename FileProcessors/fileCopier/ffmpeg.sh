#!/bin/sh
# Purpose : convert webm into wav & create json files & backup & relocate files related to kids
# Writer : ksr20612@mediazen.co.kr

# 파일 변환 
YESTERDAY="$(date '+%Y-%m-%d' -d '1 day ago')"

#comExtra
FILESC="$(sudo find /mnt/data1/mz/f1soft/2021-03-08 -name '*.m4a')"

for FILEC in $FILESC
do
  FILENAMEC="${FILEC%.*}"
  ffmpeg -i $FILEC $FILENAMEC.wav
  if [ -f "$FILENAMEC.wav" ]; then
    sudo rm $FILEC
  fi
done


