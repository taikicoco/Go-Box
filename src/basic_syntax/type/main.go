package main

import "fmt"

func main() {
	fmt.Println("基本型")
	fmt.Println("===ゼロ値===")
	{
		var x int
		var y float32
		var z string
		fmt.Println("x:",x)//0
		fmt.Println("y:",y)//0
		fmt.Println("z:",z)//空白
	}

	fmt.Println("===リテラル===")
	fmt.Println("整数リテラル")
	fmt.Println(1_234_567)
	fmt.Println("小数点リテラル")
	fmt.Println(6.03e23)
	fmt.Println(12.3_45)

	fmt.Println("===論理型===")
	{
		var flag bool // 値が代入されないとfalseになる
		var isAwesome = true
		fmt.Println(flag)  // false
		fmt.Println(isAwesome)  // true
		flag = true
		fmt.Println(flag)	//true

	}

	fmt.Println("===複素数===")
	{
		x := complex(2.5, 3.1)
		y := complex(10.2, 2)
		fmt.Println(x + y) // (12.7+5.1i)
		fmt.Println(real(x)) // 2.5
		fmt.Println(imag(x)) // 3.1	
	}

	fmt.Println("===文字列型===")
	fmt.Println("hello" + "string")

	fmt.Println("===変数の宣言===")
	{
		var x, y int = 10, 20
		fmt.Println(x, y)  // 10 20
	}
	{
		var (
			x    int
			y        = 20
			z    int = 30
			d, e     = 40, "hello"
			f, g string
		)
		fmt.Println(x, y, z, d, e) // 0 20 30 40 hello

		fmt.Println("f=|", f, "| g=|", g, "|" ) // f=|  | g=|  |
	}

	fmt.Println("===定数===")
	{
		const x int64 = 10

		const (
			idKey   = "id"
			nameKey = "name"
		)

		const z = 20 * 10
		const y = "hello"

		fmt.Println(x)
		fmt.Println(y)

		//error
		//x = x + 1
		//y = "bye"

		fmt.Println(x)
		fmt.Println(y)
	}

	fmt.Println("===未使用変数===")
	{
		//error
		//x := 10
		y := 20
		fmt.Println(y)
	}
	{
		x := 30
		fmt.Println(x)
		x = 40
	}
}

