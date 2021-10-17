package queue

import "fmt"

type Queue []int

func (root *Queue) Print() {
	fmt.Println(*root)

}
