#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-MAR-18
# Description : error handling functions 
# Usage       : ./04_err-fn.sh
##################################################

handle_minor_error() {
  echo "A minor error has occured, please check the output."
}

handle_fatal_error() {
  echo "A critical error has occured, stopping script."
  exit 1
}

ls -l /tmp/ || handle_minor_error
ls -l /root/ || handle_minor_error

cat /etc/shadow || handle_fatal_error
cat /etc/passwd || handle_fatal_error

