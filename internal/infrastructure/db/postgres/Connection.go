package postgres

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewConnection() (*gorm.DB, error) {
	return gorm.Open("postgres", "User ID=fximrwid;Password=R-30tgm9-g5nxXgppwJuKvbay_zh1Nrm;Host=castor.db.elephantsql.com;Port=5432;Database=fximrwid;Pooling=true;Min Pool Size=0;Max Pool Size=100;")
}
