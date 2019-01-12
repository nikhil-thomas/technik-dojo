#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-12
# Description : Script to take yes/no input
# Usage       : ./09_yes-no.sh
##################################################

read -p 'Yes/No: ' yesNo

if [[ ${yesNo,,} = 'y' || ${yesNo,,} = 'yes' ]]
then
  echo Yes
  exit 0
fi

if [[ ${yesNo^^} = 'N' || ${yesNo^^} = 'NO' ]]
then
  echo No
  exit 0
fi

echo "Invalid input : ${yesNo}"
exit 1

