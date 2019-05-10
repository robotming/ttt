package regions

import (
	. "weather/models"

	"github.com/astaxie/beego/orm"
)

func OneByIdInDB(id int64) *Regions {
	if id == 0 {
		return nil
	}

	c := Regions{Id: id}
	err := orm.NewOrm().Read(&c, "Id")

	if err != nil {
		return nil
	}
	return &c
}

func AllInDB() []Regions {
	o := orm.NewOrm()
	// qs := o.QueryTable("regions")
	var maps []Regions
	// _, err := qs.All(&maps)

	_, err := o.Raw("select id, name, parent_id from regions").QueryRows(&maps)

	if err != nil {
		return nil
	}
	return maps
}
