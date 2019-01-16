#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-16
# Description : Script to check 'if file or if directory
# Usage       : ./02_if-elif.sh <name>
##################################################

if [ $# = 0 ]
then
  echo "usage ./02_if-elif.sh <name>"
  exit 1
fi

input=$1
cd $(dirname $0)

if [[ -f ${input} ]]
then
  cat ${input} || { echo "cannot print file : ${input}"; exit 1; }
elif [[ -d ${input} ]]
then
  ls -l ${input} || { echo "cannot list directory : ${input}"; exit 1; }
else
  echo "${input} is neither file nor directory"
  exit 1
fi

