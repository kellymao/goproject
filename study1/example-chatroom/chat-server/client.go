package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)
import "bytes"

type Client struct{


	conn net.Conn
	buf [8192]byte



}



func (p *Client) readPackage() (msg Message, err error){


	n,err:=p.conn.Read(p.buf[0:4])
	if n!=4 {

		err = errors.New("read header failed")
		return

	}

	buffer := bytes.NewBuffer(p.buf[0:4])

	var packlen uint32
	binary.Read(buffer,binary.BigEndian,&packlen)

	if err!=nil{

		fmt.Println()
		return


	}

	n,err = p.conn.Read(p.buf[0:packlen])

	if n!=int(packlen){

		err = errors.New("read body failed")
		return

	}


	err = json.Unmarshal(p.buf[0:packlen],&msg)

	if err!=nil{

		fmt.Println("unmarshal failed",err)
		return
	}
	return msg,nil


}


func (p *Client) Process() error {

	for {

		msg,err := p.readPackage()

		if err!=nil{
			return err
		}

		err = p.Processmsg(msg)
		if err!=nil{

			return err
		}
	}

}


func (p *Client) Processmsg(msg Message) (err error){


	switch (msg.Cmd){

	case Userlogin:
		err = p.login(msg)

	case Userregister:
		//err= p.register(msg)
	default:
		err = errors.New("fail message")
		return

	}

	return


}


func (p *Client) login(msg Message) (err error){

	//var cmd Logincmd

	//err = json.Unmarshal([])


}