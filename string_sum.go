package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	errorInvalidInput   = errors.New("invalid input")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf
const (
	errFormat = "err: %w"
)

func StringSum(input string) (output string, err error) {

	in := []rune(strings.ReplaceAll(input, " ", ""))
	if len(in) == 0 {
		return "", fmt.Errorf(errFormat, errorEmptyInput)

	}
	var listNumbers []int
	var num []rune

	isDigitalBefore := false
	isDigitalNow := false

	for i := 0; i < len(in); i++ {
		first := false
		AddToLists := false
		x, _ := regexp.MatchString("[0-9+-]", string(in[i]))

		if !x {
			return "", fmt.Errorf(errFormat, errorInvalidInput)
		}

		if len(num) == 0 {
			first = true
		}

		num = append(num, in[i])

		if unicode.IsDigit(in[i]) {
			isDigitalNow = true
		} else {
			isDigitalNow = false
		}

		if !first && isDigitalNow == false && isDigitalBefore || i == (len(in)-1) && unicode.IsDigit(in[i]) {
			AddToLists = true
		}

		if AddToLists {
			var number int
			var err error

			if i == len(in)-1 {
				number, err = strconv.Atoi(string(num))
			} else {
				number, err = strconv.Atoi(string(num[0 : len(num)-1]))
			}
			if err != nil {
				return "", fmt.Errorf(errFormat, err)
			}

			listNumbers = append(listNumbers, number)
			num = nil
			num = append(num, in[i])
		}

		if unicode.IsDigit(in[i]) {
			isDigitalBefore = true
		} else {
			isDigitalBefore = false
		}
	}
	if len(listNumbers) != 2 {
		return "", fmt.Errorf(errFormat, errorNotTwoOperands)
	}

	return strconv.Itoa(listNumbers[0] + listNumbers[1]), nil
}
