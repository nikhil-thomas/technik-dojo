#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 01-JAN-2019
# Description: A script runs a command until it succeeds
# Usage: ./12_try-until-success.sh
##################################################

function repeat() {

# 'true' is a binary in most systems
# while true will spawn a process in each iteration
# while true
# it can be avoided by using
# while :
# ':' always returns exit code 0
while :
  do
    $@ && return
    sleep 1
  done
}

repeat $@

