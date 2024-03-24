package repository

import (
	"context"
	"database/sql"

	"github.com/Shteyd/wallet-app/src/backend/internal/core/contract"
)

var _ contract.RepositoryManager = (*Manager)(nil)

type Manager struct {
	db *sql.DB
	tx *sql.Tx
}

func NewManager(db *sql.DB, tx *sql.Tx) *Manager {
	return &Manager{
		db: db,
		tx: tx,
	}
}

// Begin implements contract.RepositoryManager.
func (m *Manager) Begin(ctx context.Context) (contract.RepositoryManager, error) {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return NewManager(m.db, tx), nil
}

func (m *Manager) clearTx() {
	m.tx = nil
}

// Commit implements contract.RepositoryManager.
func (m *Manager) Commit(_ context.Context) error {
	defer m.clearTx()

	return m.tx.Commit()
}

// Rollback implements contract.RepositoryManager.
func (m *Manager) Rollback(ctx context.Context) error {
	defer m.clearTx()

	return m.tx.Rollback()
}

func (m *Manager) conn() contract.DBTX {
	if m.tx != nil {
		return m.tx
	}

	return m.db
}

// UserRepo implements contract.RepositoryManager.
func (m *Manager) UserRepo() contract.UserRepository {
	panic("unimplemented")
}
