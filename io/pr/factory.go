package pr

type PR struct {
	data       []byte
	Base       map[string]interface{}
	UserData   map[string]interface{}
	Characters []map[string]interface{}
}

var pr *PR

func NewPR() *PR {
	if pr == nil {
		pr = &PR{
			Base:       make(map[string]interface{}),
			UserData:   make(map[string]interface{}),
			Characters: make([]map[string]interface{}, 40),
		}
	}
	return pr
}
