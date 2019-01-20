#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 2019-JAN-20
# Description: A script to print a random number between min and max
# Usage: ./06_rand.sh [max=100] [min=0]
##################################################

min=0
max=100

if [[ $# -eq 2 ]]; then
  min=$2
  max=$1
elif [[ $# -eq 1 ]]; then
  max=$1
fi

div=$(( $max - $min ))
echo $(( $min + $(( $RANDOM % $div )) + 1 ))

