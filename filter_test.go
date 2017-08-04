package addrutil

import (
	"testing"

	ma "github.com/multiformats/go-multiaddr"
)

func TestIsFDCostly(t *testing.T) {
	good := []ma.Multiaddr{
		newMultiaddr(t, "/ip4/127.0.0.1/tcp/1234"),
		newMultiaddr(t, "/ip4/0.0.0.0/tcp/1234"),
		newMultiaddr(t, "/ip6/::1/tcp/1234"),
		newMultiaddr(t, "/ip6/::/tcp/1234"),
		newMultiaddr(t, "/ip6/fe80::1/tcp/1234"),
		newMultiaddr(t, "/ip6/fe80::/tcp/1234"),
		newMultiaddr(t, "/ip6/fe80::/tcp/1234/ws"),
		newMultiaddr(t, "/ip4/127.0.0.1/tcp/1234/ws"),
	}

	bad := []ma.Multiaddr{
		newMultiaddr(t, "/ip4/127.0.0.1/udp/1234"),
		newMultiaddr(t, "/ip4/0.0.0.0/udp/1234/utp"),
		newMultiaddr(t, "/ip6/::1/udp/1234"),
		newMultiaddr(t, "/ip6/::/udp/1234"),
	}

	for _, a := range bad {
		if IsFDCostlyTransport(a) {
			t.Errorf("addr %s should not be fd costly", a)
		}
	}

	for _, a := range good {
		if !IsFDCostlyTransport(a) {
			t.Errorf("addr %s should be fd costly", a)
		}
	}
}

func TestIsFdCostlyMalformed(t *testing.T) {
	bad := []ma.Multiaddr{
		newMultiaddr(t, "/ip4/127.0.0.1/"),
		newMultiaddr(t, "/ip4/0.0.0.0/"),
		newMultiaddr(t, "/ip6/::1/"),
		newMultiaddr(t, "/ip6/::/"),
	}

	for _, a := range bad {
		if IsFDCostlyTransport(a) {
			t.Errorf("addr %s should not be fd costly", a)
		}
	}
}
