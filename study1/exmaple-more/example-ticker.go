package main



import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int, 10)

	go func() {
		for {
			time.Sleep(time.Second)
			ch <- 100
		}
	}()

	idleDuration := 3 * time.Second
	idleDelay := time.NewTicker(idleDuration)
	defer idleDelay.Stop()

	for {


		select {
		case x := <- ch:
			fmt.Println(x)
		case t1:= <-idleDelay.C:
			fmt.Println(t1)

			//return
		}
	}
}


/*
100
100
2019-11-19 17:57:20.624775358 +0800 CST m=+3.000467237
100
100
100
2019-11-19 17:57:23.625567424 +0800 CST m=+6.001259193
100
100
100
2019-11-19 17:57:26.625772103 +0800 CST m=+9.001463870
100
100
100
2019-11-19 17:57:29.626079469 +0800 CST m=+12.001771263
100
100
 */