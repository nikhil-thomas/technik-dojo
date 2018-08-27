#!/bin/bash
# validApphaNum--Ensures that input contains
# only alphabetical and numeric characters

validAlphaNum()
{
  # Validate arg: returns 0 if all upper+lower+digits; 1 otherwise

  # Remove all unacceptable characters
  validchars="$(echo $1 | sed -e 's/[^[:alnum:]]//g')"

  if [ "$validchars" = "$1"  ] ; then
    return 0
  else
    return 1
  fi
}

echo -n "Enter input: "
read input

# input validation
if ! validAlphaNum "$input" ; then
  echo "Please enter only letters and numbers." >&2
  exit 1
else
  echo "Input is valid."
fi

exit 0

