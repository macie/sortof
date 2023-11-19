//go:build openbsd && !linux

package main

import "golang.org/x/sys/unix"

// IsHardened reports whether security sandbox is enabled.
const IsHardened = true

// Sandbox restrict application access to necessary system calls needed by
// network connections and standard i/o.
//
// See: https://man.openbsd.org/pledge.2
func Sandbox() error {
	return unix.PledgePromises("stdio rpath")
}
