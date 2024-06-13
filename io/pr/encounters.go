package pr

// func (p *PR) CompleteAllEncounters(saveDir string) (err error) {
// 	const file = "dp3fS2vqP7GDj8eF72YKqbT7FIAF=e7Shy2CsTITm2E="
// 	var (
// 		out     []byte
// 		s       string
// 		names   []unicodeNameReplace
// 		toFile  = path.Join(saveDir, file)
// 		temp    = filepath.Join(global.PWD, "temp")
// 		saveCmd = exec.Command("python", "./io/python/io.py", "obfuscateFile", toFile, temp)
// 	)
// 	// p.names = make([][]rune, 0, 40)
// 	if out, err = p.readFile(toFile); err != nil {
// 		return
// 	}
//
// 	s = strings.ReplaceAll(s, `\\r\\n`, "")
// 	s = p.fixEscapeCharsForLoad(s)
// 	if strings.Contains(s, "\\x") {
// 		b := []byte(s)
// 		if b, err = p.replaceUnicodeNames(b, &names); err != nil {
// 			return
// 		}
// 		s = string(b)
// 	}
//
// 	if len(os.Args) >= 2 && os.Args[1] == "print" {
// 		if _, err = os.Stat("encounters.json"); errors.Is(err, os.ErrNotExist) {
// 			if _, err = os.Create("encounters.json"); err != nil {
// 			}
// 		}
// 		t := strings.ReplaceAll(s, `\`, ``)
// 		t = strings.ReplaceAll(t, `"{`, `{`)
// 		t = strings.ReplaceAll(t, `}"`, `}`)
// 		var prettyJSON bytes.Buffer
// 		err = json.Indent(&prettyJSON, []byte(t), "", "\t")
// 		if err == nil {
// 			err = ioutil.WriteFile("encounters.json", prettyJSON.Bytes(), 0755)
// 		}
// 	}
//
// 	base := ordered.NewOrderedMap()
// 	base.UnmarshalJSON([]byte(s))
// 	md := ordered.NewOrderedMap()
// 	if err = p.unmarshalFrom(base, "monsterDefeats", md); err != nil {
// 		return errors.New("unable to parse encounters file, monsterDefeats")
// 	}
// 	keys := md.Get("keys")
// 	if keys == nil {
// 		return errors.New("unable to parse encounters file, monsterDefeats.keys")
// 	}
// 	sl := keys.([]interface{})
// 	jn := make([]json.Number, len(sl))
// 	definedIDs := make(map[int16]bool)
// 	for i := range sl {
// 		jn[i] = sl[i].(json.Number)
// 		j, _ := jn[i].Int64()
// 		definedIDs[int16(j)] = true
// 	}
// 	var idsToAdd []int16
// 	for _, id := range models.MonsterIDs {
// 		if _, found := definedIDs[id]; !found {
// 			idsToAdd = append(idsToAdd, id)
// 			jn = append(jn, json.Number(fmt.Sprintf("%d", id)))
// 		}
// 	}
// 	md.Set("keys", jn)
//
// 	vs := md.Get("values")
// 	if vs == nil {
// 		return errors.New("unable to parse encounters file, monsterDefeats.values")
// 	}
// 	sl = vs.([]interface{})
// 	jn = make([]json.Number, len(sl)+len(idsToAdd))
// 	total := 2000
// 	for i := range sl {
// 		if i == 0 {
// 			jn[i] = "2000"
// 			continue
// 		}
// 		if sl[i].(json.Number) == "0" {
// 			jn[i] = "1"
// 			total++
// 		} else {
// 			jn[i] = sl[i].(json.Number)
// 			j, _ := jn[i].Int64()
// 			total += int(j)
// 		}
// 	}
// 	for i := 0; i < len(idsToAdd); i++ {
// 		jn[i+len(sl)] = "1"
// 	}
// 	md.Set("values", jn)
// 	base.Set("totalSubjugationCount", total)
//
// 	var sf interface{}
// 	if sf, err = p.getFromTarget(base, "scenarioFlags"); err != nil {
// 		return
// 	}
// 	sl = sf.([]interface{})
// 	for i := range sl {
// 		if sl[i].(json.Number) == "0" {
// 			sl[i] = json.Number("1")
// 		}
// 	}
// 	if err = p.setTarget(base, "scenarioFlags", sl); err != nil {
// 		return
// 	}
//
// 	if err = p.marshalTo(base, "monsterDefeats", md); err != nil {
// 		return
// 	}
//
// 	var data []byte
// 	if data, err = json.Marshal(base); err != nil {
// 		return
// 	}
//
// 	if _, err = os.Stat(temp); errors.Is(err, os.ErrNotExist) {
// 		if _, err = os.Create(temp); err != nil {
// 			return fmt.Errorf("failed to create save file %s: %v", toFile, err)
// 		}
// 	}
// 	defer os.Remove(temp)
//
// 	if err = os.WriteFile(temp, data, 0755); err != nil {
// 		return fmt.Errorf("failed to create temp file %s: %v", toFile, err)
// 	}
//
// 	if _, err = os.Stat(toFile); errors.Is(err, os.ErrNotExist) {
// 		if _, err = os.Create(toFile); err != nil {
// 			return fmt.Errorf("failed to create save file %s: %v", toFile, err)
// 		}
// 	}
//
// 	if out, err = saveCmd.Output(); err != nil {
// 		if ee, ok := err.(*exec.ExitError); ok {
// 			err = errors.New(string(ee.Stderr))
// 		} else {
// 			err = fmt.Errorf("%s: %v", string(out), err)
// 		}
// 	}
// 	return
// }
