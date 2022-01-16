package main
import "fmt"

func main()  {
	linguagens := [] string{
		"Java",
		"PHP",
		"Angular",
		"Go",
	}
	for _,value :=range linguagens{
		fmt.Println(value)
	}

	arr := [4] int{
		1,2,3,4,
	}		

	for _,value :=range arr{
		fmt.Println(value)
	}

	mapper:= map[int]string{
		 1:"Java",
		 2 :"PHP",
		 3 :"Angular",
		 4: "Go",

	}
	for key,value :=range mapper{
		fmt.Printf("A chave-> \"%d\" e o valor \"%s\"\n",key ,value)
	}

}


