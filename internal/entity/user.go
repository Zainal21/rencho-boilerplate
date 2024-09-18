package entity

import "time"

type User struct {
	ID          int    `db:"id" json:"id"`
	UID         string `db:"uid" json:"uid"`
	FirebaseUID string `db:"firebase_uid" json:"firebase_uid"`
	Email       string `db:"email" json:"email"`
	Name        string `db:"name" json:"name"`
	Phone       string `db:"phone" json:"phone"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
