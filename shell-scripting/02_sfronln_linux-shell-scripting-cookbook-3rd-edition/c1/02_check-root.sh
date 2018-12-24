#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 24-DEC-2018
# Description: A script to check whether the user us root
# Usage: ./02_check-root.sh
##################################################

if [ $UID -ne 0  ];
then
  echo "Non root user. Please run as root."
else
  echo "Root user"
fi

