#!/bin/bash

##################################################
# Author      : Nikhil Thomas
# Version     : v1.0.0
# Date        : 2019-JAN-26
# Description : A script supplying input using heredoc
# Usage       : ./07-2_heredoc-input.sh
##################################################

cd $(dirname $0)

bash ./07-1_while-riddle.sh << EOF
mouse # Try 1
monitor # Try 2
keyboard # Try 3 (right answer)
printer
EOF

