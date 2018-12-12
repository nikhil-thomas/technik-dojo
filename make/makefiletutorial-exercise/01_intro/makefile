# The default target is some_binary, because it is first
some_binary:
	echo "hello world"

# This file will make file1 the first time,
# and the second time notice it’s already made,
# resulting in make: 'file1' is up to date.
file1:
	touch file1

# Alternative syntax: same line
file2: ; touch file2

# \ gives us multilines
file3: ; touch \
	file3


# Will call other.txt target if it is newer than the some_binary file,
# or it doesn’t exist.
# It will call the other.txt rule first.
file4: other.txt
	touch file4

other.txt:
	echo "other.txt target" 
	touch other.txt

# This will always make both targets,
# because some_binary depends on other.txt,
# which is never created.
file5: task1
	touch file5

task1:
	echo "task1 target"

# “clean” is not a special word.
# If there’s a file called clean that is made,
# then “make clean” won’t have to do anything.
# Similarly, if the clean file is older than the some_binary file,
# the clean rule will not be called.
clean:
	touch clean

actual_clean:
	rm some_binary file1 file2 file3 file4 other.txt file5
	rm clean

# Adding PHONY to a target will prevent make from confusing the phony target with a file name.
# In this example, if clean is created, make clean will still be run.
.PHONY: clean
clean:
	rm some_binary file1 file2 file3 file4 other.txt file5
	rm clean

