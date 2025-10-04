package countdown

type SpySleepCountDown struct{
	Calls []string
}

func (s *SpySleepCountDown) Sleep(){
	s.Calls = append(s.Calls,sleep)
}


func (s *SpySleepCountDown) Write(p []byte)(n int,err error){
	s.Calls = append(s.Calls,write)
	return
}

const sleep = "sleep"
const write = "write"
