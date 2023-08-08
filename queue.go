package main
import "fmt"

type Queue struct{
	front int
	rear int 
	size int
}
var queue= make([]int, 100)

func (q *Queue) dequeue()int{
    if q.front==-1||q.front>q.rear{
	fmt.Printf("Empty Queue")
    return -1
}

    var element int

    element=queue[q.front]
    q.front++
    return element
}

func (q *Queue) enqueue(n int){
    if q.rear==q.size-1{
    fmt.Printf("size exceed")
}
    if q.front==-1{
    q.front=0
}
    q.rear++
    queue[q.rear]=n

}

func main() {
	q := Queue{-1,-1,100}
    q.enqueue(10) 
    q.enqueue(20)
    fmt.Printf("%d ",q.dequeue())
    fmt.Printf("%d ",q.dequeue())
}
