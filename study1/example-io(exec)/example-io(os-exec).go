package main

import "os/exec"
import "fmt"
import "io/ioutil"
import "os"


	func print1(){

		/*

		延迟输出，最后统一输出结果

		 */

		//创建一个cmd
		cmd3 := exec.Command("ping","-c","3", "www.baidu.com")
		//获取命令在start后标准输出管道
		out3, err := cmd3.StdoutPipe()

		if err != nil {
			fmt.Printf("stdoutpipe Error: %s\n", err)
			return
		}
		//执行命令
		if err = cmd3.Start(); err != nil { //表示命令正确执行了
			fmt.Printf("startError: %s\n", err)
		}
		//读取管道中所有数据
		data3, _ := ioutil.ReadAll(out3)
		//等待命令执行完成
		if err := cmd3.Wait(); err != nil {
			fmt.Printf("waitError: %s\n", err)
		}
		fmt.Println(string(data3))

	}



	func print2(){
		/*

		实时的输出


		 */

		//创建一个cmd
		cmd3 := exec.Command("ping","-c","3", "www.google.com")
		//获取命令在start后标准输出管道
		//out3, err := cmd3.StdoutPipe()

		cmd3.Stdout = os.Stdout
		cmd3.Stderr = os.Stderr


		//执行命令
		if err := cmd3.Start(); err != nil { //表示命令正确执行了
			fmt.Printf("startError: %s\n", err)
		}

		if err := cmd3.Wait(); err != nil {
			fmt.Printf("waitError: %s\n", err)
		}


	   //https://blog.csdn.net/zhuxinquan61/article/details/89716301

	}

func main(){



	print2()

}