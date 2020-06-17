package main

import (
	"fmt"

	"github.com/babamba/learngo/mydict"
)

func main() {
	//-----------
	// dictionary := mydict.Dictionary{"first": "First word"}
	// // dictionary["hello"] = "hello"
	// // dictionary.Add()
	// //fmt.Println(dictionary["first"])
	// //definition, err := dictionary.Search("second")
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }
	//-----------

	// dictionary := mydict.Dictionary{}
	// word := "hello"
	// definition := "Greeting"

	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// hello, _ := dictionary.Search(word)
	// fmt.Println("Found : ", word, " def : ", hello)

	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	//-----------

	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)

	// Challenge

	isDeleteErr := dictionary.Delete(baseWord)
	if isDeleteErr != nil {
		fmt.Println(isDeleteErr)
	}
	fmt.Println(dictionary)

	//-----------
}

// Account
//-----------
// func main() {
// 	// account := banking.Account{Owner: "nico"}
// 	// fmt.Println("test :", account)
// 	//-----------
// 	account := accounts.NewAccount("nico")
// 	account.Deposit(10)
// 	// fmt.Println(account.Balance())
// 	// err := account.Withdraw(20)
// 	// if err != nil {
// 	// 	//log.Fatalln(err) //프린트를 호출하고 프로그램을 종료시킨다.
// 	// 	fmt.Println(err)
// 	// }
// 	//fmt.Println(account.Balance(), account.Owner())
// 	//-----------
// 	fmt.Println(account)
// }
//-----------
// name := false  / func 밖에서는 축약형선언이 되지않는다.
// var name string = "nico" func 밖에서는 이런식으로 선언가능

// Struct 를 사용하기 위한 구조체 명세
// type person struct {
// 	name    string
// 	age     int
// 	favFood []string
// }

// func main() {
// __     __         _       _     _         _              _  ____            _              _
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

// &  - address
// *  - see through

// a := 2
// b := &a        // a의 메모리 주소를 바라본다.
// *b = 20        // a의 메모리를 바라보고있는 b의 value값을 바꾸므로 a의 메모리주소에 담긴 값이 변경 a=20
// fmt.Println(a) // * 은 메모리주소를 살펴보는 기능 (2, 2)

// ----------
//     _                                 _              _      ____  _ _
//    / \   _ __ _ __ __ _ _   _        / \   _ __   __| |    / ___|| (_) ___ ___
//   / _ \ | '__| '__/ _` | | | |_____ / _ \ | '_ \ / _` |____\___ \| | |/ __/ _ \
//  / ___ \| |  | | | (_| | |_| |_____/ ___ \| | | | (_| |_____|__) | | | (_|  __/
// /_/   \_\_|  |_|  \__,_|\__, |    /_/   \_\_| |_|\__,_|    |____/|_|_|\___\___|
//                         |___/

// names := [5]string{"nico", "lynn", "dal"} // 크기 5로 제한
// names[3] = "alala"
// names[4] = "alala"
// names[5] = "alala"  // 길이가 5이기 때문에 6번째 요소를 넣으려면 error

//names := []string{"nico", "lynn", "dal"} // slice 요소는 Array와 같지만 지정된 길이 가 없음
// names[3] = "lala" // error
//names = append(names, "flynn") // append()는 새로운 값이 추가된 slice를 return (es의 map, filter같이)
// Go lang 은 array.push 같은걸 지원하지 않는다.
//fmt.Println("test : ", names)
// ----------
// __  __
//|  \/  | __ _ _ __  ___
//| |\/| |/ _` | '_ \/ __|
//| |  | | (_| | |_) \__ \
//|_|  |_|\__,_| .__/|___/
//             |_|

//nico := map[string]string{"name": "nico", "age": "12"} // only string. 다른타입을 넣으려면 Struct 사용

//map도 반복가능
// for key, value := range nico {
// 	fmt.Println(key, value)
// }

//fmt.Println("test : ", nico)

// ----------
// ____  _                   _
/// ___|| |_ _ __ _   _  ___| |_ ___
//\___ \| __| '__| | | |/ __| __/ __|
// ___) | |_| |  | |_| | (__| |_\__ \
//|____/ \__|_|   \__,_|\___|\__|___/

// favFood := []string{"Kimchi", "ramen"}
// //nico := person{"nico", 18, favFood}
// nico := person{name: "nico", age: 18, favFood: favFood}
// fmt.Println(nico)
// fmt.Println(nico.name)
// fmt.Println(nico.age)
// fmt.Println(nico.favFood)

// ----------
// }

//   __                  _   _
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

//   __                  _   _                   _
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

//   __
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

//               _ _       _
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
