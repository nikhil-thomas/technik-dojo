#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 03-JAN-2019
# Description: A script to output YES for y/Y and NO for n/N
# Usage: ./01_looping-and-skipping.sh
##################################################

read yn
if [[ $yn = "Y" || $yn = "y"  ]]
then
  echo YES
elif [[ $yn = "N" || $yn = "n"  ]]
then
  echo NO
fi

