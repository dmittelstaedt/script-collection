#!/usr/bin/env bash

BASE_DIR=/home/david/Music
DIR_OUT=/home/david/Documents/music-dirs.txt
FILES_OUT=/home/david/Documents/music-files.txt

printf "Dump files and directories...\n"

IFS=$'\n'
dirs=($(find ${BASE_DIR} -mindepth 2 -maxdepth 2 -type d | shuf | head -n 15))

for dir in "${dirs[@]}"; do
  cd ${dir}
  directory=$(pwd | cut -d / -f4-9)
  printf "${directory}\n" >> ${DIR_OUT}
  files=($(find -type f -printf "%f\n" | sort))
  for file in "${files[@]}"; do
    printf "${directory}/${file}\n" >> ${FILES_OUT}
  done
done

printf "Finished dumping files and directories!\n"

exit 0
