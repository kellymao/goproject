package invoicedata

import "io"
import _ "os"
import "time"


type Invoice struct {

	ID int
	CUSTOMER int
	RAISED time.Time
	DUE  time.Time
	PAID bool
	ITEM []*Item


}


type Item struct {

	ID string
	PRICE float32
	QUANTITY int
	NOTE string



}
type Marshainvoice interface {


	Invoice_write(io.Writer, []*Invoice)


}


type Unmarshainvoice interface {


	Invoice_read(io.Reader)([]*Invoice ,error)
}
