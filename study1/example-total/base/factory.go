package base

// 定义一个借口Class , 将结构体返回给接口即可

type Class interface {

Know()
}

var clsmap map[string]func()Class = make(map[string]func()Class)


func Registermap(name string, f func()Class){

	clsmap[name] = f

}


func Callmap(name string){

	f_func := clsmap[name]

	cls:=f_func()
	cls.Know()

}


