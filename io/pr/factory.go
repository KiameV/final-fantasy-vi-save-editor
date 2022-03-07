package pr

type PR struct {
	Base       map[string]interface{}
	UserData   map[string]interface{}
	Characters []map[string]interface{}
}

func NewPR() *PR {
	return &PR{
		Base:       make(map[string]interface{}),
		UserData:   make(map[string]interface{}),
		Characters: make([]map[string]interface{}, 40),
	}
}
