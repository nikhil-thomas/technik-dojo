#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 31-DEC-2018
# Description: A script to check the status of a command
# Usage: ./10_cmd-status.sh
##################################################

CMD=$@
eval $CMD

if [ $? -eq 0  ];
then
  echo "$CMD" : executed successfully
else
  echo "$CMD" : terminated unsuccessfully
fi

