package database

func DetectUniqueRowColName(str string) string {
	index := 0

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '.' {
			index = i + 1
			break
		}
	}

	return str[index:]
}
