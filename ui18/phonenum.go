package ui18

import (
	"strconv"

	"github.com/nyaruka/phonenumbers"
)

// ParsePhoneNum 解析国际手机号码：861234567890
func ParsePhoneNum(phone string) (countryCode, mobile string, err error) {
	parseSend, err := phonenumbers.Parse("+"+phone, "")
	if err != nil {
		return "", "", err
	}
	countryCode = strconv.Itoa(int(parseSend.GetCountryCode()))
	mobile = strconv.FormatUint(parseSend.GetNationalNumber(), 10)
	return countryCode, mobile, nil
}
