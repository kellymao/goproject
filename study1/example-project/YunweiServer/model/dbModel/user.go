package dbModel

import (
	"fmt"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	//"main/controller/servers"
	"study1/example-project/YunweiServer/init/qmsql"
	//"main/model/modelInterface"
	"study1/example-project/YunweiServer/tools"
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
