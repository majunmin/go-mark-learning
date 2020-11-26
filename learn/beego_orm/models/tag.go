package models

const TagTableName = "t_tag"

type Tag struct {
	Id   int
	Name string
	Post []*Post `orm:"reverse(many)"`
}

func (t *Tag) TableName() string {
	return TagTableName
}
