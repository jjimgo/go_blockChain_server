// Code generated by sqlc. DO NOT EDIT.
// source: evmLaunchpad.sql

package launchpad

import (
	"context"
	"database/sql"
)

const createEvmLaunchpad = `-- name: CreateEvmLaunchpad :execresult
INSERT INTO evmLaunchpad (
    eoa_address,
    ca_address,
    chain_id,
    created_at
) VALUES (
   ?, ?, ?, ?
)
`

type CreateEvmLaunchpadParams struct {
	EoaAddress string `json:"eoa_address"`
	CaAddress  string `json:"ca_address"`
	ChainID    int32  `json:"chain_id"`
	CreatedAt  string `json:"created_at"`
}

func (q *Queries) CreateEvmLaunchpad(ctx context.Context, arg CreateEvmLaunchpadParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createEvmLaunchpad,
		arg.EoaAddress,
		arg.CaAddress,
		arg.ChainID,
		arg.CreatedAt,
	)
}

const getMyAllLaunchpad = `-- name: GetMyAllLaunchpad :many
SELECT id, eoa_address, ca_address, chain_id, price, created_at FROM evmLaunchpad
`

func (q *Queries) GetMyAllLaunchpad(ctx context.Context) ([]EvmLaunchpad, error) {
	rows, err := q.db.QueryContext(ctx, getMyAllLaunchpad)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []EvmLaunchpad{}
	for rows.Next() {
		var i EvmLaunchpad
		if err := rows.Scan(
			&i.ID,
			&i.EoaAddress,
			&i.CaAddress,
			&i.ChainID,
			&i.Price,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMyLaunchpad = `-- name: GetMyLaunchpad :one
SELECT id, eoa_address, ca_address, chain_id, price, created_at FROM evmLaunchpad
WHERE eoa_address = ? LIMIT 1
`

func (q *Queries) GetMyLaunchpad(ctx context.Context, eoaAddress string) (EvmLaunchpad, error) {
	row := q.db.QueryRowContext(ctx, getMyLaunchpad, eoaAddress)
	var i EvmLaunchpad
	err := row.Scan(
		&i.ID,
		&i.EoaAddress,
		&i.CaAddress,
		&i.ChainID,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}
