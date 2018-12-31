#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 31-DEC-2018
# Description: A script which defines and uses a function
# Usage: ./09_func.sh
##################################################

function fname() {
  echo $1, $2, $3
  echo $0
  echo "$@"
  echo "$*"

  echo '$*'
  for c in "$*"
  do
    echo $c
  done

  echo '$@'
  for c in "$@"
  do
    echo $c
  done


  return 0
}

fname $1 $2 $3

