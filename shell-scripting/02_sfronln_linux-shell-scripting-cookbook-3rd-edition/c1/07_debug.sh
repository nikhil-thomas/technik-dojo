#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 30-DEC-2018
# Description: A script which set and unset debug option within
# Usage: ./07_debug.sh
##################################################

for i in {1..6}
do
  # Set debug on
  set -x
  echo $i

  # Set debug off
  set +x
  sleep 0.5
done
printf "\n"

