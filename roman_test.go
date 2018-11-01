package roman

import "testing"

func TestConvertRomanNumeralTo1(t *testing.T) {
	r := NewRoman()

	n := r.ToNumber("I")

	if n != 1 {
		t.Errorf("I should return 1, but %v", n)
	}
}

func TestConvertRomanNumeralTo2(t *testing.T) {
	r := NewRoman()

	n := r.ToNumber("II")

	if n != 2 {
		t.Errorf("II should return 2, but %v", n)
	}
}

func TestConvertRomanNumeralTo3(t *testing.T) {
	r := NewRoman()

	n := r.ToNumber("III")

	if n != 3 {
		t.Errorf("III should return 3, but %v", n)
	}
}

func TestConvertRomanNumeralTo4(t *testing.T) {
	r := NewRoman()

	n := r.ToNumber("IV")

	if n != 4 {
		t.Errorf("IV should return 4, but %v", n)
	}
}

func TestConvertRomanNumeralTo1996(t *testing.T) {
	r := NewRoman()

	n := r.ToNumber("MCMXCVI")

	if n != 1996 {
		t.Errorf("MCMXCVI should return 1996, but %v", n)
	}
}

func TestConvertDecimalToI(t *testing.T) {
	r := NewRoman()

	n := r.ToRoman(1)

	if n != "I" {
		t.Errorf("1 should return I, but %v", n)
	}
}

func TestFindingHighestDecimalFor1996(t *testing.T) {
	n := highestDecimal(1996)

	if n != 1000 {
		t.Errorf("1996 should return 1000, but %v", n)
	}
}

func TestFindingHighestDecimalFor996(t *testing.T) {
	n := highestDecimal(996)

	if n != 900 {
		t.Errorf("996 should return 900, but %v", n)
	}
}

func TestConvertDecimalToII(t *testing.T) {
	r := NewRoman()

	n := r.ToRoman(2)

	if n != "II" {
		t.Errorf("2 should return II, but %v", n)
	}
}

func TestConvertDecimalToIII(t *testing.T) {
	r := NewRoman()

	n := r.ToRoman(3)

	if n != "III" {
		t.Errorf("3 should return III, but %v", n)
	}
}

func TestConvertDecimalToIV(t *testing.T) {
	r := NewRoman()

	n := r.ToRoman(4)

	if n != "IV" {
		t.Errorf("4 should return IV, but %v", n)
	}
}

func TestConvertDecimalToV(t *testing.T) {
	r := NewRoman()

	n := r.ToRoman(5)

	if n != "V" {
		t.Errorf("5 should return V, but %v", n)
	}
}

func TestConvertDecimalToIX(t *testing.T) {
	r := NewRoman()

	n := r.ToRoman(9)

	if n != "IX" {
		t.Errorf("9 should return IX, but %v", n)
	}
}

func TestConvertDecimalToXI(t *testing.T) {
	r := NewRoman()

	n := r.ToRoman(11)

	if n != "XI" {
		t.Errorf("11 should return XI, but %v", n)
	}
}

func TestConvertDecimalToMCMXCVI(t *testing.T) {
	r := NewRoman()

	n := r.ToRoman(1996)

	if n != "MCMXCVI" {
		t.Errorf("1996 should return MCMXCVI, but %v", n)
	}
}
