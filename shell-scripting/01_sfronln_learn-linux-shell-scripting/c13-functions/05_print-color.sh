#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-MAR-19
# Description : function to to print colored text 
# Usage       : ./05_print-color.sh
##################################################

print_color() {
  if [[ $# -ne 2 ]]; then
    echo "print_color function needs two arguments"
    exit 1
  fi

  local string=$1
  local color=$2

  case ${color} in
  red)
    local color_code="\e[31m";;
  blue)
    local color_code="\e[34m";;
  green)
    local color_code="\e[32m";;  
  *)
    local color_code="\e[39m";;  
  esac

  echo -e ${color_code}${string}"\e[39m"
}

print_color "Hello world!" "red"
print_color "Hello world!" "blue"
print_color "Hello world!" "green"
print_color "Hello world!" "magenta"

