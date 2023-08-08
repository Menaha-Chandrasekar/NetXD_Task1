package main

import "fmt"

type Node struct {
	prev  *Node
	value int
	next  *Node
}
var head *Node = nil

func CreateNewNode(value int) *Node {
	var d Node
	d.next = nil
	d.value = value
	d.prev = nil
	return &d
}
func insert(value int) {
	d := CreateNewNode(value)
	if head == nil {
		head = d
	} else {
		d.next = head
		head.prev = d
		head = d
	}
}
func delete() {
	if head == nil {
		return
	}
	if head.next == nil {
		head = nil
		println("nill")
	} else {
		head = head.next
		head.prev = nil
	}

}
func display(){
	for temp := head;temp != nil;temp = temp.next{
	   fmt.Printf("%d ", temp.value)
	}
	fmt.Println()
 }

func main() {

	insert(10)
	insert(20)
	insert(30)
  delete()
  display()
	
}
