#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 25-DEC-2018
# Description: A script to print Christmas Greetings
# Usage: ./01_printf.sh
##################################################

name=$1

if test -z ${name};
then
  name=All
fi

echo Merry Christmas ${name}!

