package db

import (
	"context"
)

type OtpDomain struct {
	ID     int    `json:"id"`
	Domain string `json:"domain"`
}

type NewOtpDomain struct {
	Domain string `json:"domain"`
}

func GetOtpDomains() ([]OtpDomain, error) {
	rows, err := Pool.Query(context.Background(), "SELECT id, domain FROM otp_domains")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var domains []OtpDomain
	for rows.Next() {
		var d OtpDomain
		if err := rows.Scan(&d.ID, &d.Domain); err != nil {
			return nil, err
		}
		domains = append(domains, d)
	}
	return domains, nil
}

func CreateOtpDomain(domain NewOtpDomain) error {
	_, err := Pool.Exec(context.Background(), "INSERT INTO otp_domains (domain) VALUES ($1)", domain.Domain)
	return err
}

func DeleteOtpDomain(id string) error {
	_, err := Pool.Exec(context.Background(), "DELETE FROM otp_domains WHERE id=$1", id)
	return err
}

