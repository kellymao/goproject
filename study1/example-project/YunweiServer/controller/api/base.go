package api




import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	//uuid "github.com/satori/go.uuid"
	"study1/example-project/YunweiServer/controller/servers"
	"study1/example-project/YunweiServer/middleware"
	"study1/example-project/YunweiServer/model/dbModel"
	//"study1/example-project/YunweiServer/model/modelInterface"
	//"mime/multipart"
	"time"
)

var (
	USER_HEADER_IMG_PATH string = "http://qmplusimg.henrongyi.top"
	USER_HEADER_BUCKET   string = "qm-plus-img"
)

type RegistAndLoginStuct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func Login(c *gin.Context) {
	var L RegistAndLoginStuct
	_ = c.BindJSON(&L)
	U := &dbModel.User{Username: L.Username, Password: L.Password}
	if err, user := U.Login(); err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("用户名密码错误或%v", err), gin.H{})
	} else {
		tokenNext(c, *user)
	}
}


func tokenNext(c *gin.Context, user dbModel.User) {
	j := &middleware.JWT{
		[]byte("qmPlus"), // 唯一签名
	}
	clams := middleware.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24*7), // 过期时间 一周
			Issuer:    "qmPlus",                              //签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		servers.ReportFormat(c, false, "获取token失败", gin.H{})
	} else {
		servers.ReportFormat(c, true, "登录成功", gin.H{"user": user, "token": token, "expiresAt": clams.StandardClaims.ExpiresAt * 1000})
	}
}
