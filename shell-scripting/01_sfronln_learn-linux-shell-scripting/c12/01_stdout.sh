#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-20
# Description : Script to redirect stdout to file
# Usage       : ./01_square.sh <number>
##################################################

cd $(dirname $0)

read -p "Type : " input

echo ${input} >> redirect

