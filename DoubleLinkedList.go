package main
import "fmt"
type Node struct {
   prev *Node
   value int
   next *Node
}
func CreateNewNode(value int) *Node{
   var d Node
   d.next = nil
   d.value = value
   d.prev = nil
   return &d
}
func Traverse(node1 * Node){

   temp := node1
   for temp!= nil{
      fmt.Printf("%d ", temp.value)
      temp = temp.next
   }
   fmt.Println()
}
func main(){
   
   node1 := CreateNewNode(10)
   node2 :=CreateNewNode(11)
   node3 :=CreateNewNode(12)
   node1.next=node2
   node2.prev=node1
   node2.next=node3
   node3.prev=node2
   Traverse(node1)
} 

                            