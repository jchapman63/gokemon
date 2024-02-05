package pokemon

// the pokemon struct will need a move set. Moves will be a struct of their own.  I will start off just worrying about
// damage based moves to get me started.

// General pokemon type
// lowercase fields are not exposed, so these are uppercase
type Pokemon struct {
	Name  string       `json:"pokemon-name"`
	Hp    int          `json:"hp"`
	Moves []DamageMove `json:"moves"`
}

func (p *Pokemon) Attack(o *Pokemon, attack DamageMove) {
	o.Hp -= attack.Power
}

type DamageMove struct {
	Name  string
	Power int
}
