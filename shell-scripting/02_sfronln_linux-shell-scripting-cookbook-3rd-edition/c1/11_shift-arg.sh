#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 31-DEC-2018
# Description: A script which uses shift command
# Usage: ./11_shift-arg.sh
##################################################

for i in `seq 1 $#`
do
  echo arg $i is $1
  shift
done

