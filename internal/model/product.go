package model

import (
	"database/sql"
	"time"
)

// PartnerVendor represents a row in the partner_vendor table.
type PartnerVendor struct {
	ID            int            `db:"id"`
	Account       int            `db:"account"` // Assuming `account` references another table's ID.
	Name          string         `db:"name"`
	Description   string         `db:"description"`
	VideoLink     sql.NullString `db:"video_link"` // Nullable fields can be represented using sql.NullString
	ImageLink     string         `db:"image_link"`
	ThumbnailLink string         `db:"thumbnail_link"`
	Gallery       []string       `db:"gallery"` // PostgreSQL specific array represented as a slice in Go.
	Featured      int            `db:"featured"`
	ContactInfo   int            `db:"contact_info"` // Assuming `contact_info` references another table's ID.
	Created       time.Time      `db:"created"`
	Edited        time.Time      `db:"edited"`
}
