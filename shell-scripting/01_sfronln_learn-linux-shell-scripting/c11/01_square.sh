#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-15
# Description : Script to square a number
# Usage       : ./01_square.sh <number>
##################################################

if [ $# = 0 ]
then
  echo "usage ./01_square.sh <number>"
  exit 1
fi

input=$1

# Check whether the input is number
# =~ operator checks a value against a regular expression
if [[ ! ${input} =~ ^[[:digit:]]+$ ]]
then
  echo wrong type of input
  exit 1
fi

square=$(( ${input} * ${input} ))
echo ${square}

