package factory

type Class interface{

	Do()
}


var clsmap map[string]func()Class  = make(map[string]func()Class)



func Register_cls(name string, f func()Class){



	clsmap[name] = f
}


func Create(name string) Class {



	if rel,ok:= clsmap[name]; ok{

		return rel()
	}else{

		panic("name is not found")
	}
}