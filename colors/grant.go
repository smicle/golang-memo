package colors

import "github.com/logrusorgru/aurora"

// blackを付与
func black(v aurora.Value, bright, bg bool) aurora.Value {
	if bright && bg {
		return aurora.BgBrightBlack(v)
	} else if bright {
		return aurora.BrightBlack(v)
	} else if bg {
		return aurora.BgBlack(v)
	} else {
		return aurora.Black(v)
	}
}

// redを付与
func red(v aurora.Value, bright, bg bool) aurora.Value {
	if bright && bg {
		return aurora.BgBrightRed(v)
	} else if bright {
		return aurora.BrightRed(v)
	} else if bg {
		return aurora.BgRed(v)
	} else {
		return aurora.Red(v)
	}
}

// greenを付与
func green(v aurora.Value, bright, bg bool) aurora.Value {
	if bright && bg {
		return aurora.BgBrightGreen(v)
	} else if bright {
		return aurora.BrightGreen(v)
	} else if bg {
		return aurora.BgGreen(v)
	} else {
		return aurora.Green(v)
	}
}

// yellowの付与
func yellow(v aurora.Value, bright, bg bool) aurora.Value {
	if bright && bg {
		return aurora.BgBrightYellow(v)
	} else if bright {
		return aurora.BrightYellow(v)
	} else if bg {
		return aurora.BgYellow(v)
	} else {
		return aurora.Yellow(v)
	}
}

// blueの付与
func blue(v aurora.Value, bright, bg bool) aurora.Value {
	if bright && bg {
		return aurora.BgBrightBlue(v)
	} else if bright {
		return aurora.BrightBlue(v)
	} else if bg {
		return aurora.BgBlue(v)
	} else {
		return aurora.Blue(v)
	}
}

// magentaの付与
func magenta(v aurora.Value, bright, bg bool) aurora.Value {
	if bright && bg {
		return aurora.BgBrightMagenta(v)
	} else if bright {
		return aurora.BrightMagenta(v)
	} else if bg {
		return aurora.BgMagenta(v)
	} else {
		return aurora.Magenta(v)
	}
}

// cyanの付与
func cyan(v aurora.Value, bright, bg bool) aurora.Value {
	if bright && bg {
		return aurora.BgBrightCyan(v)
	} else if bright {
		return aurora.BrightCyan(v)
	} else if bg {
		return aurora.BgCyan(v)
	} else {
		return aurora.Cyan(v)
	}
}

// whiteの付与
func white(v aurora.Value, bright, bg bool) aurora.Value {
	if bright && bg {
		return aurora.BgBrightWhite(v)
	} else if bright {
		return aurora.BrightWhite(v)
	} else if bg {
		return aurora.BgWhite(v)
	} else {
		return aurora.White(v)
	}
}
