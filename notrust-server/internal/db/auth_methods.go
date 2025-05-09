package db

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type AuthMethod struct {
	Method  string `json:"method"`
	Enabled bool   `json:"enabled"`
}

type AuthMethodUpdate struct {
	Method  string `json:"method"`
	Enabled bool   `json:"enabled"`
}

func GetAuthMethods() ([]AuthMethod, error) {
	rows, err := Pool.Query(context.Background(), "SELECT method, enabled FROM auth_methods")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var methods []AuthMethod
	for rows.Next() {
		var m AuthMethod
		if err := rows.Scan(&m.Method, &m.Enabled); err != nil {
			return nil, err
		}
		methods = append(methods, m)
	}
	return methods, nil
}

func UpdateAuthMethods(updates []AuthMethodUpdate) error {
	batch := &pgx.Batch{}
	for _, update := range updates {
		batch.Queue("UPDATE auth_methods SET enabled=$1 WHERE method=$2", update.Enabled, update.Method)
	}
	br := Pool.SendBatch(context.Background(), batch)
	return br.Close()
}

