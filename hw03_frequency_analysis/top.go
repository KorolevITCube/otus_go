package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`((\s+-\s+)|[\"?!;:.,\s])`)

type Word struct {
	text  string
	count int
}

func Top10(str string) []string {
	lowerStr := strings.ToLower(str)
	words := re.Split(lowerStr, -1)
	counter := make(map[string]int)
	for _, s := range words {
		if v, ok := counter[s]; ok {
			counter[s] = v + 1
		} else {
			counter[s] = 1
		}
	}

	sortSlice := make([]Word, 0, len(counter))
	for k, v := range counter {
		sortSlice = append(sortSlice, Word{k, v})
	}

	sort.Slice(sortSlice, func(i, j int) bool {
		if sortSlice[i].count != sortSlice[j].count {
			return sortSlice[i].count > sortSlice[j].count
		}
		return sortSlice[i].text < sortSlice[j].text
	})

	var res []string
	for i := 0; i < 11 && i < len(sortSlice); i++ {
		if sortSlice[i].text != "" {
			res = append(res, sortSlice[i].text)
		}
	}

	return res
}
