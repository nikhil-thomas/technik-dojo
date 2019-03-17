#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-26
# Description : A riddle using while loop
# Usage       : ./07-1_while-riddle.sh
##################################################

while true
do
  read -p "keys but no locks, space but no room, enter but no exit :" ans
  if [[ ! ${ans} =~ [Kk]eyboard ]]
  then
    echo "wrong answer try again"
  else
    break
  fi
done

echo "right answer"

