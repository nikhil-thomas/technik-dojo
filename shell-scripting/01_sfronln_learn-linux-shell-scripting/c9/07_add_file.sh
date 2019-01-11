#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-11
# Description : Script to create a file and write contents
# Usage       : ./07_add_file.sh
##################################################

if [[ $# -ne 3 ]];
then
  echo "usage: ./07_add_file.sh <directory> <file> <contents>"
  exit 1
fi

directory=$1
file=$2
contents=$3

absolutePath=${directory}/${file}

if [[ ! -d ${directory} ]];
then
  mkdir ${directory} || { echo "cannot create directory ${directory}"; exit 1; }
fi

if [[ ! -f ${absolutePath} ]];
then
  touch ${file} || { echo "cannot create file ${file}"; exit 1; }
fi

echo ${contents} > ${absolutePath}

