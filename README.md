##### go-lang
Learning GoLang

##### 单元测试
> 参考：https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/11.3.md

Run `go test -v`

##### 压力测试
* Run `go test -test.bench=".*"`

##### Usage
* Run `go mod init go-lang` to create a go.mod file in the root of the project
* Run `go run basic/book.go` to run the book.go file where the main function is located
* Run `go get github.com/go-sql-driver/mysql` to install the mysql driver or any other package

##### More to do
- [ ] Add `GORM` as the ORM to operate on the database
- [ ] Add `gin` as the web framework to handle the web requests
- [ ] Add `golang-jwt` as the JWT middleware to handle the authentication