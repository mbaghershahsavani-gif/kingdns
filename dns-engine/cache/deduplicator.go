package cache

func UniqueAnswers(records []string) []string {
	result := []string{}

	seen := map[string]bool{}

	for _, record := range records {
		if !seen[record] {
			result = append(result, record)
			seen[record] = true
		}
	}

	return result
}
