#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 2018-DEC-23
# Description: Script to prompt input from user if commandline arguments are not supplied
# Usage: ./05_interactive-2 <name> <location> <food>
##################################################

name=$1
location=$2
food=$3

# Prompt the user for input
if test -z ${name}; then read -p "Enter Name: " name; fi
if test -z ${location}; then read -p "Enter Location: " location; fi
if test -z ${food}; then read -p "Enter food: " food; fi

# Print sentence
echo ${name} eats ${food} from ${location}

