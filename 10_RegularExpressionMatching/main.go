package main

import (
	"fmt"
	"strings"
)

func Split(r rune) bool {
	return r == '*' || r == '.'
}

type Checker struct {
	PList     []string
	PIdx      int
	Word      string
	Character string
	Count     int
}

func (c *Checker) Next() {
	fmt.Printf("Next: %+v\n", c)

	if c.PIdx == 0 && c.Word == "" {
		c.Word = c.PList[0]
	}

	for {
		c.PIdx += 1
		switch p := c.PList[c.PIdx]; p {
		case "*":
			c.Character = p
			break

		case ".":
			c.Character = p
			break
		default:
			c.Word = p
			c.Count = 0
		}

	}
}

func (c *Checker) Check(s string) bool {
	fmt.Printf("Check: %v, %+v\n", s, c)

	if c.Word == s {
		if c.Count == 0 {
			c.Count += 1
			return true
		}
		if c.Count > 0 {
			if c.Character == "*" {
				c.Count += 1
				return true
			}
			if c.Character == "." {
				return false
			}
		}
	} else {
		c.Next()
	}
	return true
}

func isMatch(s string, p string) bool {
	sList := strings.Fields(s)
	fmt.Println(sList)

	checker := Checker{
		PList: strings.FieldsFunc(p, Split),
		PIdx:  0,
		Word:  "",
	}

	for idx, ss := range sList {
		fmt.Printf("idx: %v, ss: %v\n", idx, ss)
		b := checker.Check(ss)
		if !b {
			return false
		}

	}

	return true
}
