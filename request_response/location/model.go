package location

import "time"

const (
	StatusPending   = "pending"
	StatusPublished = "published"
	StatusFailed    = "failed"
	StatusConsumed  = "consumed"
)

type LocationModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TenantID  string    `gorm:"index;not null" json:"tenant_id"`
	UserID    string    `gorm:"index;not null" json:"user_id"`
	Latitude  float64   `gorm:"not null" json:"latitude"`
	Longitude float64   `gorm:"not null" json:"longitude"`
	Timestamp time.Time `gorm:"not null" json:"timestamp"`
	Status    string    `gorm:"not null;default:'pending'" json:"status"`
}

func (LocationModel) TableName() string {
	return "location_data"
}
