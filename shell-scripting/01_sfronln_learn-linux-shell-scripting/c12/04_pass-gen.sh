#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-23
# Description : Script to generate a password of given length
#               with given length
# Usage       : ./04_pass-gen.sh <length>
##################################################

if [[ $# -ne 1 ]]; then
  echo "usage : ./04_pass-gen.sh <length>"
  exit 1
fi

length=$1

if [[ ! ${length} =~ ^[[:digit:]]+$ ]]; then
  echo "invalid input : ${length}"
  echo "usage : ./04_pass-gen.sh <length>"
  exit 1
fi

password=$(tr -dc '[A-Za-z0-1]' < /dev/urandom | head -c ${length})
printf "\npassword : %s\n\n" ${password}

