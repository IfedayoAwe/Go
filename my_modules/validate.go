package mymodules

func main() {

}

// Vaidate phone number
// type StagStr struct {
// 	Value string
// }

// func ValidateNumber(number string) (string, error) {
// 	if number == "" {
// 		return "", fmt.Errorf("empty number")
// 	}
// 	stringData := StagStr{Value: number}
// 	includeLandLine := true
// 	if stringData.contains("+") || stringData.contains("(") || len(number) > 11 {
// 		return number, nil
// 	} else if string(number[0]) == "0" {
// 		return strings.Replace(number, "0", "234", 1), nil
// 	} else if country := phonenumber.GetISO3166ByNumber("09032092434", includeLandLine); country.CountryName != "" {
// 		return number, nil
// 	} else {
// 		return number, fmt.Errorf("wrong format")
// 	}
// }

// func (str StagStr) contains(val string) bool {
// 	for _, v := range str.Value {
// 		if string(v) == val {
// 			return true
// 		}
// 	}
// 	return false
// }
