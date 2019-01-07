#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-07
# Description : Script to show functional checks
# Usage       : ./02_functional-check.sh
##################################################

mkdir /tmp/temp-dir
status1=$?
test -d /tmp/temp-dir
status2=$?
echo "status1: $status1"
echo "status2: $status2"
printf "\n"

# [ ] and [[ ]] are short verisons of test command
# if there are whitespaces in inout [ ] interprets as delimiter
# hence, it is safer to use [[ ]]
[ -d /tmp/temp-dir ]
echo $?
[[ -d /tmp/temp-dir ]]
echo $?

