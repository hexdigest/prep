package prep

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.8
The original interface "PreparedStatement" can be found in github.com/hexdigest/prep
*/
import (
	context "context"
	sql "database/sql"
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"

	testify_assert "github.com/stretchr/testify/assert"
)

//PreparedStatementMock implements github.com/hexdigest/prep.PreparedStatement
type PreparedStatementMock struct {
	t minimock.Tester

	ExecFunc    func(p ...interface{}) (r sql.Result, r1 error)
	ExecCounter uint64
	ExecMock    mPreparedStatementMockExec

	ExecContextFunc    func(p context.Context, p1 ...interface{}) (r sql.Result, r1 error)
	ExecContextCounter uint64
	ExecContextMock    mPreparedStatementMockExecContext

	QueryFunc    func(p ...interface{}) (r *sql.Rows, r1 error)
	QueryCounter uint64
	QueryMock    mPreparedStatementMockQuery

	QueryContextFunc    func(p context.Context, p1 ...interface{}) (r *sql.Rows, r1 error)
	QueryContextCounter uint64
	QueryContextMock    mPreparedStatementMockQueryContext

	QueryRowFunc    func(p ...interface{}) (r *sql.Row)
	QueryRowCounter uint64
	QueryRowMock    mPreparedStatementMockQueryRow

	QueryRowContextFunc    func(p context.Context, p1 ...interface{}) (r *sql.Row)
	QueryRowContextCounter uint64
	QueryRowContextMock    mPreparedStatementMockQueryRowContext
}

//NewPreparedStatementMock returns a mock for github.com/hexdigest/prep.PreparedStatement
func NewPreparedStatementMock(t minimock.Tester) *PreparedStatementMock {
	m := &PreparedStatementMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ExecMock = mPreparedStatementMockExec{mock: m}
	m.ExecContextMock = mPreparedStatementMockExecContext{mock: m}
	m.QueryMock = mPreparedStatementMockQuery{mock: m}
	m.QueryContextMock = mPreparedStatementMockQueryContext{mock: m}
	m.QueryRowMock = mPreparedStatementMockQueryRow{mock: m}
	m.QueryRowContextMock = mPreparedStatementMockQueryRowContext{mock: m}

	return m
}

type mPreparedStatementMockExec struct {
	mock             *PreparedStatementMock
	mockExpectations *PreparedStatementMockExecParams
}

//PreparedStatementMockExecParams represents input parameters of the PreparedStatement.Exec
type PreparedStatementMockExecParams struct {
	p []interface{}
}

//Expect sets up expected params for the PreparedStatement.Exec
func (m *mPreparedStatementMockExec) Expect(p ...interface{}) *mPreparedStatementMockExec {
	m.mockExpectations = &PreparedStatementMockExecParams{p}
	return m
}

//Return sets up a mock for PreparedStatement.Exec to return Return's arguments
func (m *mPreparedStatementMockExec) Return(r sql.Result, r1 error) *PreparedStatementMock {
	m.mock.ExecFunc = func(p ...interface{}) (sql.Result, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of PreparedStatement.Exec method
func (m *mPreparedStatementMockExec) Set(f func(p ...interface{}) (r sql.Result, r1 error)) *PreparedStatementMock {
	m.mock.ExecFunc = f
	return m.mock
}

//Exec implements github.com/hexdigest/prep.PreparedStatement interface
func (m *PreparedStatementMock) Exec(p ...interface{}) (r sql.Result, r1 error) {
	defer atomic.AddUint64(&m.ExecCounter, 1)

	if m.ExecMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ExecMock.mockExpectations, PreparedStatementMockExecParams{p},
			"PreparedStatement.Exec got unexpected parameters")

		if m.ExecFunc == nil {

			m.t.Fatal("No results are set for the PreparedStatementMock.Exec")

			return
		}
	}

	if m.ExecFunc == nil {
		m.t.Fatal("Unexpected call to PreparedStatementMock.Exec")
		return
	}

	return m.ExecFunc(p...)
}

//ExecMinimockCounter returns a count of PreparedStatement.Exec invocations
func (m *PreparedStatementMock) ExecMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ExecCounter)
}

type mPreparedStatementMockExecContext struct {
	mock             *PreparedStatementMock
	mockExpectations *PreparedStatementMockExecContextParams
}

//PreparedStatementMockExecContextParams represents input parameters of the PreparedStatement.ExecContext
type PreparedStatementMockExecContextParams struct {
	p  context.Context
	p1 []interface{}
}

//Expect sets up expected params for the PreparedStatement.ExecContext
func (m *mPreparedStatementMockExecContext) Expect(p context.Context, p1 ...interface{}) *mPreparedStatementMockExecContext {
	m.mockExpectations = &PreparedStatementMockExecContextParams{p, p1}
	return m
}

//Return sets up a mock for PreparedStatement.ExecContext to return Return's arguments
func (m *mPreparedStatementMockExecContext) Return(r sql.Result, r1 error) *PreparedStatementMock {
	m.mock.ExecContextFunc = func(p context.Context, p1 ...interface{}) (sql.Result, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of PreparedStatement.ExecContext method
func (m *mPreparedStatementMockExecContext) Set(f func(p context.Context, p1 ...interface{}) (r sql.Result, r1 error)) *PreparedStatementMock {
	m.mock.ExecContextFunc = f
	return m.mock
}

//ExecContext implements github.com/hexdigest/prep.PreparedStatement interface
func (m *PreparedStatementMock) ExecContext(p context.Context, p1 ...interface{}) (r sql.Result, r1 error) {
	defer atomic.AddUint64(&m.ExecContextCounter, 1)

	if m.ExecContextMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ExecContextMock.mockExpectations, PreparedStatementMockExecContextParams{p, p1},
			"PreparedStatement.ExecContext got unexpected parameters")

		if m.ExecContextFunc == nil {

			m.t.Fatal("No results are set for the PreparedStatementMock.ExecContext")

			return
		}
	}

	if m.ExecContextFunc == nil {
		m.t.Fatal("Unexpected call to PreparedStatementMock.ExecContext")
		return
	}

	return m.ExecContextFunc(p, p1...)
}

//ExecContextMinimockCounter returns a count of PreparedStatement.ExecContext invocations
func (m *PreparedStatementMock) ExecContextMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ExecContextCounter)
}

type mPreparedStatementMockQuery struct {
	mock             *PreparedStatementMock
	mockExpectations *PreparedStatementMockQueryParams
}

//PreparedStatementMockQueryParams represents input parameters of the PreparedStatement.Query
type PreparedStatementMockQueryParams struct {
	p []interface{}
}

//Expect sets up expected params for the PreparedStatement.Query
func (m *mPreparedStatementMockQuery) Expect(p ...interface{}) *mPreparedStatementMockQuery {
	m.mockExpectations = &PreparedStatementMockQueryParams{p}
	return m
}

//Return sets up a mock for PreparedStatement.Query to return Return's arguments
func (m *mPreparedStatementMockQuery) Return(r *sql.Rows, r1 error) *PreparedStatementMock {
	m.mock.QueryFunc = func(p ...interface{}) (*sql.Rows, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of PreparedStatement.Query method
func (m *mPreparedStatementMockQuery) Set(f func(p ...interface{}) (r *sql.Rows, r1 error)) *PreparedStatementMock {
	m.mock.QueryFunc = f
	return m.mock
}

//Query implements github.com/hexdigest/prep.PreparedStatement interface
func (m *PreparedStatementMock) Query(p ...interface{}) (r *sql.Rows, r1 error) {
	defer atomic.AddUint64(&m.QueryCounter, 1)

	if m.QueryMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.QueryMock.mockExpectations, PreparedStatementMockQueryParams{p},
			"PreparedStatement.Query got unexpected parameters")

		if m.QueryFunc == nil {

			m.t.Fatal("No results are set for the PreparedStatementMock.Query")

			return
		}
	}

	if m.QueryFunc == nil {
		m.t.Fatal("Unexpected call to PreparedStatementMock.Query")
		return
	}

	return m.QueryFunc(p...)
}

//QueryMinimockCounter returns a count of PreparedStatement.Query invocations
func (m *PreparedStatementMock) QueryMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.QueryCounter)
}

type mPreparedStatementMockQueryContext struct {
	mock             *PreparedStatementMock
	mockExpectations *PreparedStatementMockQueryContextParams
}

//PreparedStatementMockQueryContextParams represents input parameters of the PreparedStatement.QueryContext
type PreparedStatementMockQueryContextParams struct {
	p  context.Context
	p1 []interface{}
}

//Expect sets up expected params for the PreparedStatement.QueryContext
func (m *mPreparedStatementMockQueryContext) Expect(p context.Context, p1 ...interface{}) *mPreparedStatementMockQueryContext {
	m.mockExpectations = &PreparedStatementMockQueryContextParams{p, p1}
	return m
}

//Return sets up a mock for PreparedStatement.QueryContext to return Return's arguments
func (m *mPreparedStatementMockQueryContext) Return(r *sql.Rows, r1 error) *PreparedStatementMock {
	m.mock.QueryContextFunc = func(p context.Context, p1 ...interface{}) (*sql.Rows, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of PreparedStatement.QueryContext method
func (m *mPreparedStatementMockQueryContext) Set(f func(p context.Context, p1 ...interface{}) (r *sql.Rows, r1 error)) *PreparedStatementMock {
	m.mock.QueryContextFunc = f
	return m.mock
}

//QueryContext implements github.com/hexdigest/prep.PreparedStatement interface
func (m *PreparedStatementMock) QueryContext(p context.Context, p1 ...interface{}) (r *sql.Rows, r1 error) {
	defer atomic.AddUint64(&m.QueryContextCounter, 1)

	if m.QueryContextMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.QueryContextMock.mockExpectations, PreparedStatementMockQueryContextParams{p, p1},
			"PreparedStatement.QueryContext got unexpected parameters")

		if m.QueryContextFunc == nil {

			m.t.Fatal("No results are set for the PreparedStatementMock.QueryContext")

			return
		}
	}

	if m.QueryContextFunc == nil {
		m.t.Fatal("Unexpected call to PreparedStatementMock.QueryContext")
		return
	}

	return m.QueryContextFunc(p, p1...)
}

//QueryContextMinimockCounter returns a count of PreparedStatement.QueryContext invocations
func (m *PreparedStatementMock) QueryContextMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.QueryContextCounter)
}

type mPreparedStatementMockQueryRow struct {
	mock             *PreparedStatementMock
	mockExpectations *PreparedStatementMockQueryRowParams
}

//PreparedStatementMockQueryRowParams represents input parameters of the PreparedStatement.QueryRow
type PreparedStatementMockQueryRowParams struct {
	p []interface{}
}

//Expect sets up expected params for the PreparedStatement.QueryRow
func (m *mPreparedStatementMockQueryRow) Expect(p ...interface{}) *mPreparedStatementMockQueryRow {
	m.mockExpectations = &PreparedStatementMockQueryRowParams{p}
	return m
}

//Return sets up a mock for PreparedStatement.QueryRow to return Return's arguments
func (m *mPreparedStatementMockQueryRow) Return(r *sql.Row) *PreparedStatementMock {
	m.mock.QueryRowFunc = func(p ...interface{}) *sql.Row {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of PreparedStatement.QueryRow method
func (m *mPreparedStatementMockQueryRow) Set(f func(p ...interface{}) (r *sql.Row)) *PreparedStatementMock {
	m.mock.QueryRowFunc = f
	return m.mock
}

//QueryRow implements github.com/hexdigest/prep.PreparedStatement interface
func (m *PreparedStatementMock) QueryRow(p ...interface{}) (r *sql.Row) {
	defer atomic.AddUint64(&m.QueryRowCounter, 1)

	if m.QueryRowMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.QueryRowMock.mockExpectations, PreparedStatementMockQueryRowParams{p},
			"PreparedStatement.QueryRow got unexpected parameters")

		if m.QueryRowFunc == nil {

			m.t.Fatal("No results are set for the PreparedStatementMock.QueryRow")

			return
		}
	}

	if m.QueryRowFunc == nil {
		m.t.Fatal("Unexpected call to PreparedStatementMock.QueryRow")
		return
	}

	return m.QueryRowFunc(p...)
}

//QueryRowMinimockCounter returns a count of PreparedStatement.QueryRow invocations
func (m *PreparedStatementMock) QueryRowMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.QueryRowCounter)
}

type mPreparedStatementMockQueryRowContext struct {
	mock             *PreparedStatementMock
	mockExpectations *PreparedStatementMockQueryRowContextParams
}

//PreparedStatementMockQueryRowContextParams represents input parameters of the PreparedStatement.QueryRowContext
type PreparedStatementMockQueryRowContextParams struct {
	p  context.Context
	p1 []interface{}
}

//Expect sets up expected params for the PreparedStatement.QueryRowContext
func (m *mPreparedStatementMockQueryRowContext) Expect(p context.Context, p1 ...interface{}) *mPreparedStatementMockQueryRowContext {
	m.mockExpectations = &PreparedStatementMockQueryRowContextParams{p, p1}
	return m
}

//Return sets up a mock for PreparedStatement.QueryRowContext to return Return's arguments
func (m *mPreparedStatementMockQueryRowContext) Return(r *sql.Row) *PreparedStatementMock {
	m.mock.QueryRowContextFunc = func(p context.Context, p1 ...interface{}) *sql.Row {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of PreparedStatement.QueryRowContext method
func (m *mPreparedStatementMockQueryRowContext) Set(f func(p context.Context, p1 ...interface{}) (r *sql.Row)) *PreparedStatementMock {
	m.mock.QueryRowContextFunc = f
	return m.mock
}

//QueryRowContext implements github.com/hexdigest/prep.PreparedStatement interface
func (m *PreparedStatementMock) QueryRowContext(p context.Context, p1 ...interface{}) (r *sql.Row) {
	defer atomic.AddUint64(&m.QueryRowContextCounter, 1)

	if m.QueryRowContextMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.QueryRowContextMock.mockExpectations, PreparedStatementMockQueryRowContextParams{p, p1},
			"PreparedStatement.QueryRowContext got unexpected parameters")

		if m.QueryRowContextFunc == nil {

			m.t.Fatal("No results are set for the PreparedStatementMock.QueryRowContext")

			return
		}
	}

	if m.QueryRowContextFunc == nil {
		m.t.Fatal("Unexpected call to PreparedStatementMock.QueryRowContext")
		return
	}

	return m.QueryRowContextFunc(p, p1...)
}

//QueryRowContextMinimockCounter returns a count of PreparedStatement.QueryRowContext invocations
func (m *PreparedStatementMock) QueryRowContextMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.QueryRowContextCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *PreparedStatementMock) ValidateCallCounters() {

	if m.ExecFunc != nil && atomic.LoadUint64(&m.ExecCounter) == 0 {
		m.t.Fatal("Expected call to PreparedStatementMock.Exec")
	}

	if m.ExecContextFunc != nil && atomic.LoadUint64(&m.ExecContextCounter) == 0 {
		m.t.Fatal("Expected call to PreparedStatementMock.ExecContext")
	}

	if m.QueryFunc != nil && atomic.LoadUint64(&m.QueryCounter) == 0 {
		m.t.Fatal("Expected call to PreparedStatementMock.Query")
	}

	if m.QueryContextFunc != nil && atomic.LoadUint64(&m.QueryContextCounter) == 0 {
		m.t.Fatal("Expected call to PreparedStatementMock.QueryContext")
	}

	if m.QueryRowFunc != nil && atomic.LoadUint64(&m.QueryRowCounter) == 0 {
		m.t.Fatal("Expected call to PreparedStatementMock.QueryRow")
	}

	if m.QueryRowContextFunc != nil && atomic.LoadUint64(&m.QueryRowContextCounter) == 0 {
		m.t.Fatal("Expected call to PreparedStatementMock.QueryRowContext")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *PreparedStatementMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *PreparedStatementMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *PreparedStatementMock) MinimockFinish() {

	if m.ExecFunc != nil && atomic.LoadUint64(&m.ExecCounter) == 0 {
		m.t.Fatal("Expected call to PreparedStatementMock.Exec")
	}

	if m.ExecContextFunc != nil && atomic.LoadUint64(&m.ExecContextCounter) == 0 {
		m.t.Fatal("Expected call to PreparedStatementMock.ExecContext")
	}

	if m.QueryFunc != nil && atomic.LoadUint64(&m.QueryCounter) == 0 {
		m.t.Fatal("Expected call to PreparedStatementMock.Query")
	}

	if m.QueryContextFunc != nil && atomic.LoadUint64(&m.QueryContextCounter) == 0 {
		m.t.Fatal("Expected call to PreparedStatementMock.QueryContext")
	}

	if m.QueryRowFunc != nil && atomic.LoadUint64(&m.QueryRowCounter) == 0 {
		m.t.Fatal("Expected call to PreparedStatementMock.QueryRow")
	}

	if m.QueryRowContextFunc != nil && atomic.LoadUint64(&m.QueryRowContextCounter) == 0 {
		m.t.Fatal("Expected call to PreparedStatementMock.QueryRowContext")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *PreparedStatementMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *PreparedStatementMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.ExecFunc == nil || atomic.LoadUint64(&m.ExecCounter) > 0)
		ok = ok && (m.ExecContextFunc == nil || atomic.LoadUint64(&m.ExecContextCounter) > 0)
		ok = ok && (m.QueryFunc == nil || atomic.LoadUint64(&m.QueryCounter) > 0)
		ok = ok && (m.QueryContextFunc == nil || atomic.LoadUint64(&m.QueryContextCounter) > 0)
		ok = ok && (m.QueryRowFunc == nil || atomic.LoadUint64(&m.QueryRowCounter) > 0)
		ok = ok && (m.QueryRowContextFunc == nil || atomic.LoadUint64(&m.QueryRowContextCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.ExecFunc != nil && atomic.LoadUint64(&m.ExecCounter) == 0 {
				m.t.Error("Expected call to PreparedStatementMock.Exec")
			}

			if m.ExecContextFunc != nil && atomic.LoadUint64(&m.ExecContextCounter) == 0 {
				m.t.Error("Expected call to PreparedStatementMock.ExecContext")
			}

			if m.QueryFunc != nil && atomic.LoadUint64(&m.QueryCounter) == 0 {
				m.t.Error("Expected call to PreparedStatementMock.Query")
			}

			if m.QueryContextFunc != nil && atomic.LoadUint64(&m.QueryContextCounter) == 0 {
				m.t.Error("Expected call to PreparedStatementMock.QueryContext")
			}

			if m.QueryRowFunc != nil && atomic.LoadUint64(&m.QueryRowCounter) == 0 {
				m.t.Error("Expected call to PreparedStatementMock.QueryRow")
			}

			if m.QueryRowContextFunc != nil && atomic.LoadUint64(&m.QueryRowContextCounter) == 0 {
				m.t.Error("Expected call to PreparedStatementMock.QueryRowContext")
			}

			m.t.Fatalf("Some mocks were not called on time: %s", timeout)
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

//AllMocksCalled returns true if all mocked methods were called before the execution of AllMocksCalled,
//it can be used with assert/require, i.e. assert.True(mock.AllMocksCalled())
func (m *PreparedStatementMock) AllMocksCalled() bool {

	if m.ExecFunc != nil && atomic.LoadUint64(&m.ExecCounter) == 0 {
		return false
	}

	if m.ExecContextFunc != nil && atomic.LoadUint64(&m.ExecContextCounter) == 0 {
		return false
	}

	if m.QueryFunc != nil && atomic.LoadUint64(&m.QueryCounter) == 0 {
		return false
	}

	if m.QueryContextFunc != nil && atomic.LoadUint64(&m.QueryContextCounter) == 0 {
		return false
	}

	if m.QueryRowFunc != nil && atomic.LoadUint64(&m.QueryRowCounter) == 0 {
		return false
	}

	if m.QueryRowContextFunc != nil && atomic.LoadUint64(&m.QueryRowContextCounter) == 0 {
		return false
	}

	return true
}
