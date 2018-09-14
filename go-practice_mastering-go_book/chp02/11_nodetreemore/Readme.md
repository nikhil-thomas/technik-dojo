## To see go node tree

#### full node tree
> go tool compile -W main.go

#### functionOne in node tree
> go tool compile -W main.go | grep functionOne | uniq
