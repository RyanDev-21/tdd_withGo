package integers

import (
	"fmt"
	"testing"
)


func TestAddr(t *testing.T){
	sum := Add(1,2)
	got := 3

	if sum != got{
		t.Errorf("got %d want %d", sum ,got)
	}
	
}


func ExampleAdd(){
	sum := Add(1,5)
	fmt.Println(sum)
	//Ouput :6
}
