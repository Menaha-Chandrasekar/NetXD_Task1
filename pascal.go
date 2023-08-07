package main
import "fmt"


func pascal1(n int)int{
    var f int
    for f=1;n>1;n--{
        f=f*n
    }
    return f
}
func pascal(n,r int)int{
    return pascal1(n)/(pascal1(n-r)*pascal1(r))
}

func main() {
  var i,j,n int
  n=5
  
  for i=0;i<=n;i++{
      for j=0;j<=n-i;j++{
          fmt.Printf("  ")
      } 
      for j=0;j<=i;j++{
          fmt.Printf(" %3d",pascal(i,j))
      }
      fmt.Println()
  }
}