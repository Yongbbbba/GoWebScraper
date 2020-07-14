//컴파일 하기 위해서는 main.go라고 이름을 지어야함, main  package는 오로지 컴파일하기 위해 사용하는 패키지
//main을 정의하는 것은 프로그램의 시작을 알리는 것

//package main

import (
	"fmt"

	"github.com/yongbbbba/nomadcoders/mydict"
)

//array & slice

// func main() {
// 	names := []string{"nico", "lynn", "dal"}
// 	names = append(names, "dustin")
// 	fmt.Println(names)
// }

// map, python의 딕셔너리와 비슷

// func main() {
// 	nico := map[string]string{"name": "nico", "age": "12"}
// 	for key, value := range nico {
// 		fmt.Println(key, value)
// 	}
// }

// func main() {
// 	account := accounts.NewAccount("nico")
// 	account.Deposit(10)
// 	err := account.Withdraw(20)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(account.Balance(), account.Owner())
// }

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(word)

}
