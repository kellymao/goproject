package main


import "bufio"
import "os"
import "fmt"
import "io"

// 从终端读取

func r_stdin(){

	var reader *bufio.Reader
	reader = bufio.NewReader(os.Stdin)

	str,err:=reader.ReadString('\n')

	if err!= nil{

		fmt.Println(err)
		return
	}

	fmt.Printf("read str success: %s", str)


}

// 从文件读取

func r_file(){

	afile,err:= os.Open("/tmp/aaa.log")


	if err!= nil {

		fmt.Println(err)
		return
	}

	defer afile.Close()

	reader := bufio.NewReader(afile)
	str,err:=reader.ReadString('\n')

	if err!= nil{

		fmt.Println(err)
		return
	}

	fmt.Printf("read str success: %s", str)


}


type CharCount struct{

	Strcount int
	Numcount int
	Space int
	Other int


}

func charcount_file(){

	afile,err:= os.Open("/tmp/aaa.log")


	if err!= nil {

		fmt.Println(err)
		return
	}

	defer afile.Close()

	reader := bufio.NewReader(afile)


	var count CharCount

	for {

		str,err:=reader.ReadString('\n')
		if err == io.EOF {

			break
		}

		if err != nil {
			fmt.Println(err)
			break
		}


		runestr:= []rune(str)

		for _,v := range runestr{

			switch {

			case v>='a' && v <='z':
				fallthrough
			case v>='A' && v<='Z':
				count.Strcount++
			case v == ' ' || v == '\t':
				count.Space++
			case v>='0' && v <='9':
				count.Numcount++



			}




		}


	}


	fmt.Printf("char count is %d",count.Strcount)

	fmt.Printf("space count is %d",count.Space)
	fmt.Printf("number count is %d",count.Numcount)



}


// 字符串读取

type Student struct{

	Name string
	Age int
	Score float32

}

func read_str(){


	var str="stu01 18 98.6"
	var stu *Student = new(Student)

	fmt.Sscanf(str,"%s %d %f", &stu.Name,&stu.Age,&stu.Score)
	fmt.Println(*stu)

	fmt.Sscanln(str,"%s %d %f", &stu.Name,&stu.Age,&stu.Score)
	fmt.Println(*stu)

	var firstname, lastname string

	fmt.Scan(&firstname,&lastname)
	fmt.Println("firstname:",firstname)
	fmt.Println("lastname:",lastname)


	var firstname1, lastname1 string
	var age int

	fmt.Scanln(&firstname1,&lastname1)
	fmt.Println("firstname1:",firstname1)
	fmt.Println("lastname1:",lastname1)

	fmt.Scanln(&age)

	fmt.Println(age)

	/*

	./example-rw
	{stu01 18 98.6}
	{stu01 18 98.6}
	abc
	ef
	firstname: abc
	lastname: ef
	abc df
	firstname1: abc
	lastname1: df
	10
	10


	 */

	fmt.Scanf("the student name is %s , %s ,age is %d",&firstname1,&lastname1,&age)


	fmt.Println("firstname1:",firstname1)
	fmt.Println("lastname1:",lastname1)



	fmt.Println(age)








}


func read_str1(){


	var firstname1, lastname1 string
	var age int


	fmt.Scanf("the student name is %s , %s ,age is %d",&firstname1,&lastname1,&age)


	fmt.Println("firstname1:",firstname1)
	fmt.Println("lastname1:",lastname1)



	fmt.Println(age)

	/*

	root@pts/1 # ./example-rw
	the student name is 1hj , fds ,age is 19
	firstname1: 1hj
	lastname1: fds
	19

	 */

}
func main(){

	//charcount_file()

	read_str()

}
