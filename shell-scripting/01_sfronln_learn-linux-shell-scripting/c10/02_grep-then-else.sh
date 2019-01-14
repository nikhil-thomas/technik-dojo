#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-14
# Description : Script to delete a file if it contains 'keyword', not append 'keyword to file
# Usage       : ./02_grep-then-else.sh
##################################################

FILE_NAME="/tmp/grep-then-else"

touch ${FILE_NAME}
checkPattern='J.n'

grep -q ${checkPattern} ${FILE_NAME}
grep_rc=$?

if [[ ${grep_rc} -eq 0 ]]
then
  rm ${FILE_NAME}
else
  echo "${checkPattern}" >> ${FILE_NAME}
fi

