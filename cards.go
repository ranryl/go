package main
import (
    "fmt"
)
// 一副从1到n的牌，每次从牌堆顶取一张放桌子上，再取一张放牌堆底，直到手机没牌，最
// 后桌子上的牌是从1到n有序，设计程序，输入n，输出牌堆的顺序数组
func main() {
    fmt.Println(cards(1000))
}
func cards(n int) []int {
    list := make([]int, 100, 200)
    if n == 1 {
        return []int{1}
    }
    if n == 2 {
        return []int{1, 2}
    }
    list = []int{1, 2}
    for i := 3; i <= n; i++ {
        list = append(list[1:], list[0])
        list = append(list, i)
    }
    return list
}

