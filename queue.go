package main
import "fmt"



func dequeue()int{
if front==-1||front>rear{
	fmt.Printf("Empty Queue")
    return -1
}

var element int
element=queue[front]
front++
return element
}

func enqueue(n int){
if rear==100-1{
fmt.Printf("size exceed")
}
if front==-1{
front=0
}
rear++
queue[rear]=n

}

var front int=-1
var rear int=-1
var queue[100]int

func main() {
enqueue(10)
enqueue(20)
fmt.Printf("%d ",dequeue())
fmt.Printf("%d ",dequeue())
}