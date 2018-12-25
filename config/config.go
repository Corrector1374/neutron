// Parses the configuration file.
package config

import (
	"github.com/Corrector1374/neutron/backend/imap"
	"github.com/Corrector1374/neutron/backend/smtp"
	"github.com/Corrector1374/neutron/backend/disk"
)

// Configuration for all backends.
// Backends omitted or set to null won't be activated.
type Config struct {
	// Memory config.
	Memory *MemoryConfig

	// IMAP config.
	Imap *ImapConfig

	// SMTP config.
	Smtp *SmtpConfig

	// Disk config.
	Disk *DiskConfig
}

type BackendConfig struct {
	Enabled bool
}

type MemoryConfig struct {
	*BackendConfig
	Populate bool
	Domains []string
}

type ImapConfig struct {
	*BackendConfig
	*imap.Config
}

type SmtpConfig struct {
	*BackendConfig
	*smtp.Config
}

type DiskConfig struct {
	*BackendConfig
	*disk.Config

	Contacts *DiskConfig
	Keys *DiskConfig
	UsersSettings *DiskConfig
	Addresses *DiskConfig
}
