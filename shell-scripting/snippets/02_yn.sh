#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 26-DEC-2018
# Description: A script to provide y/n option to user 
# Usage: ./02_yn.sh
##################################################

while true;
do
  read -p "Do you wish to perform action (y/n):" yn
  case $yn in
      [Yy]* ) echo "perform action"; break;;
      [Nn]* ) exit;;
      * ) echo "please answer y or n";;
  esac
done
  
