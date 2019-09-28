package main
import "fmt"
import "encoding/json"


func init(){


	fmt.Println("jl")

}

func main(){


	fmt.Println(pool)
	pool = "why"
	fmt.Println(pool)


	var person *Person = &Person {

		Name:"zhangsan",
		Age:18,
		Fav:Favorite{
			Own:"man",
			Like:"women",
		},


	}


	fmt.Println(person)

	data,err:=json.Marshal(person)  // 嵌套结构体如何序列化
	if err!=nil{

		fmt.Println(err)
	}

	fmt.Println(string(data)) //  {"name":"zhangsan","age":18,"fav":{"own":"man","like":"women"}}


	var person1  *Person = &Person{}
	err = json.Unmarshal(data,person1)
	if err!=nil{

		fmt.Println(err)
	}

	fmt.Println(*person1)  // {zhangsan 18 {man women}}




}

/*

包可以共享不同文件的全局变量，但是 import 必须是各自文件的

root@pts/0 # /workspace/bin/commom
jl
test
abc123
why


 */



