package budget

// type CSVManager struct {
// 	db *gorm.DB
// }

// func (m *CSVManager) getFileConfig(fileName string) {
// 	// var result Configuration
// 	// db.Raw(`
// 	// 	with temp (id, is_match) as
// 	// 	(
// 	// 		select id,
// 	// 			regexp_matches(?,  pattern)
// 	// 		from
// 	// 			public.csv_file_configurations cfc
// 	// 	)
// 	// 	select id from temp where is_match is not null limit 1;
// 	// `, fileName).Scan(&result)

// 	// fmt.Println(result)

// 	// for _, conf := range m.configs {
// 	// 	if conf.FilePattern == "" {
// 	// 		continue
// 	// 	}

// 	// 	match, err := regexp.MatchString(conf.FilePattern, filename)

// 	// 	if err != nil {
// 	// 		return conf, err
// 	// 	}

// 	// 	if !match {
// 	// 		continue
// 	// 	}

// 	// 	if match {
// 	// 		return conf, nil
// 	// 	}
// 	// }

// 	// return Configuration{}, errors.New("could not find config for file")
// }

// func (m *CSVManager) getVal(col CSVColumn, row CSVRowData) (interface{}, error) {
// 	v := row[col.Ordinal-1]

// 	switch col.Type {
// 	case Float64Column:
// 		return m.parseString(v)
// 	case StrColumn:
// 		return m.parseString(v)
// 	case TimestampColumn:
// 		return m.parseTimestamp(v, col.Args)
// 	default:
// 		return nil, errors.New("invalid column type")
// 	}
// }

// func (m *CSVManager) parseTimestamp(v interface{}, args JSONB) (t time.Time, err error) {
// 	s, err := m.parseString(v)

// 	if err != nil {
// 		return t, err
// 	}

// 	// Parse timestamp format from args
// 	format, ok := args["timestamp_format"].(string)

// 	if !ok {
// 		format = "2006-01-02"
// 	}

// 	t, err = time.Parse(format, s)

// 	return t, err
// }

// func (m *CSVManager) parseString(v interface{}) (string, error) {
// 	return fmt.Sprintf("%v", v), nil
// }

// func (m *CSVManager) parseFloat(v interface{}) (float64, error) {
// 	s, err := m.parseString(v)

// 	if err != nil {
// 		return 0, err
// 	}

// 	return strconv.ParseFloat(s, 64)
// }

// func (m *CSVManager) parseRow(conf CSVFileConfiguration, row CSVRowData) (Record, error) {
// 	record := make(map[string]interface{})
// 	code := ""

// 	for _, col := range conf.CSVColumns {
// 		if col.Ordinal-1 > len(record) {
// 			return nil, fmt.Errorf("Column ordinal outside record index")
// 		}
// 		key := col.Name
// 		val, err := m.getVal(col, row)

// 		if err != nil {
// 			return nil, err
// 		}

// 		record[key] = val

// 		if col.IsKeyColumn {
// 			del := ""
// 			if code != "" {
// 				del = "_"
// 			}
// 			code = fmt.Sprintf("%s%s%s", code, del, val)
// 		}
// 	}
// 	record["code"] = code
// 	record["_meta"] = map[string]interface{}{
// 		"config_id": conf.ID,
// 	}

// 	return record, nil
// }

// func (m *CSVManager) ParseFile(fileName string) (chan Record, chan error) {
// 	// make channels for sending data / errors to consumer
// 	outCh := make(chan Record)
// 	errCh := make(chan error)
// 	var f *os.File

// 	go func() {
// 		ext := filepath.Ext(fileName)

// 		defer func() {
// 			if f != nil {
// 				f.Close()
// 			}
// 			close(outCh)
// 			close(errCh)
// 		}()

// 		if ext != ".csv" {
// 			errCh <- fmt.Errorf("file of type \"%s\" is not a valid CSV", ext)
// 			return
// 		}

// 		// Read file
// 		f, err := os.Open(fileName)

// 		if err != nil {
// 			errCh <- fmt.Errorf("Unable to read input file %v", err)
// 			return
// 		}

// 		r := csv.NewReader(f)
// 		codes := make(map[string]int)
// 		// for {
// 		// 	record, err := r.Read()
// 		// 	if err == io.EOF {
// 		// 		return
// 		// 	}
// 		// 	if err != nil {
// 		// 		errCh <- err
// 		// 		return
// 		// 	}
// 		// 	datum, err := m.parseRow(config, record)
// 		// 	if err != nil {
// 		// 		errCh <- err
// 		// 	} else {
// 		// 		code := fmt.Sprintf("%s", datum["code"])
// 		// 		val, ok := codes[code]
// 		// 		val++
// 		// 		if !ok {
// 		// 			codes[code] = val
// 		// 		}
// 		// 		code = fmt.Sprintf("%s:%d", code, val)
// 		// 		datum["code"] = code
// 		// 		outCh <- datum
// 		// 	}
// 		// }
// 	}()

// 	return outCh, errCh
// }

// func NewManager(db *gorm.DB) {
// 	// m = &Manager{db: db}

// 	// Load all csv file configurations
// 	// db.Preload("Columns").Find(&m.configs)

// 	// return m, nil
// }
