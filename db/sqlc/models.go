// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type UsersRole string

const (
	UsersRoleAdmin  UsersRole = "admin"
	UsersRoleMember UsersRole = "member"
)

func (e *UsersRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UsersRole(s)
	case string:
		*e = UsersRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UsersRole: %T", src)
	}
	return nil
}

type NullUsersRole struct {
	UsersRole UsersRole `json:"users_role"`
	Valid     bool      `json:"valid"` // Valid is true if UsersRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUsersRole) Scan(value interface{}) error {
	if value == nil {
		ns.UsersRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UsersRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUsersRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UsersRole), nil
}

type UsersStatus string

const (
	UsersStatusActive   UsersStatus = "active"
	UsersStatusInactive UsersStatus = "inactive"
)

func (e *UsersStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UsersStatus(s)
	case string:
		*e = UsersStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for UsersStatus: %T", src)
	}
	return nil
}

type NullUsersStatus struct {
	UsersStatus UsersStatus `json:"users_status"`
	Valid       bool        `json:"valid"` // Valid is true if UsersStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUsersStatus) Scan(value interface{}) error {
	if value == nil {
		ns.UsersStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UsersStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUsersStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UsersStatus), nil
}

type User struct {
	ID                    uint32       `json:"id"`
	UserStatusID          uint32       `json:"user_status_id"`
	UserTypeID            uint32       `json:"user_type_id"`
	FullName              string       `json:"full_name"`
	Email                 string       `json:"email"`
	Password              string       `json:"password"`
	PhoneNumber           string       `json:"phone_number"`
	Role                  UsersRole    `json:"role"`
	Status                UsersStatus  `json:"status"`
	EmailVerifiedAt       sql.NullTime `json:"email_verified_at"`
	PhoneNumberVerifiedAt sql.NullTime `json:"phone_number_verified_at"`
	ApiKey                string       `json:"api_key"`
	CreatedAt             time.Time    `json:"created_at"`
	UpdatedAt             time.Time    `json:"updated_at"`
	DeletedAt             sql.NullTime `json:"deleted_at"`
	Version               int32        `json:"version"`
}

type UserStatus struct {
	ID          uint32       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
	Version     int32        `json:"version"`
}

type UserType struct {
	ID          uint32       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
	Version     int32        `json:"version"`
}
