package main

import "fmt"
import "runtime"

var data = make(chan string)

func main(){
	runtime.GOMAXPROCS(3)	
	
	go tampilkan_pesan("a")
	go tampilkan_pesan("b")
	go tampilkan_pesan("c")
	
	var message1 = <-data
	fmt.Println(message1);
	
	var message2 = <-data
	fmt.Println(message2);
	
	var message3 = <-data
	fmt.Println(message3);
}

func tampilkan_pesan(pesan string){
	data <- pesan
}
