package schemas

import (
    "time"

    "gorm.io/gorm"
)

type Works struct {
    gorm.Model
    Role     string
    Company  string
    Location string
    Remote   bool
    Link     string
    Salary   int64
}

type WorksResponse struct {
    ID        uint      `json:"id"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    DeletedAt time.Time `json:"deletedAt"`
    Role      string    `json:"role"`
    Company   string    `json:"company"`
    Location  string    `json:"location"`
    Remote    bool      `json:"remote"`
    Link      string    `json:"link"`
    Salary    int64     `json:"salary"`
}
