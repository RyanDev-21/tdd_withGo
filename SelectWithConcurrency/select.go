package selectwithconcurrency

import (
	"net/http"
	"time"
	"fmt"
)


var timeout = 3 * time.Millisecond

func Racer(first ,second string,timeout time.Duration) (string,error){
	select{
	case <- ping(first):
	 return first,nil
	case <- ping(second):
	return second,nil
	case <- time.After(timeout):
	return "",fmt.Errorf("timed out waiting for %s and %s",first ,second)
	}

}


func ping(url string)chan struct{}{
	ch := make(chan struct{})
	go func(){
		http.Get(url)
		close(ch)
	}()
	return ch
}



