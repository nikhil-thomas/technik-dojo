#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-20
# Description : Script to find a file (discard all error messages)
# Usage       : ./02_find-file.sh <filename>
##################################################

if [[ $# -ne 1 ]]; then
  echo "usage : ./02_find-file <filename>"
  exit 1
fi

filename=$1
find / -name ${filename} 2> /dev/null

