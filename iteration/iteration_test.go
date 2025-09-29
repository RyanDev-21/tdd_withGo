package iteration

import (
	"fmt"
	"testing"
)






func TestIterate(t *testing.T){
	t.Run("Has to repeat the inpout four times",func(t *testing.T) {
		repeated := Iterate("a",10)
		expected := "aaaaaaaaaa"

		if repeated != expected{
			t.Errorf("got : %q expected: %q ",repeated,expected)
		}
	
	})
}

func BenchmarkIterate(b *testing.B){
	for b.Loop(){
		Iterate("a",10)
	}
}

func ExampleIterate(){
	repeated := Iterate("a",5)
	fmt.Println(repeated)
	//Output : "aaaaa"		
}
