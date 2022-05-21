package Observer

type Observer interface {
	Update(subject *Subject)
}
