package  dictionarytest 


type Dictionary map[string]string

type DictionaryErr string 


func (dE DictionaryErr) Error()string{
	return string(dE)
}



const (NotFoundError = DictionaryErr("could not find the word you were looking for")
 WordAlreadyExistError = DictionaryErr("word already exist")
 WordNotFoundToOperate = DictionaryErr("Cannot perform operation on this word, word doesn't exist in dictionary yet")
)



func (dic Dictionary) Search(word string)(string,error){
	defination , ok := dic[word]
	if !ok{
		return "", NotFoundError
	}
	return defination,nil
}


func (dic Dictionary) Add(word ,defination string)error{
	// _, err := dic.Search(word)
	//
	// switch err {
	// 	case NotFoundError: 
	// 		dic[word] = defination
	// 		return nil
	// case nil:
	// 	return WordAlreadyExistError
	//
	// default:
	// 	return err
	//
	// }

	if dic[word] == ""{
		dic[word] = defination
		return nil
	}
	return WordAlreadyExistError
}


func (dic Dictionary) Update(word, newDef string)error{
	_,err:= dic.Search(word)
	switch err {
	case NotFoundError:
		return WordNotFoundToOperate
	case nil :
		dic[word] = newDef
		return nil
	default: 
	return err
	}

}

func (dic Dictionary) Delete(word string)error{
	_,err := dic.Search(word)
	switch err{
	case NotFoundError:
		return WordNotFoundToOperate
	case nil:
		delete(dic,word)
		return nil
	default:
	return err
		
	}
}
