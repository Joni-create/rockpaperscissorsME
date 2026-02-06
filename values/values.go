package values

type Value int

const (
	Rock Value = iota
	Paper
	Scissors
)

// String liefert den Wert als String.
func (v Value) String() string {
	switch v {
	case Rock:
		return "Stein"
	case Paper:
		return "Papier"
	case Scissors:
		return "Schere"
	default:
		return "Unbekannt"
	}
}

// Beats gibt an, ob der Wert v den Wert w schlägt.
func (v Value) Beats(w Value) bool {
	switch v {
	case Rock:
		return w == Scissors
	case Paper:
		return w == Rock
	case Scissors:
		return w == Paper
	default:
		return false
	}
}
