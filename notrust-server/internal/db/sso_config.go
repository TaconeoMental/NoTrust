package db

import (
	"context"
)

type SsoConfig struct {
	ID           int    `json:"id"`
	Provider     string `json:"provider"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
}

func GetSsoConfigs() ([]SsoConfig, error) {
	rows, err := Pool.Query(context.Background(), "SELECT id, provider, client_id, client_secret, redirect_uri FROM sso_configs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var configs []SsoConfig
	for rows.Next() {
		var c SsoConfig
		if err := rows.Scan(&c.ID, &c.Provider, &c.ClientID, &c.ClientSecret, &c.RedirectURI); err != nil {
			return nil, err
		}
		configs = append(configs, c)
	}
	return configs, nil
}

func SaveSsoConfig(config SsoConfig) error {
	_, err := Pool.Exec(context.Background(), `
		INSERT INTO sso_configs (provider, client_id, client_secret, redirect_uri)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (provider) DO UPDATE SET client_id=$2, client_secret=$3, redirect_uri=$4
	`, config.Provider, config.ClientID, config.ClientSecret, config.RedirectURI)
	return err
}

