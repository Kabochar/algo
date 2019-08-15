package CartoonAlgo

// 求一个两数的最大公约数

// 辗转相除+更相减损
// Time: O(log(max(a,b)))

func gcd(ap, bp int) int {
	if ap == bp {
		return ap
	}

	var big, small int
	// // a 偶，b 偶，gcd(a, b) = gcd(a>>1, b)/2 = gcd(a>>1, b)<<1
	if (ap&1) == 0 && (bp&1) == 0 {
		return gcd(ap>>1, bp>>1) << 1

		// a 偶，b 奇，gcd(a, b) = gcd(a>>1, b)
	} else if (ap&1) == 0 && (bp&1) != 0 {
		return gcd(ap>>1, bp)

		// a 奇，b 偶，gcd(a, b) = gcd(a, b>>1)
	} else if (ap&1) != 0 && (bp&1) == 0 {
		return gcd(ap, bp>>1)

		// a 奇，b 奇，更相减损术运算一次，再 gcd(a, b) = gcd(b, a-b)
	} else {
		if ap > bp {
			big, small = ap, bp
		} else {
			big, small = bp, ap
		}

		return gcd(big-small, small)
	}
}
