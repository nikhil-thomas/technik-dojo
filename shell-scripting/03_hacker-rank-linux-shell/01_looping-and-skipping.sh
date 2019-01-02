#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 02-JAN-2019
# Description: A script to loop and skip even numbers
# Usage: ./01_looping-and-skipping.sh
##################################################

for i in {1..99}
do
  if [ $(( $i%2 )) !=  0 ]
  then
    echo $i
  fi
done

