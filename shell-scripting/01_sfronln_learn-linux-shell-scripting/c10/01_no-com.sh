#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-13
# Description : Script to print a file without comments and newlines
# Usage       : ./01_no-com.sh <filename> [comment character/keyword]
##################################################

if [ $# = 0 ]
then
  echo usage ./01_no-com.sh <filename> [comment character/keyword]
  exit 1
fi

commentKey=#
fileName=$1

if [ -n "$2" ]
then
  commentKey=$2
fi

commentRegEx=\'^${commentKey}\'

# use echo and bash to use '' in command string
echo "grep -v -e ${commentRegEx} -e ^$ ${fileName}" | bash

