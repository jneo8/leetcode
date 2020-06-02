package main

// RomanNumM .
type RomanNumM struct {
	OneThousand  int
	FiveHundreds int
	OneHundred   int
	Fivty        int
	Ten          int
	One          int
}

func intToRoman(num int) string {

	s := ""
	n1000 := num / 1000
	for i := 0; i < n1000; i++ {
		s = s + "M"
	}
	num = num % 1000

	if num >= 900 {
		s = s + "CM"
		num -= 900
	}

	n500 := num / 500
	for i := 0; i < n500; i++ {
		s = s + "D"
	}
	num = num % 500

	if num >= 400 {
		s = s + "CD"
		num -= 400
	}

	n100 := num / 100
	for i := 0; i < n100; i++ {
		s = s + "C"
	}
	num = num % 100

	if num >= 90 {
		s = s + "XC"
		num -= 90
	}

	n50 := num / 50
	for i := 0; i < n50; i++ {
		s = s + "L"
	}
	num = num % 50

	if num >= 40 {
		s = s + "XL"
		num -= 40
	}

	n10 := num / 10
	for i := 0; i < n10; i++ {
		s = s + "X"
	}
	num = num % 10

	if num >= 9 {
		s = s + "IX"
		num -= 9
	}

	n5 := num / 5
	for i := 0; i < n5; i++ {
		s = s + "V"
	}
	num = num % 5

	if num >= 4 {
		s = s + "IV"
		num -= 4
	}

	for i := 0; i < num; i++ {
		s = s + "I"
	}
	return s
}
