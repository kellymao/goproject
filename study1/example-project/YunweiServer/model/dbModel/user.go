package dbModel

import (
	"fmt"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	//"main/controller/servers"
	"study1/example-project/YunweiServer/init/qmsql"
	//"main/model/modelInterface"
	"study1/example-project/YunweiServer/tools"
	"study1/example-project/YunweiServer/model/modelInterface"

)


type User struct {
	gorm.Model
	UUID        uuid.UUID `json:"uuid"`
	Username    string    `json:"userName"`
	Password    string    `json:"-"`
	NickName    string    `json:"nickName" gorm:"default:'QMPlusUser'"`     // 表字段名默认以大写字母来分隔为下划线
	HeaderImg   string    `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	Authority   Authority `json:"authority" gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId"`
	AuthorityId string    `json:"-" gorm:"default:888"`
	//Propertie                //	多余属性自行添加
	//PropertieId uint  // 自动关联 Propertie 的Id 附加属性过多 建议创建一对一关系
}




//用户登录
func (u *User) Login() (err error, userInter *User) {
	var user User
	u.Password = tools.MD5V(u.Password)
	err = qmsql.DEFAULTDB.Where("username = ? AND password = ? ", u.Username,u.Password).First(&user).Error
	//err = qmsql.DEFAULTDB.Where("authority_id = ?", user.AuthorityId).First(&user.Authority).Error

	fmt.Printf("%+v\n",user)
	fmt.Println(err)
	return err, &user
}


func (a *User) GetInfoList(info modelInterface.PageInfo) (err error, userList []User, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	//err, db, total := servers.PagingServer(a, info)
	//if err != nil {
	//	return
	//} else {
	//	var apiList []Authority
	//
	//	//err = qmsql.DEFAULTDB.Where("path LIKE ?", "%"+a.AuthorityName+"%").Find(&apiList).Count(&total).Error
	//	//db.Where(&User{Name: "jinzhu", Age: 20}).First(&user) /// SELECT * FROM users ..
	//
	//
	//	if  a.AuthorityId != ""  ||  a.AuthorityName != "" {
	//
	//		fmt.Println("true")
	//		err = db.Where("authority_id = ?" , a.AuthorityId).Or("authority_name = ?",a.AuthorityName).Find(&apiList).Count(&total).Error
	//
	//
	//	} else{
	//		err = db.Find(&apiList).Count(&total).Error
	//
	//
	//	}
	//
	//
	//	if err!=nil{
	//		return err, apiList, total
	//	}
	//	return nil, apiList, total
	//}

	var db *gorm.DB
	if a.Username != ""  ||  a.AuthorityId != "" {
		db = qmsql.DEFAULTDB.Model(&User{}).Where("authority_id = ?" , a.AuthorityId).Or("username = ?",a.Username).Count(&total)
	}else{
		db = qmsql.DEFAULTDB.Model(&User{}).Count(&total)


	}

	if err = db.Error;err!=nil{
		fmt.Println("获取userlist total失败！",err )
		return
	}

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	if limit >0 {
		// 客户端获取分页
		err=db.Limit(limit).Offset(offset).Order("id ").Find(&userList).Error
	}else{
		// 不获取分页
		err=db.Order("id ").Find(&userList).Error

	}

	for i:=0; i<len(userList);i++{

		err = qmsql.DEFAULTDB.Model(&Authority{}).Where("authority_id = ?",userList[i].AuthorityId).Find(&userList[i].Authority).Error
	}

	if err!=nil{
		fmt.Println("获取userlist 分页失败！",err )
		return
	}

	return





}