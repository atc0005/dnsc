// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package dqrs

// defaultDNSPort is the default UDP and TCP port used for incoming DNS
// requests
const defaultDNSPort = "53"

// Supported Request types
// TODO: Duplicated in config package
const (
	RequestTypeA       string = "A"
	RequestTypeAAAA    string = "AAAA"
	RequestTypeCNAME   string = "CNAME"
	RequestTypeMX      string = "MX"
	RequestTypeNS      string = "NS"
	RequestTypePTR     string = "PTR"
	RequestTypeSRV     string = "SRV"
	RequestTypeUnknown string = "UNKNOWN"
)

const recordValueUnknown string = "type unknown"

// Results summary output display options
// TODO: Duplicated in config package
const (
	ResultsOutputSingleLine string = "single-line"
	ResultsOutputMultiLine  string = "multi-line"
)
