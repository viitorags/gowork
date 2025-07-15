package config

import (
    _ "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var (
    db *gorm.DB
)

func Init() error {
    return nil
}
