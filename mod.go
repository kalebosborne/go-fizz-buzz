package mod


func Modulus(num,div float64)  float64{
	var a float64
	var b float64
	for num >= 0 {
		a = num - div
		if a - div >= 0 {
			num -= div
		} else if a == 0 {
			b = 0
			break
		} else {
			b = a
			break
		}
	}
		return b
	}