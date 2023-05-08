### Errors
Go programs express error state with `error` values.
The error type is a built-in interface similar to `fmt.Stringer`:
```
type error interface {
    Error() string
}
```