package main

import (
	"fmt"
)

// 1.  设计一个立方体类BOX，它能计算并输出立方体的体积和表面积
// -------------- 1 --------------------------
//type Box struct {
//	length float32
//}

//func (b Box) volume() float32 {
//	return b.length * b.length * b.length
//}
//func (b Box) area() float32 {
//	return b.length * b.length * 6
//}
//func main() {
//	b := Box{100}
//	fmt.Printf("%f\n", b.volume())
//	fmt.Printf("%f\n", b.area())
//}
// ---------------------------------------------

// 2. 有5个学生，每个学生的数据包括学号、姓名、三门课成绩，从键盘输入5个学生的数据，要求计算并输出。
// --------------------------2------------------
//type Student struct {
//	id      int
//	name    string
//	chinese float32
//	math    float32
//	english float32
//	total   float32
//}

//func main() {
//	stu := make([]Student, 10)
//	var n int
//	fmt.Println("input the number of student")
//	fmt.Scanf("%d", &n)
//	fmt.Println("input student's id name chinese math english")
//	for i := 0; i < n; i++ {
//		fmt.Scanf("%d%s%f%f%f", &stu[i].id, &stu[i].name, &stu[i].chinese, &stu[i].math, &stu[i].english)
//	}
//	fmt.Printf("学号\t姓名\t语文\t数学\t英语\t总分\n")
//	for i := 0; i < n; i++ {
//		total := stu[i].chinese + stu[i].math + stu[i].english
//		fmt.Printf("%d\t%s\t%.2f\t%.2f\t%.2f\t%.2f\n", stu[i].id, stu[i].name, stu[i].chinese, stu[i].math, stu[i].english, total)

//	}

//}

// ---------------------------------------------

// 3. 假定居民的基本数据包括身份证号、姓名、性别和出生日期，而居民中的成年人又多项数据：最高学历和职业，成年人中的党员又多一项数据：党员类别。现要求建立三个类，让成年人类继承居民类，而党员类继承成年人类，并要求在每个类中都提供有数据输入和输出的功能。
// ---------------------3-----------------------
//type People struct {
//	id       string
//	name     string
//	sex      string
//	birthday string
//}

//func (p *People) getInfo() {
//	fmt.Println("id: " + p.id)
//	fmt.Println("name: " + p.name)
//	fmt.Println("sex: " + p.sex)
//	fmt.Println("birthday: " + p.birthday)
//}

//type Adult struct {
//	People
//	education string
//	job       string
//}

//func (a *Adult) getInfo() {
//	// 调用父类方法
//	a.People.getInfo()
//	fmt.Println("education: " + a.education)
//	fmt.Println("job: " + a.job)
//}

//type Party struct {
//	Adult
//	parties string
//}

//func (p *Party) getInfo() {
//	p.Adult.getInfo()
//	fmt.Println("parties: " + p.parties)
//}
//func main() {
//	peo := People{"422802", "zhangsan", "male", "1994-10-10"}
//	peo.getInfo()
//	fmt.Println()
//	// 用peo 初始化 adult
//	adult := Adult{peo, "硕士", "工程师"}
//	adult.getInfo()
//	fmt.Println()
//	// 用adult 初始化party
//	party := Party{adult, "党员"}
//	party.getInfo()
//}

// ---------------------------------------------
