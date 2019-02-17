package main
import "fmt"
func main() {
    arr := []int{1, 100, 2, 432, 234, 56, 90, 34, 24}
    //bubbleSort(arr)
    //selectSort(arr)
    //insertSort(arr)
    quickSort(arr, 0, len(arr))
    fmt.Println(arr)
    fmt.Println(int('a'))
}
// 冒泡排序
func bubbleSort(arr []int) {
    for i := 0; i < len(arr); i++ {
        for j := i + 1; j < len(arr); j++ {
            if arr[i] > arr[j] {
                arr[i], arr[j] = arr[j], arr[i]
            }
        }
    }
}
// 选择排序
func selectSort(arr []int) {
    for i := 0; i < len(arr) -1 ; i++ {
        minIndex := i
        for j := i+1; j < len(arr); j++ {
            if arr[minIndex] > arr[j] {
                minIndex = j
            }
        }
        arr[i], arr[minIndex] = arr[minIndex], arr[i]
    }
}
// 插入排序
func insertSort(arr []int) {
    for i := 1; i < len(arr); i++ {
        for j := 0; j < i; j++ {
            if arr[i] < arr[j] {
                arr[i], arr[j] = arr[j], arr[i]
            }
        }
    }
}
// 快速排序
func quickSort(arr []int, start, end int) {
    if start >= end {
        return
    }
    x := start
    index := arr[x]
    i, j := start, end-1
    fmt.Println(arr)
    for ; i < j; {
        for arr[j] >= index && i < j {
            j--
        }
        for arr[i] <= index && i < j {
            i++
        }
        arr[j], arr[i] = arr[i], arr[j]
    }
    arr[x], arr[i] = arr[i], arr[x]
    quickSort(arr, start, i)
    quickSort(arr, i+1, end)
}
