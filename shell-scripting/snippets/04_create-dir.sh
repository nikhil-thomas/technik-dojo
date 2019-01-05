#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 05-JAN-2019
# Description: A script to create a directory only if it does not exist
# Usage: ./04_create-dir.sh/
##################################################

read -p "Enter directory name: " dir
if [ -d $dir ];
then
  echo "Directory exists"
else
  mkdir $dir
  echo $dir - directory created
fi

