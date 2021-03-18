package main

func main() {
		// ------------------
		// Zero value concept
		// ------------------

		// Every single value we create must be initialized. If we don't specify it, it will be set to
	  // the zero value. The entire allocation of memory, we reset that bit to 0.	
		
		// - Boolean false
		// - Integer 0
		// - Floating Point 0
		// - Complex 0i
		// - String "" (empty string)
		// - Pointer nil

		///////////// Khai báo 1 ///////////////// 
		var a int
		var b string
		var c float64
		var d bool

		fmt.Printf("var a int \t %T [%v]\n", a, a)
		fmt.Printf("var b string \t %T [%v]\n", b, b)
		fmt.Printf("var c float64 \t %T [%v]\n", c, c)
		fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

		///////////// Khai báo 2 //////////////////
		var aa int = 1
		var bb string = "hello"
		var cc float64 = 3.14
		var dd bool = true

		fmt.Printf("var aa int \t %T [%v]\n", aa, aa)
		fmt.Printf("var bb string \t %T [%v]\n", bb, bb)
		fmt.Printf("var cc float64 \t %T [%v]\n", cc, cc)
		fmt.Printf("var dd bool \t %T [%v]\n\n", dd, dd)

		///////////// Khai báo 3 ////////////////// 
		aaa := 10
		bbb := "hello"
		ccc := 3.14159
		ddd := true
	
		fmt.Printf("aaa := 10 \t %T [%v]\n", aaa, aaa)
		fmt.Printf("bbb := \"hello\" \t %T [%v]\n", bbb, bbb)
		fmt.Printf("ccc := 3.14159 \t %T [%v]\n", ccc, ccc)
		fmt.Printf("ddd := true \t %T [%v]\n\n", ddd, ddd)
}
