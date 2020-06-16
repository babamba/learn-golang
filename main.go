package main

import "fmt"

// name := false  / func 밖에서는 축약형선언이 되지않는다.
// var name string = "nico" func 밖에서는 이런식으로 선언가능

func main() {
	//  __     __         _       _     _         _              _  ____            _              _
	// \ \   / /_ _ _ __(_) __ _| |__ | | ___   / \   _ __   __| |/ ___|___  _ __ | |_ __ _ _ __ | |_
	//  \ \ / / _` | '__| |/ _` | '_ \| |/ _ \ / _ \ | '_ \ / _` | |   / _ \| '_ \| __/ _` | '_ \| __|
	//   \ V / (_| | |  | | (_| | |_) | |  __// ___ \| | | | (_| | |__| (_) | | | | || (_| | | | | |_
	//    \_/ \__,_|_|  |_|\__,_|_.__/|_|\___/_/   \_\_| |_|\__,_|\____\___/|_| |_|\__\__,_|_| |_|\__|
	//name := false // 변수 선언이되 처음선언된 변수의 값으로 타입이 지정되는 축약버전 변수선언(함수안에서만 가능)
	//name = "lynn" // 처음선언된 bool 타입이 아니라 string으로 선언되었기때문에 에러
	//	---------
	// var name string = "nico" // 변수선언
	// name = "lynn"
	//	---------
	//const name string = "nico" 상수선언
	//	---------
	//fmt.Println(name)
	//	---------
	//fmt.Println(multiply(2, 2))
	//totalLength, upperName := lenAndUpper("nico") // go는 무엇인가 만들고 안쓰면 에러냄
	//fmt.Println(totalLength, upperName)
	// totalLength, _ := lenAndUpper("nico") // 원하는것만 리턴받고싶다면 언더바
	// fmt.Println(totalLength)
	//	---------
	// repeatMe("nico", "nico")
	//	---------
	// totalLength, up := lenAndUpper("nico")
	// fmt.Println(totalLength, up)
	// superAdd(1, 2, 3, 4, 5, 6)
	//	---------

	// result := superAdd(1, 2, 3, 4, 5, 6)
	// fmt.Println(result)
	//	---------
	//fmt.Println(canIDrink(16))
	//	--------- Pointer
	//  ____       _       _
	// |  _ \ ___ (_)_ __ | |_ ___ _ __ ___
	// | |_) / _ \| | '_ \| __/ _ \ '__/ __|
	// |  __/ (_) | | | | | ||  __/ |  \__ \
	// |_|   \___/|_|_| |_|\__\___|_|  |___/

	// a := 2
	// b := a
	// a = 10 //10 10 이 나와야하는데? 10 2가 나옴
	// fmt.Println(a, b)
	//	---------

	// a := 2
	// b := 5
	// fmt.Println(&a, &b) // &을 붙이면 메모리 주소를 볼 수 있음.
	//0xc000014082 0xc000014090

	a := 2
	b := &a            // a의 메모리 주소를 바라본다.
	fmt.Println(a, *b) // * 은 메모리주소를 살펴보는 기능 (2, 2)
}

//    __                  _   _
//  / _|_   _ _ __   ___| |_(_) ___  _ __         ___  _ __   ___
// | |_| | | | '_ \ / __| __| |/ _ \| '_ \ _____ / _ \| '_ \ / _ \
// |  _| |_| | | | | (__| |_| | (_) | | | |_____| (_) | | | |  __/
// |_|  \__,_|_| |_|\___|\__|_|\___/|_| |_|      \___/|_| |_|\___|

// func multiply(a, b int) int { // 파라미터 타입 종류 선언 / 리턴타입 선언
// 	//func multiply(a int, b int) int { // 파라미터 타입 종류 선언 / 리턴타입 선언
// 	return a * b
// }

// func repeatMe(words ...string) { // 단순 반복 . 여러개의 arguments(Array)를 받을땐 ... 처리
// 	fmt.Println(words)
// }

//    __                  _   _                   _
//  / _|_   _ _ __   ___| |_(_) ___  _ __       | |___      _____
// | |_| | | | '_ \ / __| __| |/ _ \| '_ \ _____| __\ \ /\ / / _ \
// |  _| |_| | | | | (__| |_| | (_) | | | |_____| |_ \ V  V / (_) |
// |_|  \__,_|_| |_|\___|\__|_|\___/|_| |_|      \__| \_/\_/ \___/

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name) // 여러가지로 리턴가능
// }

// Naked Return 방식
// return 할 variable을 굳이 명시 하지 않아도 되는 방법
// func lenAndUpper(name string) (length int, uppercase string) { // 리턴에 어떤 varialbe을 리턴할지 정의 해두었기때문
// 	defer fmt.Println("I'm done")
// 	length = len(name)
// 	uppercase = strings.ToUpper(name)
// 	//return length, uppercase
// 	return
// }

//  __
//  / _| ___  _ __
// | |_ / _ \| '__|
// |  _| (_) | |
// |_|  \___/|_|

// func superAdd(numbers ...int) int {
// 		fmt.Println(numbers)
//	--------- or range ignore param
// 	total := 0
// 	for _, number := range numbers { // value 를 _ 로 ignore 처리
// 		total += number
// 	}
// 	return total

//	--------- for
// 	// for i := 0; i < len(numbers); i++ {
// 	// 	fmt.Println(numbers[i])
// 	// }
// 	// return 1
//	--------- for range
// 	// for index, number := range numbers { //첫번째는 index 두번째가 value
// 	// 	fmt.Println(index, number)
// 	// }
// 	//return 1
//	---------

// }
//  _  __            _
// (_)/ _|       ___| |___  ___
// | | |_ _____ / _ \ / __|/ _ \
// | |  _|_____|  __/ \__ \  __/
// |_|_|        \___|_|___/\___|
//	--------- if-else
// func canIDrink(age int) bool {
// 	if koreaAge := age + 2; koreaAge < 18 { // 세미콜론까지 적은부분이 variable expression
// 		// if age < 18 {
// 		return false // 리턴과 함께 끝나기때문에 else 를 적을 필요가 없다고 lint에서 추천해줌.
// 	}
// 	return true
// }
//	---------

//                _ _       _
//  _____      _(_) |_ ___| |__
// / __\ \ /\ / / | __/ __| '_ \
//\__ \\ V  V /| | || (__| | | |
//|___/ \_/\_/ |_|\__\___|_| |_|

// func canIDrink(age int) bool {
// 	switch koreaAge := age + 2; koreaAge {
// 	case 10:
// 		return false
// 	case 18:
// 		return true
// 	}
// 	return false
// }
//	---------
