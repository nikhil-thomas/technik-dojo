#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-12
# Description : Script to copy a file to script location
# Usage       : ./08_copy-file-to-script-location.sh
##################################################

if [[ $# -ne 1 ]];
then
  echo "usage: ./08_copy-file-to-script-location.sh <file>"
  exit 1
fi

directory=backup
file=$1

cd $(dirname $0)

if [[ ! -d ${directory} ]];
then
  mkdir ${directory} || { echo "cannot create directory ${directory}"; exit 1; }
fi

echo $(pwd)

cp ${file} ${directory}/ || { echo "Cannot copy ${file} into ${directory}"; exit 1; }

