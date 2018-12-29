#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 29-DEC-2018
# Description: A script which counts down and cretes a delay
# Usage: ./06_delay.sh
##################################################

# Store current cursor position
tput sc

# Loop for 10 seconds
for count in `seq 0 10`
do
  # Restore cursor position
  tput rc
  # Clear to the end of the line
  tput ed
  echo -n $count
  sleep 1
done
printf "\n"

