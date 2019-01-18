#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-18
# Description : for loop usage
# Usage       : ./07_for-s.sh
##################################################

words="cat dog horse lion"

for word in ${words}
do
  echo "animal : ${word}"
done

for ((i=1; i<=5; i++))
do
  echo "val : ${i}"
done

for i in {1..5}
do
  echo "val : ${i}"
done

echo {1..5}
echo {a..z}
echo {A..z}

echo {1..100}
echo {1..100..10}
echo {0..100..10}

# for ((;;))
# do
#   echo hello
#   sleep 0.25
# done

for ((i=0;;i++))
do
  sleep 0.25
  echo ${i}
done

