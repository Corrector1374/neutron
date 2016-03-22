package imap

import (
	"errors"

	"github.com/mxk/go-imap/imap"
	"github.com/emersion/neutron/backend"
	"github.com/emersion/neutron/backend/memory"
)

func (b *Backend) GetUser(id string) (user *backend.User, err error) {
	user, ok := b.users[id]
	if !ok {
		err = errors.New("No such user")
	}
	return
}

func (b *Backend) Auth(username, password string) (session *backend.Session, err error) {
	c, err := imap.DialTLS(b.config.Host(), nil)
	if err != nil {
		return
	}

	email := username + b.config.Suffix
	_, err = c.Login(email, password)
	if err != nil {
		return
	}
	c.Data = nil

	user := &backend.User{
		ID: username,
		Name: username,
		DisplayName: username,
		Addresses: []*backend.Address{
			&backend.Address{
				ID: username,
				Email: email,
				Send: 1,
				Receive: 1,
				Status: 1,
				Type: 1,
				Keys: []*backend.Keypair{
					&backend.Keypair{
						ID: username,
						PublicKey: memory.DefaultPublicKey(),
						PrivateKey: memory.DefaultPrivateKey(),
					},
				},
			},
		},
	}

	session, err = b.InsertSession(&backend.Session{User: user})
	if err != nil {
		return
	}

	b.users[user.ID] = user
	b.passwords[user.ID] = password
	b.insertConn(user.ID, c)

	return
}

func (b *Backend) IsUsernameAvailable(username string) (bool, error) {
	return false, errors.New("Cannot check if a username is available with IMAP backend")
}

func (b *Backend) InsertUser(u *backend.User, password string) (*backend.User, error) {
	return nil, errors.New("Cannot register new user with IMAP backend")
}

func (b *Backend) UpdateUser(update *backend.UserUpdate) error {
	return errors.New("Cannot update user with IMAP backend")
}

func (b *Backend) UpdateUserPassword(id, current, new string) error {
	return errors.New("Cannot update user password with IMAP backend")
}

func (b *Backend) UpdateKeypair(id, password string, keypair *backend.Keypair) error {
	return errors.New("Not yet implemented")
}

func (b *Backend) GetPublicKey(email string) (string, error) {
	// TODO
	return "", nil
}

// Allow other backends (e.g. a SMTP backend) to access users' password.
func (b *Backend) GetPassword(user string) (string, error) {
	if password, ok := b.passwords[user]; ok {
		return password, nil
	}
	return "", errors.New("No password stored for such user")
}
