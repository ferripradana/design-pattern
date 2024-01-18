package gun

type Ak47 struct {
	Gun
}

func NewAk47() IGun {
	return &Ak47{Gun{
		Name:  "AK-47",
		Power: 4,
	}}
}
