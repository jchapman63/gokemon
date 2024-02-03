package pokemon

// General pokemon type
// lowercase fields are not exposed, so these are uppercase
type Pokemon struct {
	Name string
	Hp   int
}

func (p *Pokemon) Damage(loss int) {
	p.Hp -= loss
}
