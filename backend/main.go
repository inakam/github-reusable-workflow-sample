// バックエンドアプリケーションのメインファイル
package main

import (
	"fmt"
)

// ユーザーへの挨拶を生成する
func GreetUser(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to our backend API.", name)
}

// 2つの整数の積を計算する
func Multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(GreetUser("World"))
	fmt.Printf("2 * 3 = %d\n", Multiply(2, 3))
}
