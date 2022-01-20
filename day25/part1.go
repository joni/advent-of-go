package day25

const base = 20201227

func decrypt(publicKey int) int {
	subject := 7
	value := 1
	for privateKey := 1; privateKey < base; privateKey++ {
		value = value * subject % base
		if value == publicKey {
			return privateKey
		}
	}
	return -1
}

func encrypt(key, exponent int) int {
	ans := 1
	for i := 0; i < exponent; i++ {
		ans = ans * key % base
	}
	return ans
}

func solve1(cardPublic, doorPublic int) int {
	cardPrivate := decrypt(cardPublic)
	return encrypt(doorPublic, cardPrivate)
}
