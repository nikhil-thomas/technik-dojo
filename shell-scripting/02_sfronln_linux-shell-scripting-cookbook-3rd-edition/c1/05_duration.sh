#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 29-DEC-2018
# Description: A script which prints time taken for its execution
# Usage: ./05_duration.sh
##################################################

p () {
  sleep 0.5
  printf "%1s" $1
}

greeting="Hello, World!"

start=$(date +%s)
printf "%-15s : script started\n" $start

while read -n 1 c
do
  p "$c"
done <<< "$greeting"
printf "\n"

end=$(date +%s)
printf "%-15s : script ended\n" $end
d=$(($end - $start))
printf "\n%-15s : ${d}s\n" 'duration'

