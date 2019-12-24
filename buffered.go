package main

import "fmt"
import "runtime"

func main(){
	runtime.GOMAXPROCS(2)	
	
	var data = make(chan int, 3)
	
	go func(){
		for {
			i := <- data 
			fmt.Println("terima data",i)
		}
	}()
	
	for i:=0;i<5;i++ {
		fmt.Println("kirim data",i)
		data <- i
	}
	
	
}

