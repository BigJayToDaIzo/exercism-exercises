package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	candidate := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, phoneNumber)
	canLen := len(candidate)

	if canLen < 10 || canLen > 11 {
		return "", errors.New("invalid length")
	}
	if canLen == 11 && candidate[0] != '1' {
		return "", errors.New("invalid country code")
	}
	if canLen == 11 {
		if !checkAreaExch(string(candidate[1])) {
			return "", errors.New("invalid area code")
		}
		if !checkAreaExch(string(candidate[4])) {
			return "", errors.New("invalid exchange code")
		}
		return candidate[1:], nil
	}
	if !checkAreaExch(string(candidate[0])) {
		return "", errors.New("invalid area code")
	}
	if !checkAreaExch(string(candidate[3])) {
		return "", errors.New("invalid exchange code")
	}
	return candidate, nil
}

func AreaCode(phoneNumber string) (string, error) {
	parsedNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return parsedNumber[:3], nil
}

func Format(phoneNumber string) (string, error) {
	parsedNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", parsedNumber[:3], parsedNumber[3:6], parsedNumber[6:]), nil
}

func checkAreaExch(s string) bool {
	return s != "0" && s != "1"
}
