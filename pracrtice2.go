package main
import "fmt"
import "errors"
func main(){
	//loops
	const height = 4
	if height > 4{
		fmt.Println("good")
	} else if height == 4{
		fmt.Println("okay")
	} else {
		fmt.Println("smol")
	}
	//ALT**
	// if height := 4; height > 2 { //height = 4 GIVES ERROR
	// 	fmt.Println("good")
	// }

	result := sub(1,2)
	fmt.Println(result) //GIVES ADDRESS FOR (sub)

	x:=5
	// inc(x) //prints 5 w/o pointer
	inc(&x) 
	fmt.Println(x) 

	sendsSoFar := 430
	const sendsToAdd = 25
	incSends(sendsSoFar, sendsToAdd)
	sendsSoFar = incSends(sendsSoFar, sendsToAdd)
	fmt.Println("You've sent", sendsSoFar, "messages")

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
//main closed

//functiosn (Other than main)
//:= TO DECLARE AND INIT NEW VAR, NOT FOR ONES ALREADY DECLARED
func sub(x int, y int) int{
	x=5; y=3 //= as x and y are already declared
	return x-y
	//FN CALL INSIDE MAIN FN ONLIIIIIII
}

func inc(x *int){
	(*x)++
}

func incSends(sendsSoFar, sendsToAdd int) int{
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar //return value
}

//EARLY RETURN depends upon condition
//**guard clauses** are a coding pattern used to exit early from a function when a certain condition is met
//import "errors"
func divide(dividend, divisor int) (int, error){
	if divisor ==0{
		return 0, errors.New("Can't divide by zero")
	}
	return dividend/divisor, nil
}
