// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/dnsc
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package dqrs

// DNSRecord represents a record returned as part of a query response.
type DNSRecord struct {

	// Value represents the record value such as smtp1.example.com.
	Value string

	// Type is the record type, such as A, MX or CNAME.
	Type string

	// TTL is the lifetime of the record, such as 300.
	TTL uint32
}
