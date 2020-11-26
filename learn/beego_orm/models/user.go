package models

const UserTableName = "t_user"

type User struct {
	Id       int      `json:"id"`
	Username string   `json:"username"`
	Status   string   `json:"status"`
	Profile  *Profile `json:"profile" orm:"rel(one)"`    // 一对一的关系
	Posts    []*Post  `json:"posts" orm:"reverse(many)"` // 一对多的反向关系
}

func (context *User) TableName() string {
	return UserTableName
}
