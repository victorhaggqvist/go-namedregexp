package namedregexp

import "regexp"

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
