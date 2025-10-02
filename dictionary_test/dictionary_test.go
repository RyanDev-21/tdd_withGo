package dictionarytest 


import "testing"


func TestSearch(t * testing.T){	
	
	dictionary := Dictionary{"test":"this is just a test"}
	t.Run("test with key ",func(t *testing.T) {


		got,_ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t,got,want)

	})


	t.Run("test with unknown",func(t *testing.T) {
		_,got := dictionary.Search("unknown")
		if got  == nil{
			t.Fatal("expect to get an error")
		}
		assertErrors(t,got,NotFoundError)
	})

}

func TestAdd(t *testing.T){
	word := "test"
	defination := "this is just a test"

	t.Run("test with new word",func(t *testing.T) {

	dictionary := Dictionary{}
	err :=dictionary.Add(word,defination)

	assertErrors(t,err,nil)
	assertDefination(t,dictionary,word,defination)
		
	})

	t.Run("testing with value already exists",func(t *testing.T) {
		dictionary := Dictionary{word: defination}

		err :=dictionary.Add(word,"new test")	

		assertErrors(t,err,WordAlreadyExistError)
		assertDefination(t,dictionary,word,defination)
	
	

	})


	

}


func TestUpdate(t *testing.T){
	word := "test"
	defination := "new update"

	dictionary := Dictionary{"test": "old value"}
	t.Run("normal update",func(t *testing.T) {



		err := dictionary.Update("test",defination)
		assertErrors(t, err, nil)	
		assertDefination(t,dictionary,word,defination)	
	})

	t.Run("testing for unknown word",func(t *testing.T) {
		word := "unknown"
		err :=dictionary.Update(word,defination)
		
		assertErrors(t,err,WordNotFoundToOperate)	
	})

}


func TestDelete(t *testing.T){
	word := "test"
	def := "delete test"
	dic := Dictionary{word: def}

	t.Run("testing normal case",func(t *testing.T) {

		err := dic.Delete(word) 
		assertErrors(t,err,nil)
		_,err = dic.Search(word)
		assertErrors(t,err,NotFoundError)
		

	})


	t.Run("testing for word not found",func(t *testing.T) {
		err := dic.Delete("unknown")
		assertErrors(t,err,WordNotFoundToOperate)

	})
}


func assertDefination(t testing.TB,dic Dictionary,word,def string){
	t.Helper()
	got, err := dic.Search(word)
	if err !=nil{
		t.Fatalf("Should have found the word ,%q",word)
	}

	assertStrings(t,got,def)

}

func assertStrings(t testing.TB,got, want string){
	t.Helper()	
	if got != want{
		t.Errorf("got %q want %q",got ,want)
	}
}

func assertErrors(t testing.TB,got ,want error){
	t.Helper()
	if got != want{
		t.Errorf("got errro %q want %q",got ,want)
	}
}





