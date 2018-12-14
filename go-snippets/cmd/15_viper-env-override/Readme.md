#### Case 1

```
go run main.go
```
output

viper test
V_VAL_1:
V_VAL_2: val set from config.yaml

#### Case 2
```
source .env
export V_VAL_2
go run main.go

```
output

V_VAL_1:
V_VAL_2: val set from .env

