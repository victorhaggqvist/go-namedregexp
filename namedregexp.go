package namedregexp

import (
	"fmt"
	"regexp"
)

func FindNamedStringSubmatch(pattern, haystack string) (map[string]string, error) {
	result := map[string]string{}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return result, err
	}

	match := re.FindStringSubmatch(haystack)
	for i, name := range re.SubexpNames() {
		// result[name] = match[i]
		if i != 0 {
			result[name] = match[i]
		}
	}

	return result, nil
}

func FindAllNamedStringSubmatch(pattern, haystack string, n int) ([]map[string]string, error) {
	result := []map[string]string{}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return result, err
	}

	match := re.FindAllStringSubmatch(haystack, n)

	renames := re.SubexpNames()
	for _, matchset := range match {
		resultset := map[string]string{}

		if len(renames) != len(matchset) {
			return result, fmt.Errorf("result size missmatch, number of names and matches are not equal, name=%d matches=%d", len(renames), len(matchset))
		}

		for i, name := range renames {
			if name == "" {
				continue
			}

			resultset[name] = matchset[i]
		}
		result = append(result, resultset)
	}

	return result, nil
}
