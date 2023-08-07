package main
import "fmt"

func main() {
  string:="ABCDEFGHIJKL"
  
  var i int
   temp:=[]byte(string)
  for i=2;i<len(temp)-1;i+=4{
           temp[i], temp[i+1] = temp[i+1], temp[i] 
  }
  
  fmt.Printf("%s\n",temp)
}