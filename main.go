package main

import (
	"fmt"

	"github.com/miekg/dns"
)

// Possible usage (not sure of flag names yet):
//
// dnsc -servers DNS1,DNS2,DNS3 -query server1.example.com -answer 192.168.1.250
// dnsc -servers DNS1,DNS2,DNS3 -q server1.example.com -e 192.168.1.250
// dnsc -servers DNS1,DNS2,DNS3 -ask server1.example.com -answer 192.168.1.250
// dnsc -servers DNS1,DNS2,DNS3 -ask server1.example.com -expect 192.168.1.250
//
// Thus far, I like these long/short options:
//
// dnsc -servers DNS1,DNS2,DNS3 -q server1.example.com -e 192.168.1.250
// dnsc -servers DNS1,DNS2,DNS3 -query server1.example.com -expect 192.168.1.250
//
// The "query" word is hard to avoid when it comes to DNS-related "things"
//
// Result: Should print small table of details with a final pass/fail verdict

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
