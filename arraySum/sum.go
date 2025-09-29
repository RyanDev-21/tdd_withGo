package arraySum

func Sum(array []int)int{
	var sum int
	for _,v:=range array{
		sum +=v 
	}
	return sum
}


func SumAll(numbersToSum ...[]int) []int{
	var slice []int	
	for _,numbers := range numbersToSum{
		slice= append(slice,Sum(numbers))
	}
	return slice


}


func SumAllTails(numbersToTail ...[]int)[]int{
	var slice []int
	for _,slices := range numbersToTail{
		if len(slices) == 0{
			slice= append(slice,0)
		}else{


			tail := slices[1:]
			slice = append(slice,Sum(tail))
		}

	}
	return slice
}
