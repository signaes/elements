package attributes

type Attributes map[string]string

func (a Attributes) String() string {
	if len(a) == 0 {
		return ""
	}

	result := ""

	for k, v := range a {
		if len(v) > 0 {
			result += " " + k + "=" + "\"" + v + "\""
		} else {
			result += " " + k
		}
	}

	return result
}

func Merge(l []map[string]string) Attributes {
	m := map[string]string{}

	for _, attrs := range l {
		for k, v := range attrs {
			m[k] = v
		}
	}

	return m
}
