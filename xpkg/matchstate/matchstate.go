package matchstate

type MatchState int

const (
	NotStarted MatchState = iota
	FirstHalf
	HalfTime
	SecondHalf
	FullTime
	Paused
	Abandoned
	Postponed
)

func (m MatchState) ToString() string {
	return [...]string{
		"Not started",
		"First half",
		"Half time",
		"Second half",
		"Full time",
		"Paused",
		"Abandoned",
		"Postponed",
	}[m]
}
