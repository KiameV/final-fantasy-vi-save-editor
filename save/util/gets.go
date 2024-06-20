package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	jo "gitlab.com/c0b/go-ordered-json"
	"pixel-remastered-save-editor/models"
)

func GetString(c *jo.OrderedMap, key string) (s string, err error) {
	j, ok := c.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}
	if s, ok = j.(string); !ok {
		err = fmt.Errorf("unable to parse field %s value %v ", key, j)
	}
	return
}

func GetBool(c *jo.OrderedMap, key string) (b bool, err error) {
	j, ok := c.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}
	if b, ok = j.(bool); !ok {
		err = fmt.Errorf("unable to parse field %s value %v", key, j)
	}
	return
}

func GetInt(c *jo.OrderedMap, key string) (i int, err error) {
	j, ok := c.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}

	k := reflect.ValueOf(j)
	switch k.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(k.Int()), nil
	case reflect.Float32, reflect.Float64:
		return int(k.Float()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int(k.Uint()), nil
	case reflect.String:
		var l int64
		l, err = strconv.ParseInt(k.String(), 10, 32)
		if err == nil {
			i = int(l)
		}
	default:
		err = fmt.Errorf("unable to parse field %s value %v ", key, j)
	}
	return
}

func GetUint(c *jo.OrderedMap, key string) (i uint64, err error) {
	j, ok := c.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}

	k := reflect.ValueOf(j)
	switch k.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(k.Int()), nil
	case reflect.Float32, reflect.Float64:
		return uint64(k.Float()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uint64(k.Uint()), nil
	case reflect.String:
		var l int64
		l, err = strconv.ParseInt(k.String(), 10, 64)
		if err == nil {
			i = uint64(l)
		}
	default:
		err = fmt.Errorf("unable to parse field %s value %v ", key, j)
	}
	return
}

func GetFloat(c *jo.OrderedMap, key string) (f float64, err error) {
	j, ok := c.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}

	k := reflect.ValueOf(j)
	switch k.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return k.Float(), nil
	case reflect.Float32, reflect.Float64:
		return k.Float(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return k.Float(), nil
	case reflect.String:
		f, err = strconv.ParseFloat(k.String(), 64)
	default:
		err = fmt.Errorf("unable to parse field %s value %v ", key, j)
	}
	return
}

func GetJsonInts(data *jo.OrderedMap, key string) (ints []interface{}, err error) {
	j, ok := data.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}
	ints, ok = j.([]interface{})
	if !ok {
		err = fmt.Errorf("unable to load %s", key)
	}
	return
}

func GetFlag(m *jo.OrderedMap, key string) (b bool, err error) {
	var i int
	if i, err = GetInt(m, key); err != nil {
		return
	}
	if i == 0 {
		b = false
	} else {
		b = true
	}
	return
}

func GetIntFromSlice(from *jo.OrderedMap, key string) (v int, err error) {
	var (
		i, ok = from.GetValue(key)
		sl    []interface{}
		i64   int64
	)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
		return
	}

	if sl, ok = i.([]interface{}); !ok || len(sl) < 9 {
		err = fmt.Errorf("unable to load cursed shield battle count")
		return
	}

	if i64, err = sl[9].(json.Number).Int64(); err != nil {
		return
	}
	v = int(i64)
	return
}

func UnmarshalFrom(from *jo.OrderedMap, key string, m *jo.OrderedMap) (err error) {
	i, ok := from.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}
	switch t := i.(type) {
	case string:
		err = m.UnmarshalJSON([]byte(t))
	default:
		err = fmt.Errorf("cannot unmarshal unkown type %v", i)
	}
	return
}

func Unmarshal(i interface{}, m *map[string]interface{}) error {
	// s := strings.ReplaceAll(i.(string), `\\"`, `\"`)
	return json.Unmarshal([]byte(i.(string)), m)
}

func UnmarshalEquipment(m *jo.OrderedMap) (idCounts []*models.IdCount, err error) {
	i, ok := m.GetValue(EquipmentList)
	if !ok {
		return nil, fmt.Errorf("%s not found", EquipmentList)
	}

	eq := jo.NewOrderedMap()
	if err = eq.UnmarshalJSON([]byte(i.(string))); err != nil {
		return
	}

	if i, ok = eq.GetValue("values"); ok && i != nil {
		l, ok := i.([]interface{})
		if !ok {
			return nil, errors.New("unable to load equipment, try saving again")
		}
		idCounts = make([]*models.IdCount, len(l))
		for j, v := range i.([]interface{}) {
			if err = json.Unmarshal([]byte(v.(string)), &idCounts[j]); err != nil {
				return
			}
		}
	}
	return
}

func FromTarget(data *jo.OrderedMap, key string) (i interface{}, err error) {
	var (
		slTarget = jo.NewOrderedMap()
		ok       bool
	)
	if err = UnmarshalFrom(data, key, slTarget); err != nil {
		return
	}
	if i, ok = slTarget.GetValue(TargetKey); !ok {
		err = fmt.Errorf("unable to find %s", TargetKey)
	}
	return
}
