#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-19
# Description : A script to demo 'case' statement
# Usage       : ./10_case.sh
##################################################

for i in {1..110}; do
  echo ${i}
  sleep 0.25
  case ${i} in
  [1-6]*)
    echo "case [1-6]*"
  ;;
  [7-8]*)
    echo 'case [7-8]*'
  ;;
  9[1-8])
    echo 'case 9[1-8]'
  ;;
  99)
    echo 'case 99'
  ;;
  *)
    echo 'case *'
  ;;
esac
done

