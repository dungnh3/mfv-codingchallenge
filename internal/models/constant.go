package models

type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInActive UserStatus = "inactive"
)

func (s UserStatus) String() string { return string(s) }
