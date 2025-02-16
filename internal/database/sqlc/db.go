// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createAgentStmt, err = db.PrepareContext(ctx, createAgent); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAgent: %w", err)
	}
	if q.createRemoteClientStmt, err = db.PrepareContext(ctx, createRemoteClient); err != nil {
		return nil, fmt.Errorf("error preparing query CreateRemoteClient: %w", err)
	}
	if q.deleteAgentStmt, err = db.PrepareContext(ctx, deleteAgent); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteAgent: %w", err)
	}
	if q.deleteRemoteClientStmt, err = db.PrepareContext(ctx, deleteRemoteClient); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteRemoteClient: %w", err)
	}
	if q.getAgentStmt, err = db.PrepareContext(ctx, getAgent); err != nil {
		return nil, fmt.Errorf("error preparing query GetAgent: %w", err)
	}
	if q.getAgentByNameStmt, err = db.PrepareContext(ctx, getAgentByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetAgentByName: %w", err)
	}
	if q.getRemoteClientStmt, err = db.PrepareContext(ctx, getRemoteClient); err != nil {
		return nil, fmt.Errorf("error preparing query GetRemoteClient: %w", err)
	}
	if q.getRemoteClientByNameStmt, err = db.PrepareContext(ctx, getRemoteClientByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetRemoteClientByName: %w", err)
	}
	if q.listAgentsStmt, err = db.PrepareContext(ctx, listAgents); err != nil {
		return nil, fmt.Errorf("error preparing query ListAgents: %w", err)
	}
	if q.listRemoteClientsStmt, err = db.PrepareContext(ctx, listRemoteClients); err != nil {
		return nil, fmt.Errorf("error preparing query ListRemoteClients: %w", err)
	}
	if q.updateAgentStatusStmt, err = db.PrepareContext(ctx, updateAgentStatus); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateAgentStatus: %w", err)
	}
	if q.updateRemoteClientStatusStmt, err = db.PrepareContext(ctx, updateRemoteClientStatus); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateRemoteClientStatus: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createAgentStmt != nil {
		if cerr := q.createAgentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAgentStmt: %w", cerr)
		}
	}
	if q.createRemoteClientStmt != nil {
		if cerr := q.createRemoteClientStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createRemoteClientStmt: %w", cerr)
		}
	}
	if q.deleteAgentStmt != nil {
		if cerr := q.deleteAgentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteAgentStmt: %w", cerr)
		}
	}
	if q.deleteRemoteClientStmt != nil {
		if cerr := q.deleteRemoteClientStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteRemoteClientStmt: %w", cerr)
		}
	}
	if q.getAgentStmt != nil {
		if cerr := q.getAgentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAgentStmt: %w", cerr)
		}
	}
	if q.getAgentByNameStmt != nil {
		if cerr := q.getAgentByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAgentByNameStmt: %w", cerr)
		}
	}
	if q.getRemoteClientStmt != nil {
		if cerr := q.getRemoteClientStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRemoteClientStmt: %w", cerr)
		}
	}
	if q.getRemoteClientByNameStmt != nil {
		if cerr := q.getRemoteClientByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRemoteClientByNameStmt: %w", cerr)
		}
	}
	if q.listAgentsStmt != nil {
		if cerr := q.listAgentsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listAgentsStmt: %w", cerr)
		}
	}
	if q.listRemoteClientsStmt != nil {
		if cerr := q.listRemoteClientsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listRemoteClientsStmt: %w", cerr)
		}
	}
	if q.updateAgentStatusStmt != nil {
		if cerr := q.updateAgentStatusStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateAgentStatusStmt: %w", cerr)
		}
	}
	if q.updateRemoteClientStatusStmt != nil {
		if cerr := q.updateRemoteClientStatusStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateRemoteClientStatusStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                           DBTX
	tx                           *sql.Tx
	createAgentStmt              *sql.Stmt
	createRemoteClientStmt       *sql.Stmt
	deleteAgentStmt              *sql.Stmt
	deleteRemoteClientStmt       *sql.Stmt
	getAgentStmt                 *sql.Stmt
	getAgentByNameStmt           *sql.Stmt
	getRemoteClientStmt          *sql.Stmt
	getRemoteClientByNameStmt    *sql.Stmt
	listAgentsStmt               *sql.Stmt
	listRemoteClientsStmt        *sql.Stmt
	updateAgentStatusStmt        *sql.Stmt
	updateRemoteClientStatusStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                           tx,
		tx:                           tx,
		createAgentStmt:              q.createAgentStmt,
		createRemoteClientStmt:       q.createRemoteClientStmt,
		deleteAgentStmt:              q.deleteAgentStmt,
		deleteRemoteClientStmt:       q.deleteRemoteClientStmt,
		getAgentStmt:                 q.getAgentStmt,
		getAgentByNameStmt:           q.getAgentByNameStmt,
		getRemoteClientStmt:          q.getRemoteClientStmt,
		getRemoteClientByNameStmt:    q.getRemoteClientByNameStmt,
		listAgentsStmt:               q.listAgentsStmt,
		listRemoteClientsStmt:        q.listRemoteClientsStmt,
		updateAgentStatusStmt:        q.updateAgentStatusStmt,
		updateRemoteClientStatusStmt: q.updateRemoteClientStatusStmt,
	}
}
