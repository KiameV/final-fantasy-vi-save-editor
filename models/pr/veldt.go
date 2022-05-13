package pr

type Veldt struct {
	Encounters []bool
}

var veldt *Veldt

func GetVeldt() *Veldt {
	if veldt == nil {
		veldt = &Veldt{}
	}
	return veldt
}
