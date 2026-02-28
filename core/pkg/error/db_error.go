package myerror

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	// Map constraint names/suffixes to friendly display names
	constraintMap = map[string]string{
		"uni_news_title":        "judul news",
		"uni_news_slug":         "slug news",
		"uni_departments_name":  "nama department",
		"uni_members_name":      "nama member",
		"uni_progendas_name":    "nama progenda",
		"uni_monthly_events_title": "judul event",
		"uni_nrp_whitelists_nrp": "NRP",
	}

	// Regex to extract conflicting value from Postgres error Detail
	// Example Detail: Key (name)=(Department A) already exists.
	pgValueRegex = regexp.MustCompile(`Key \(.+\)=\((.+)\) already exists`)
)

func ParseDBError(err error, entityDisplayName string) error {
	if err == nil {
		return nil
	}

	// Check if it's a Postgres error
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch pgErr.Code {
		case "23505": // unique_violation
			return handleUniqueViolation(pgErr, entityDisplayName)
		case "23503": // foreign_key_violation
			return New(fmt.Sprintf("gagal simpan %s: data terkait tidak ditemukan", entityDisplayName), http.StatusBadRequest)
		}
	}

	return err
}

func handleUniqueViolation(pgErr *pgconn.PgError, entityDisplayName string) error {
	fieldName := ""
	
	// Try to find a friendly name from the constraint map
	for constraint, friendlyName := range constraintMap {
		if strings.Contains(pgErr.ConstraintName, constraint) {
			fieldName = friendlyName
			break
		}
	}

	// Fallback: if no specific mapping, try to infer from constraint name
	if fieldName == "" {
		// common GORM naming: uni_table_field
		parts := strings.Split(pgErr.ConstraintName, "_")
		if len(parts) >= 3 {
			fieldName = parts[len(parts)-1]
		} else {
			fieldName = "atribut"
		}
	}

	// Extract the conflicting value if possible
	conflictValue := ""
	matches := pgValueRegex.FindStringSubmatch(pgErr.Detail)
	if len(matches) > 1 {
		conflictValue = matches[1]
	}

	var msg string
	if conflictValue != "" {
		msg = fmt.Sprintf("%s '%s' sudah ada (harus unik)", fieldName, conflictValue)
	} else {
		msg = fmt.Sprintf("%s sudah ada (harus unik)", fieldName)
	}

	return New(msg, http.StatusBadRequest)
}
