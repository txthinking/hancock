// Copyright (c) 2020-present Cloud <cloud@txthinking.com>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of version 3 of the GNU General Public
// License as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package hancock

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Instance struct {
	Client *ssh.Client
}

func NewInstance(server, user, password string, privateKey []byte) (*Instance, error) {
	l := make([]ssh.AuthMethod, 0)
	if password != "" {
		l = append(l, ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
			if len(questions) == 0 {
				return []string{}, nil
			}
			return []string{password}, nil
		}))
		l = append(l, ssh.Password(password))
	}
	if privateKey != nil && len(privateKey) != 0 {
		signer, err := ssh.ParsePrivateKey(privateKey)
		if err != nil {
			return nil, err
		}
		l = append(l, ssh.PublicKeys(signer))
	}

	config := &ssh.ClientConfig{
		User:            user,
		Auth:            l,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		// Timeout:         10 * time.Second,
	}
	client, err := ssh.Dial("tcp", server, config)
	if err != nil {
		return nil, err
	}
	return &Instance{
		Client: client,
	}, nil
}

func (i *Instance) Run(cmd string) error {
	s, err := i.Client.NewSession()
	if err != nil {
		return err
	}
	defer s.Close()
	s.Stdout = os.Stdout
	s.Stderr = os.Stderr
	err = s.Run("sudo -H -u root sh -c 'PATH=/root/.nami/bin:$PATH " + cmd + "'")
	if err != nil {
		return err
	}
	return nil
}

func (i *Instance) Start(cmd string) error {
	s, err := i.Client.NewSession()
	if err != nil {
		return err
	}
	defer s.Close()
	return s.Start("sudo -H -u root sh -c 'PATH=/root/.nami/bin:$PATH " + cmd + "'")
}

func (i *Instance) HasNami() (bool, error) {
	s, err := i.Client.NewSession()
	if err != nil {
		return false, err
	}
	defer s.Close()
	s1, err := s.CombinedOutput("sudo -H -u root sh -c '[ -f /root/.nami/bin/nami ] && echo 1'")
	if err != nil {
		return false, nil
	}
	if strings.TrimSpace(string(s1)) != "1" {
		return false, nil
	}
	return true, nil
}

func (i *Instance) HasJoker() (bool, error) {
	s, err := i.Client.NewSession()
	if err != nil {
		return false, err
	}
	defer s.Close()
	s1, err := s.CombinedOutput("sudo -H -u root sh -c '[ -f /root/.nami/bin/joker ] && echo 1'")
	if err != nil {
		return false, nil
	}
	if strings.TrimSpace(string(s1)) != "1" {
		return false, nil
	}
	return true, nil
}

func (i *Instance) InstallNami() error {
	err := i.Run("curl https://bash.ooo/nami.sh | bash")
	if err != nil {
		return err
	}
	return nil
}

func (i *Instance) InstallJoker() error {
	err := i.Run("nami install joker")
	if err != nil {
		return err
	}
	return nil
}

func (i *Instance) Upload(file string) error {
	src, err := os.Open(file)
	if err != nil {
		return err
	}
	defer src.Close()
	c, err := sftp.NewClient(i.Client)
	if err != nil {
		return err
	}
	defer c.Close()
	dst, err := c.Create("/tmp/" + filepath.Base(file))
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	if err := c.Chmod("/tmp/"+filepath.Base(file), 0777); err != nil {
		return err
	}
	err = i.Run("mv /tmp/" + filepath.Base(file) + " /root/.nami/bin/" + filepath.Base(file))
	if err != nil {
		return err
	}
	return nil
}
