package models

import "github.com/astaxie/beego/orm"

// 我的喜爱
type Favorite struct {
	Id				 	int		`json:"id"`
	FavoriteId			int		`json:"favorite_id"`		// 喜爱相关联的 id
	FavoriteType		string	`json:"favorite_type"`		// 喜爱类型：收藏(collect)、赞(praise)
	UserName			string 	`json:"user_name"`
}




// Favorite 表设计优化
// 目前提供四种 api：AddFavorite、DelFavorite、IsFavorite、QueryFavoriteIds
// 分表表示：新增喜爱、删除喜爱、是否喜爱、查询所有喜爱 id
func DelFavorite(user_name string,favorite_id int,favorite_type string) {
	o := orm.NewOrm()
	o.QueryTable("favorite").Filter("user_name", user_name).Filter("favorite_id", favorite_id).
		Filter("favorite_type", favorite_type).Delete()
}
func AddFavorite(user_name string,favorite_id int,favorite_type string) {
	o := orm.NewOrm()
	var favorite Favorite
	favorite.FavoriteId = favorite_id
	favorite.FavoriteType = favorite_type
	favorite.UserName = user_name
	o.Insert(&favorite)
}

func IsFavorite(user_name string,favorite_id int,favorite_type string) (flag bool){
	o := orm.NewOrm()
	count, err := o.QueryTable("favorite").Filter("user_name", user_name).Filter("favorite_id", favorite_id).
		Filter("favorite_type", favorite_type).Count()
	if err == nil && count > 0{
		flag = true
	}
	return
}

func QueryFavoriteIds(user_name string, favorite_type string) (favorite_ids []int) {
	o := orm.NewOrm()
	var list orm.ParamsList
	// 查询所有喜爱的 id 信息
	count, err := o.QueryTable("favorite").Filter("user_name", user_name).
		Filter("favorite_type", favorite_type).ValuesFlat(&list, "favorite_id")
	if err == nil && count > 0{
		for index := range list {
			if value, ok := list[index].(int); ok {
				favorite_ids = append(favorite_ids, value)
			}
		}
	}
	return
}