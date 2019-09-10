package main

import "os"
import "fmt"
import "study1/invoicedata/processdata"


func main(){

	inputfile:="/workspace/src/study1/datadir/data"

	file,err:=os.Open(inputfile)

	if err!=nil{

		fmt.Println(err)
		return
	}
	defer file.Close()




	var reader invoicedata.Unmarshainvoice  = invoicedata.NewProcessText()

	content,err:=reader.Invoice_read(file)

	if err!=nil{

		fmt.Println(err)
		return

	}

	for _, txt:= range content{


		fmt.Printf("%v\n",txt.ID)
		fmt.Printf("%v\n",txt.CUSTOMER)
		fmt.Printf("%v\n",txt.RAISED)
		fmt.Printf("%v\n",txt.DUE)
		fmt.Printf("%v\n",txt.PAID)

			for _,item := range txt.ITEM{

				fmt.Println(*item)


			}




	}




}
