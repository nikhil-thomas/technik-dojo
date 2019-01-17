#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-17
# Description : Script to print hello in a loop
# Usage       : ./02_if-nesting.sh <name>
##################################################

if [ $# = 0 ]
then
  echo "usage ./04_while.sh <iteretions>"
  exit 1
fi

itr=$1
count=0
while [[ ${count} -lt ${itr} ]]
do
  echo hello : ${count}
  sleep 0.5
  count=$((count+1))
done

