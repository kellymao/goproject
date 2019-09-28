package main

import "sync"
import "fmt"
import "study1/example-chat/common"


var clientmgr *Clientmgr = &Clientmgr{}

type Clientmgr struct{

	lock sync.RWMutex

	clients  map[string]*Handle


}


func (p *Clientmgr) sendall(msg common.Message){


	//p.lock.RLock()
	for i ,h  := range p.clients{

		fmt.Println(i,"****",msg)
		//fmt.Printf("[%s] %p \n" , i, p.clients[i])
		//fmt.Printf("%p %v %v \n",p.clients[i].conn,p.clients[i].conn,p.clients[i].userobj)



		//p.clients[i].conn.Write([]byte("testtt\n") )
		//p.clients[i].send_msg(msg)

		h.send_msg(msg)

	}
	//p.lock.RUnlock()


}



func (p *Clientmgr) sendone(id string, msg common.Message){





}

func (p *Clientmgr) add(id string,handle *Handle){

	if p.clients == nil {

		p.clients = make(map[string]*Handle)
	}
	_,ok:=p.clients[id]
	if ! ok {


		p.clients[id] = handle

	}

	p.clients[id] = handle

	for i := range p.clients{

		fmt.Println(i,"==========",p.clients[i])
		fmt.Printf("加入后，在map里循环的handle[%s]： %v %p \n",i,p.clients[i],p.clients[i])
		//p.clients[i].conn.Write([]byte("test\n") )




	}

}



func (p *Clientmgr) del(id string){

	p.lock.Lock()
	delete(p.clients,id)
	p.lock.Unlock()

}