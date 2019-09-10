package invoicedata

import "io"
import "fmt"
import "bufio"
import "strings"
import "time"


type ProcessText struct{}

func NewProcessText() *ProcessText{


	return &ProcessText{}
}

func (p *ProcessText) Invoice_read(file io.Reader)([]*Invoice ,error){

	var invoices []*Invoice = make([]*Invoice,0)

	r := bufio.NewReader(file)

	for {

		strline,err:=r.ReadString('\n')
		if err == io.EOF{

			break
		}
		if err !=nil {

			return nil, fmt.Errorf("has a error ,%s \n",err.Error())
		}


		if strings.HasPrefix(strline,"INVOICE") {

			invoice := process_invoice(strline)
			invoices = append(invoices,invoice)
		}else if  strings.HasPrefix(strline,"ITEM"){
			process_item(strline,invoices[len(invoices)-1])   // 切片取最后的元素不能用 -1
		}


	}

	return invoices ,nil

}

func (p *ProcessText) Invoice_write(w io.Writer, invoices []*Invoice){






}



func process_invoice(line string) *Invoice {

	var invoice *Invoice = &Invoice{}
	var  RAISED,DUE string

	if _,err:= fmt.Sscanf(line,"INVOICE ID=%d CUSTOMER=%d RAISED=%s DUE=%s PAID=%t",&invoice.ID,&invoice.CUSTOMER,&RAISED,&DUE,&invoice.PAID); err!=nil {

		fmt.Println("process invoice err %s",err)

		return nil
	}

	var err error

	invoice.RAISED,err = time.Parse("2006-01-02",RAISED)   //字符串转换成日期
	if err !=nil {

		fmt.Println("time parse error %s",err)

	}
	invoice.DUE,_ = time.Parse("2006-01-02",DUE)

	

	invoice.ITEM = make([]*Item,0) // 必须默认初始化长度，否则会报错。长度太多了会以nil填充

	//fmt.Println(*invoice)
	return invoice

}


func process_item(line string,invoice *Invoice) {




	var item *Item = &Item{}

	if _,err:= fmt.Sscanf(line,"ITEM ID=%s PRICE=%f QUANTITY=%d",&item.ID,&item.PRICE,&item.QUANTITY);err!=nil{

		fmt.Println("process item err %s",err)

		return
	}

	if idx:=strings.Index(line,":");idx > -1 {
		note:= line[idx+1:]

		item.NOTE = note
	}

	//fmt.Println(*item)
	invoice.ITEM = append(invoice.ITEM,item)

	//fmt.Println(*invoice)



}