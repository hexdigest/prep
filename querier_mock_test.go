package prep

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.8
The original interface "Querier" can be found in github.com/hexdigest/prep
*/
import (
	context "context"
	sql "database/sql"
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"

	testify_assert "github.com/stretchr/testify/assert"
)

//QuerierMock implements github.com/hexdigest/prep.Querier
type QuerierMock struct {
	t minimock.Tester

	ExecFunc    func(p string, p1 ...interface{}) (r sql.Result, r1 error)
	ExecCounter uint64
	ExecMock    mQuerierMockExec

	ExecContextFunc    func(p context.Context, p1 string, p2 ...interface{}) (r sql.Result, r1 error)
	ExecContextCounter uint64
	ExecContextMock    mQuerierMockExecContext

	PrepareFunc    func(p string) (r *sql.Stmt, r1 error)
	PrepareCounter uint64
	PrepareMock    mQuerierMockPrepare

	PrepareContextFunc    func(p context.Context, p1 string) (r *sql.Stmt, r1 error)
	PrepareContextCounter uint64
	PrepareContextMock    mQuerierMockPrepareContext

	QueryFunc    func(p string, p1 ...interface{}) (r *sql.Rows, r1 error)
	QueryCounter uint64
	QueryMock    mQuerierMockQuery

	QueryContextFunc    func(p context.Context, p1 string, p2 ...interface{}) (r *sql.Rows, r1 error)
	QueryContextCounter uint64
	QueryContextMock    mQuerierMockQueryContext

	QueryRowFunc    func(p string, p1 ...interface{}) (r *sql.Row)
	QueryRowCounter uint64
	QueryRowMock    mQuerierMockQueryRow

	QueryRowContextFunc    func(p context.Context, p1 string, p2 ...interface{}) (r *sql.Row)
	QueryRowContextCounter uint64
	QueryRowContextMock    mQuerierMockQueryRowContext
}

//NewQuerierMock returns a mock for github.com/hexdigest/prep.Querier
func NewQuerierMock(t minimock.Tester) *QuerierMock {
	m := &QuerierMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ExecMock = mQuerierMockExec{mock: m}
	m.ExecContextMock = mQuerierMockExecContext{mock: m}
	m.PrepareMock = mQuerierMockPrepare{mock: m}
	m.PrepareContextMock = mQuerierMockPrepareContext{mock: m}
	m.QueryMock = mQuerierMockQuery{mock: m}
	m.QueryContextMock = mQuerierMockQueryContext{mock: m}
	m.QueryRowMock = mQuerierMockQueryRow{mock: m}
	m.QueryRowContextMock = mQuerierMockQueryRowContext{mock: m}

	return m
}

type mQuerierMockExec struct {
	mock             *QuerierMock
	mockExpectations *QuerierMockExecParams
}

//QuerierMockExecParams represents input parameters of the Querier.Exec
type QuerierMockExecParams struct {
	p  string
	p1 []interface{}
}

//Expect sets up expected params for the Querier.Exec
func (m *mQuerierMockExec) Expect(p string, p1 ...interface{}) *mQuerierMockExec {
	m.mockExpectations = &QuerierMockExecParams{p, p1}
	return m
}

//Return sets up a mock for Querier.Exec to return Return's arguments
func (m *mQuerierMockExec) Return(r sql.Result, r1 error) *QuerierMock {
	m.mock.ExecFunc = func(p string, p1 ...interface{}) (sql.Result, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Querier.Exec method
func (m *mQuerierMockExec) Set(f func(p string, p1 ...interface{}) (r sql.Result, r1 error)) *QuerierMock {
	m.mock.ExecFunc = f
	return m.mock
}

//Exec implements github.com/hexdigest/prep.Querier interface
func (m *QuerierMock) Exec(p string, p1 ...interface{}) (r sql.Result, r1 error) {
	defer atomic.AddUint64(&m.ExecCounter, 1)

	if m.ExecMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ExecMock.mockExpectations, QuerierMockExecParams{p, p1},
			"Querier.Exec got unexpected parameters")

		if m.ExecFunc == nil {

			m.t.Fatal("No results are set for the QuerierMock.Exec")

			return
		}
	}

	if m.ExecFunc == nil {
		m.t.Fatal("Unexpected call to QuerierMock.Exec")
		return
	}

	return m.ExecFunc(p, p1...)
}

//ExecMinimockCounter returns a count of Querier.Exec invocations
func (m *QuerierMock) ExecMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ExecCounter)
}

type mQuerierMockExecContext struct {
	mock             *QuerierMock
	mockExpectations *QuerierMockExecContextParams
}

//QuerierMockExecContextParams represents input parameters of the Querier.ExecContext
type QuerierMockExecContextParams struct {
	p  context.Context
	p1 string
	p2 []interface{}
}

//Expect sets up expected params for the Querier.ExecContext
func (m *mQuerierMockExecContext) Expect(p context.Context, p1 string, p2 ...interface{}) *mQuerierMockExecContext {
	m.mockExpectations = &QuerierMockExecContextParams{p, p1, p2}
	return m
}

//Return sets up a mock for Querier.ExecContext to return Return's arguments
func (m *mQuerierMockExecContext) Return(r sql.Result, r1 error) *QuerierMock {
	m.mock.ExecContextFunc = func(p context.Context, p1 string, p2 ...interface{}) (sql.Result, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Querier.ExecContext method
func (m *mQuerierMockExecContext) Set(f func(p context.Context, p1 string, p2 ...interface{}) (r sql.Result, r1 error)) *QuerierMock {
	m.mock.ExecContextFunc = f
	return m.mock
}

//ExecContext implements github.com/hexdigest/prep.Querier interface
func (m *QuerierMock) ExecContext(p context.Context, p1 string, p2 ...interface{}) (r sql.Result, r1 error) {
	defer atomic.AddUint64(&m.ExecContextCounter, 1)

	if m.ExecContextMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ExecContextMock.mockExpectations, QuerierMockExecContextParams{p, p1, p2},
			"Querier.ExecContext got unexpected parameters")

		if m.ExecContextFunc == nil {

			m.t.Fatal("No results are set for the QuerierMock.ExecContext")

			return
		}
	}

	if m.ExecContextFunc == nil {
		m.t.Fatal("Unexpected call to QuerierMock.ExecContext")
		return
	}

	return m.ExecContextFunc(p, p1, p2...)
}

//ExecContextMinimockCounter returns a count of Querier.ExecContext invocations
func (m *QuerierMock) ExecContextMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ExecContextCounter)
}

type mQuerierMockPrepare struct {
	mock             *QuerierMock
	mockExpectations *QuerierMockPrepareParams
}

//QuerierMockPrepareParams represents input parameters of the Querier.Prepare
type QuerierMockPrepareParams struct {
	p string
}

//Expect sets up expected params for the Querier.Prepare
func (m *mQuerierMockPrepare) Expect(p string) *mQuerierMockPrepare {
	m.mockExpectations = &QuerierMockPrepareParams{p}
	return m
}

//Return sets up a mock for Querier.Prepare to return Return's arguments
func (m *mQuerierMockPrepare) Return(r *sql.Stmt, r1 error) *QuerierMock {
	m.mock.PrepareFunc = func(p string) (*sql.Stmt, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Querier.Prepare method
func (m *mQuerierMockPrepare) Set(f func(p string) (r *sql.Stmt, r1 error)) *QuerierMock {
	m.mock.PrepareFunc = f
	return m.mock
}

//Prepare implements github.com/hexdigest/prep.Querier interface
func (m *QuerierMock) Prepare(p string) (r *sql.Stmt, r1 error) {
	defer atomic.AddUint64(&m.PrepareCounter, 1)

	if m.PrepareMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.PrepareMock.mockExpectations, QuerierMockPrepareParams{p},
			"Querier.Prepare got unexpected parameters")

		if m.PrepareFunc == nil {

			m.t.Fatal("No results are set for the QuerierMock.Prepare")

			return
		}
	}

	if m.PrepareFunc == nil {
		m.t.Fatal("Unexpected call to QuerierMock.Prepare")
		return
	}

	return m.PrepareFunc(p)
}

//PrepareMinimockCounter returns a count of Querier.Prepare invocations
func (m *QuerierMock) PrepareMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.PrepareCounter)
}

type mQuerierMockPrepareContext struct {
	mock             *QuerierMock
	mockExpectations *QuerierMockPrepareContextParams
}

//QuerierMockPrepareContextParams represents input parameters of the Querier.PrepareContext
type QuerierMockPrepareContextParams struct {
	p  context.Context
	p1 string
}

//Expect sets up expected params for the Querier.PrepareContext
func (m *mQuerierMockPrepareContext) Expect(p context.Context, p1 string) *mQuerierMockPrepareContext {
	m.mockExpectations = &QuerierMockPrepareContextParams{p, p1}
	return m
}

//Return sets up a mock for Querier.PrepareContext to return Return's arguments
func (m *mQuerierMockPrepareContext) Return(r *sql.Stmt, r1 error) *QuerierMock {
	m.mock.PrepareContextFunc = func(p context.Context, p1 string) (*sql.Stmt, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Querier.PrepareContext method
func (m *mQuerierMockPrepareContext) Set(f func(p context.Context, p1 string) (r *sql.Stmt, r1 error)) *QuerierMock {
	m.mock.PrepareContextFunc = f
	return m.mock
}

//PrepareContext implements github.com/hexdigest/prep.Querier interface
func (m *QuerierMock) PrepareContext(p context.Context, p1 string) (r *sql.Stmt, r1 error) {
	defer atomic.AddUint64(&m.PrepareContextCounter, 1)

	if m.PrepareContextMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.PrepareContextMock.mockExpectations, QuerierMockPrepareContextParams{p, p1},
			"Querier.PrepareContext got unexpected parameters")

		if m.PrepareContextFunc == nil {

			m.t.Fatal("No results are set for the QuerierMock.PrepareContext")

			return
		}
	}

	if m.PrepareContextFunc == nil {
		m.t.Fatal("Unexpected call to QuerierMock.PrepareContext")
		return
	}

	return m.PrepareContextFunc(p, p1)
}

//PrepareContextMinimockCounter returns a count of Querier.PrepareContext invocations
func (m *QuerierMock) PrepareContextMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.PrepareContextCounter)
}

type mQuerierMockQuery struct {
	mock             *QuerierMock
	mockExpectations *QuerierMockQueryParams
}

//QuerierMockQueryParams represents input parameters of the Querier.Query
type QuerierMockQueryParams struct {
	p  string
	p1 []interface{}
}

//Expect sets up expected params for the Querier.Query
func (m *mQuerierMockQuery) Expect(p string, p1 ...interface{}) *mQuerierMockQuery {
	m.mockExpectations = &QuerierMockQueryParams{p, p1}
	return m
}

//Return sets up a mock for Querier.Query to return Return's arguments
func (m *mQuerierMockQuery) Return(r *sql.Rows, r1 error) *QuerierMock {
	m.mock.QueryFunc = func(p string, p1 ...interface{}) (*sql.Rows, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Querier.Query method
func (m *mQuerierMockQuery) Set(f func(p string, p1 ...interface{}) (r *sql.Rows, r1 error)) *QuerierMock {
	m.mock.QueryFunc = f
	return m.mock
}

//Query implements github.com/hexdigest/prep.Querier interface
func (m *QuerierMock) Query(p string, p1 ...interface{}) (r *sql.Rows, r1 error) {
	defer atomic.AddUint64(&m.QueryCounter, 1)

	if m.QueryMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.QueryMock.mockExpectations, QuerierMockQueryParams{p, p1},
			"Querier.Query got unexpected parameters")

		if m.QueryFunc == nil {

			m.t.Fatal("No results are set for the QuerierMock.Query")

			return
		}
	}

	if m.QueryFunc == nil {
		m.t.Fatal("Unexpected call to QuerierMock.Query")
		return
	}

	return m.QueryFunc(p, p1...)
}

//QueryMinimockCounter returns a count of Querier.Query invocations
func (m *QuerierMock) QueryMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.QueryCounter)
}

type mQuerierMockQueryContext struct {
	mock             *QuerierMock
	mockExpectations *QuerierMockQueryContextParams
}

//QuerierMockQueryContextParams represents input parameters of the Querier.QueryContext
type QuerierMockQueryContextParams struct {
	p  context.Context
	p1 string
	p2 []interface{}
}

//Expect sets up expected params for the Querier.QueryContext
func (m *mQuerierMockQueryContext) Expect(p context.Context, p1 string, p2 ...interface{}) *mQuerierMockQueryContext {
	m.mockExpectations = &QuerierMockQueryContextParams{p, p1, p2}
	return m
}

//Return sets up a mock for Querier.QueryContext to return Return's arguments
func (m *mQuerierMockQueryContext) Return(r *sql.Rows, r1 error) *QuerierMock {
	m.mock.QueryContextFunc = func(p context.Context, p1 string, p2 ...interface{}) (*sql.Rows, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Querier.QueryContext method
func (m *mQuerierMockQueryContext) Set(f func(p context.Context, p1 string, p2 ...interface{}) (r *sql.Rows, r1 error)) *QuerierMock {
	m.mock.QueryContextFunc = f
	return m.mock
}

//QueryContext implements github.com/hexdigest/prep.Querier interface
func (m *QuerierMock) QueryContext(p context.Context, p1 string, p2 ...interface{}) (r *sql.Rows, r1 error) {
	defer atomic.AddUint64(&m.QueryContextCounter, 1)

	if m.QueryContextMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.QueryContextMock.mockExpectations, QuerierMockQueryContextParams{p, p1, p2},
			"Querier.QueryContext got unexpected parameters")

		if m.QueryContextFunc == nil {

			m.t.Fatal("No results are set for the QuerierMock.QueryContext")

			return
		}
	}

	if m.QueryContextFunc == nil {
		m.t.Fatal("Unexpected call to QuerierMock.QueryContext")
		return
	}

	return m.QueryContextFunc(p, p1, p2...)
}

//QueryContextMinimockCounter returns a count of Querier.QueryContext invocations
func (m *QuerierMock) QueryContextMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.QueryContextCounter)
}

type mQuerierMockQueryRow struct {
	mock             *QuerierMock
	mockExpectations *QuerierMockQueryRowParams
}

//QuerierMockQueryRowParams represents input parameters of the Querier.QueryRow
type QuerierMockQueryRowParams struct {
	p  string
	p1 []interface{}
}

//Expect sets up expected params for the Querier.QueryRow
func (m *mQuerierMockQueryRow) Expect(p string, p1 ...interface{}) *mQuerierMockQueryRow {
	m.mockExpectations = &QuerierMockQueryRowParams{p, p1}
	return m
}

//Return sets up a mock for Querier.QueryRow to return Return's arguments
func (m *mQuerierMockQueryRow) Return(r *sql.Row) *QuerierMock {
	m.mock.QueryRowFunc = func(p string, p1 ...interface{}) *sql.Row {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Querier.QueryRow method
func (m *mQuerierMockQueryRow) Set(f func(p string, p1 ...interface{}) (r *sql.Row)) *QuerierMock {
	m.mock.QueryRowFunc = f
	return m.mock
}

//QueryRow implements github.com/hexdigest/prep.Querier interface
func (m *QuerierMock) QueryRow(p string, p1 ...interface{}) (r *sql.Row) {
	defer atomic.AddUint64(&m.QueryRowCounter, 1)

	if m.QueryRowMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.QueryRowMock.mockExpectations, QuerierMockQueryRowParams{p, p1},
			"Querier.QueryRow got unexpected parameters")

		if m.QueryRowFunc == nil {

			m.t.Fatal("No results are set for the QuerierMock.QueryRow")

			return
		}
	}

	if m.QueryRowFunc == nil {
		m.t.Fatal("Unexpected call to QuerierMock.QueryRow")
		return
	}

	return m.QueryRowFunc(p, p1...)
}

//QueryRowMinimockCounter returns a count of Querier.QueryRow invocations
func (m *QuerierMock) QueryRowMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.QueryRowCounter)
}

type mQuerierMockQueryRowContext struct {
	mock             *QuerierMock
	mockExpectations *QuerierMockQueryRowContextParams
}

//QuerierMockQueryRowContextParams represents input parameters of the Querier.QueryRowContext
type QuerierMockQueryRowContextParams struct {
	p  context.Context
	p1 string
	p2 []interface{}
}

//Expect sets up expected params for the Querier.QueryRowContext
func (m *mQuerierMockQueryRowContext) Expect(p context.Context, p1 string, p2 ...interface{}) *mQuerierMockQueryRowContext {
	m.mockExpectations = &QuerierMockQueryRowContextParams{p, p1, p2}
	return m
}

//Return sets up a mock for Querier.QueryRowContext to return Return's arguments
func (m *mQuerierMockQueryRowContext) Return(r *sql.Row) *QuerierMock {
	m.mock.QueryRowContextFunc = func(p context.Context, p1 string, p2 ...interface{}) *sql.Row {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Querier.QueryRowContext method
func (m *mQuerierMockQueryRowContext) Set(f func(p context.Context, p1 string, p2 ...interface{}) (r *sql.Row)) *QuerierMock {
	m.mock.QueryRowContextFunc = f
	return m.mock
}

//QueryRowContext implements github.com/hexdigest/prep.Querier interface
func (m *QuerierMock) QueryRowContext(p context.Context, p1 string, p2 ...interface{}) (r *sql.Row) {
	defer atomic.AddUint64(&m.QueryRowContextCounter, 1)

	if m.QueryRowContextMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.QueryRowContextMock.mockExpectations, QuerierMockQueryRowContextParams{p, p1, p2},
			"Querier.QueryRowContext got unexpected parameters")

		if m.QueryRowContextFunc == nil {

			m.t.Fatal("No results are set for the QuerierMock.QueryRowContext")

			return
		}
	}

	if m.QueryRowContextFunc == nil {
		m.t.Fatal("Unexpected call to QuerierMock.QueryRowContext")
		return
	}

	return m.QueryRowContextFunc(p, p1, p2...)
}

//QueryRowContextMinimockCounter returns a count of Querier.QueryRowContext invocations
func (m *QuerierMock) QueryRowContextMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.QueryRowContextCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *QuerierMock) ValidateCallCounters() {

	if m.ExecFunc != nil && atomic.LoadUint64(&m.ExecCounter) == 0 {
		m.t.Fatal("Expected call to QuerierMock.Exec")
	}

	if m.ExecContextFunc != nil && atomic.LoadUint64(&m.ExecContextCounter) == 0 {
		m.t.Fatal("Expected call to QuerierMock.ExecContext")
	}

	if m.PrepareFunc != nil && atomic.LoadUint64(&m.PrepareCounter) == 0 {
		m.t.Fatal("Expected call to QuerierMock.Prepare")
	}

	if m.PrepareContextFunc != nil && atomic.LoadUint64(&m.PrepareContextCounter) == 0 {
		m.t.Fatal("Expected call to QuerierMock.PrepareContext")
	}

	if m.QueryFunc != nil && atomic.LoadUint64(&m.QueryCounter) == 0 {
		m.t.Fatal("Expected call to QuerierMock.Query")
	}

	if m.QueryContextFunc != nil && atomic.LoadUint64(&m.QueryContextCounter) == 0 {
		m.t.Fatal("Expected call to QuerierMock.QueryContext")
	}

	if m.QueryRowFunc != nil && atomic.LoadUint64(&m.QueryRowCounter) == 0 {
		m.t.Fatal("Expected call to QuerierMock.QueryRow")
	}

	if m.QueryRowContextFunc != nil && atomic.LoadUint64(&m.QueryRowContextCounter) == 0 {
		m.t.Fatal("Expected call to QuerierMock.QueryRowContext")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *QuerierMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *QuerierMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *QuerierMock) MinimockFinish() {

	if m.ExecFunc != nil && atomic.LoadUint64(&m.ExecCounter) == 0 {
		m.t.Fatal("Expected call to QuerierMock.Exec")
	}

	if m.ExecContextFunc != nil && atomic.LoadUint64(&m.ExecContextCounter) == 0 {
		m.t.Fatal("Expected call to QuerierMock.ExecContext")
	}

	if m.PrepareFunc != nil && atomic.LoadUint64(&m.PrepareCounter) == 0 {
		m.t.Fatal("Expected call to QuerierMock.Prepare")
	}

	if m.PrepareContextFunc != nil && atomic.LoadUint64(&m.PrepareContextCounter) == 0 {
		m.t.Fatal("Expected call to QuerierMock.PrepareContext")
	}

	if m.QueryFunc != nil && atomic.LoadUint64(&m.QueryCounter) == 0 {
		m.t.Fatal("Expected call to QuerierMock.Query")
	}

	if m.QueryContextFunc != nil && atomic.LoadUint64(&m.QueryContextCounter) == 0 {
		m.t.Fatal("Expected call to QuerierMock.QueryContext")
	}

	if m.QueryRowFunc != nil && atomic.LoadUint64(&m.QueryRowCounter) == 0 {
		m.t.Fatal("Expected call to QuerierMock.QueryRow")
	}

	if m.QueryRowContextFunc != nil && atomic.LoadUint64(&m.QueryRowContextCounter) == 0 {
		m.t.Fatal("Expected call to QuerierMock.QueryRowContext")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *QuerierMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *QuerierMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.ExecFunc == nil || atomic.LoadUint64(&m.ExecCounter) > 0)
		ok = ok && (m.ExecContextFunc == nil || atomic.LoadUint64(&m.ExecContextCounter) > 0)
		ok = ok && (m.PrepareFunc == nil || atomic.LoadUint64(&m.PrepareCounter) > 0)
		ok = ok && (m.PrepareContextFunc == nil || atomic.LoadUint64(&m.PrepareContextCounter) > 0)
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
				m.t.Error("Expected call to QuerierMock.Exec")
			}

			if m.ExecContextFunc != nil && atomic.LoadUint64(&m.ExecContextCounter) == 0 {
				m.t.Error("Expected call to QuerierMock.ExecContext")
			}

			if m.PrepareFunc != nil && atomic.LoadUint64(&m.PrepareCounter) == 0 {
				m.t.Error("Expected call to QuerierMock.Prepare")
			}

			if m.PrepareContextFunc != nil && atomic.LoadUint64(&m.PrepareContextCounter) == 0 {
				m.t.Error("Expected call to QuerierMock.PrepareContext")
			}

			if m.QueryFunc != nil && atomic.LoadUint64(&m.QueryCounter) == 0 {
				m.t.Error("Expected call to QuerierMock.Query")
			}

			if m.QueryContextFunc != nil && atomic.LoadUint64(&m.QueryContextCounter) == 0 {
				m.t.Error("Expected call to QuerierMock.QueryContext")
			}

			if m.QueryRowFunc != nil && atomic.LoadUint64(&m.QueryRowCounter) == 0 {
				m.t.Error("Expected call to QuerierMock.QueryRow")
			}

			if m.QueryRowContextFunc != nil && atomic.LoadUint64(&m.QueryRowContextCounter) == 0 {
				m.t.Error("Expected call to QuerierMock.QueryRowContext")
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
func (m *QuerierMock) AllMocksCalled() bool {

	if m.ExecFunc != nil && atomic.LoadUint64(&m.ExecCounter) == 0 {
		return false
	}

	if m.ExecContextFunc != nil && atomic.LoadUint64(&m.ExecContextCounter) == 0 {
		return false
	}

	if m.PrepareFunc != nil && atomic.LoadUint64(&m.PrepareCounter) == 0 {
		return false
	}

	if m.PrepareContextFunc != nil && atomic.LoadUint64(&m.PrepareContextCounter) == 0 {
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
