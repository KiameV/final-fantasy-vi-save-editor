package util

import (
	"encoding/json"
	"fmt"

	"gitlab.com/c0b/go-ordered-json"
	"pixel-remastered-save-editor/models"
)

func SetIntInSlice(to *ordered.OrderedMap, key string, value int) (err error) {
	var (
		i, ok = to.GetValue(key)
		sl    []interface{}
	)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
		return
	}

	if sl, ok = i.([]interface{}); !ok || len(sl) < 9 {
		err = fmt.Errorf("unable to load cursed shield battle count")
		return
	}
	sl[9] = value
	to.Set("global", sl)
	return
}

func SetValue(to *ordered.OrderedMap, key string, value interface{}) (err error) {
	if !to.Has(key) {
		err = fmt.Errorf("unable to find %s", key)
	}
	to.Set(key, value)
	return
}

func SetFlag(to *ordered.OrderedMap, key string, value bool) error {
	var i int
	if value {
		i = 1
	}
	return SetValue(to, key, i)
}

func MarshalTo(to *ordered.OrderedMap, key string, value *ordered.OrderedMap) error {
	if !to.Has(key) {
		return fmt.Errorf("unable to find %s", key)
	}
	if v, err := value.MarshalJSON(); err != nil {
		return err
	} else {
		to.Set(key, string(v))
	}
	return nil
}

func Floor0(i int) int {
	if i < 0 {
		return 0
	}
	return i
}

func InvCount(eq *[]string, counts map[int]int, addedItems *[]int, id int, emptyID int) {
	var i models.IdCount
	if id == 0 {
		i.ContentID = emptyID
		i.Count = counts[emptyID]
	}

	if count, ok := counts[id]; ok {
		i.ContentID = id
		i.Count = count
	} else {
		// *addedItems = append(*addedItems, id)
		i.ContentID = id
		i.Count = 1
	}
	b, _ := json.Marshal(&i)
	*eq = append(*eq, string(b))
}

func SetTarget(d *ordered.OrderedMap, key string, value []interface{}) (err error) {
	var (
		t = ordered.NewOrderedMap()
		b []byte
	)
	if value != nil {
		t.Set(TargetKey, value)
	} else {
		t.Set(TargetKey, make([]interface{}, 0))
	}
	b, err = t.MarshalJSON()
	return SetValue(d, key, string(b))
}
