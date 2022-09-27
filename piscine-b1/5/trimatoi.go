package piscine

func TrimAtoi(s string) int {
	final := 0
	len := AlphaCount(s) - 1
	onoff := 0
	k := 0
	for _, v := range s {
		multi := puissance(10, len-k)
		if v == 45 && final == 0 {
			onoff = 1
		}
		if v >= 48 && v <= 57 {
			final += (int(v) - 48) * multi
			if final == 0 && v == 48 {
			} else {
				k++
			}
		}
	}
	if onoff == 1 {
		final *= -1
	}
	return final
}

// Fait la puissance
func puissance(n int, k int) (o int) {
	if k < 1 {
		return 1
	}
	k -= 1
	val0 := n
	for i := 0; i != k; i++ {
		n = n * val0
	}
	o = n
	return
}

func AlphaCount(s string) int {
	len := 0
	for _, v := range s {
		if len == 0 && int(v) == 48 {
		} else {
			if int(v) >= 48 && int(v) <= 57 {
				len += 1
			}
		}
	}
	return len
}
