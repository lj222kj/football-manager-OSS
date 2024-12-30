package positions

type Position int

const (
	Goalkeeper Position = iota
	CentreBack
	LeftBack
	RightBack
	DefensiveMidfielder
	CentralMidfielder
	AttackingMidfielder
	LeftWinger
	RightWinger
	Striker
)

func (p Position) ToString() string {
	return [...]string{
		"Goalkeeper",
		"Centre Back",
		"Left Back",
		"Right Back",
		"Defensive Midfielder",
		"Central Midfielder",
		"Attacking Midfielder",
		"Left Winger",
		"Right Winger",
		"Striker",
	}[p]
}
