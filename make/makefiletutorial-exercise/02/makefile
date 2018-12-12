# Variables can only be strings
files = file1 file2
some_binary: $(files)
	echo "files variable: " $(files)
	touch some_binary

file1:
	touch file1

file2:
	touch file2

# If we have a target that is a “.c” file,
# there is an implicit command that will be “cc -c file.c -o file.o”
hello.o:

# print $ sign
dollar:
	echo $$

# We can use wildcards in the target, prerequisites, or commands.
# Valid wildcards are *, ?, [...]
task1: *.c
	echo "task1"
*.c:
	touch f1.c
	touch f2.c

all: one two three
.PHONY: all

one:
	touch one

two:
	touch two

three:
	touch three

clean:
	rm hello.o
	rm file1 file2 some_binary
	rm *.c
	rm one two three

