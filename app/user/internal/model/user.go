package model

type UserInfo struct {
	Name      string
	Phone     string
	IsAdmin   bool
	IsTeacher bool
	IsStudent bool
	Id        int32
}
type UserModel struct {
	Name     string
	Phone    string
	Password string
}

type UserCheck struct {
	Phone    string
	Password string
}
