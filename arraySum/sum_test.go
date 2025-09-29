package arraySum

import (
	"reflect"
	"testing"
)


func TestSum(t *testing.T){
	t.Run("sum the number of any array size", func(t *testing.T) {
		
		numbers := []int{1,2,3}
		sum := Sum(numbers)
		expected := 6
		if sum!=expected{
			t.Errorf("got : %d want: %d  array: %v",sum,expected,numbers)
		}


	})

	t.Run("sum all the nubmers containing in the slice and return a new one", func(t *testing.T) {
		sliceOne := []int{2,3,5}
		sliceTwo := []int{1,2}
		sum := SumAll(sliceOne,sliceTwo)
		expected := []int{10,3}
		
		if !reflect.DeepEqual(sum,expected){
			t.Errorf("got : %d want: %d ",sum,expected )
		}


	})

}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB,got,want []int){
		if !reflect.DeepEqual(got,want){
			t.Errorf("got : %v want: %v", got ,want)
		}
	}


	t.Run("make the sum of some slices",func(t *testing.T) {

		sliceOne := []int{2,3,5}
		sliceTwo := []int{10,23,45}
		sum := SumAllTails(sliceOne,sliceTwo)
		expected := []int{8,68}

		checkSums(t,sum,expected)	

	})


	t.Run("safely sum empty slices",func(t *testing.T) {
		sliceOne := []int{}
		sliceTwo := []int{23,45,54}

		sum := SumAllTails(sliceOne,sliceTwo)

		expected := []int{0,99}

		checkSums(t,sum,expected)
	})




}


func BenchmarkSum(b *testing.B){
	numbers := []int{1,39,32,53}	
	for b.Loop(){
		Sum(numbers)	
	}	
}

func BenchmarkSumAll(b *testing.B){
	for b.Loop(){

		sliceOne := []int{2,34,5,10}
		sliceTwo := []int{23,53,12,3}

		SumAll(sliceOne,sliceTwo)
	}
}

func BenchmarkSumAllTails(b *testing.B){
	for b.Loop(){

		sliceOne := []int{2,34,5,10}
		sliceTwo := []int{23,53,12,3}

		SumAllTails(sliceOne,sliceTwo)
	}

}


