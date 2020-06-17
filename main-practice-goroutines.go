package main

import (
	"fmt"
	"time"
)

func main() { // main 함수는 go routines를 기다려주지않는다.

	// 채널 === 파이프라인이라고 생각하면될까?
	c := make(chan string) // 채널을 통해 어떤 타입을 보낼지 정의 / 채널을 하나만든다.

	people := []string{"nico", "flynn", "dal", "larry", "babamba"}
	for _, person := range people {
		go isSexy(person, c) // 채널도 인수로 보내준다
		// 함수가 리턴해주는게 아니라 채널을 통해 메세지를 받는다라고 생각하면된다.
	}

	//resultOne := <-c // 채널의 메세지로 결과를 받는다.
	//resultTwo := <-c // 채널의 메세지로 결과를 받는다.

	//fmt.Println(resultOne)
	//fmt.Println(resultTwo)
	// 2개의 고루틴이 돌아가고있다면, 혹은 끝났다면 그이상 메시지를 받기위해서 채널은 기다리지않음

	// fmt.Println("Waitng this message")
	// fmt.Println("Received this message : ", <-c)
	// fmt.Println("Received this message : ", <-c)

	// 하지만 일일히 갯수만큼 채널에 메시지를 작성할순없고.
	// 이렇게 메시지 리시버를 작성
	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for : ", i)
		fmt.Println(<-c)
		// 메시지를 받는다는건 blocking operation
		// Rule = 먼저해야되는걸 먼저 해치운다. 채널과 고루틴 제일먼저 하지만 함수가 끝나면 고루틴이건 뭐건 없다.
		// 받을 데이터에 대해서 어떤 타입을 받을지 구체적으로 명시해줘야한다.
		// 메세지를 채널로 보내는 방법은 Arrow <- 를 이용하여 채널을 향하게 작성.
	}
	// 병렬 프로그래밍이 되서 순서는 종료되는데로 처리 할때마다 다름.

	// go sexyCount("nico") // go 를 추가해주는것만으로 병렬 프로그래밍이 처리된다.
	// sexyCount("flynn")
	//go sexyCount("flynn") // 두번째까지 go routine 이 들어가면 메인함수가 종료되버린다.

	// Go routines 는 프로그램이 작동하는 동안만 유효하다.
}

// 일반적인 탑다운 형식의 프로그래밍
func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func isSexy(person string, c chan string) { // 채널 타입에 채널을 통해 어떤 타입의 데이터를 주고받을지 Go에 알려줘야함
	time.Sleep(time.Second * 5)
	//fmt.Println(person)
	c <- person + "  is sexy" //채널을 통해 리턴값을 리턴해준다. (<- send)
}
