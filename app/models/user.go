package models

import "time"

type UserType int

const (
	Student    UserType = 0
	Counsellor UserType = 1
	Admin      UserType = 2
)

type User struct {
	// ID作为账号的唯一标识符
	ID         string // 账号唯一标识符
	UserID     string
	Username   string // 用户名
	Password   string // 数据库
	Type       UserType
	School     string    // 学校
	From       string    // 生源地
	Grade      int       // 年级
	Profession string    // 专业
	Email      string    // 邮箱
	Telephone  string    // 手机号码
	CreateTime time.Time // 创建时间
}
