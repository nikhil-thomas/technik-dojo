#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 04-JAN-2019
# Description: A script to compare two numbers
# Usage: ./04_compare-numbers.sh
##################################################

read x
read y

dxy=$(( $x - $y ))

if [ $dxy -eq 0 ];
then
  # print 'X' not value ($x) as required
  # in the hacker rank challenge
  echo X is equal to Y
elif [ $dxy -gt 0 ];
then
  echo X is greater than Y
else
 echo X is less than Y
fi

