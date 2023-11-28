package coveralls

func Even(in int64) bool {
	return in%2 == 0
}

func Odd(in int64) bool {
	return !Even(in)
}

func IsTwo(in int64) bool {
	return in == 2
}
