### Build Steps

#### build go code as a c-shared-library
```
go build -o usedByC.o -buildmode=c-shared main.go
```

#### compile C code
```
gcc -o willUseGo willUseGo.c ./usedByC.o
```