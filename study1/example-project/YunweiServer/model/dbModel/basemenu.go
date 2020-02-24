package dbModel

import "github.com/jinzhu/gorm"

type Meta struct {

	Title string `json:"title"`
	Icon  string `json:"icon"`
}

type BaseMenu struct {

	gorm.Model
	MenuLevel uint   `json:"-"`
	ParentId  string `json:"parentId"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
	Meta      `json:"meta"`
	NickName  string     `json:"nickName"`
	Children  []BaseMenu `json:"children"`
}

