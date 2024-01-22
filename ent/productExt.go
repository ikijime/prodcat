package ent

func (pr *Product) SayHello() string {
	return pr.Name + " says Hello!"
}
