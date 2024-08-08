package models

import (
    "time"

    "github.com/google/uuid"
)

type User struct {
    UUID      uuid.UUID `db:"uuid" json:"uuid"`
    Username  string    `db:"username" json:"username"`
    Email     string    `db:"email" json:"email"`
    Password  string    `db:"password" json:"-"`
    FirstName string    `db:"firstname" json:"firstname"`
    LastName  string    `db:"lastname" json:"lastname"`
    CreatedAt time.Time `db:"created_at" json:"created_at"`
    UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
