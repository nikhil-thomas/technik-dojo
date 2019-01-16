#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: JAN-16-2019
# Description: A script to print user information
# Usage: ./05_userinfo.sh/
##################################################

while IFS=: read -r f1 f2 f3 f4 f5 f6 f7
do
  echo "${f1} uses ${f7} shell"
  sleep 0.5
done < /etc/passwd

