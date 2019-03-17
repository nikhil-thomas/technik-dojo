#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-22
# Description : Script to read/print a file line by line
# Usage       : ./03_line-print.sh <filename>
##################################################

if [[ $# -ne 1 ]]; then
  echo "usage : ./03_line-print.sh <filename>"
  exit 1
fi

filename=$1

while read line; do
  echo ${line}
  sleep 0.25
done < ${filename}

