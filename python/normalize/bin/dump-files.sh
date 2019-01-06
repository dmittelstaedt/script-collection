#!/usr/bin/env bash
#
# Creates executables to create directory structures
#
# Author: David Mittelstaedt  <cpt.dave@web.de>
# Date: 2019-01-06

BASE_DIR=/home/david/Music

BASE_DOCKER_DIR=/data

# Current directory
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

DIR_OUT=${CURRENT_DIR}/create-dirs.sh
FILES_OUT=${CURRENT_DIR}/create-files.sh

printf "Dump files and directories...\n"

echo "#!/usr/bin/env bash" > ${DIR_OUT}
echo "#!/usr/bin/env bash" > ${FILES_OUT}

IFS=$'\n'
dirs=($(find ${BASE_DIR} -mindepth 2 -maxdepth 2 -type d | shuf | head -n 5))

for dir in "${dirs[@]}"; do
  cd ${dir}
  directory=$(pwd | cut -d / -f4-9)
  printf "mkdir -p '${BASE_DOCKER_DIR}/${directory}'\n" >> ${DIR_OUT}
  files=($(find -type f -printf "%f\n" | sort))
  for file in "${files[@]}"; do
    printf "touch $\"${BASE_DOCKER_DIR}/${directory}/${file}\"\n" >> ${FILES_OUT}
  done
done

sed -i 's/(/\\\(/g' ${FILES_OUT}
sed -i 's/)/\\\)/g' ${FILES_OUT}
sed -i 's/(/\\\(/g' ${DIR_OUT}
sed -i 's/)/\\\)/g' ${DIR_OUT}

chmod 750 ${DIR_OUT} ${FILES_OUT}

printf "Finished dumping files and directories!\n"

exit 0


# sed -i.bak 's/(/\\\(/g' test.tx
