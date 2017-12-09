package prep

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

type (
	executer interface {
		Exec(query string, args ...interface{}) (sql.Result, error)
	}
	executerWithContext interface {
		ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	}

	preparer interface {
		Prepare(query string) (*sql.Stmt, error)
	}

	preparerWithContext interface {
		PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	}

	rowQuerier interface {
		QueryRow(query string, args ...interface{}) *sql.Row
	}

	rowQuerierWithContext interface {
		QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	}

	querier interface {
		Query(query string, args ...interface{}) (*sql.Rows, error)
	}

	querierWithContext interface {
		QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	}
)

// PreparedStatement is implemented by *sql.Stmt
type PreparedStatement interface {
	Exec(args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error)
	QueryRow(args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row
	Query(args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, args ...interface{}) (*sql.Rows, error)
}

// Querier contains methods of the database/sql.DB that sending queries to the database
type Querier interface {
	executer
	executerWithContext
	preparer
	preparerWithContext
	rowQuerier
	rowQuerierWithContext
	querier
	querierWithContext
}

//rester contains all the rest methods of the database/sql.DB
type rester interface {
	Begin() (*sql.Tx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Close() error
	Driver() driver.Driver
	Ping() error
	PingContext(ctx context.Context) error
	SetConnMaxLifetime(d time.Duration)
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	Stats() sql.DBStats
}

// Connector contains all exportable methods of the database/sql.DB
type Connector interface {
	Querier
	rester
}

// preparedStatements is a map holding prepared statements for SQL statements
type preparedStatements map[string]PreparedStatement

// Connection is a wrapper for Connector instrumented with prepared SQL statements
type Connection struct {
	Querier
	rester
	statements preparedStatements
}

var errInvalidStatementType = errors.New("invalid statement type")

// NewConnection creates prepared statements for all statements in the given list
// and returns an implementation of the Connector interface instrumented with
// prepared statements
func NewConnection(c Connector, statements []string) (*Connection, error) {
	ps := preparedStatements{}

	for _, s := range statements {
		stmt, err := c.Prepare(s)
		if err != nil {
			return nil, fmt.Errorf("failed to prepare statement %s: %v", s, err)
		}

		ps[s] = stmt
	}

	return &Connection{Querier: c, rester: c, statements: ps}, nil
}

// PrepareContext implements preparerWithContext
func (c *Connection) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	if stmt, found := c.statements[query]; found {
		sqlStmt, ok := stmt.(*sql.Stmt)
		if !ok {
			return nil, errInvalidStatementType
		}
		return sqlStmt, nil
	}

	return c.Querier.PrepareContext(ctx, query)
}

// Prepare implements preparer
func (c *Connection) Prepare(query string) (*sql.Stmt, error) {
	if stmt, found := c.statements[query]; found {
		sqlStmt, ok := stmt.(*sql.Stmt)
		if !ok {
			return nil, errInvalidStatementType
		}
		return sqlStmt, nil
	}

	return c.Querier.Prepare(query)
}

// ExecContext implements executerWithContext
func (c *Connection) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if stmt, found := c.statements[query]; found {
		return stmt.ExecContext(ctx, args...)
	}

	return c.Querier.ExecContext(ctx, query, args...)
}

// Exec implements executer
func (c *Connection) Exec(query string, args ...interface{}) (sql.Result, error) {
	if stmt, found := c.statements[query]; found {
		return stmt.Exec(args...)
	}

	return c.Querier.Exec(query, args...)
}

// QueryContext implements querierWithContext
func (c *Connection) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if stmt, found := c.statements[query]; found {
		return stmt.QueryContext(ctx, args...)
	}

	return c.Querier.QueryContext(ctx, query, args...)
}

// Query implements querier
func (c *Connection) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if stmt, found := c.statements[query]; found {
		return stmt.Query(args...)
	}

	return c.Querier.Query(query, args...)
}

// QueryRowContext implements querierWithContext
func (c *Connection) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if stmt, found := c.statements[query]; found {
		return stmt.QueryRowContext(ctx, args...)
	}

	return c.Querier.QueryRowContext(ctx, query, args...)
}

// QueryRow implements rowQuerier
func (c *Connection) QueryRow(query string, args ...interface{}) *sql.Row {
	if stmt, found := c.statements[query]; found {
		return stmt.QueryRow(args...)
	}

	return c.Querier.QueryRow(query, args...)
}
