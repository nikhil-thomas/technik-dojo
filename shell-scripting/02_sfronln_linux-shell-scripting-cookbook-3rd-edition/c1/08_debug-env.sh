#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 30-DEC-2018
# Description: A script which print debug based on _DEBUG=on env var
# Usage: ./08_debug-env.sh
##################################################

function DEBUG()
{
  [ "$_DEBUG" == "on"  ] && $@ || :
}

for i in {1..10}
do
  DEBUG echo "I : ${i}"
done

