package piscine

func Abort(a, b, c, d, e int) int {
	var aPosition int = 5
	var bPosition int = 5
	var cPosition int = 5
	var dPosition int = 5
	var ePosition int = 5
	if a > b {
		aPosition--
	}
	if a > c {
		aPosition--
	}
	if a > d {
		aPosition--
	}
	if a > e {
		aPosition--
	}
	if b > a {
		bPosition--
	}
	if b > c {
		bPosition--
	}
	if b > d {
		bPosition--
	}
	if b > e {
		bPosition--
	}
	if c > a {
		cPosition--
	}
	if c > b {
		cPosition--
	}
	if c > d {
		cPosition--
	}
	if c > e {
		cPosition--
	}
	if d > a {
		dPosition--
	}
	if d > b {
		dPosition--
	}
	if d > c {
		dPosition--
	}
	if d > e {
		dPosition--
	}
	if e > a {
		ePosition--
	}
	if e > b {
		ePosition--
	}
	if e > c {
		ePosition--
	}
	if e > d {
		ePosition--
	}
	if aPosition == 3 {
		return a
	} else {
		if bPosition == 3 {
			return b
		} else {
			if cPosition == 3 {
				return c
			} else {
				if dPosition == 3 {
					return d
				} else {
					if ePosition == 3 {
						return e
					}
				}
			}
		}
	}
	return a
}
