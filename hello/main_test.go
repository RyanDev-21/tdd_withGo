package main

import "testing"


func TestHello(t *testing.T){
	t.Run("saying hello to people",func(t *testing.T){
		got := Hello("Chris","English")
		want := "Hello , Chris"
		
		assertCorrectMessage(t,got,want)
		
	})

	t.Run("if the string argument is empty print Hello,World",func(t *testing.T){
		got := Hello("","English")
		want:= "Hello , World"

			
		assertCorrectMessage(t,got,want)

	})

	t.Run("saying hello in spanish",func(t *testing.T){
		got := Hello("Elodie","Spanish")
		want := "Hola , Elodie"

		assertCorrectMessage(t,got,want)
	})

	t.Run("saying Hello in French", func(t *testing.T){
		got := Hello("HtetLin","French")
		want := "Bonjour , HtetLin"
		
		assertCorrectMessage(t, got ,want)
	})
}

func assertCorrectMessage(t testing.TB,g ,want string){
	t.Helper()
	if g !=want {
		t.Errorf("got %q want %q",g ,want)
	}
}

