package main

import "fmt"

import "time"

func channel(){

	var chanint chan int = make(chan int)

	chanint <- 10

	a:= <-chanint
	fmt.Println(a)



}



func write_chan(ch chan int){


	for i:=0;i<=1000;i++{


		ch <- i
	}

	close(ch)
}

func read_chan(ch chan int){



	for {

		out:= <-ch
		fmt.Println(out)
	}

}

func calc(ch,result chan int){


	for v := range ch {

		flag:=true
		for i:=2;i<v;i++{

			if v%i ==0 {

				flag = false
				break
			}


		}

		if flag {

			result <- v

		}
	}



}
func main (){




	intchan :=make(chan int)

	resultchan:=make(chan int)

	go write_chan(intchan)

	for i:=0;i<8;i++{

		go calc(intchan,resultchan)

	}

	go func(){

		defer func(){
			fmt.Println("exit")
		}()
	for v:=range resultchan{

		fmt.Println(v)
	}
	}()

	time.Sleep(10 * time.Second)

}
