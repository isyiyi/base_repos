package knowledge

type Num int

func (n Num) Add(n2 Num) Num {
	return n + n2
}

func (n Num) Sub(n2 Num) Num {
	return n - n2
}

func (n Num) Mod(n2 Num) Num {
	return n % n2
}

func (n Num) Pro(n2 Num) Num {
	return n * n2
}
