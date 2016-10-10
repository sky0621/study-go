package main

import "github.com/sky0621/study-go/temp/struct/family"

func main() {
	ss := &family.Dad{Name: "Taro", Age: 38}
	ss.Like()

	ss2 := &family.Dad{Name: "Taro2", Age: 48}
	ss2.Like()

	ss.Like()

	sk := &family.Mom{Name: "Hanako", Age: 36}
	sk.Like()

	sy := &family.Child{Name: "Kazuo", Age: 4}
	sy.Like()
}
