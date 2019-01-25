#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-25
# Description : A script to set a new passsword for a user
# Usage       : ./06_set-password.sh <length>
##################################################

if [[ $(id -u) -ne 0 ]]; then
  echo "please run as root user"
  exit 1
fi

if [[ $# -ne 1 ]]; then
  echo "usage : ./06_set-password.sh <username>"
  exit 1
fi

username=$1

id ${username} &> /dev/null

if [[ $? -ne 0  ]]; then
  useradd -m ${username}
fi

password=$(tr -cd '[A-Za-z0-9]' < /dev/urandom | head -c 10)
echo "${username}:${password}" | chpasswd

echo ${password}

