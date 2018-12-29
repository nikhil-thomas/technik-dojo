#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 29-DEC-2018
# Description: A script to print fortunes
# Usage: ./03_fortune.sh
##################################################

while true
do
  clear
  lines=$(tput lines)
  #cols=$(tput cols)
  linesOffset=$(( ${lines}/3  ))
  colsOffset=0
  #colsOffset=$(( ${cols}/4  ))
  tput cup $linesOffset $colsOffset
  fortune
  read nextOrExit
  if [ "$nextOrExit" = 'exit' ];
  then
    clear
    tput cup 0 0
    break;
  fi
done

