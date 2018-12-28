#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 28-DEC-2018
# Description: A script which disables echo while reading password from stdin
# Usage: ./04_disable-password-echo.sh
##################################################

echo -e "Enter password: "
# disable echo before reading password
stty -echo
read password
# enable echo after reading password
stty echo
echo
echo password read

