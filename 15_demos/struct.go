package main

import "fmt"

type I interface{
    setAge(int8)
    getAge() int8
    getName() string
}

type People struct{
    name string
    age int8
}

func NewPeople(name string, age int8) *People {
    return &People{name,age}
}

func(p *People) setName(name string) {
    p.name = name
}

func(p *People) setAge(age int8) {
    p.age = age
}

func(p People) getName() string{
    return p.name
}

func(p People) getAge() int8{
    return p.age
}

func getinfo(obj I){
   fmt.Println(obj.getName())
   fmt.Println(obj.getAge())
}

func main(){
   p := NewPeople("zhangsan", 15)
   p.setName("lisi")
   getinfo(p)
}
