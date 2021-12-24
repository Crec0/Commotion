package noise

type Xoroshiro128 struct {
	low uint64;
	high uint64;
}

func Xoroshiro(lo uint64, hi uint64) Xoroshiro128 {
	var xoroshiro Xoroshiro128
	
	xoroshiro = Xoroshiro128{lo, hi};
	
	if ((xoroshiro.low | xoroshiro.high) == 0) {
		xoroshiro = Xoroshiro128{
			low: 11400714819323198485, // -7046029254386353131L 
			high: 7640891576956012809,
		}
	}

	return xoroshiro;
}


func NextLong(xorshiro *Xoroshiro128) int64 {
	var l uint64;
	var h uint64;
    var res uint64;
	
	l = xorshiro.low;
	h = xorshiro.high;
	res = leftRotate(l + h, 17) + l;
	h ^= l;
	xorshiro.low = leftRotate(l, 49) ^ h ^ h << 21;
	xorshiro.high = leftRotate(h, 28);
	return int64(res);
}


func leftRotate(num uint64, distance uint64) uint64 {
	return (num << distance) | (num >> (64 - distance));
}
