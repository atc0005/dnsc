package main

import (
	"fmt"

	"github.com/miekg/dns"
)

// Intended usage:
//
// dnsc -servers DNS1,DNS2,DNS3 -query server1.example.com -answer 192.168.1.250
//
// Should back a small table of results with a final pass/fail verdict

func main() {
	var msg dns.Msg
	fqdn := dns.Fqdn("stacktitan.com")
	msg.SetQuestion(fqdn, dns.TypeA)
	in, err := dns.Exchange(&msg, "8.8.8.8:53")
	if err != nil {
		panic(err)
	}
	if len(in.Answer) < 1 {
		fmt.Println("No records")
		return
	}
	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			fmt.Println(a.A)
		}
	}
}
