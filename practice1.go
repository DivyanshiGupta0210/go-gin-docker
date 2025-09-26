package main
import "fmt"
func main(){
	// fmt.Print ln("hehe")

	//print() and println()-formatted o/p : built-in
	//fmt.Print(): std; fmt.Println: std + adds newline; fmt.Printf() std + formatted o/p i.e, :
	//age := 30
    //fmt.Printf("I am %d years old.\n", age)

	// cost := 0.02
	// total := cost * 4
	// fmt.Println("starting Textio server", total) o/p=0.08
	
	//var user string = "abc"
	// var pass string = "123456"
	// fmt.Println("Auth basic", user+":"+pass) o/p=Auth basic abc:123456

	//dt: bool, string, numerical dt 
	//NOTE var=variable (declaration) here :)
	// var smsSendingLimit int //var varName dt

	// *num dt*: 1. int (signed+unsigned) 2. uint (unsigned) 3. float 4. complex 
	//int int8-upto 64 int16-upto 65000 int32 int64
	//uint uint8-BYTE uint16-RUNE (REP A UNICODE CODE PT) uint32 uint64 uintptr
	//float32 float64
	//float 32 is 32 bit; byte=8bit; nibble=4bit
	//complex64 complex 128

	//%v = Default format (value) i/p:42, "hello" o/p:42, hello
	//%s = string
	//%d = Decimal integer
	//%f = Floating point (decimal)
	//%.2f = Floating point with 2 decimals
	//%q = Double-quoted string with escapes eg. "hello"
	//%t = Boolean
	//%b = Binary format
	//%x = Hexadecimal (lowercase)
	//%p = Pointer (address)
	//** %T shows dt
	penny := 2
	fmt.Printf("The type of pennies is %T\n", penny)

	yo := int16(200) //***TYPECASTINGGGGG***
	//float to int typecasting: TRUNCATES the floating val/portion
	oi:=2.6
	oiInt := int(oi)
	fmt.Println("val of", oi, "as an int is", oiInt, "you see")

	fmt.Printf("The type of large pennies is %T\n", yo)
	coin := 0.2
	fmt.Printf("The type of coins is %T\n", coin)

	//null values:
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string	
	fmt.Printf(
		"\n%v %f %v %q\n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username,
	)

	//var empty string IS SAME AS empty := ""
	congrats := "\nhappy bday\n"
	fmt.Println(congrats)

	//declare multiple val in same line
	mileage, company := 80276, "Tesla"
	//i.e, mileage := 80276 andd company := "Tesla"
	fmt.Println(mileage, company)
	fmt.Println("") //NOTEE fmt.println gives ERROR
	//println("") is also CORRECT

	//constt var: immutable val
	//diff: in js: cant redefine constt var but can calc at run time
	//whereas in go: val stored in constt musttt be known at compile time
	const sec = 60
	fmt.Println(sec)

	//fmt.Printf : prints formatted string to std out
	//fmt.Sprintf() : returns the formatted string
	const name = "divi"
	const rate = 30.5
	msg := fmt.Sprintf(
		"\nHi %s, your open rate is %.1f percent",
		name,
		rate,
	)
	fmt.Println(msg)
}
