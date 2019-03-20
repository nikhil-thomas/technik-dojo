#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-MAR-20
# Description : function to reverse a string
# Usage       : .06-reverse-func.sh
##################################################

if [[ $# -eq 0 ]]; then
  echo "incorrect number of arguments"
  echo "usage: ${0} <inout-string>"
  exit 1
fi

input="_${1}_"

reverser() {
  if [[ $# -eq 0 ]]; then
    echo "no arguments givenin reverser"
    exit 1
  fi

  rev <<< ${1}
}

echo "input: ${1}"
reversed=$(reverser ${input})
echo "reversed: ${reversed}"

