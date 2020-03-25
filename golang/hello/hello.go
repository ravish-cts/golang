package main 

import (
	"fmt"
	"os"
	"time" 
	"errors"
) 

func getGreeting(timeNow int)(string,error){
 var message string
 if timeNow < 7	{
	 message = "To early time"
	 err := errors.New("To early time")
	 return message,err	
	}else {
	 message = "Good Day" 
	} 
	 return message,nil	
}
func main() {
	var timeNow int; 
	if len(os.Args) > 1{
	fmt.Println(os.Args[1]);
	} else {
	fmt.Println("Hello, my first hello from Golang !");
	}
	timeNow = time.Now().Hour();
        Msg,err := getGreeting(timeNow)
	if err != nil {
	  Msg = "This is too early !"
	  os.Exit(1); 
	}else {
	 fmt.Println("Wait for greeting");
	} 
	fmt.Println(Msg);
} 
