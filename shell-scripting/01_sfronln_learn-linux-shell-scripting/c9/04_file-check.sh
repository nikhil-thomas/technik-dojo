#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-09
# Description : Script to check if a file exists
# Usage       : ./04_file-check.sh
##################################################

FILE=/tmp/random_file.txt

if [ ! -f ${FILE} ];
then
  echo "File does not exist"
  exit 1
fi

cat ${FILE}

