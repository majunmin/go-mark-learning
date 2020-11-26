/**
  @Author: majm@ushareit.com
  @date: 2020/11/25
  @note:
**/
package models

const PostTagTableName = "t_post_tag"

type PostTag struct {
	Id   int
	Post *Post `orm:"rel(fk)"`
	Tag  *Tag  `orm:"rel(fk)"`
}

func (t PostTag) TableName() string {
	return PostTagTableName
}
