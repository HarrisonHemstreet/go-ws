package model

import (
	"database/sql"
	"time"
)

// PartnerVendor represents a row in the partner_vendor table.
type PartnerVendor struct {
	Created       time.Time      `db:"created"`
	Edited        time.Time      `db:"edited"`
	VideoLink     sql.NullString `db:"video_link"`
	Name          string         `db:"name"`
	Description   string         `db:"description"`
	ImageLink     string         `db:"image_link"`
	ThumbnailLink string         `db:"thumbnail_link"`
	Gallery       []string       `db:"gallery"`
	ID            int            `db:"id"`
	Account       int            `db:"account"`
	Featured      int            `db:"featured"`
	ContactInfo   int            `db:"contact_info"`
}
