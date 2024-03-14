package main

import "fmt"

// 结构体
// type student struct {
// 	name string
// 	age  int
// }
// func main() {
// m := make(map[string]*student)
// fmt.Println(m)
// stus := []student{
// 	{name: "pprof.cn", age: 18},
// 	{name: "测试", age: 23},
// 	{name: "博客", age: 28},
// }
// fmt.Println(stus)
// for _, stu := range stus {
// 	fmt.Println(&stu)
// 	m[stu.name] = &stu
// }
// for k, v := range m {
// 	fmt.Println(k, "=>", v.name)
// }
// }

// 方法
/*
    func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
        函数体
    }
1.接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名。例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
2.接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
3.方法名、参数列表、返回参数：具体格式与函数定义相同。
*/
// type Person struct {
// 	name string
// 	age  int8
// }

// func NewPerson(name string, age int8) *Person {
// 	return &Person{
// 		name: name,
// 		age:  age,
// 	}
// }
// func (p Person) Dream() {
// 	fmt.Printf("%s的梦想是学好GO语言\n", p.name)
// }
// func main() {
// 	p1 := NewPerson("测试", 25)
// 	p1.Dream()
// }

// 结构体与JSON序列化
// type Student struct {
// 	ID     int
// 	Gender string
// 	Name   string
// }
// type Class struct {
// 	Title   string
// 	Student []*Student
// }

// func main() {
// 	c := &Class{
// 		Title:   "101",
// 		Student: make([]*Student, 0, 200),
// 	}
// 	fmt.Println(c)
// 	for i := 0; i < 10; i++ {
// 		stu := &Student{
// 			Name:   fmt.Sprintf("stu%02d", i),
// 			Gender: "男",
// 			ID:     i,
// 		}
// 		c.Student = append(c.Student, stu)
// 	}
// 	fmt.Println(c)
// 	// JSON序列化: 结构体 --> JSON格式的字符串
// 	data, err := json.Marshal(c)
// 	if err != nil {
// 		fmt.Println("json marshal failed")
// 		return
// 	}
// 	fmt.Printf("##########json:%s\n", data)
// 	// JSON反序列化: JSON格式的字符串 --> 结构体
// 	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
// 	c1 := &Class{}
// 	err = json.Unmarshal([]byte(str), c1)
// 	if err != nil {
// 		fmt.Printf("json unmarshal failed")
// 		return
// 	}
// 	fmt.Printf("%#v\n", c1)
// }

// type student struct {
// 	id   int
// 	name string
// 	age  int
// }

// func demo(ce []student) {
// 	// 切片是引用传递，是可以改变值的
// 	ce[1].age = 999
// 	fmt.Println("***************", ce)
// }
// func main() {
// 	var ce []student // 定义一个切片类型的结构体
// 	ce = []student{
// 		student{1, "xiaoming", 23},
// 		student{2, "xiaozhang", 33},
// 	}
// 	fmt.Println("&&&&&", ce)
// 	demo(ce)
// 	fmt.Println("$$$$$$$$", ce)
// }
// &&&&& [{1 xiaoming 23} {2 xiaozhang 33}]
// *************** [{1 xiaoming 23} {2 xiaozhang 999}]
// $$$$$$$$ [{1 xiaoming 23} {2 xiaozhang 999}]

// 删除map类型的结构体
type student struct {
	id   int
	name string
	age  int
}

func main() {
	ce := make(map[int]student)
	ce[1] = student{1, "xiaolizi", 22}
	ce[2] = student{2, "matt", 23}
	fmt.Println(ce)
	delete(ce, 2)
	fmt.Println(ce)
}
