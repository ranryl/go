package main
import (
    "fmt"
    "sync"
)
type Manager struct{}
var m *Manager
var lock *sync.Mutex = &sync.Mutex{}
// 饿汉模式,在类加载时就创建实例，线程安全
/*var m *Manager = &Manager{}
func GetInstance() *Manager {
    return m
}*/

// 懒汉模式, 在需要使用时才创建实例，线程不安全
/*func GetInstance() *Manager {
    lock.Lock()
    defer lock.Unlock()
    if m == nil {
        m = &Manager{}
    }
    return m
}*/
// 双重锁，懒汉模式的改进，线程安全
func GetInstance() *Manager {
    if m == nil {
        lock.Lock()
        defer lock.Unlock()
        if m == nil {
            m = &Manager{}
        }
    }
    return m
}
var once sync.Once
// 使用once
func GetInstance() *Manager {
    once.Do(func(){
        m = &Manager{}
    })
    return m
}
func main() {
    GetInstance()
    fmt.Println(m)
}

