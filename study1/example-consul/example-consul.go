package main


import "fmt"
import consulapi "github.com/hashicorp/consul/api"
import "time"
import "strings"


// 服务注册方式一
func register_service(client *consulapi.Client){

	registration := new(consulapi.AgentServiceRegistration)

	registration.ID = "webId"
	registration.Name = "web"
	registration.Port = 9999
	registration.Tags = []string{"rails"}
	registration.Address = "127.0.0.1"

	//  //增加check。
	//  check := new(consulapi.AgentServiceCheck)
	//  check.HTTP = fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/check")
	//  //设置超时 5s。
	//  check.Timeout = "5s"
	//  //设置间隔 5s。
	//  check.Interval = "5s"
	//  //注册check服务。
	//  registration.Check = check
	//  log.Println("get check.HTTP:", check)
	//
	//  err = client.Agent().ServiceRegister(registration)
	//
	//  if err != nil {
	//      log.Fatal("register server error : ", err)
	//  }

	//增加check。
	check := new(consulapi.AgentServiceCheck)

	check.Args = []string{"/tmp/tt.sh"}    //设置超时 5s。
	check.Timeout = "5s"
	//设置间隔 5s。
	check.Interval = "5s"
	//注册check服务。
	registration.Check = check

	err := client.Agent().ServiceRegister(registration)
	if err != nil {
		fmt.Println(time.Now())
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("Registered service %s in consul with tags %s \n", "web", strings.Join([]string{"rails"}, ","))





}

// 服务注册方式二
func register_service1(client *consulapi.Client){

	monitor_addr:="www.baidu.com"

	service := &consulapi.AgentServiceRegistration{
		ID:      "web1Id",
		Name:    "web1",
		Port:    9998,
		Address: "127.0.0.1",
		Tags:    []string{},
		Check: &consulapi.AgentServiceCheck{
			HTTP:     "http://" + monitor_addr ,
			Interval: "5s",
			Timeout:  "1s",
		},
	}

	if err := client.Agent().ServiceRegister(service); err != nil {
		panic(err)
	}
	fmt.Printf("Registered service %s in consul with tags %s \n", "web1", strings.Join([]string{}, ","))


}

func register_service2(client *consulapi.Client){



	service := &consulapi.AgentServiceRegistration{
		ID:      "web2Id",
		Name:    "web2",
		Port:    9998,
		Address: "127.0.0.1",
		Tags:    []string{"tcp"},
		Check: &consulapi.AgentServiceCheck{
			TCP:     "127.0.0.1:22" ,
			Interval: "5s",
			Timeout:  "1s",
		},
	}

	if err := client.Agent().ServiceRegister(service); err != nil {
		panic(err)
	}
	fmt.Printf("Registered service %s in consul with tags %s \n", "web2", strings.Join([]string{}, ","))


}

func register_service3(client *consulapi.Client){



	service := &consulapi.AgentServiceRegistration{
		ID:      "ttId",
		Name:    "tt",
		Port:    99998,
		Address: "127.0.0.1",
		Tags:    []string{"script"},
		Check: &consulapi.AgentServiceCheck{
			Args:     []string{"sh","/tmp/tt.sh"} ,

			Interval: "5s",
			Timeout:  "1s",
		},
	}

	if err := client.Agent().ServiceRegister(service); err != nil {
		panic(err)
	}
	fmt.Printf("Registered service %s in consul with tags %s \n", "tt", strings.Join([]string{"script"}, ","))


}


//kv

func kv(){

	config:=consulapi.DefaultConfig()

	client,_:=consulapi.NewClient(config)

	kv:=client.KV()

	pair :=&consulapi.KVPair{
		Key:"cuser",
		Value: []byte("10"),

	}

	if _,err:=kv.Put(pair,nil); err!=nil{
		panic(err)
	}


	if p,_,err:=kv.Get(pair.Key,nil); err!=nil{

		panic(err)
	}else{
		fmt.Println(p.Key,p.Value)
	}


	if ps,_,err:=kv.List(pair.Key,nil); err!=nil{

		panic(err)
	}else{
		for _,p := range ps{

			fmt.Println(p.Key,p.Value)
		}

	}

	if ks,_,err:=kv.Keys(pair.Key,"/",nil); err!=nil{

		panic(err)
	} else {
		for _,k := range ks{
			fmt.Println(k)
		}
	}

	if _,err:=kv.Delete(pair.Key,nil);err!=nil{
		panic(err)
	}else{
		fmt.Println("delete success")
	}




}

func store_key(client *consulapi.Client){


	kv:=&consulapi.KVPair{

		Key:"user/usermy1",
		Flags: 0,
		Value:[]byte("123"),


	}

	_,err:=client.KV().Put(kv,nil)

	if err!=nil{
		panic(err)
	}

	fmt.Println("store key success",kv.Key,kv.Value)
	fetch_key(client)



}

func fetch_key(client *consulapi.Client){


	kv, _, err := client.KV().Get("user/usermy123", nil) // 一个KVPair键值对
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%+v \n==%s \n==%v\n",kv,kv.Key,string(kv.Value))


	list, _, err := client.KV().List("user/", nil) // 包含很多KVPair键值对的列表
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%+v \n",list)

	ks, _, err := client.KV().Keys("user/","/", nil) // key 值的列表
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%v \n",ks) // [user/usermy1 user/usermy123]
}

func delete_service(client *consulapi.Client){

	if err := client.Agent().ServiceDeregister("ttId"); err != nil {
		panic(err)
	}

}

func delete_key(client *consulapi.Client){


}

// 定时任务
func  DoDiscover(client *consulapi.Client){

	t := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-t.C:
			query_services(client)
		}
	}


}

func check_health_service(client *consulapi.Client,service_name string){

	servicesData, _, err := client.Health().Service(service_name, "", true,
		&consulapi.QueryOptions{})

	if err!=nil{
		panic(err)
	}

	/*

	root@pts/0 #  curl -s  http://192.168.1.82:8500/v1/health/service/tt | python -m json.tool
	[
	    {
	        "Checks": [
	            {
	                "CheckID": "serfHealth",
	                "CreateIndex": 3882536,
	                "ModifyIndex": 3882536,
	                "Name": "Serf Health Status",
	                "Node": "node1",
	                "Notes": "",
	                "Output": "Agent alive and reachable",
	                "ServiceID": "",
	                "ServiceName": "",
	                "ServiceTags": [],
	                "Status": "passing"
	            }
	        ],
	        "Node": {
	            "Address": "192.168.1.82",
	            "CreateIndex": 3882536,
	            "Datacenter": "dc1",
	            "ID": "ddbb824b-6c33-5ba8-f31a-ed288d9b7bae",
	            "Meta": {},
	            "ModifyIndex": 3882538,
	            "Node": "node1",
	            "TaggedAddresses": {
	                "lan": "192.168.1.82",
	                "wan": "192.168.1.82"
	            }
	        },
	        "Service": {
	            "Address": "127.0.0.1",
	            "CreateIndex": 3889659,
	            "EnableTagOverride": false,
	            "ID": "ttId",
	            "ModifyIndex": 3889659,
	            "Port": 99998,
	            "Service": "tt",
	            "Tags": [
	                "script"
	            ]
	        }
	    }
	]
	 */
	fmt.Printf("%s \n",servicesData)

	for _,data := range servicesData{

			//fmt.Printf("%s \n",data)

			/*
			        "Service": {
			            "Address": "127.0.0.1",
			            "CreateIndex": 3889659,
			            "EnableTagOverride": false,
			            "ID": "ttId",
			            "ModifyIndex": 3889659,
			            "Port": 99998,
			            "Service": "tt",
			            "Tags": [
			                "script"
			            ]
			        }
			 */
			fmt.Println("输出服务的信息")
			fmt.Println(data.Service.Address,data.Service.Port)
			fmt.Println(data.Service.Service)


	}


}

func query_services(client *consulapi.Client){

	services, _, err := client.Catalog().Services(&consulapi.QueryOptions{})

	if err!=nil{
		panic(err)
	}

	fmt.Printf("%+v\n",services)
	fmt.Printf("%s\n",services) // map[userService:[primary v1] web:[rails] web1:[] web2:[tcp] consul:[] tt:[script]]

	for service,_ := range services{


		if service == "tt"{

			check_health_service(client,service)
		}




	}
}

func main(){


	config:=consulapi.DefaultConfig()

	//fmt.Printf("%+v",config)
	/*
	&{Address:127.0.0.1:8500
	Scheme:http
	Datacenter:
	Transport:0xc0000aad80
	HttpClient:<nil>
	HttpAuth:<nil>
	WaitTime:0s
	Token:
	TokenFile:
	TLSConfig:{Address: CAFile: CAPath: CertFile: KeyFile: InsecureSkipVerify:false}}
	 */

	config.Address = "192.168.1.82:8500"

	client,err:=consulapi.NewClient(config)

	if err!=nil{

		panic(err)

	}


	//register_service3(client)

	//register_service1(client)
	//register_service2(client)


	//store_key(client)

	query_services(client)


	go DoDiscover(client) //每10秒执行一次

	select {}  // 阻塞main函数 https://studygolang.com/articles/12992?fr=sidebar




}

