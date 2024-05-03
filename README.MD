# poc json



- test with message 
```go
go test -v ./...
```

- result 

```bash
$ go test -v ./...
?       poc_json        [no test files]
?       poc_json/configs        [no test files]
?       poc_json/dto    [no test files]
?       poc_json/handlers       [no test files]
?       poc_json/models [no test files]
?       poc_json/services       [no test files]
?       poc_json/utils  [no test files]
=== RUN   TestHelperCreateDocumentVersion
=== RUN   TestHelperCreateDocumentVersion/success
    doc_repo_test.go:50: 2024-05-03 14:18:35.6198498 +0700 +07
    doc_repo_test.go:51: Success case passed
=== RUN   TestHelperCreateDocumentVersion/error
--- PASS: TestHelperCreateDocumentVersion (0.00s)
    --- PASS: TestHelperCreateDocumentVersion/success (0.00s)
    --- PASS: TestHelperCreateDocumentVersion/error (0.00s)
PASS
ok      poc_json/repositories   0.049s
```