package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Regions struct {
	Id       int64
	Name     string
	ParentId string
}

type ChildrenRegions struct {
	Id       int64
	Name     string
	ParentId string
	Children []interface{}
}

func (*Regions) TableEngine() string {
	return engine()
}

func engine() string {
	return "INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci"
}

func init() {
	orm.RegisterModelWithPrefix("", new(Regions))
}
