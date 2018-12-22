#!/bin/bash

##################################################
# Author: Nikhil Thomas
# Version: v1.0.0
# Date: 2018-DEC-22
# Description: Script to prompt input from user and print inputs
# Usage: ./04_interactive
##################################################

# Prompt the user for input
read -p "Enter name: " name
read -p "Enter location: " location
read -p "Enter food: " food

# Print sentence
echo ${name} eats ${food} from ${location}

