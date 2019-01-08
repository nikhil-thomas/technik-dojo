#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-08
# Description : Script to show use of variables in tests
# Usage       : ./03_test-var.sh
##################################################

DIRECTORY=/tmp/
if [ -n "$1" ];
then
  DIRECTORY=$1
fi
echo $DIRECTORY

# Test if /tmp directory exists using full command
test -d ${DIRECTORY}
test_rc=$?

# Test using shorthand
[ -d ${DIRECTORY} ]
simple_rc=$?

# Test using extended shorthand
[[ -d ${DIRECTORY} ]]
extended_rc=$?

printf "%-9s : %2s\n" 'test' $test_rc
printf "%-9s : %2s\n" 'simple' $simple_rc
printf "%-9s : %2s\n" 'extended' $extended_rc

