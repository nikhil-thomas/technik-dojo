#!/bin/bash
# inpath--Verifies whether a specified program is either valid as is
# or can be found in the PATH directory

in_path()
{
	# For a given command and path
	# return 0 if found executable, 1 if not
	#
	# Note: this script modifies IFS (internal field separator) temporarily
	# but restores IFS upon completion

	cmd=$1
	ourpath=$2
	result=1
	oldIFS=$IFS
	IFS=":"

  	for directory in ${ourpath}
  	do
		echo ${directory}/${cmd}
		if [ -x "$directory/$cmd" ] ; then
			result=0
		fi
	done
	IFS=$oldIFS
	return $result
}

checkForCmdInPath()
{
	var=$1

	if [ "$var" != ""  ]; then
		if [ "${var:0:1}" = "/"  ] ; then
			if [ ! -x $var  ] ; then
				return 1
			fi
		elif ! in_path $var "$PATH" ; then
			return 2
		fi
	fi
}

if [ $# -ne 1  ] ; then
	echo "Usage: $0 command" >&2
	exit 1
fi

checkForCmdInPath "$1"
case "$?" in
	0 ) echo "$1 found in PATH"	        ;;
	1 ) echo "$1 not found or executable"	;;
	2 ) echo "$1 not found in PATH"		;;
esac

exit 0


