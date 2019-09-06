package main


import "bufio"
import "os"
import "fmt"

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

	file,err:= os.Open("/tmp/aaa.log")


	if err!= nil {

		fmt.Println(err)
		return
	}

	defer file.close()

	reader := bufio.NewReader(file)
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

	file,err:= os.Open("/tmp/aaa.log")


	if err!= nil {

		fmt.Println(err)
		return
	}

	defer file.close()

	reader := bufio.NewReader(file)


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
				count.Spacecount++
			case v>='0' && v <='9':
				count.Numcount++



			}




		}


	}


	fmt.Printf("char count is %d",count.Strcount)

	fmt.Printf("space count is %d",count.Spacecount)
	fmt.Printf("number count is %d",count.Numcount)



}


func main(){

	charcount_file()

}
