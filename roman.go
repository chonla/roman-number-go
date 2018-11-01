package roman

var num = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var numInv = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var maxTable = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

// Roman is roman
type Roman struct{}

// NewRoman to create a Roman
func NewRoman() *Roman {
	return &Roman{}
}

// ToNumber to covert roman numeral to decimal
func (r *Roman) ToNumber(n string) int {
	out := 0
	ln := len(n)
	for i := 0; i < ln; i++ {
		c := string(n[i])
		vc := num[c]
		if i < ln-1 {
			cnext := string(n[i+1])
			vcnext := num[cnext]
			if vc < vcnext {
				out += vcnext - vc
				i++
			} else {
				out += vc
			}
		} else {
			out += vc
		}
	}
	return out
}

// ToRoman is to convert decimal number to roman numeral
func (r *Roman) ToRoman(n int) string {
	out := ""
	for n > 0 {
		v := highestDecimal(n)
		out += numInv[v]
		n -= v
	}
	return out
}

func highestDecimal(n int) int {
	for _, v := range maxTable {
		if v <= n {
			return v
		}
	}
	return 1
}
