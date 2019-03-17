#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-MAR-17
# Description : function to pring cpu mem and disk
# Usage       : ./02_cpu-mem-disk.sh
##################################################

print_cpu_mem_disk() {
  date
  echo "CPU is use: $(top -bn1 | grep Cpu | awk '{print $2}')"
  echo "Memory in use: $(free -h | grep Mem | awk '{print $3}')"
  echo "Disk space available on /: $(df -kh / | grep / | awk '{print $4}')"
  echo #
}

for ((i=0; i<5; i++)); do
  print_cpu_mem_disk
  sleep 2
done

