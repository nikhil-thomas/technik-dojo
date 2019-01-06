#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-06
# Description : Script to execute a command and return status code
# Usage       : source 01_exit-status.sh
#               ./01_exit-status.sh will return error
#               'return: can only `return' from a function or sourced script'
# Reference   : http://mywiki.wooledge.org/BashFAQ/055
##################################################

# initial exit code will be
# 0      : if a new wubshell is forked (./01_exit-status.sh)
# 0 or 1 : if script is run in the parent shell using source (source 01_exit-status.sh)
#          ie, exit status of command previously executed in the parent shell
exitStatus=$?
printf "%5s : %5s\n" "start" "$exitStatus"

# after successful command execution exit status ($?) will be 0
ls > /dev/null 2>&1
exitStatus=$?
printf "%5s : %5s\n" "exp 1" "$exitStatus"


# after unsuccessful command execution exit status ($?) will be 1
mkdir /home > /dev/null 2>&1
exitStatus=$?
printf "%5s : %5s\n" "exp 2" "$exitStatus"

# if the exit status is not returned the parent shell
# will not get the final exit status from this script
return $exitStatus

