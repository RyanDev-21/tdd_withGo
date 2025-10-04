package main

import (
	"os"
	"time"

	cd "example.com/printName/CountDown"
)


func main(){
	sleeper := cd.NewConfigurationSleeper(1*time.Second,time.Sleep)
	cd.Countdown(os.Stdout,sleeper)
}



// func main(){
// 	log.Fatal(http.ListenAndServe(":5001",http.HandlerFunc(dependencyinjection.MyGreetHandler)))
// }
