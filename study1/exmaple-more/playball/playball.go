package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 打棒球的程序




var ball chan int =make(chan int)

var wg sync.WaitGroup

func play(name string){

	defer wg.Done()

	for {

		count := <-ball

		if count == 0 { // 通道关闭后，值为0

			fmt.Printf("%s win the game\n", name)
			return

		}

		if rand.Intn(100)%13 == 0 {
			fmt.Printf("%s loss the game\n", name)
			close(ball)
			return

		}

		time.Sleep(time.Second)
		fmt.Printf("%s playing the ball , num is %d .. \n", name, count)
		ball <- count+1

	}


}

func main(){




	fmt.Println("比赛开始了")


	wg.Add(2)
	go play("张三")
	go play("王五")
	ball <-1

	wg.Wait()


}

