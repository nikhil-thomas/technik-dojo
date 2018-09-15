## Passinng a slice to a function

```go
func main() {
    nums := []int{1, 2, 3}
    add10(nums)
    fmt.Println(num)
    // output 11, 12, 13

    nums = []int{1, 2, 3}
    add3Items(nums)
    fmt.Println(num)
    // output
    // 1, 2, 3
}
```
#### case 1
```go
func add10(vals []int) {
    for i := 0; i < len(vals); i++ {
        vals[i] += 10
    }
}
```

slice type is passwed by value to add10.
However, within the function slice vals refers 
to the same underlying array

therefore any update on slice without changing the underlying array 
will be reflected outside the function. Hence, add10 will have 
its effect persisted even after add10 is completes execution.

#### case 2
```go
func add3Items(vals []int) {
	for i := 100; i <= 300; i += 100 {
		vals = append(vals, i)
	}
}
```
If the underlying array has to change (recreated/replaced) within the function the changes will not be available outside the function. This is because outside the function the slice will be still refering to the old array as the underlying type.

In other words, withing the function if cap or len changes, if the underlying array has to expand to accomodate a new value the change will not be reflected.
