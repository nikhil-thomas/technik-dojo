#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 03-JAN-2019
# Description: A script to output YES for y/Y and NO for n/N
# Usage: ./01_looping-and-skipping.sh
##################################################

read x
read y
read z

dxy=$(( $x - $y ))
dyz=$(( $y - $z ))
dzx=$(( $z - $x ))

if [ $x = $y ] && [ $y = $z ]
then
  echo EQUILATERAL
elif [ $dxy = 0 ] || [ $dyz = 0 ] || [ $dzx = 0 ];
then
  echo ISOSCELES

else
 echo SCALENE
fi

