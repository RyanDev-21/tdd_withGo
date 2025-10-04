package concurrency

import (
	"reflect"
	"testing"
	"time"
)


func MockWebsiteChecker(url string)bool{
	return url != "waat://furfurterwee.geds"
}

func TestCheckWebsites(t *testing.T){
	websites := []string{
		"https://google.com",
		"https://blog.go.test.com",
		"waat://furfurterwee.geds",

	}	
	want := map[string]bool{
		"https://google.com":true,
		"https://blog.go.test.com":true,
		"waat://furfurterwee.geds":false,

	}

	got := CheckWebsites(MockWebsiteChecker,websites)

	if !reflect.DeepEqual(want,got){
		t.Errorf("want %v got %v",want,got)
	}

}

func slowWebsiteChecker(_ string)bool{
	time.Sleep(20*time.Millisecond)
	return true
}


func BenchmarkCheckWebsites(b *testing.B){
	urls := make([]string,100)	

	for i:= range urls{
		urls[i] = "a url"	
	}
	for b.Loop(){
		CheckWebsites(slowWebsiteChecker,urls)
	}
	


}
