#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-24
# Description : Script to loop until password of length less than
#               given length is found
#               with given length
# Usage       : ./05_pass-gen-len-loop.sh <length>
##################################################

if [[ $# -ne 1 ]]; then
  echo "usage : ./05_pass-gen-len-loop.sh <length>"
  exit 1
fi

length=$1

if [[ ! ${length} =~ ^[[:digit:]]+$ ]]; then
  echo "invalid input : ${length}"
  echo "usage : ./04_pass-gen.sh <length>"
  exit 1
fi

while true; do
  password=$(head -n 4 /dev/urandom | tr -dc '[A-Za-z0-1]' | head -c ${length})
  printf "\n length: %3s : password : %s\n\n" ${#password} ${password}
  if [[ ${#password} -lt ${length} ]]; then
    echo "short password : exiting"
    break
  fi
done

