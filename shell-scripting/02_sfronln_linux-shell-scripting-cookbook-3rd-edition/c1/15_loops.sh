#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 01-JAN-2019
# Description: A script which  
# Usage: ./15_loops.sh
##################################################

count=0

echo while
# while loops as long as the condition remains true
while [ $count -ne 10 ];
do
  echo $count
  let count++
done

echo -e "\nuntil"
# untli loops till condition becomes true
until [ $count -eq 0 ];
do
 echo $count
 let count--
done

