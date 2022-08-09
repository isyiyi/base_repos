package reflect

import "strconv"

func Sprintf(s interface{}) string {
	switch s := s.(type) {
	case string:
		return s
	case int:
		return strconv.Itoa(s)
	default:
		return "???"
	}
}