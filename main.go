package main

import "fmt"

func main() {
	var myPointer int
	fmt.Println("my valueis : ", myPointer)
	//0
	fmt.Println("my address is : ", &myPointer) //ambil alamat myPointer
	//0xc000100010

	myPointer = 50                          //assign value
	fmt.Println("my valueis : ", myPointer) //ambil value myPointer
	//50
	fmt.Println("my address is : ", &myPointer) //ambil alamat myPointer
	//0xc000100010

	yourPointer := &myPointer //copy alamat
	//yourPointer mereferensikan alamat dari myPointer
	fmt.Println("your address is : ", yourPointer)
	//0xc000100010
	fmt.Println("your value is : ", *yourPointer)
	// 50

	*yourPointer = 70
	fmt.Println("your address is : ", yourPointer)
	//0xc000100010
	fmt.Println("your value is : ", *yourPointer)
	//70
	fmt.Println("my valueis : ", myPointer)
	//70
	fmt.Println("my address is : ", &myPointer) //0xc000100010
	// result = 2 values in 1 memory
}
