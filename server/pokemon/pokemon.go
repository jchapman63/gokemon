package pokemon

// General pokemon type
type Pokemon struct {
	name string
	hp   int
}

func (p *Pokemon) Damage(loss int) {
	p.hp -= loss
}
