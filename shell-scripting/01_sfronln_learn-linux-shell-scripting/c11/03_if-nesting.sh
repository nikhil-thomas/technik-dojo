#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-16
# Description : Script to check 'if file or if directory
# Usage       : ./02_if-nesting.sh <name>
##################################################

if [ $# = 0 ]
then
  echo "usage ./02_if-elif.sh <name>"
  exit 1
fi

input=$1
cd $(dirname $0)

if [[ -r ${input} ]]
then
  if [[ -f ${input} ]]
  then
    cat ${input}
  elif [[ -d ${input} ]]
  then
    ls -l ${input}
  else
    echo "${input} is neither file nor directory"
    exit 1
  fi
else
  echo "cannot read file/directory : ${input}"
  exit 1
fi

