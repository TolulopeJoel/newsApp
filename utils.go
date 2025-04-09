package main

import (
	"database/sql"
	"time"
)

// Convert sql.NullString to string
func nullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

// Convert sql.NullTime to time.Time
func nullTimeToTime(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{} // Zero time
}
