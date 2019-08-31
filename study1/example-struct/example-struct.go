package main

import "fmt"
import "math/rand"
import "encoding/json"

import "time"

type Student struct{

	Name string
	Age int
	Score float32

}

type Student_single_list struct{

	Name string
	Age int
	Score float32
	Next *Student_single_list

}

// head 是链表第一个元素
func insertbefore(head *Student_single_list){


	for j:=0 ; j<2;j++{

		/*
		变量不可以重复定义， 但是for 循环里面可以重复定义，应该是每次for循环里面的同名变量都是一个局部变量 。
		 */
		var a *int = new(int)
		*a = 3
		fmt.Println(a)
	}


	for i:=0 ;i<100;i++{
		var p *Student_single_list = new(Student_single_list)
		p.Name = fmt.Sprintf("stu_%d",i)
		p.Age = rand.Intn(100)
		p.Score = float32(rand.Float64() * 100)
		head.Next = p
		head = p

	}



}

// head 是链表最后一个元素
func inserttail(head *(*Student_single_list)) {




	for i:=0 ;i<100;i++{
		var p *Student_single_list = new(Student_single_list)
		p.Name = fmt.Sprintf("stu_%d",i)
		p.Age = rand.Intn(100)
		p.Score = float32(rand.Float64() * 100)
		p.Next = *head
		*head = p

	}
	fmt.Println(*head)


	fmt.Println(head)
	fmt.Printf("%p\n",head)



}

// 删除链表特定的节点

func delnode(stu *Student_single_list, name string){

	prev := stu
	for stu !=nil {


		if stu.Name == name{

			prev.Next = stu.Next
			break

		}
		prev = stu
		stu = stu.Next

	}
}

func trans(stu *Student_single_list){
	for stu != nil{

		fmt.Println(stu)
		stu = stu.Next
	}

}
// 插入新节点到链表特定位置

func insnode(stu *Student_single_list, name string, newstu *Student_single_list){
	for stu != nil {
		fmt.Println(stu)
		if stu.Name == name {
			newstu.Next = stu.Next
			stu.Next = newstu

		}

		stu = stu.Next

	}

}
// 二叉树

type stutree struct{

	name string
	age int
	score int
	left  *stutree
	right *stutree
}

func funtree(){

	left2 := &stutree{"left2",30,93,nil,nil}
	left1 := &stutree{"left1",20,91,left2,nil}

	right1 := &stutree{"right1",20,91,nil,nil}

	var root *stutree = new(stutree)
	root.name ="root"
	root.age = 10
	root.score = 90

	root.left = left1
	root.right  = right1

	fmt.Println("先遍历根节点，在从左到右遍历")
	bianlitree_center(root)
	fmt.Println("先遍历left节点，在遍历root,在遍历right")
	bianlitree_left(root)
	fmt.Println("先遍历在遍历right节点，在遍历root,在遍历left节点")
	bianlitree_right(root)



}


func bianlitree_center(root *stutree){

	if root == nil {

		return

	}

	fmt.Println(root)
	bianlitree_center(root.left)
	bianlitree_center(root.right)

}

func bianlitree_left(root *stutree){

	if root == nil {

		return

	}


	bianlitree_left(root.left)
	fmt.Println(root)
	bianlitree_left(root.right)

}

func bianlitree_right(root *stutree){

	if root == nil {

		return

	}
	bianlitree_right(root.right)
	fmt.Println(root)
	bianlitree_right(root.left)


}
// 结构体如何实现构造函数
func Newstudent(name string,age int ,score int) *student_json{

	return &student_json{
		Name:name,
		Age:age,
		Score:score}
}

func newinit(){

	st := Newstudent("zhagnwu",30,86)

	fmt.Println(st)
}

// 结构体tag 与json序列化

type student_json struct{
	Name string `json:"sutdent_name"`
	Age int		`json:"age"`
	Score int	`json:"score"`

}

func funjson(){

	var stu student_json = student_json{"zhangsan",
		10,98}


	data,err:=json.Marshal(stu)
	if err!= nil {

		fmt.Println("json encoding err",err)
		return
	}
	fmt.Println(data)  // [123 125]，json 默认是[] byte 类型
	fmt.Println(string(data)) // {}  , 因此变量都是小写，json包无法访问里面的变量

	/*
	必须把字段改成大写，json包才可以访问
	那么如果我就想json出来都是小写如何解决呢？ tag中定义json字段。通过反射来取的

	type student_json struct{
		name string `json:"sutdent_name"`
		age int		`json:"age"`
		score int	`json:"score"`

	}

	 */


}

// 匿名字段的访问

func funnoname(){

	type Car struct{
		name string
		age int
	}

	type Strain struct{
		Car
		int
		start time.Time
	}

	var t Strain
	t.name = "train"
	t.age = 10
	t.int = 30

	fmt.Println(t) // {{train 10} 30 {0 0 <nil>}}



}

func lianbiao(){

	var head Student_single_list
	head.Age = 18
	head.Score = 80
	head.Name = "zhuzhu"

	var stu1 Student_single_list = Student_single_list{"zhagnsan",90,100,nil}
	var stu2 *Student_single_list = &Student_single_list{"wangwu",90,100,nil}


	head.Next = &stu1
	stu1.Next = stu2


	var p *Student_single_list = &head

	for p != nil{
		fmt.Println(p.Name)

		if p.Next == nil {

			break
		}
		p = p.Next
	}

}

func main(){

	var stu Student
	stu.Age = 18
	stu.Score = 80
	stu.Name = "zhuzhu"

	fmt.Println(stu)
	fmt.Printf("name:%p\n",&stu.Name)
	fmt.Printf("age:%p\n",&stu.Age)
	fmt.Printf("score:%p\n",&stu.Score)

	var stu1 Student = Student{"zhagnsan",90,100}

	var stu2 *Student = &Student{}
	stu2.Name = "wang"
	fmt.Println(stu1,stu2)

	lianbiao()


	var head *Student_single_list = new(Student_single_list)
	head.Age = 18
	head.Score = 80
	head.Name = "zhuzhu"


	insertbefore(head)


	fmt.Println("zhuzhu 是链表的第一个元素")
	trans(head)
	fmt.Println("========")
	fmt.Println("")

	fmt.Println("删除一个元素")
	delnode(head,"stu_6")




	trans(head)
	fmt.Println("========")
	fmt.Println("")

	fmt.Println("插入一个元素")
	insnode(head,"stu_5",&Student_single_list{"why",10,60,nil})
	trans(head)


	var headtail *Student_single_list = new(Student_single_list)

	headtail.Age = 18
	headtail.Score = 80
	headtail.Name = "zhuzhu"



	inserttail(&headtail)
	fmt.Printf("%p\n",headtail)


	fmt.Println("zhuzhu 是链表的最后一个元素")
	for headtail != nil {
		fmt.Println(headtail)

		if headtail.Next == nil{
			break
		}

		headtail = headtail.Next

	}

	funjson()
	funnoname()
	newinit()

	fmt.Println("二叉树的遍历")
	funtree()

}