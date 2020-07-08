//컴파일 하기 위해서는 main.go라고 이름을 지어야함, main  package는 오로지 컴파일하기 위해 사용하는 패키지
package main

import (
	"fmt"

	"github.com/yongbbbba/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
