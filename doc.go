/*
Package prep instruments instruments db connection with prepared statements.
It allows you to benefit from the prepared SQL statements almost without any changes to your code.

Prep consists of two parts:

1. A command line tool that finds all SQL statements in your code

2. A package that instruments your code with prepared SQL statements using the found ones

Firstly you should generate a list of SQL statements used in your application

	$ cat github.com/hexdigest/prepdemo/example.go
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

Let's generate a list of the SQL statements used in your package:

	$ prep -f github.com/hexdigest/prepdemo
	$ cat prepared_statements.go
	//go:generate prep -f github.com/hexdigest/prepdemo

	package main

	var prepStatements = []string{
		"SELECT CONCAT(\"Hello \", ?, \"!\")",
	}

Using prepared statements:

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

Take a look at the line:
	db, err := prep.NewConnection(sqlDB, prepStatements)

It instruments your connection with prepared statements found by the generator.
The generated code already contains //go:generate instruction, so in order to update the statements list you can simply run:

	$ go generate

Some synthetic benchmarks:

	$ go test -bench=.
	BenchmarkPostgresWithoutPreparedStatements-4   	   20000	     59941 ns/op	    1183 B/op	      32 allocs/op
	BenchmarkPostgresWithPreparedStatements-4      	   50000	     41560 ns/op	    1021 B/op	      26 allocs/op
	BenchmarkMySQLWithoutPreparedStatements-4      	   50000	     26454 ns/op	     827 B/op	      23 allocs/op
	BenchmarkMySQLWithPreparedStatements-4         	  200000	      9509 ns/op	     634 B/op	      19 allocs/op
	PASS
	ok  	github.com/hexdigest/prep	7.884s
*/
package prep
