// 고루틴과 채널 학습
// 동작을 순서대로가 아니라 동시에 수행할 수 있게 해줌.
// 채널은 기본적으로 blocking operation의 개념, 뭔가를 받기 전에는 동작을 멈추는 것

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "hoony", "jam", "hong"}
	for _, person := range people {
		go isSexy(person, c) //go쓰레드를 생성하고,
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("waiting for ", i) //이 동작을 수행한 후 채널로부터 값을 받을 때까지 operation이 blocked됨
		fmt.Println(<-c)               // 채널로부터 값을 받을 때마다 출력해주기
	}

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}
