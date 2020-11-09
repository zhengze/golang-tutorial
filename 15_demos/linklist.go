package main

import "fmt"

type Object interface{}

type Node struct {
    data Object
    next *Node
}

type List struct {
    headNode *Node
}

func NewNode(data Object, next *Node) *Node {
    return &Node{data, next}
}

func (list *List) IsEmpty() bool {
    return list.headNode == nil
}

func (list *List) Add(node *Node) *List {
    headNode := list.headNode  
    if headNode.next != nil {
        node.next = headNode.next 
    } 
    headNode.next = node
    return list
}

func (list *List) Append(node *Node) *List {
    if list.IsEmpty() {
        list.headNode = node
        return list
    }
    curNode := list.headNode
    for curNode.next != nil {
        curNode = curNode.next
    }
    curNode.next = node
    return list
}

func (list *List) Insert(index int, data Object) {
    if (index >= 0 && index < list.Length()) {
        count := 0
        if !list.IsEmpty() {
            curNode := list.headNode
            for curNode != nil && count < index {
                count++
                curNode = curNode.next
            }
            node := NewNode(data, curNode.next)
            curNode.next = node
        }
    } 
}

func (list *List) Remove(index int) {
    if (index >= 0 && index < list.Length()) {
        count := 0
        if !list.IsEmpty() {
            curNode := list.headNode
            for curNode != nil && count < index-1 {
                count++
                curNode = curNode.next
            }
            curNode.next = curNode.next.next
        }
    } 
}

func PrintList(list *List) {
    cur := list.headNode 
    for cur != nil {
        fmt.Println(cur.data)
        cur = cur.next
    }
}

func (list *List) Length() int {
    var length int
    curNode := list.headNode
    for curNode != nil {
        length++
        curNode = curNode.next
    }
    return length
}

func ReverseList(head *Node) *Node {
    cur := head
    var pre *Node = nil
    for cur != nil {
        pre, cur, cur.next = cur, cur.next, pre
    }
    return cur
}

func main(){
    fmt.Println("NewNode======================")
    node := NewNode(1, nil)
    fmt.Println(node.data)
    list := &List{node}
    PrintList(list)
    node2 := NewNode("a", nil)
    list.Append(node2)
    fmt.Println("Add======================")
    PrintList(list)
    node1 := NewNode(2, nil)
    list.Add(node1)
    fmt.Println("Length======================")
    PrintList(list)
    fmt.Println(list.Length())
    fmt.Println("Insert======================")
    list.Insert(1, 4)
    PrintList(list)
    fmt.Println("Remove======================")
    list.Remove(2)
    PrintList(list)
    fmt.Println("ReverseList======================")
    ReverseList(node)
    PrintList(list)
}
