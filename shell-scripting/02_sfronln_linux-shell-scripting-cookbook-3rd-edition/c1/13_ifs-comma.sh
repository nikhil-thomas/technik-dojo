#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 01-JAN-2019
# Description: A script which iteratest over comma separated data 
# Usage: ./13_ifs-comma.sh
##################################################

data="name gender rollnumber location"
for item in $data;
do
  echo item: $item
done

data="name,gender,rollnumber,location"
for item in $data;
do
  echo item: $item
done

oldIFS=$IFS
IFS=,

for item in $data;
do
  echo item: $item
done
IFS=$oldIFS

