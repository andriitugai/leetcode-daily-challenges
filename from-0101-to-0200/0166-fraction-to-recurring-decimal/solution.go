func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {
        return "0"
    }
    var result bytes.Buffer
    if numerator < 0 && denominator > 0 || numerator > 0 && denominator < 0 {
        result.WriteString("-")
    }
    if numerator < 0 {
        numerator = -numerator
    }
    if denominator < 0 {
        denominator = -denominator
    }
    result.WriteString(fmt.Sprintf("%d", numerator/denominator))
    numerator %= denominator
    if numerator == 0 {
        return result.String()
    }
    result.WriteString(".")
    remainder := make(map[int]int)
	remainder[numerator] = result.Len()
	for numerator != 0 {
		numerator *= 10
		result.WriteString(fmt.Sprintf("%d", numerator/denominator))
		numerator %= denominator
		if index, ok := remainder[numerator]; ok {
			res := result.String()
			return res[:index] + "(" + res[index:] + ")"
		}
		remainder[numerator] = result.Len()
	}
	return result.String()
}