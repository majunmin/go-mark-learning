/**
  @Author: majm@ushareit.com
  @date: 2020/11/24
  @note:
**/
package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"go-mark-learning/learn/beego_orm/models"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/learn?charset=utf8", 30)
}

func main() {
	//orm.Debug = true
	//o := orm.NewOrm()
	//
	//post := models.Post{Id: 1}
	//m2m := o.QueryM2M(&post, "Tags")
	//
	//tag := &models.Tag{Id: 14, Name: "golang"}
	//
	//num, err := m2m.Add(tag)
	//if err != nil {
	//	fmt.Println("err,", err)
	//}
	//fmt.Println("add Nums: ", num)

	orm.Debug = true
	o := orm.NewOrm()

	user := &models.User{}
	// 关联查询到 一对一的表 数据
	// [SELECT T0.`id`, T0.`username`, T0.`status`, T0.`profile_id`, T1.`id`, T1.`age`
	//  FROM `t_user` T0
	//  INNER JOIN `t_profile` T1 ON T1.`id` = T0.`profile_id`
	//  WHERE T0.`id` = ? LIMIT 1] - `1`
	o.QueryTable(models.UserTableName).Filter("Id", 1).RelatedSel().One(user)

	o.LoadRelated(user, "Posts")

	profile := &models.Profile{}
	o.QueryTable(models.ProfileTableName).Filter("User__Id", 1).RelatedSel().One(profile)

	var posts []*models.Post
	num, err := o.QueryTable(models.PostTableName).Filter("User", 1).RelatedSel().All(&posts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

	// [SELECT T0.`id`, T0.`title`, T0.`user_id` FROM `t_post` T0
	//  INNER JOIN `t_post_tag` T1 ON T1.`post_id` = T0.`id`
	//  INNER JOIN `t_tag` T2 ON T2.`id` = T1.`tag_id`
	// WHERE T2.`name` = ? ] - `golang`
	num, err = o.QueryTable(models.PostTableName).Filter("Tags__Tag__Name", "golang").All(&posts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

	var userTmp *models.User
	o.Raw("select * from t_user where id = 1").QueryRow(&userTmp)
	fmt.Println(userTmp)

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From(models.TagTableName + " tag").
		InnerJoin(models.PostTagTableName + " pt").On("pt.tag_id = tag.id").
		InnerJoin(models.PostTableName + " post").On("pt.post_id = post.id").
		Where("post.user_id = 1")

	fmt.Println(qb.String())

}
