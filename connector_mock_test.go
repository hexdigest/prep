package prep

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.8
The original interface "Connector" can be found in github.com/hexdigest/prep
*/
import (
	context "context"
	sql "database/sql"
	driver "database/sql/driver"
	"sync/atomic"
	time "time"

	"github.com/gojuno/minimock"

	testify_assert "github.com/stretchr/testify/assert"
)

//ConnectorMock implements github.com/hexdigest/prep.Connector
type ConnectorMock struct {
	t minimock.Tester

	BeginFunc    func() (r *sql.Tx, r1 error)
	BeginCounter uint64
	BeginMock    mConnectorMockBegin

	BeginTxFunc    func(p context.Context, p1 *sql.TxOptions) (r *sql.Tx, r1 error)
	BeginTxCounter uint64
	BeginTxMock    mConnectorMockBeginTx

	CloseFunc    func() (r error)
	CloseCounter uint64
	CloseMock    mConnectorMockClose

	DriverFunc    func() (r driver.Driver)
	DriverCounter uint64
	DriverMock    mConnectorMockDriver

	ExecFunc    func(p string, p1 ...interface{}) (r sql.Result, r1 error)
	ExecCounter uint64
	ExecMock    mConnectorMockExec

	ExecContextFunc    func(p context.Context, p1 string, p2 ...interface{}) (r sql.Result, r1 error)
	ExecContextCounter uint64
	ExecContextMock    mConnectorMockExecContext

	PingFunc    func() (r error)
	PingCounter uint64
	PingMock    mConnectorMockPing

	PingContextFunc    func(p context.Context) (r error)
	PingContextCounter uint64
	PingContextMock    mConnectorMockPingContext

	PrepareFunc    func(p string) (r *sql.Stmt, r1 error)
	PrepareCounter uint64
	PrepareMock    mConnectorMockPrepare

	PrepareContextFunc    func(p context.Context, p1 string) (r *sql.Stmt, r1 error)
	PrepareContextCounter uint64
	PrepareContextMock    mConnectorMockPrepareContext

	QueryFunc    func(p string, p1 ...interface{}) (r *sql.Rows, r1 error)
	QueryCounter uint64
	QueryMock    mConnectorMockQuery

	QueryContextFunc    func(p context.Context, p1 string, p2 ...interface{}) (r *sql.Rows, r1 error)
	QueryContextCounter uint64
	QueryContextMock    mConnectorMockQueryContext

	QueryRowFunc    func(p string, p1 ...interface{}) (r *sql.Row)
	QueryRowCounter uint64
	QueryRowMock    mConnectorMockQueryRow

	QueryRowContextFunc    func(p context.Context, p1 string, p2 ...interface{}) (r *sql.Row)
	QueryRowContextCounter uint64
	QueryRowContextMock    mConnectorMockQueryRowContext

	SetConnMaxLifetimeFunc    func(p time.Duration)
	SetConnMaxLifetimeCounter uint64
	SetConnMaxLifetimeMock    mConnectorMockSetConnMaxLifetime

	SetMaxIdleConnsFunc    func(p int)
	SetMaxIdleConnsCounter uint64
	SetMaxIdleConnsMock    mConnectorMockSetMaxIdleConns

	SetMaxOpenConnsFunc    func(p int)
	SetMaxOpenConnsCounter uint64
	SetMaxOpenConnsMock    mConnectorMockSetMaxOpenConns

	StatsFunc    func() (r sql.DBStats)
	StatsCounter uint64
	StatsMock    mConnectorMockStats
}

//NewConnectorMock returns a mock for github.com/hexdigest/prep.Connector
func NewConnectorMock(t minimock.Tester) *ConnectorMock {
	m := &ConnectorMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.BeginMock = mConnectorMockBegin{mock: m}
	m.BeginTxMock = mConnectorMockBeginTx{mock: m}
	m.CloseMock = mConnectorMockClose{mock: m}
	m.DriverMock = mConnectorMockDriver{mock: m}
	m.ExecMock = mConnectorMockExec{mock: m}
	m.ExecContextMock = mConnectorMockExecContext{mock: m}
	m.PingMock = mConnectorMockPing{mock: m}
	m.PingContextMock = mConnectorMockPingContext{mock: m}
	m.PrepareMock = mConnectorMockPrepare{mock: m}
	m.PrepareContextMock = mConnectorMockPrepareContext{mock: m}
	m.QueryMock = mConnectorMockQuery{mock: m}
	m.QueryContextMock = mConnectorMockQueryContext{mock: m}
	m.QueryRowMock = mConnectorMockQueryRow{mock: m}
	m.QueryRowContextMock = mConnectorMockQueryRowContext{mock: m}
	m.SetConnMaxLifetimeMock = mConnectorMockSetConnMaxLifetime{mock: m}
	m.SetMaxIdleConnsMock = mConnectorMockSetMaxIdleConns{mock: m}
	m.SetMaxOpenConnsMock = mConnectorMockSetMaxOpenConns{mock: m}
	m.StatsMock = mConnectorMockStats{mock: m}

	return m
}

type mConnectorMockBegin struct {
	mock *ConnectorMock
}

//Return sets up a mock for Connector.Begin to return Return's arguments
func (m *mConnectorMockBegin) Return(r *sql.Tx, r1 error) *ConnectorMock {
	m.mock.BeginFunc = func() (*sql.Tx, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.Begin method
func (m *mConnectorMockBegin) Set(f func() (r *sql.Tx, r1 error)) *ConnectorMock {
	m.mock.BeginFunc = f
	return m.mock
}

//Begin implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) Begin() (r *sql.Tx, r1 error) {
	defer atomic.AddUint64(&m.BeginCounter, 1)

	if m.BeginFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.Begin")
		return
	}

	return m.BeginFunc()
}

//BeginMinimockCounter returns a count of Connector.Begin invocations
func (m *ConnectorMock) BeginMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.BeginCounter)
}

type mConnectorMockBeginTx struct {
	mock             *ConnectorMock
	mockExpectations *ConnectorMockBeginTxParams
}

//ConnectorMockBeginTxParams represents input parameters of the Connector.BeginTx
type ConnectorMockBeginTxParams struct {
	p  context.Context
	p1 *sql.TxOptions
}

//Expect sets up expected params for the Connector.BeginTx
func (m *mConnectorMockBeginTx) Expect(p context.Context, p1 *sql.TxOptions) *mConnectorMockBeginTx {
	m.mockExpectations = &ConnectorMockBeginTxParams{p, p1}
	return m
}

//Return sets up a mock for Connector.BeginTx to return Return's arguments
func (m *mConnectorMockBeginTx) Return(r *sql.Tx, r1 error) *ConnectorMock {
	m.mock.BeginTxFunc = func(p context.Context, p1 *sql.TxOptions) (*sql.Tx, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.BeginTx method
func (m *mConnectorMockBeginTx) Set(f func(p context.Context, p1 *sql.TxOptions) (r *sql.Tx, r1 error)) *ConnectorMock {
	m.mock.BeginTxFunc = f
	return m.mock
}

//BeginTx implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) BeginTx(p context.Context, p1 *sql.TxOptions) (r *sql.Tx, r1 error) {
	defer atomic.AddUint64(&m.BeginTxCounter, 1)

	if m.BeginTxMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.BeginTxMock.mockExpectations, ConnectorMockBeginTxParams{p, p1},
			"Connector.BeginTx got unexpected parameters")

		if m.BeginTxFunc == nil {

			m.t.Fatal("No results are set for the ConnectorMock.BeginTx")

			return
		}
	}

	if m.BeginTxFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.BeginTx")
		return
	}

	return m.BeginTxFunc(p, p1)
}

//BeginTxMinimockCounter returns a count of Connector.BeginTx invocations
func (m *ConnectorMock) BeginTxMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.BeginTxCounter)
}

type mConnectorMockClose struct {
	mock *ConnectorMock
}

//Return sets up a mock for Connector.Close to return Return's arguments
func (m *mConnectorMockClose) Return(r error) *ConnectorMock {
	m.mock.CloseFunc = func() error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.Close method
func (m *mConnectorMockClose) Set(f func() (r error)) *ConnectorMock {
	m.mock.CloseFunc = f
	return m.mock
}

//Close implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) Close() (r error) {
	defer atomic.AddUint64(&m.CloseCounter, 1)

	if m.CloseFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.Close")
		return
	}

	return m.CloseFunc()
}

//CloseMinimockCounter returns a count of Connector.Close invocations
func (m *ConnectorMock) CloseMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.CloseCounter)
}

type mConnectorMockDriver struct {
	mock *ConnectorMock
}

//Return sets up a mock for Connector.Driver to return Return's arguments
func (m *mConnectorMockDriver) Return(r driver.Driver) *ConnectorMock {
	m.mock.DriverFunc = func() driver.Driver {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.Driver method
func (m *mConnectorMockDriver) Set(f func() (r driver.Driver)) *ConnectorMock {
	m.mock.DriverFunc = f
	return m.mock
}

//Driver implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) Driver() (r driver.Driver) {
	defer atomic.AddUint64(&m.DriverCounter, 1)

	if m.DriverFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.Driver")
		return
	}

	return m.DriverFunc()
}

//DriverMinimockCounter returns a count of Connector.Driver invocations
func (m *ConnectorMock) DriverMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.DriverCounter)
}

type mConnectorMockExec struct {
	mock             *ConnectorMock
	mockExpectations *ConnectorMockExecParams
}

//ConnectorMockExecParams represents input parameters of the Connector.Exec
type ConnectorMockExecParams struct {
	p  string
	p1 []interface{}
}

//Expect sets up expected params for the Connector.Exec
func (m *mConnectorMockExec) Expect(p string, p1 ...interface{}) *mConnectorMockExec {
	m.mockExpectations = &ConnectorMockExecParams{p, p1}
	return m
}

//Return sets up a mock for Connector.Exec to return Return's arguments
func (m *mConnectorMockExec) Return(r sql.Result, r1 error) *ConnectorMock {
	m.mock.ExecFunc = func(p string, p1 ...interface{}) (sql.Result, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.Exec method
func (m *mConnectorMockExec) Set(f func(p string, p1 ...interface{}) (r sql.Result, r1 error)) *ConnectorMock {
	m.mock.ExecFunc = f
	return m.mock
}

//Exec implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) Exec(p string, p1 ...interface{}) (r sql.Result, r1 error) {
	defer atomic.AddUint64(&m.ExecCounter, 1)

	if m.ExecMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ExecMock.mockExpectations, ConnectorMockExecParams{p, p1},
			"Connector.Exec got unexpected parameters")

		if m.ExecFunc == nil {

			m.t.Fatal("No results are set for the ConnectorMock.Exec")

			return
		}
	}

	if m.ExecFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.Exec")
		return
	}

	return m.ExecFunc(p, p1...)
}

//ExecMinimockCounter returns a count of Connector.Exec invocations
func (m *ConnectorMock) ExecMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ExecCounter)
}

type mConnectorMockExecContext struct {
	mock             *ConnectorMock
	mockExpectations *ConnectorMockExecContextParams
}

//ConnectorMockExecContextParams represents input parameters of the Connector.ExecContext
type ConnectorMockExecContextParams struct {
	p  context.Context
	p1 string
	p2 []interface{}
}

//Expect sets up expected params for the Connector.ExecContext
func (m *mConnectorMockExecContext) Expect(p context.Context, p1 string, p2 ...interface{}) *mConnectorMockExecContext {
	m.mockExpectations = &ConnectorMockExecContextParams{p, p1, p2}
	return m
}

//Return sets up a mock for Connector.ExecContext to return Return's arguments
func (m *mConnectorMockExecContext) Return(r sql.Result, r1 error) *ConnectorMock {
	m.mock.ExecContextFunc = func(p context.Context, p1 string, p2 ...interface{}) (sql.Result, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.ExecContext method
func (m *mConnectorMockExecContext) Set(f func(p context.Context, p1 string, p2 ...interface{}) (r sql.Result, r1 error)) *ConnectorMock {
	m.mock.ExecContextFunc = f
	return m.mock
}

//ExecContext implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) ExecContext(p context.Context, p1 string, p2 ...interface{}) (r sql.Result, r1 error) {
	defer atomic.AddUint64(&m.ExecContextCounter, 1)

	if m.ExecContextMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ExecContextMock.mockExpectations, ConnectorMockExecContextParams{p, p1, p2},
			"Connector.ExecContext got unexpected parameters")

		if m.ExecContextFunc == nil {

			m.t.Fatal("No results are set for the ConnectorMock.ExecContext")

			return
		}
	}

	if m.ExecContextFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.ExecContext")
		return
	}

	return m.ExecContextFunc(p, p1, p2...)
}

//ExecContextMinimockCounter returns a count of Connector.ExecContext invocations
func (m *ConnectorMock) ExecContextMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ExecContextCounter)
}

type mConnectorMockPing struct {
	mock *ConnectorMock
}

//Return sets up a mock for Connector.Ping to return Return's arguments
func (m *mConnectorMockPing) Return(r error) *ConnectorMock {
	m.mock.PingFunc = func() error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.Ping method
func (m *mConnectorMockPing) Set(f func() (r error)) *ConnectorMock {
	m.mock.PingFunc = f
	return m.mock
}

//Ping implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) Ping() (r error) {
	defer atomic.AddUint64(&m.PingCounter, 1)

	if m.PingFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.Ping")
		return
	}

	return m.PingFunc()
}

//PingMinimockCounter returns a count of Connector.Ping invocations
func (m *ConnectorMock) PingMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.PingCounter)
}

type mConnectorMockPingContext struct {
	mock             *ConnectorMock
	mockExpectations *ConnectorMockPingContextParams
}

//ConnectorMockPingContextParams represents input parameters of the Connector.PingContext
type ConnectorMockPingContextParams struct {
	p context.Context
}

//Expect sets up expected params for the Connector.PingContext
func (m *mConnectorMockPingContext) Expect(p context.Context) *mConnectorMockPingContext {
	m.mockExpectations = &ConnectorMockPingContextParams{p}
	return m
}

//Return sets up a mock for Connector.PingContext to return Return's arguments
func (m *mConnectorMockPingContext) Return(r error) *ConnectorMock {
	m.mock.PingContextFunc = func(p context.Context) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.PingContext method
func (m *mConnectorMockPingContext) Set(f func(p context.Context) (r error)) *ConnectorMock {
	m.mock.PingContextFunc = f
	return m.mock
}

//PingContext implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) PingContext(p context.Context) (r error) {
	defer atomic.AddUint64(&m.PingContextCounter, 1)

	if m.PingContextMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.PingContextMock.mockExpectations, ConnectorMockPingContextParams{p},
			"Connector.PingContext got unexpected parameters")

		if m.PingContextFunc == nil {

			m.t.Fatal("No results are set for the ConnectorMock.PingContext")

			return
		}
	}

	if m.PingContextFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.PingContext")
		return
	}

	return m.PingContextFunc(p)
}

//PingContextMinimockCounter returns a count of Connector.PingContext invocations
func (m *ConnectorMock) PingContextMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.PingContextCounter)
}

type mConnectorMockPrepare struct {
	mock             *ConnectorMock
	mockExpectations *ConnectorMockPrepareParams
}

//ConnectorMockPrepareParams represents input parameters of the Connector.Prepare
type ConnectorMockPrepareParams struct {
	p string
}

//Expect sets up expected params for the Connector.Prepare
func (m *mConnectorMockPrepare) Expect(p string) *mConnectorMockPrepare {
	m.mockExpectations = &ConnectorMockPrepareParams{p}
	return m
}

//Return sets up a mock for Connector.Prepare to return Return's arguments
func (m *mConnectorMockPrepare) Return(r *sql.Stmt, r1 error) *ConnectorMock {
	m.mock.PrepareFunc = func(p string) (*sql.Stmt, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.Prepare method
func (m *mConnectorMockPrepare) Set(f func(p string) (r *sql.Stmt, r1 error)) *ConnectorMock {
	m.mock.PrepareFunc = f
	return m.mock
}

//Prepare implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) Prepare(p string) (r *sql.Stmt, r1 error) {
	defer atomic.AddUint64(&m.PrepareCounter, 1)

	if m.PrepareMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.PrepareMock.mockExpectations, ConnectorMockPrepareParams{p},
			"Connector.Prepare got unexpected parameters")

		if m.PrepareFunc == nil {

			m.t.Fatal("No results are set for the ConnectorMock.Prepare")

			return
		}
	}

	if m.PrepareFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.Prepare")
		return
	}

	return m.PrepareFunc(p)
}

//PrepareMinimockCounter returns a count of Connector.Prepare invocations
func (m *ConnectorMock) PrepareMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.PrepareCounter)
}

type mConnectorMockPrepareContext struct {
	mock             *ConnectorMock
	mockExpectations *ConnectorMockPrepareContextParams
}

//ConnectorMockPrepareContextParams represents input parameters of the Connector.PrepareContext
type ConnectorMockPrepareContextParams struct {
	p  context.Context
	p1 string
}

//Expect sets up expected params for the Connector.PrepareContext
func (m *mConnectorMockPrepareContext) Expect(p context.Context, p1 string) *mConnectorMockPrepareContext {
	m.mockExpectations = &ConnectorMockPrepareContextParams{p, p1}
	return m
}

//Return sets up a mock for Connector.PrepareContext to return Return's arguments
func (m *mConnectorMockPrepareContext) Return(r *sql.Stmt, r1 error) *ConnectorMock {
	m.mock.PrepareContextFunc = func(p context.Context, p1 string) (*sql.Stmt, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.PrepareContext method
func (m *mConnectorMockPrepareContext) Set(f func(p context.Context, p1 string) (r *sql.Stmt, r1 error)) *ConnectorMock {
	m.mock.PrepareContextFunc = f
	return m.mock
}

//PrepareContext implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) PrepareContext(p context.Context, p1 string) (r *sql.Stmt, r1 error) {
	defer atomic.AddUint64(&m.PrepareContextCounter, 1)

	if m.PrepareContextMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.PrepareContextMock.mockExpectations, ConnectorMockPrepareContextParams{p, p1},
			"Connector.PrepareContext got unexpected parameters")

		if m.PrepareContextFunc == nil {

			m.t.Fatal("No results are set for the ConnectorMock.PrepareContext")

			return
		}
	}

	if m.PrepareContextFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.PrepareContext")
		return
	}

	return m.PrepareContextFunc(p, p1)
}

//PrepareContextMinimockCounter returns a count of Connector.PrepareContext invocations
func (m *ConnectorMock) PrepareContextMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.PrepareContextCounter)
}

type mConnectorMockQuery struct {
	mock             *ConnectorMock
	mockExpectations *ConnectorMockQueryParams
}

//ConnectorMockQueryParams represents input parameters of the Connector.Query
type ConnectorMockQueryParams struct {
	p  string
	p1 []interface{}
}

//Expect sets up expected params for the Connector.Query
func (m *mConnectorMockQuery) Expect(p string, p1 ...interface{}) *mConnectorMockQuery {
	m.mockExpectations = &ConnectorMockQueryParams{p, p1}
	return m
}

//Return sets up a mock for Connector.Query to return Return's arguments
func (m *mConnectorMockQuery) Return(r *sql.Rows, r1 error) *ConnectorMock {
	m.mock.QueryFunc = func(p string, p1 ...interface{}) (*sql.Rows, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.Query method
func (m *mConnectorMockQuery) Set(f func(p string, p1 ...interface{}) (r *sql.Rows, r1 error)) *ConnectorMock {
	m.mock.QueryFunc = f
	return m.mock
}

//Query implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) Query(p string, p1 ...interface{}) (r *sql.Rows, r1 error) {
	defer atomic.AddUint64(&m.QueryCounter, 1)

	if m.QueryMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.QueryMock.mockExpectations, ConnectorMockQueryParams{p, p1},
			"Connector.Query got unexpected parameters")

		if m.QueryFunc == nil {

			m.t.Fatal("No results are set for the ConnectorMock.Query")

			return
		}
	}

	if m.QueryFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.Query")
		return
	}

	return m.QueryFunc(p, p1...)
}

//QueryMinimockCounter returns a count of Connector.Query invocations
func (m *ConnectorMock) QueryMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.QueryCounter)
}

type mConnectorMockQueryContext struct {
	mock             *ConnectorMock
	mockExpectations *ConnectorMockQueryContextParams
}

//ConnectorMockQueryContextParams represents input parameters of the Connector.QueryContext
type ConnectorMockQueryContextParams struct {
	p  context.Context
	p1 string
	p2 []interface{}
}

//Expect sets up expected params for the Connector.QueryContext
func (m *mConnectorMockQueryContext) Expect(p context.Context, p1 string, p2 ...interface{}) *mConnectorMockQueryContext {
	m.mockExpectations = &ConnectorMockQueryContextParams{p, p1, p2}
	return m
}

//Return sets up a mock for Connector.QueryContext to return Return's arguments
func (m *mConnectorMockQueryContext) Return(r *sql.Rows, r1 error) *ConnectorMock {
	m.mock.QueryContextFunc = func(p context.Context, p1 string, p2 ...interface{}) (*sql.Rows, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.QueryContext method
func (m *mConnectorMockQueryContext) Set(f func(p context.Context, p1 string, p2 ...interface{}) (r *sql.Rows, r1 error)) *ConnectorMock {
	m.mock.QueryContextFunc = f
	return m.mock
}

//QueryContext implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) QueryContext(p context.Context, p1 string, p2 ...interface{}) (r *sql.Rows, r1 error) {
	defer atomic.AddUint64(&m.QueryContextCounter, 1)

	if m.QueryContextMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.QueryContextMock.mockExpectations, ConnectorMockQueryContextParams{p, p1, p2},
			"Connector.QueryContext got unexpected parameters")

		if m.QueryContextFunc == nil {

			m.t.Fatal("No results are set for the ConnectorMock.QueryContext")

			return
		}
	}

	if m.QueryContextFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.QueryContext")
		return
	}

	return m.QueryContextFunc(p, p1, p2...)
}

//QueryContextMinimockCounter returns a count of Connector.QueryContext invocations
func (m *ConnectorMock) QueryContextMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.QueryContextCounter)
}

type mConnectorMockQueryRow struct {
	mock             *ConnectorMock
	mockExpectations *ConnectorMockQueryRowParams
}

//ConnectorMockQueryRowParams represents input parameters of the Connector.QueryRow
type ConnectorMockQueryRowParams struct {
	p  string
	p1 []interface{}
}

//Expect sets up expected params for the Connector.QueryRow
func (m *mConnectorMockQueryRow) Expect(p string, p1 ...interface{}) *mConnectorMockQueryRow {
	m.mockExpectations = &ConnectorMockQueryRowParams{p, p1}
	return m
}

//Return sets up a mock for Connector.QueryRow to return Return's arguments
func (m *mConnectorMockQueryRow) Return(r *sql.Row) *ConnectorMock {
	m.mock.QueryRowFunc = func(p string, p1 ...interface{}) *sql.Row {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.QueryRow method
func (m *mConnectorMockQueryRow) Set(f func(p string, p1 ...interface{}) (r *sql.Row)) *ConnectorMock {
	m.mock.QueryRowFunc = f
	return m.mock
}

//QueryRow implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) QueryRow(p string, p1 ...interface{}) (r *sql.Row) {
	defer atomic.AddUint64(&m.QueryRowCounter, 1)

	if m.QueryRowMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.QueryRowMock.mockExpectations, ConnectorMockQueryRowParams{p, p1},
			"Connector.QueryRow got unexpected parameters")

		if m.QueryRowFunc == nil {

			m.t.Fatal("No results are set for the ConnectorMock.QueryRow")

			return
		}
	}

	if m.QueryRowFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.QueryRow")
		return
	}

	return m.QueryRowFunc(p, p1...)
}

//QueryRowMinimockCounter returns a count of Connector.QueryRow invocations
func (m *ConnectorMock) QueryRowMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.QueryRowCounter)
}

type mConnectorMockQueryRowContext struct {
	mock             *ConnectorMock
	mockExpectations *ConnectorMockQueryRowContextParams
}

//ConnectorMockQueryRowContextParams represents input parameters of the Connector.QueryRowContext
type ConnectorMockQueryRowContextParams struct {
	p  context.Context
	p1 string
	p2 []interface{}
}

//Expect sets up expected params for the Connector.QueryRowContext
func (m *mConnectorMockQueryRowContext) Expect(p context.Context, p1 string, p2 ...interface{}) *mConnectorMockQueryRowContext {
	m.mockExpectations = &ConnectorMockQueryRowContextParams{p, p1, p2}
	return m
}

//Return sets up a mock for Connector.QueryRowContext to return Return's arguments
func (m *mConnectorMockQueryRowContext) Return(r *sql.Row) *ConnectorMock {
	m.mock.QueryRowContextFunc = func(p context.Context, p1 string, p2 ...interface{}) *sql.Row {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.QueryRowContext method
func (m *mConnectorMockQueryRowContext) Set(f func(p context.Context, p1 string, p2 ...interface{}) (r *sql.Row)) *ConnectorMock {
	m.mock.QueryRowContextFunc = f
	return m.mock
}

//QueryRowContext implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) QueryRowContext(p context.Context, p1 string, p2 ...interface{}) (r *sql.Row) {
	defer atomic.AddUint64(&m.QueryRowContextCounter, 1)

	if m.QueryRowContextMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.QueryRowContextMock.mockExpectations, ConnectorMockQueryRowContextParams{p, p1, p2},
			"Connector.QueryRowContext got unexpected parameters")

		if m.QueryRowContextFunc == nil {

			m.t.Fatal("No results are set for the ConnectorMock.QueryRowContext")

			return
		}
	}

	if m.QueryRowContextFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.QueryRowContext")
		return
	}

	return m.QueryRowContextFunc(p, p1, p2...)
}

//QueryRowContextMinimockCounter returns a count of Connector.QueryRowContext invocations
func (m *ConnectorMock) QueryRowContextMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.QueryRowContextCounter)
}

type mConnectorMockSetConnMaxLifetime struct {
	mock             *ConnectorMock
	mockExpectations *ConnectorMockSetConnMaxLifetimeParams
}

//ConnectorMockSetConnMaxLifetimeParams represents input parameters of the Connector.SetConnMaxLifetime
type ConnectorMockSetConnMaxLifetimeParams struct {
	p time.Duration
}

//Expect sets up expected params for the Connector.SetConnMaxLifetime
func (m *mConnectorMockSetConnMaxLifetime) Expect(p time.Duration) *mConnectorMockSetConnMaxLifetime {
	m.mockExpectations = &ConnectorMockSetConnMaxLifetimeParams{p}
	return m
}

//Return sets up a mock for Connector.SetConnMaxLifetime to return Return's arguments
func (m *mConnectorMockSetConnMaxLifetime) Return() *ConnectorMock {
	m.mock.SetConnMaxLifetimeFunc = func(p time.Duration) {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.SetConnMaxLifetime method
func (m *mConnectorMockSetConnMaxLifetime) Set(f func(p time.Duration)) *ConnectorMock {
	m.mock.SetConnMaxLifetimeFunc = f
	return m.mock
}

//SetConnMaxLifetime implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) SetConnMaxLifetime(p time.Duration) {
	defer atomic.AddUint64(&m.SetConnMaxLifetimeCounter, 1)

	if m.SetConnMaxLifetimeMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.SetConnMaxLifetimeMock.mockExpectations, ConnectorMockSetConnMaxLifetimeParams{p},
			"Connector.SetConnMaxLifetime got unexpected parameters")

		if m.SetConnMaxLifetimeFunc == nil {

			m.t.Fatal("No results are set for the ConnectorMock.SetConnMaxLifetime")

			return
		}
	}

	if m.SetConnMaxLifetimeFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.SetConnMaxLifetime")
		return
	}

	m.SetConnMaxLifetimeFunc(p)
}

//SetConnMaxLifetimeMinimockCounter returns a count of Connector.SetConnMaxLifetime invocations
func (m *ConnectorMock) SetConnMaxLifetimeMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SetConnMaxLifetimeCounter)
}

type mConnectorMockSetMaxIdleConns struct {
	mock             *ConnectorMock
	mockExpectations *ConnectorMockSetMaxIdleConnsParams
}

//ConnectorMockSetMaxIdleConnsParams represents input parameters of the Connector.SetMaxIdleConns
type ConnectorMockSetMaxIdleConnsParams struct {
	p int
}

//Expect sets up expected params for the Connector.SetMaxIdleConns
func (m *mConnectorMockSetMaxIdleConns) Expect(p int) *mConnectorMockSetMaxIdleConns {
	m.mockExpectations = &ConnectorMockSetMaxIdleConnsParams{p}
	return m
}

//Return sets up a mock for Connector.SetMaxIdleConns to return Return's arguments
func (m *mConnectorMockSetMaxIdleConns) Return() *ConnectorMock {
	m.mock.SetMaxIdleConnsFunc = func(p int) {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.SetMaxIdleConns method
func (m *mConnectorMockSetMaxIdleConns) Set(f func(p int)) *ConnectorMock {
	m.mock.SetMaxIdleConnsFunc = f
	return m.mock
}

//SetMaxIdleConns implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) SetMaxIdleConns(p int) {
	defer atomic.AddUint64(&m.SetMaxIdleConnsCounter, 1)

	if m.SetMaxIdleConnsMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.SetMaxIdleConnsMock.mockExpectations, ConnectorMockSetMaxIdleConnsParams{p},
			"Connector.SetMaxIdleConns got unexpected parameters")

		if m.SetMaxIdleConnsFunc == nil {

			m.t.Fatal("No results are set for the ConnectorMock.SetMaxIdleConns")

			return
		}
	}

	if m.SetMaxIdleConnsFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.SetMaxIdleConns")
		return
	}

	m.SetMaxIdleConnsFunc(p)
}

//SetMaxIdleConnsMinimockCounter returns a count of Connector.SetMaxIdleConns invocations
func (m *ConnectorMock) SetMaxIdleConnsMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SetMaxIdleConnsCounter)
}

type mConnectorMockSetMaxOpenConns struct {
	mock             *ConnectorMock
	mockExpectations *ConnectorMockSetMaxOpenConnsParams
}

//ConnectorMockSetMaxOpenConnsParams represents input parameters of the Connector.SetMaxOpenConns
type ConnectorMockSetMaxOpenConnsParams struct {
	p int
}

//Expect sets up expected params for the Connector.SetMaxOpenConns
func (m *mConnectorMockSetMaxOpenConns) Expect(p int) *mConnectorMockSetMaxOpenConns {
	m.mockExpectations = &ConnectorMockSetMaxOpenConnsParams{p}
	return m
}

//Return sets up a mock for Connector.SetMaxOpenConns to return Return's arguments
func (m *mConnectorMockSetMaxOpenConns) Return() *ConnectorMock {
	m.mock.SetMaxOpenConnsFunc = func(p int) {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.SetMaxOpenConns method
func (m *mConnectorMockSetMaxOpenConns) Set(f func(p int)) *ConnectorMock {
	m.mock.SetMaxOpenConnsFunc = f
	return m.mock
}

//SetMaxOpenConns implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) SetMaxOpenConns(p int) {
	defer atomic.AddUint64(&m.SetMaxOpenConnsCounter, 1)

	if m.SetMaxOpenConnsMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.SetMaxOpenConnsMock.mockExpectations, ConnectorMockSetMaxOpenConnsParams{p},
			"Connector.SetMaxOpenConns got unexpected parameters")

		if m.SetMaxOpenConnsFunc == nil {

			m.t.Fatal("No results are set for the ConnectorMock.SetMaxOpenConns")

			return
		}
	}

	if m.SetMaxOpenConnsFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.SetMaxOpenConns")
		return
	}

	m.SetMaxOpenConnsFunc(p)
}

//SetMaxOpenConnsMinimockCounter returns a count of Connector.SetMaxOpenConns invocations
func (m *ConnectorMock) SetMaxOpenConnsMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SetMaxOpenConnsCounter)
}

type mConnectorMockStats struct {
	mock *ConnectorMock
}

//Return sets up a mock for Connector.Stats to return Return's arguments
func (m *mConnectorMockStats) Return(r sql.DBStats) *ConnectorMock {
	m.mock.StatsFunc = func() sql.DBStats {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Connector.Stats method
func (m *mConnectorMockStats) Set(f func() (r sql.DBStats)) *ConnectorMock {
	m.mock.StatsFunc = f
	return m.mock
}

//Stats implements github.com/hexdigest/prep.Connector interface
func (m *ConnectorMock) Stats() (r sql.DBStats) {
	defer atomic.AddUint64(&m.StatsCounter, 1)

	if m.StatsFunc == nil {
		m.t.Fatal("Unexpected call to ConnectorMock.Stats")
		return
	}

	return m.StatsFunc()
}

//StatsMinimockCounter returns a count of Connector.Stats invocations
func (m *ConnectorMock) StatsMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.StatsCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *ConnectorMock) ValidateCallCounters() {

	if m.BeginFunc != nil && atomic.LoadUint64(&m.BeginCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.Begin")
	}

	if m.BeginTxFunc != nil && atomic.LoadUint64(&m.BeginTxCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.BeginTx")
	}

	if m.CloseFunc != nil && atomic.LoadUint64(&m.CloseCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.Close")
	}

	if m.DriverFunc != nil && atomic.LoadUint64(&m.DriverCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.Driver")
	}

	if m.ExecFunc != nil && atomic.LoadUint64(&m.ExecCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.Exec")
	}

	if m.ExecContextFunc != nil && atomic.LoadUint64(&m.ExecContextCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.ExecContext")
	}

	if m.PingFunc != nil && atomic.LoadUint64(&m.PingCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.Ping")
	}

	if m.PingContextFunc != nil && atomic.LoadUint64(&m.PingContextCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.PingContext")
	}

	if m.PrepareFunc != nil && atomic.LoadUint64(&m.PrepareCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.Prepare")
	}

	if m.PrepareContextFunc != nil && atomic.LoadUint64(&m.PrepareContextCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.PrepareContext")
	}

	if m.QueryFunc != nil && atomic.LoadUint64(&m.QueryCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.Query")
	}

	if m.QueryContextFunc != nil && atomic.LoadUint64(&m.QueryContextCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.QueryContext")
	}

	if m.QueryRowFunc != nil && atomic.LoadUint64(&m.QueryRowCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.QueryRow")
	}

	if m.QueryRowContextFunc != nil && atomic.LoadUint64(&m.QueryRowContextCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.QueryRowContext")
	}

	if m.SetConnMaxLifetimeFunc != nil && atomic.LoadUint64(&m.SetConnMaxLifetimeCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.SetConnMaxLifetime")
	}

	if m.SetMaxIdleConnsFunc != nil && atomic.LoadUint64(&m.SetMaxIdleConnsCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.SetMaxIdleConns")
	}

	if m.SetMaxOpenConnsFunc != nil && atomic.LoadUint64(&m.SetMaxOpenConnsCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.SetMaxOpenConns")
	}

	if m.StatsFunc != nil && atomic.LoadUint64(&m.StatsCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.Stats")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *ConnectorMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *ConnectorMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *ConnectorMock) MinimockFinish() {

	if m.BeginFunc != nil && atomic.LoadUint64(&m.BeginCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.Begin")
	}

	if m.BeginTxFunc != nil && atomic.LoadUint64(&m.BeginTxCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.BeginTx")
	}

	if m.CloseFunc != nil && atomic.LoadUint64(&m.CloseCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.Close")
	}

	if m.DriverFunc != nil && atomic.LoadUint64(&m.DriverCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.Driver")
	}

	if m.ExecFunc != nil && atomic.LoadUint64(&m.ExecCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.Exec")
	}

	if m.ExecContextFunc != nil && atomic.LoadUint64(&m.ExecContextCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.ExecContext")
	}

	if m.PingFunc != nil && atomic.LoadUint64(&m.PingCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.Ping")
	}

	if m.PingContextFunc != nil && atomic.LoadUint64(&m.PingContextCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.PingContext")
	}

	if m.PrepareFunc != nil && atomic.LoadUint64(&m.PrepareCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.Prepare")
	}

	if m.PrepareContextFunc != nil && atomic.LoadUint64(&m.PrepareContextCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.PrepareContext")
	}

	if m.QueryFunc != nil && atomic.LoadUint64(&m.QueryCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.Query")
	}

	if m.QueryContextFunc != nil && atomic.LoadUint64(&m.QueryContextCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.QueryContext")
	}

	if m.QueryRowFunc != nil && atomic.LoadUint64(&m.QueryRowCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.QueryRow")
	}

	if m.QueryRowContextFunc != nil && atomic.LoadUint64(&m.QueryRowContextCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.QueryRowContext")
	}

	if m.SetConnMaxLifetimeFunc != nil && atomic.LoadUint64(&m.SetConnMaxLifetimeCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.SetConnMaxLifetime")
	}

	if m.SetMaxIdleConnsFunc != nil && atomic.LoadUint64(&m.SetMaxIdleConnsCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.SetMaxIdleConns")
	}

	if m.SetMaxOpenConnsFunc != nil && atomic.LoadUint64(&m.SetMaxOpenConnsCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.SetMaxOpenConns")
	}

	if m.StatsFunc != nil && atomic.LoadUint64(&m.StatsCounter) == 0 {
		m.t.Fatal("Expected call to ConnectorMock.Stats")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *ConnectorMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *ConnectorMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.BeginFunc == nil || atomic.LoadUint64(&m.BeginCounter) > 0)
		ok = ok && (m.BeginTxFunc == nil || atomic.LoadUint64(&m.BeginTxCounter) > 0)
		ok = ok && (m.CloseFunc == nil || atomic.LoadUint64(&m.CloseCounter) > 0)
		ok = ok && (m.DriverFunc == nil || atomic.LoadUint64(&m.DriverCounter) > 0)
		ok = ok && (m.ExecFunc == nil || atomic.LoadUint64(&m.ExecCounter) > 0)
		ok = ok && (m.ExecContextFunc == nil || atomic.LoadUint64(&m.ExecContextCounter) > 0)
		ok = ok && (m.PingFunc == nil || atomic.LoadUint64(&m.PingCounter) > 0)
		ok = ok && (m.PingContextFunc == nil || atomic.LoadUint64(&m.PingContextCounter) > 0)
		ok = ok && (m.PrepareFunc == nil || atomic.LoadUint64(&m.PrepareCounter) > 0)
		ok = ok && (m.PrepareContextFunc == nil || atomic.LoadUint64(&m.PrepareContextCounter) > 0)
		ok = ok && (m.QueryFunc == nil || atomic.LoadUint64(&m.QueryCounter) > 0)
		ok = ok && (m.QueryContextFunc == nil || atomic.LoadUint64(&m.QueryContextCounter) > 0)
		ok = ok && (m.QueryRowFunc == nil || atomic.LoadUint64(&m.QueryRowCounter) > 0)
		ok = ok && (m.QueryRowContextFunc == nil || atomic.LoadUint64(&m.QueryRowContextCounter) > 0)
		ok = ok && (m.SetConnMaxLifetimeFunc == nil || atomic.LoadUint64(&m.SetConnMaxLifetimeCounter) > 0)
		ok = ok && (m.SetMaxIdleConnsFunc == nil || atomic.LoadUint64(&m.SetMaxIdleConnsCounter) > 0)
		ok = ok && (m.SetMaxOpenConnsFunc == nil || atomic.LoadUint64(&m.SetMaxOpenConnsCounter) > 0)
		ok = ok && (m.StatsFunc == nil || atomic.LoadUint64(&m.StatsCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.BeginFunc != nil && atomic.LoadUint64(&m.BeginCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.Begin")
			}

			if m.BeginTxFunc != nil && atomic.LoadUint64(&m.BeginTxCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.BeginTx")
			}

			if m.CloseFunc != nil && atomic.LoadUint64(&m.CloseCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.Close")
			}

			if m.DriverFunc != nil && atomic.LoadUint64(&m.DriverCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.Driver")
			}

			if m.ExecFunc != nil && atomic.LoadUint64(&m.ExecCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.Exec")
			}

			if m.ExecContextFunc != nil && atomic.LoadUint64(&m.ExecContextCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.ExecContext")
			}

			if m.PingFunc != nil && atomic.LoadUint64(&m.PingCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.Ping")
			}

			if m.PingContextFunc != nil && atomic.LoadUint64(&m.PingContextCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.PingContext")
			}

			if m.PrepareFunc != nil && atomic.LoadUint64(&m.PrepareCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.Prepare")
			}

			if m.PrepareContextFunc != nil && atomic.LoadUint64(&m.PrepareContextCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.PrepareContext")
			}

			if m.QueryFunc != nil && atomic.LoadUint64(&m.QueryCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.Query")
			}

			if m.QueryContextFunc != nil && atomic.LoadUint64(&m.QueryContextCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.QueryContext")
			}

			if m.QueryRowFunc != nil && atomic.LoadUint64(&m.QueryRowCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.QueryRow")
			}

			if m.QueryRowContextFunc != nil && atomic.LoadUint64(&m.QueryRowContextCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.QueryRowContext")
			}

			if m.SetConnMaxLifetimeFunc != nil && atomic.LoadUint64(&m.SetConnMaxLifetimeCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.SetConnMaxLifetime")
			}

			if m.SetMaxIdleConnsFunc != nil && atomic.LoadUint64(&m.SetMaxIdleConnsCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.SetMaxIdleConns")
			}

			if m.SetMaxOpenConnsFunc != nil && atomic.LoadUint64(&m.SetMaxOpenConnsCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.SetMaxOpenConns")
			}

			if m.StatsFunc != nil && atomic.LoadUint64(&m.StatsCounter) == 0 {
				m.t.Error("Expected call to ConnectorMock.Stats")
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
func (m *ConnectorMock) AllMocksCalled() bool {

	if m.BeginFunc != nil && atomic.LoadUint64(&m.BeginCounter) == 0 {
		return false
	}

	if m.BeginTxFunc != nil && atomic.LoadUint64(&m.BeginTxCounter) == 0 {
		return false
	}

	if m.CloseFunc != nil && atomic.LoadUint64(&m.CloseCounter) == 0 {
		return false
	}

	if m.DriverFunc != nil && atomic.LoadUint64(&m.DriverCounter) == 0 {
		return false
	}

	if m.ExecFunc != nil && atomic.LoadUint64(&m.ExecCounter) == 0 {
		return false
	}

	if m.ExecContextFunc != nil && atomic.LoadUint64(&m.ExecContextCounter) == 0 {
		return false
	}

	if m.PingFunc != nil && atomic.LoadUint64(&m.PingCounter) == 0 {
		return false
	}

	if m.PingContextFunc != nil && atomic.LoadUint64(&m.PingContextCounter) == 0 {
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

	if m.SetConnMaxLifetimeFunc != nil && atomic.LoadUint64(&m.SetConnMaxLifetimeCounter) == 0 {
		return false
	}

	if m.SetMaxIdleConnsFunc != nil && atomic.LoadUint64(&m.SetMaxIdleConnsCounter) == 0 {
		return false
	}

	if m.SetMaxOpenConnsFunc != nil && atomic.LoadUint64(&m.SetMaxOpenConnsCounter) == 0 {
		return false
	}

	if m.StatsFunc != nil && atomic.LoadUint64(&m.StatsCounter) == 0 {
		return false
	}

	return true
}
