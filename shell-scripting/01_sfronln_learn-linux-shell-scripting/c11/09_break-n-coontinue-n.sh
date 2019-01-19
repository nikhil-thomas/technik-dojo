#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-19
# Description : A script to demo 'continue n' and 'break n'
# Usage       : ./09_break-n-coontinue-n.sh
##################################################

for i in {1..3}; do
  #echo "l1 : $i"
  for j in {1..3}; do
   # echo "l2 : $j"
    for k in {1..3}; do
    #  echo "l3 : $k"
      if [[ $k -eq 2 ]]; then
        break 2
      fi
    done
  done
done

for i in {1..3}; do
  echo "l1 : $i"
  for j in {1..3}; do
    echo "l2 : $j"
    for k in {1..3}; do
      echo "l3 : $k"
      if [[ $k -eq 2 ]]; then
        continue 2
      fi
    done
  done
done

