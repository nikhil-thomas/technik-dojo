#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-18
# Description : A script to copy error log lines
# Usage       : ./08_glob-error.sh
##################################################

ERROR_DIR=/tmp/error-dir
mkdir -p ${ERROR_DIR}
if [[ $? -ne 0 ]]
then
  echo "could not create ${ERROR_DIR}"
  exit 1
fi

for file in $(ls /var/log/*.log); do
  grep --quiet [Ee]rror ${file}
  if [[ $? -eq 0 ]]; then
    echo "${file} contains errors logs"
    cp ${file} ${ERROR_DIR}/
    new_file=${ERROR_DIR}/$(basename ${file})
    sed --quiet --in-place '/[Ee]rror/p' ${new_file}
  fi
done

