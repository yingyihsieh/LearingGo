package main

func main() {
	//switch
	//var age = 20
	//switch age {
	//case 18:
	//	fmt.Println("成年囉")
	//case 20:
	//	fmt.Println("可以投票了")
	//default:
	//	fmt.Println("其他年紀")
	//}

	//循環
	//for {
	//	fmt.Println("死循環")
	//	time.Sleep(time.Second * 1)
	//}

	//var count = 0
	//for count < 5 {
	//	fmt.Println("while")
	//	count += 1
	//	time.Sleep(time.Millisecond * 100)
	//}

	//for i := 0; i < 10; i++ {
	//
	//}

	//打標籤 跳出多層回圈
	//f1:
	//	for i := 0; i < 3; i++ {
	//		for j := 0; j < 2; j++ {
	//			if j == 1 {
	//				continue f1
	//			}
	//			fmt.Println(i, j)
	//		}
	//	}

	//goto
	//	var level string
	//	fmt.Println("請輸入您的等級")
	//	fmt.Scanln(&level)
	//	if level == "vip10" {
	//		goto vip10
	//	} else if level == "vip1" {
	//		goto vip1
	//	} else {
	//	}
	//	fmt.Println("預約")
	//vip1:
	//	fmt.Println("排隊")
	//vip10:
	//	fmt.Println("進入")

	//字串格式化
	//var (
	//	name   string = "andy"
	//	age    int    = 20
	//	status bool   = true
	//)
	//
	//str := fmt.Sprintf("名字:%s\n年齡:%d\nvip:%v", name, age, status)
	//fmt.Println(str)

	//homework
	//var guessAge int
	//var realAge int = 76
	//for i := 3; i > 0; i-- {
	//	fmt.Println("請輸入猜測年齡")
	//	_, _ = fmt.Scanln(&guessAge)
	//	if guessAge != realAge {
	//		if i-1 > 0 {
	//			fmt.Printf("猜錯了 你還有%d次機會", i-1)
	//		} else {
	//			fmt.Printf("遊戲結束 你輸了")
	//		}
	//	} else {
	//		fmt.Printf("恭喜猜對了 遊戲結束")
	//		break
	//	}
	//}

	// 字符串長度 字節
	//var name = "網路"
	//fmt.Println(len(name))

}
