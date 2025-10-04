package countdown


import "time"

type Sleeper interface{
	Sleep()
}

type ConfigurationSleeper struct{
	duration time.Duration
	sleep  func(time.Duration)
}

func (c *ConfigurationSleeper) Sleep(){
	c.sleep(c.duration)
}

func NewConfigurationSleeper(duration time.Duration,sleep func(time.Duration)) *ConfigurationSleeper{
	return &ConfigurationSleeper{duration,sleep}
}

type SpyTime struct{
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration){
	s.durationSlept = duration
}
