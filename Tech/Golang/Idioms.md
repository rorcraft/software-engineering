### Idioms

__check err in `if` condition__

```golang
  if num, err := strconv.ParseInt("123"); err == nil {
    fmt.Println(num == 123)
  }
```
