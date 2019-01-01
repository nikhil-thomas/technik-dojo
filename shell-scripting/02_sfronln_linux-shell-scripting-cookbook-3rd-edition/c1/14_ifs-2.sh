#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 01-JAN-2019
# Description: A script which uses custom IFS ':' 
# Usage: ./14_ifs-2.sh
##################################################

data=root:x:0:0:root:/root:/bin/bash
echo Data: $data
oldIFS=$IFS
IFS=":"
count=0
for item in $data;
do
  [ $count -eq 0 ] && user=$item
  [ $count -eq 6 ] && shell=$item
  let count++
done;
IFS=$oldIFS

echo ${user}\'s shell is $shell;

