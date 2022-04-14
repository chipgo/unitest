### Unit Test is very important but first [the code you wrote should be testable](https://go.dev/blog/examples)

---

Test all _test.go, from current package (ctx directory)
```
    go test
```
Test for all _test.go file
```
    go test -v ./...
```
Benchmark all _test.go, from current package (ctx directory)
```
    go test -bench=.
```
```
    go test -bench .
```
Run specific Test Func
```
    go test -run=TestTwoBigFloat
```
Write a coverage profile to the file after all tests have passed
```
    go test -coverprofile cover.out .
```
Then open it
```
    go tool cover -html=cover.out
```
a
```
    go test -tags=integration
    go test -cover
```

[Test Coverage](https://go.dev/blog/cover)

[Markdown-Cheatsheet](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet)
go test -coverprofile=coverage.out ; go tool cover -func=coverage.out