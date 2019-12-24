package main

import "fmt"
import "runtime"

func main(){
	runtime.GOMAXPROCS(3)	
	
	var number = [] int{3,7,9,10,8,6,6,3,4,5,2}
	fmt.Println("Number :",number)
	
	var channe11 = make(chan float64)
	go getratarata(number, channe11)
	
	var channe12 = make(chan int)
	go nilaimaksimal(number, channe12)
	
	for i:=0;i<2;i++{
		select{
			case rata_rata := <- channe11:
				fmt.Printf("Rata rata \t : %.2f \n",rata_rata)
			case maksimal := <- channe12:
				fmt.Printf("Maksimal \t : %.d \n",maksimal)
		}
	}
}

func getratarata(number []int, ch chan float64){
	var sum = 0
	for _, e:=range number{
		sum += e
	}
	ch <- float64(sum) / float64(len(number))
}

func nilaimaksimal(number []int, ch chan int){
	var max = number[0]
	for _, e:=range number{
		if max < e{
			max = e
		}
	}
	ch <- max
}
