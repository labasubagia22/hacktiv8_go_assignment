package controllers

import "time"

type user struct {
	ID        *uint      `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Email     *string    `json:"email,omitempty"`
	Username  *string    `json:"username,omitempty"`
	Password  *string    `json:"password,omitempty"`
	Age       *uint      `json:"age,omitempty"`
}

type photo struct {
	ID        *uint      `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Title     *string    `json:"title,omitempty"`
	URL       *string    `json:"photo_url,omitempty"`
	UserID    *uint      `json:"user_id,omitempty"`
	User      *user      `json:"user,omitempty"`
	Caption   *string    `json:"caption,omitempty"`
}

type comment struct {
	ID        *uint      `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UserID    *uint      `json:"user_id,omitempty"`
	PhotoID   *uint      `json:"photo_id,omitempty"`
	Message   *string    `json:"message,omitempty"`
	User      *user      `json:"user,omitempty"`
	Photo     *photo     `json:"photo,omitempty"`
}
