package phone

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
	"regexp"
	"strconv"
)

const (
	// DefaultCountryCodeNumberID is the default country code number used for parsing using ID.
	DefaultCountryCodeNumberID = 62
)

var (
	regexPlusLeadingNumber = regexp.MustCompile(`^\+*`)
	supportedCountryCodes  = phonenumbers.GetSupportedCallingCodes()
)

// getRegexZeroLeadingNumber will panic if the regexp.MustCompile cannot compile the regex.
func getRegexZeroLeadingNumber(countryCode string) *regexp.Regexp {
	return regexp.MustCompile(`^\+0+(` + countryCode + `)?`)
}

// NormalizeID parses the phone number using the countryCode.
// It returns the normalized phone number and the country code.
// The default country code is ID.
// WARNING: This function can panic if the regex is invalid.
func NormalizeID(phoneNumber string, countryCode int) (res string) {
	res = phoneNumber
	if countryCode == 0 || !supportedCountryCodes[countryCode] {
		countryCode = DefaultCountryCodeNumberID
	}
	countryCodeStr := strconv.Itoa(countryCode)

	phoneNumber = regexPlusLeadingNumber.ReplaceAllString(phoneNumber, "+")
	phoneNumber = getRegexZeroLeadingNumber(countryCodeStr).
		ReplaceAllString(phoneNumber, "+"+countryCodeStr)
	pn, err := phonenumbers.Parse(phoneNumber, "")
	if err != nil {
		return
	}

	res = fmt.Sprintf("%d%d", pn.GetCountryCode(), pn.GetNationalNumber())
	return
}
