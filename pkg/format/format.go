package format

import (
	"fmt"
	"github.com/tbauriedel/go-telegram-icinga2/internal"
)

// HostMessage returns a formatted output string for host notifications
func HostMessage() string {
	return fmt.Sprintf("%s is %s\n\nInfo:\t%s\nAddress:\t%s\nWhen:\t%s",
		internal.Hostname,
		internal.ObjectState,
		internal.ObjectOutput,
		internal.HostAddress,
		internal.DateTime,
	)
}

// ServiceMessage returns a formatted output string for service notifications
func ServiceMessage() string {
	return fmt.Sprintf("%s on %s is %s\n\nInfo:\t%s\nAddress:\t%s\nWhen:\t%s",
		internal.ServiceName,
		internal.Hostname,
		internal.ObjectState,
		internal.ObjectOutput,
		internal.HostAddress,
		internal.DateTime,
	)
}
