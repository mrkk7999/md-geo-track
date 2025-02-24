package repository

import (
	mdgeotrack "md-geo-track/iface"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type repository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func New(db *gorm.DB, log *logrus.Logger) mdgeotrack.Repository {
	return &repository{
		db:  db,
		log: log,
	}
}
