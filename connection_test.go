package prep

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConnection(t *testing.T) {
	t.Run("prepare fails", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		connectorMock := NewConnectorMock(mc)
		connectorMock.PrepareMock.
			Expect("SELECT 'prepare failure'").
			Return(nil, errors.New("prepare failed"))

		c, err := NewConnection(connectorMock, []string{"SELECT 'prepare failure'"})
		assert.Error(t, err)
		assert.Nil(t, c)
	})

	t.Run("success", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		connectorMock := NewConnectorMock(mc)
		connectorMock.PrepareMock.
			Expect("SELECT 'success'").
			Return(&sql.Stmt{}, nil)

		c, err := NewConnection(connectorMock, []string{"SELECT 'success'"})
		assert.NoError(t, err)
		assert.NotNil(t, c)
	})
}

func TestConnection_PrepareContext(t *testing.T) {
	t.Run("statement found", func(t *testing.T) {
		c := &Connection{
			statements: preparedStatements{"SELECT 1": &sql.Stmt{}},
		}

		stmt, err := c.PrepareContext(nil, "SELECT 1")
		assert.NoError(t, err)
		assert.NotNil(t, stmt)
	})

	t.Run("invalid statement type", func(t *testing.T) {
		c := &Connection{
			statements: preparedStatements{"SELECT 1": NewPreparedStatementMock(t)},
		}

		stmt, err := c.PrepareContext(nil, "SELECT 1")
		require.Error(t, err)
		assert.Equal(t, errInvalidStatementType, err)
		assert.Nil(t, stmt)
	})

	t.Run("statement not found", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		dbQuerierMock := NewQuerierMock(mc)
		dbQuerierMock.PrepareContextMock.Expect(nil, "SELECT 1").Return(&sql.Stmt{}, nil)
		c := &Connection{
			statements: preparedStatements{},
			Querier:    dbQuerierMock,
		}

		stmt, err := c.PrepareContext(nil, "SELECT 1")
		assert.NoError(t, err)
		assert.NotNil(t, stmt)
	})
}

func TestConnection_Prepare(t *testing.T) {
	t.Run("statement found", func(t *testing.T) {
		c := &Connection{
			statements: preparedStatements{"SELECT 1": &sql.Stmt{}},
		}

		stmt, err := c.Prepare("SELECT 1")
		assert.NoError(t, err)
		assert.NotNil(t, stmt)
	})

	t.Run("invalid statement type", func(t *testing.T) {
		c := &Connection{
			statements: preparedStatements{"SELECT 1": NewPreparedStatementMock(t)},
		}

		stmt, err := c.Prepare("SELECT 1")
		require.Error(t, err)
		assert.Equal(t, errInvalidStatementType, err)
		assert.Nil(t, stmt)
	})

	t.Run("statement not found", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		dbQuerierMock := NewQuerierMock(mc)
		dbQuerierMock.PrepareMock.Expect("SELECT 1").Return(&sql.Stmt{}, nil)
		c := &Connection{
			statements: preparedStatements{},
			Querier:    dbQuerierMock,
		}

		stmt, err := c.Prepare("SELECT 1")
		assert.NoError(t, err)
		assert.NotNil(t, stmt)
	})
}

func TestConnection_ExecContext(t *testing.T) {
	t.Run("statement found", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		c := &Connection{
			statements: preparedStatements{
				"SELECT 1": NewPreparedStatementMock(mc).ExecContextMock.Expect(nil, 1, 2, 3).Return(nil, nil),
			},
		}

		result, err := c.ExecContext(nil, "SELECT 1", 1, 2, 3)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("statement not found", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		dbQuerierMock := NewQuerierMock(mc)
		dbQuerierMock.ExecContextMock.Expect(nil, "SELECT 1", 1, 2, 3).Return(nil, nil)

		c := &Connection{
			statements: preparedStatements{},
			Querier:    dbQuerierMock,
		}

		result, err := c.ExecContext(nil, "SELECT 1", 1, 2, 3)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})
}

func TestConnection_Exec(t *testing.T) {
	t.Run("statement found", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		c := &Connection{
			statements: preparedStatements{
				"SELECT 1": NewPreparedStatementMock(mc).ExecMock.Expect(1, 2, 3).Return(nil, nil),
			},
		}

		result, err := c.Exec("SELECT 1", 1, 2, 3)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("statement not found", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		dbQuerierMock := NewQuerierMock(mc)
		dbQuerierMock.ExecMock.Expect("SELECT 1", 1, 2, 3).Return(nil, nil)

		c := &Connection{
			statements: preparedStatements{},
			Querier:    dbQuerierMock,
		}

		result, err := c.Exec("SELECT 1", 1, 2, 3)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})
}

func TestConnection_QueryContext(t *testing.T) {
	t.Run("statement found", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		c := &Connection{
			statements: preparedStatements{
				"SELECT 1": NewPreparedStatementMock(mc).QueryContextMock.Expect(nil, 1, 2, 3).Return(&sql.Rows{}, nil),
			},
		}

		rows, err := c.QueryContext(nil, "SELECT 1", 1, 2, 3)
		assert.NoError(t, err)
		assert.NotNil(t, rows)
	})

	t.Run("statement not found", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		dbQuerierMock := NewQuerierMock(mc)
		dbQuerierMock.QueryContextMock.Expect(nil, "SELECT 1", 1, 2, 3).Return(&sql.Rows{}, nil)

		c := &Connection{
			statements: preparedStatements{},
			Querier:    dbQuerierMock,
		}

		rows, err := c.QueryContext(nil, "SELECT 1", 1, 2, 3)
		assert.NoError(t, err)
		assert.NotNil(t, rows)
	})
}

func TestConnection_Query(t *testing.T) {
	t.Run("statement found", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		c := &Connection{
			statements: preparedStatements{
				"SELECT 1": NewPreparedStatementMock(mc).QueryMock.Expect(1, 2, 3).Return(&sql.Rows{}, nil),
			},
		}

		rows, err := c.Query("SELECT 1", 1, 2, 3)
		assert.NoError(t, err)
		assert.NotNil(t, rows)
	})

	t.Run("statement not found", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		dbQuerierMock := NewQuerierMock(mc)
		dbQuerierMock.QueryMock.Expect("SELECT 1", 1, 2, 3).Return(&sql.Rows{}, nil)

		c := &Connection{
			statements: preparedStatements{},
			Querier:    dbQuerierMock,
		}

		rows, err := c.Query("SELECT 1", 1, 2, 3)
		assert.NoError(t, err)
		assert.NotNil(t, rows)
	})
}

func TestConnection_QueryRowContext(t *testing.T) {
	t.Run("statement found", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		c := &Connection{
			statements: preparedStatements{
				"SELECT 1": NewPreparedStatementMock(mc).QueryRowContextMock.Expect(nil, 1, 2, 3).Return(&sql.Row{}),
			},
		}

		row := c.QueryRowContext(nil, "SELECT 1", 1, 2, 3)
		assert.NotNil(t, row)
	})

	t.Run("statement not found", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		dbQuerierMock := NewQuerierMock(mc)
		dbQuerierMock.QueryRowContextMock.Expect(nil, "SELECT 1", 1, 2, 3).Return(&sql.Row{})

		c := &Connection{
			statements: preparedStatements{},
			Querier:    dbQuerierMock,
		}

		row := c.QueryRowContext(nil, "SELECT 1", 1, 2, 3)
		assert.NotNil(t, row)
	})
}

func TestConnection_QueryRow(t *testing.T) {
	t.Run("statement found", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		c := &Connection{
			statements: preparedStatements{
				"SELECT 1": NewPreparedStatementMock(mc).QueryRowMock.Expect(1, 2, 3).Return(&sql.Row{}),
			},
		}

		row := c.QueryRow("SELECT 1", 1, 2, 3)
		assert.NotNil(t, row)
	})

	t.Run("statement not found", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		dbQuerierMock := NewQuerierMock(mc)
		dbQuerierMock.QueryRowMock.Expect("SELECT 1", 1, 2, 3).Return(&sql.Row{})

		c := &Connection{
			statements: preparedStatements{},
			Querier:    dbQuerierMock,
		}

		rows := c.QueryRow("SELECT 1", 1, 2, 3)
		assert.NotNil(t, rows)
	})
}
