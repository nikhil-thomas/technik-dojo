#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-MAR-18
# Description : local variables and global variables 
# Usage       : ./03_fn-var.sh
##################################################

if [[ $# -eq 0 ]]; then
  echo "Missing argumet"
  echo "usage: ${0} <input>"
  exit 1
fi

input_variable=$1
CONSTANT_VARIABLE="constant"

hello_variable() {
  local CONSTANT_VARIABLE="may not be constant"
  var2=123
  echo "input: ${input_variable}"
  echo "constant: ${CONSTANT_VARIABLE}"
  echo "new var: ${var2}"
}

hello_variable

echo "CONSANT: ${CONSTANT_VARIABLE}"
echo "new variable: ${var2}"

