package dqrs

import (
	"fmt"
	"strings"

	"github.com/miekg/dns"
)

// RRTypeToString converts a known Resource Record type to the appropriate
// internal Resource Record string value. An error is returned if the Resource
// Record type is not a valid key in the dns.TypeToString map.
func RRTypeToString(rrType uint16) (string, error) {

	// what checks can we run against rrType aside from validating whether it
	// exists as a key in the map?

	rrString, ok := dns.TypeToString[rrType]
	if !ok {
		return "", fmt.Errorf("provided key %v not in dns.TypeToString map", rrType)
	}

	return rrString, nil

}

// RRStringToType converts a known Resource Record string value to the
// appropriate internal Resource Record type. An error is returned if the
// string value is not a valid key in the dns.StringToType map.
func RRStringToType(rrString string) (uint16, error) {

	if rrString == "" {
		return 0, fmt.Errorf("empty rrString argument given")
	}

	rrType, ok := dns.StringToType[strings.ToUpper(rrString)]
	if !ok {
		return 0, fmt.Errorf("provided key %v not in dns.StringToType map", rrString)
	}

	return rrType, nil
}
