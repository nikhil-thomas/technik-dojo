# Alias commands

1. clear screen and offset prompt by (num terminal lines)/4
   ```
   alias csr='clear; tput cup $(($(tput lines)/4 ))'
   ```

