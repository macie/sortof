//go:build linux && !openbsd

package main

import (
	seccomp "github.com/seccomp/libseccomp-golang"
)

// IsHardened reports whether security sandbox is enabled.
const IsHardened = true

// Sandbox restrict application access to necessary system calls needed by
// network connections and standard i/o.
func Sandbox() error {
	// How to create minimal whitelist:
	// 1. Create empty list of allowed syscalls
	// 2. Set `seccomp.ActLog` as default filter action
	// 3. Compile and run program
	// 4. Use `dmesg` to find logged syscalls (started with _audit_)
	// 5. Translate syscalls numbers to names and add them to allowed list
	// 6. Go to point 3 and repeat until no new audit logs
	// 7. Reset default filter action to `seccomp.ActKillProcess`
	allowedSyscalls := []string{
		// similar to stdio pledge
		"clone3", "epoll_create1", "epoll_ctl", "epoll_pwait", "exit_group",
		"fcntl", "futex", "getpid", "gettid", "mmap", "mprotect", "munmap",
		"nanosleep", "pipe2", "read", "rseq", "rt_sigprocmask", "rt_sigreturn",
		"sched_yield", "set_robust_list", "sigaltstack", "tgkill", "write",

		// similar to rpath pledge
		"openat",
	}

	// By default goroutines don't play well with seccomp. Program will hang
	// when underlying thread is terminated silently. We need to kill process -
	// see: https://github.com/golang/go/issues/3405#issuecomment-750816828
	whitelist, err := seccomp.NewFilter(seccomp.ActKillProcess)
	if err != nil {
		return err
	}

	for _, callName := range allowedSyscalls {
		callId, err := seccomp.GetSyscallFromName(callName)
		if err != nil {
			return err
		}

		whitelist.AddRule(callId, seccomp.ActAllow)
	}
	whitelist.Load()

	return nil
}
