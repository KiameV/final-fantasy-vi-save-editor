package models

type CurrentMax struct {
	Current int
	Max     int
}

func (m *CurrentMax) Fix() {
	if m.Current > m.Max {
		m.Current = m.Max
	}
}
