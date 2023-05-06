package rnet

import "net"

// IsUniversalMACAddr returns true if the given MAC address is a universal address.
// this traditionally means the "burnt in address" of a network interface, but isn't always the case.
func IsUniversalMACAddr(addr net.HardwareAddr) bool {
	return !testBit(addr[0], 6)
}

// IsMulticastMACAddr returns true if the given MAC address is a multicast address.
// this indicates that the MAC is part of a group, and may not be the only address meant to receive the message.
func IsMulticastMACAddr(addr net.HardwareAddr) bool {
	return testBit(addr[0], 7)
}

func testBit(b byte, pos uint) bool {
	val := b & (1 << pos)
	return val > 0
}
