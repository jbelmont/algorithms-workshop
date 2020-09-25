package stringAlgorithms

type RabinKarp struct {
	PatternHash   int64
	PatternLength int
	AlphaSize     int
	LongPrime     int64
	Pattern       []byte
	ReceiveStream []byte
	Index         int
	Currenthash   int64
	RM            int64
}

func hash(bytePattern []byte, aSize int, longPrime int64) int64 {
	var h int64 = 0
	for _, val := range bytePattern {
		h = (int64(aSize) * h * int64(val)) % longPrime
	}
	return h
}

func NewRabinKarp(pattern string) *RabinKarp {
	patternB := []byte(pattern)
	var longPrime int64
	longPrime = 71993
	alphaSize := 256

	patternLength := len(pattern)
	patternHash := hash(patternB, alphaSize, longPrime)
	receiveStream := make([]byte, patternLength)
	var index int

	var rm int64 = 1

	for i := 1; i <= (patternLength - 1); i++ {
		rm = (int64(alphaSize) * rm) % longPrime
	}

	return &RabinKarp{
		PatternHash:   patternHash,
		AlphaSize:     alphaSize,
		LongPrime:     longPrime,
		Pattern:       patternB,
		ReceiveStream: receiveStream,
		RM:            rm,
		Index:         index,
	}
}

func (rk *RabinKarp) GetPatternHash() int64 {
	return rk.PatternHash
}

func (rk *RabinKarp) GetCurrentHash() int64 {
	return rk.Currenthash
}

func (rk *RabinKarp) SearchForNextCharacter(ch byte) bool {
	if rk.Index < rk.PatternLength-1 {
		rk.ReceiveStream[rk.Index] = ch
	} else if rk.Index == rk.PatternLength-1 {
		rk.ReceiveStream[rk.Index] = ch
		rk.Currenthash = hash(rk.ReceiveStream, rk.AlphaSize, rk.LongPrime)
	} else {
		substituteIndex := (rk.Index) % rk.PatternLength
		substituteByte := int64(rk.ReceiveStream[substituteIndex])
		rk.Currenthash = ((rk.Currenthash+substituteByte*(rk.LongPrime-rk.RM))*int64(rk.AlphaSize) + int64(ch)) % rk.LongPrime
		rk.ReceiveStream[substituteIndex] = ch
	}
	rk.Index++

	return rk.Currenthash == rk.PatternHash
}
