package main
import "fmt"

func main(){
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	var i,num,dig,sum int
	sum=0
	for i=0;i<size;i++{
        fmt.Scanf("%d ",&arr[i]);
    }
	for i=0;i<size;i++{
		if i%2==0{
			num=arr[i]*2
			for ;num!=0; {
				dig=num%10
				sum=sum+dig
				num=num/10
			}
		}else{
			num=arr[i]*1
			sum=sum+num
		}
	}
	if sum%10==0{
		fmt.Printf("Valid \n")
	}else{
		fmt.Printf("Invalid \n")
	}

}