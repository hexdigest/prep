# Prep [![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://github.com/hexdigest/prep/blob/master/LICENSE) [![Build Status](https://travis-ci.org/hexdigest/prep.svg?branch=master)](https://travis-ci.org/hexdigest/prep) [![Coverage Status](https://coveralls.io/repos/github/hexdigest/prep/badge.svg?branch=master)](https://coveralls.io/github/hexdigest/prep?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/hexdigest/prep)](https://goreportcard.com/report/github.com/hexdigest/prep)
Prep finds all SQL statements in a Go package and instruments db connection with prepared statements.
It allows you to benefit from the prepared SQL statements almost without any changes to your code.

Prep consists of two parts:
* A command line tool that finds all SQL statements in your code
* A package that instruments your code with prepared SQL statements using the found ones

# Usage

## Generate a list of SQL statements used in your application

$ cat example.go
```go
func main() {
	db, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/mysql")
	if err != nil {
		panic(err)
	}
	const query = `SELECT CONCAT("Hello ", ?, "!")`
	var s string
	if err := db.QueryRow(query, "World").Scan(&s); err != nil {
		panic(err)
	}
	fmt.Println(s)
}
```

Let's generate a list of the SQL statements used in your package:

```bash
$ prep -f github.com/hexdigest/prepdemo
$ cat prepared_statements.go
```

```go
//go:generate prep -f github.com/hexdigest/prepdemo

package main

var prepStatements = []string{
	"SELECT CONCAT(\"Hello \", ?, \"!\")",
}
```

# Using prepared statements

```go
func main() {
	sqlDB, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/mysql")
	if err != nil {
		panic(err)
	}

	db, err := prep.NewConnection(sqlDB, prepStatements)
	if err != nil {
		panic(err)
	}

	const query = `SELECT CONCAT("Hello ", ?, "!")`
	var s string
	if err := db.QueryRow(query, "World").Scan(&s); err != nil {
		panic(err)
	}
	fmt.Println(s)
}
```

Take a look at the line:
```go
db, err := prep.NewConnection(sqlDB, prepStatements)
```

It instruments your connection with prepared statements found by the generator.
The generated code already contains //go:generate instruction, so in order to update the statements list you can simply run:

```bash
$ go generate
```

# Some synthetic benchmarks

```bash
$ go test -bench=.
BenchmarkPostgresWithoutPreparedStatements-4   	   20000	     59941 ns/op	    1183 B/op	      32 allocs/op
BenchmarkPostgresWithPreparedStatements-4      	   50000	     41560 ns/op	    1021 B/op	      26 allocs/op
BenchmarkMySQLWithoutPreparedStatements-4      	   50000	     26454 ns/op	     827 B/op	      23 allocs/op
BenchmarkMySQLWithPreparedStatements-4         	  200000	      9509 ns/op	     634 B/op	      19 allocs/op
PASS
ok  	github.com/hexdigest/prep	7.884s
```
