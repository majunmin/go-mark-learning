package models

const ProfileTableName = "t_profile"

type Profile struct {
	Id   int
	Age  int
	User *User `orm:"reverse(one)"` //设置一对一的反向关系(optional)
}

func (p *Profile) TableName() string {
	return ProfileTableName
}
