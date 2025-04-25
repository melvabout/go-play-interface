package main

type Human struct {
	Name string
}

func (h Human) Jump() string {
	return h.Name + " jumped!"
}

func (h Human) Walk() string {
	return h.Name + " walked!"
}

func (h Human) Sat() string {
	return h.Name + " sat!"
}
