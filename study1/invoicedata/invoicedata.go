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


	outputfile :="/workspace/src/study1/datadir/out.dat"

	file_w,err:=os.OpenFile(outputfile,os.O_WRONLY|os.O_CREATE,0644)

	if err!=nil {
		fmt.Println(err)
		return
	}

	defer file_w.Close()

	var writer invoicedata.Marshainvoice =  invoicedata.NewProcessText()

	writer.Invoice_write(file_w,content)





}
