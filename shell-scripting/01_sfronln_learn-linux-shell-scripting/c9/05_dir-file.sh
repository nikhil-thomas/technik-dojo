#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-09
# Description : Script to check if a directory exists, if yes create a file
# Usage       : ./04_file-check.sh
##################################################

if [[ $# -ne 2 ]];
then
  echo "usage: ./04_dir-file.sh <dir> <file>"
  exit 1
fi

directory=$1
file=$2

if [[ ! -d ${directory} ]];
then
  mkdir ${directory}
  if [[ $? -ne 0 ]];
  then
    echo "could not create ${directory}"
    exit 1
  fi
fi

touch ${file}

