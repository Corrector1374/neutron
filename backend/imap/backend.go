// Stores messages on an IMAP server.
package imap

import (
	"strconv"

	"github.com/Corrector1374/neutron/backend"
	"github.com/Corrector1374/neutron/backend/util"
	"github.com/Corrector1374/neutron/backend/events"
)

type Config struct {
	Hostname string
	Port int
	Tls bool
	Suffix string
}

func (c *Config) Host() string {
	port := c.Port
	if port <= 0 {
		if c.Tls {
			port = 993
		} else {
			port = 143
		}
	}

	return c.Hostname + ":" + strconv.Itoa(port)
}

func Use(bkd *backend.Backend, config *Config) *conns {
	conns := newConns(config)
	messages := newMessages(conns)
	conversations := util.NewDummyConversations(messages)
	users := newUsers(conns)
	evts := newEvents(conns, bkd.EventsBackend, conversations)
	labels := events.NewLabels(newLabels(conns), evts)

	bkd.Set(messages, conversations, users, labels, evts)

	// TODO: do not return conns backend
	return conns
}
