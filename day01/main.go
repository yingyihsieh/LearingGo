package main

//全局變量
//var HEADERS = "user-agent"

func main() {
	//輸出
	//fmt.Println("12345")

	// 格式化
	//fmt.Printf("%s花了%d錢讓老婆的開心指數上升%.2f百分比,四捨五入=%f", "andy", 20000, 2.45, 2.5)

	//數據類型
	//fmt.Println(1 + 2)
	//fmt.Println(1 - 2)
	//fmt.Println(1 * 2)
	//fmt.Println(1 / 2)
	//fmt.Println(1 % 2)

	//變數
	//var title string = "andy"
	//var age int = 18
	//fmt.Printf("自我介紹: %s 年齡:%d", title, age)
	//var (
	//	title  = "title"
	//	age    = 20
	//	salary = 10000000
	//	gender string
	//)
	//gender = "男"
	//fmt.Println(title, age, salary, gender)

	//全局變量不能用:=
	//fmt.Println(HEADERS)

	//常量
	//const age = 18
	//const (
	//	name  = "123"
	//	title = "456"
	//)
	//fmt.Println(age, name, title)

	// iota
	//const (
	//	a = iota + 1
	//	b = 2
	//	c
	//)
	//fmt.Println(a, b, c)

	//輸入
	// scan 數量沒輸入完會持續堵塞
	//var name string
	//var age int
	//fmt.Println("請輸入name與age:")
	//count, err := fmt.Scan(&name, &age)
	//fmt.Println(count, err)

	//scanln 回車後不再堵塞
	//var name string
	//var age int
	//fmt.Println("請輸入name與age:")
	//count, err := fmt.Scanln(&name, &age)
	//fmt.Println(count, err)

	//scanf 根據佔位符提取相對應
	//var name string
	////var age int
	//fmt.Println("請輸入name與age:")
	//count, err := fmt.Scanf("name=%s", &name)
	//fmt.Println(count, err)
	//fmt.Println(name)

	//輸入的bug 遇到空格只擷取空格之前
	//var msg string
	//fmt.Println("請輸入信息:")
	//_, _ = fmt.Scanln(&msg)
	//fmt.Println(msg)
	//解決方法 stdin
	//reader := bufio.NewReader(os.Stdin)
	//// isPrefix=false(讀完)
	//line, isPrefix, _ := reader.ReadLine()
	//data := string(line)
	//fmt.Println(isPrefix, data)

	//條件
	//var a, b = rand.Intn(100), rand.Intn(100)
	//fmt.Println("a=", a, "b=", b)
	//if a > b {
	//	fmt.Println("a > b")
	//} else {
	//	fmt.Println("a <= b")
	//}
}
