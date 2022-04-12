### Unit Test is very important but first [the code you wrote should be testable](https://go.dev/blog/examples)

---

Test all _test.go file from the current package (ctx directory)
```
    go test
```
Test for all _test.go file 
```
    go test -v ./...
```

Run specific Test Func by specify the test func's name
```
    go test -run=TestTwoBigFloat
```
a
```
    go test -tags=integration
    go test -cover
```

[Markdown-Cheatsheet](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet)
