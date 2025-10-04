package countdown

import (
	"fmt"
	"io"
)


const (
	countDownStart = 3
	finalWord = "Go!"
)

func Countdown(writer io.Writer,s Sleeper){
	for i :=countDownStart;i>0;i--{
		fmt.Fprintln(writer,i)
		s.Sleep() 
	
	}

	fmt.Fprint(writer,finalWord)


}
