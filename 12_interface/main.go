package main

import (
    "fmt"
)

type Control interface {
    Run()
}

type Cat struct{
    name string
}

func (this *Cat) GetName() string{
    return "this cat's name is:" + this.name
}

func (this *Cat) Run() {
    fmt.Println(this.name+" is running")
}

type Dog struct{
    name string
}

func(this *Dog) Run(){
    fmt.Println(this.name+" is running")
}


func main(){

    var control Control
    cat := new(Cat)
    cat.name = "tom"
    fmt.Println(*cat)
    cat.GetName()
    control = cat
    control.Run()

    dog := new(Dog)
    dog.name = "jerry"
    control = dog
    control.Run()
}
