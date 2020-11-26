package models

const PostTableName = "t_post"

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`                                                               // 一对多的关系
	Tags  []*Tag `orm:"rel(m2m);rel_through(go-mark-learning/learn/beego_orm/models.PostTag)"` //设置多对多的关系
}

func (p *Post) TableName() string {
	return PostTableName
}
