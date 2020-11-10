package transformers

type 

type CSVColumnMapping struct {
	Name string
	Ordinal int
	DataType string
}

type DiscoverCsvRow struct {

}

type DiscoverRowMapper struct {
	
}


ordinalMap := make(map[string]int)

	return func(row []string) (data *models.Activity, err error) {
		s := strings.Split(row[0], "/")
		s1 := fmt.Sprintf("%s-%s-%s", s[2], s[0], s[1])
		td, err := time.Parse("2006-01-02", s1)

		if err != nil {
			return data, err
		}

		s = strings.Split(row[1], "/")
		s1 = fmt.Sprintf("%s-%s-%s", s[2], s[0], s[1])
		pd, err := time.Parse("2006-01-02", s1)

		if err != nil {
			return data, err
		}

		amount, err := strconv.ParseFloat(row[3], 32)

		if err != nil {
			return data, err
		}

		key := strings.Join(row, "-")
		fmt.Println(key)

		_, ok := ordinalMap[key]

		if !ok {
			ordinalMap[key] = 0
		}

		ordinal := ordinalMap[key]
		ordinalMap[key]++

		code := fmt.Sprintf("discover:%s-%d", key, ordinal)

		data = &models.Activity{
			TransactionDate: td,
			PostedDate:      pd,
			Description:     row[2],
			Amount:          float32(amount),
			Category:        row[4],
			Code:            code,
			SourceID:        1,
		}

		return data, err
	}