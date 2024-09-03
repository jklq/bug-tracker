package helpers

import (
	"github.com/jackc/pgx/v5/pgtype"
)

func ParseDate(date pgtype.Timestamptz) string {
	return date.Time.Format("2. January, 2006")
}

func ParseTime(date pgtype.Timestamptz) string {
	return date.Time.Format("15:04")
}
