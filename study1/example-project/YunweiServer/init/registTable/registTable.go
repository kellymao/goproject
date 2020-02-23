package registTable

import (
	"github.com/jinzhu/gorm"
	"study1/example-project/YunweiServer/model/dbModel"
)

//注册数据库表专用
func RegistTable(db *gorm.DB) {
	//db.AutoMigrate(dbModel.User{}, dbModel.Authority{}, dbModel.Menu{}, dbModel.Api{}, dbModel.ApiAuthority{}, dbModel.BaseMenu{}, dbModel.FileUploadAndDownload{})
	db.AutoMigrate(dbModel.User{})
	db.AutoMigrate(dbModel.Authority{})
}

