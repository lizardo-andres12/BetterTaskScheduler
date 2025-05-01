package models

type User struct {
	User UserInfo
	Password PasswordInfo
}

type UserInfo struct {
	ID, TeamID uint64
	Email string
}

type PasswordInfo struct {
	Password, Salt string
}
