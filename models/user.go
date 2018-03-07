package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
)

type User struct {
	Id       int `pk`
	Username string
	Passwd   string
}

func SaveUser(user User) error {
	o := orm.NewOrm()
	count, _ := o.QueryTable("user").Filter("username",user.Username).Count()
	if count > 0{
		return errors.New("用户已注册!")
	}else{
		_, err := o.Insert(&user)
		return err
	}
	return nil
}

func QueryUser(username,passwd string) (user User, err error)  {
	o := orm.NewOrm()
	err = o.QueryTable("user").Filter("username",username).Filter("passwd",passwd).One(&user)
	return
}


