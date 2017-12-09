package prep

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

/*
$ go test -bench=.
BenchmarkPostgresWithoutPreparedStatements-4   	   20000	     59941 ns/op	    1183 B/op	      32 allocs/op
BenchmarkPostgresWithPreparedStatements-4      	   50000	     41560 ns/op	    1021 B/op	      26 allocs/op
BenchmarkMySQLWithoutPreparedStatements-4      	   50000	     26454 ns/op	     827 B/op	      23 allocs/op
BenchmarkMySQLWithPreparedStatements-4         	  200000	      9509 ns/op	     634 B/op	      19 allocs/op
PASS
ok  	github.com/hexdigest/prep	7.884s
*/

func BenchmarkPostgresWithoutPreparedStatements(b *testing.B) {
	b.ReportAllocs()

	db, err := sql.Open("postgres", "postgres://pg:pg@127.0.0.1:5432/pg?sslmode=disable")
	if err != nil {
		b.Fatal(err)
	}

	const query = "SELECT $1::text"
	var s string

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := db.QueryRow(query, "1").Scan(&s); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkPostgresWithPreparedStatements(b *testing.B) {
	b.ReportAllocs()

	var (
		db  Connector
		err error
	)

	db, err = sql.Open("postgres", "postgres://pg:pg@127.0.0.1:5432/pg?sslmode=disable")
	if err != nil {
		b.Fatal(err)
	}

	const query = "SELECT $1::text"
	var s string

	db, err = NewConnection(db, []string{query})
	if err != nil {
		b.Fatal(err)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := db.QueryRow(query, "1").Scan(&s); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkMySQLWithoutPreparedStatements(b *testing.B) {
	b.ReportAllocs()

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/mysql")
	if err != nil {
		b.Fatal(err)
	}

	const query = "SELECT ?"
	var s string

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := db.QueryRow(query, "1").Scan(&s); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkMySQLWithPreparedStatements(b *testing.B) {
	b.ReportAllocs()

	var (
		db  Connector
		err error
	)
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/mysql")
	if err != nil {
		b.Fatal(err)
	}

	const query = "SELECT ?"
	var s string

	db, err = NewConnection(db, []string{query})
	if err != nil {
		b.Fatal(err)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := db.QueryRow(query, "1").Scan(&s); err != nil {
				b.Fatal(err)
			}
		}
	})
}
