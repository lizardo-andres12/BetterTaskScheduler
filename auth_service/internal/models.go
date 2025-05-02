package models

type User struct {
	User UserInfo
	Password string
}

type UserInfo struct {
	ID, TeamID uint64
	Email string
}
