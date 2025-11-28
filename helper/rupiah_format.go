package helper

import (
	"strconv"
	"strings"
)

func FormatRupiah(amount float64) string {
	s := strconv.FormatFloat(amount, 'f', 2, 64)

	parts := strings.Split(s, ".")
	intPart := parts[0]
	decPart := parts[1]

	var result []string
	for len(intPart) > 3 {
		result = append([]string{intPart[len(intPart)-3:]}, result...)
		intPart = intPart[:len(intPart)-3]
	}
	result = append([]string{intPart}, result...)

	return "Rp " + strings.Join(result, ".") + "," + decPart
}
