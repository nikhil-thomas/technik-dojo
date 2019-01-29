#################################################
# Author: Nikhil Thomas
# version: 1.0
# Date: Jan-29-2019
# Description: A script to launch firefox with set of pages
# Usage:Usage: 07_firefox-startup-pages.sh [url1] [url2] ...
#################################################

if [[ $# -lt 1 ]]; then
  firefox --new-window &
  exit
fi

command="firefox --new-window $1"
$(${command} &)

if [[ $# -gt 1 ]]; then
  command="firefox"

  for url in $@; do
    if [[ ${url} = $1 ]]; then
      continue;
    fi
    command="${command} --new-tab ${url}
  done
  #echo $command

  # add delay so that the new window is ready for the tabs to be added
  sleep 1
  $(${command} &)
fi

