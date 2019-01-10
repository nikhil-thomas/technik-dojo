#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-10
# Description : Script to enusre that a file exists, and it is empty
# Usage       : ./06_empty-file.sh
##################################################

if [[ $# -ne 1 ]];
then
  echo "usage: ./06_empty-file.sh <file>"
  exit 1
fi

file=$1

if [[ -f ${file} ]];
then
  cp /dev/null ${file}
else
  touch ${file}
fi

if [[ $? -ne 0 ]];
then
  echo "something went wrong"
  echo "please check ${file}"
  exit 1
else
  echo "${file} is now empty"
fi

