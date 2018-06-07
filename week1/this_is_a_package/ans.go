package ans

import (
	"fmt"
	"strconv"
)

func Echo()  {
	var a int = 1
	var b int32 = 2
	var c int64 = 3
	var d string = "999"
	var e float32 = 88.8
	var f float64 = 99.9
	var x string = "I Love Golang_"

	fmt.Println("Week 1")
	fmt.Printf("1. a(int)+b(int32) = %+v \n", a+int(b))
	fmt.Printf("2. a(int)+b(int32)+c(int64) = %+v \n", a+int(b)+int(c))
	fmt.Printf("3. f(float64)/e(float32) = %+v \n", f/float64(e))
	fmt.Printf("4. a(int)+d(string) = %+v \n", a+int(strToInt(d)))
	fmt.Printf("5. x(string)+a(int) = %+v \n\n", x+strconv.Itoa(a))
}

func strToInt(num string) int64 {
	numInt, err := strconv.ParseInt(num, 10, 64); if err != nil {
		fmt.Println(err.Error())
	}

	return  numInt
}