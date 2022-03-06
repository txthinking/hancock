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
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/olekukonko/tablewriter"
	"go.etcd.io/bbolt"
)

type Hancock struct {
	DB *bbolt.DB
}

type SSH struct {
	Name     string
	Server   string
	User     string
	Password string
	Key      []byte
}

func NewHancock() (*Hancock, error) {
	s, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	db, err := bbolt.Open(filepath.Join(s, ".hancock"), 0644, nil)
	if err != nil {
		return nil, err
	}
	return &Hancock{
		DB: db,
	}, nil
}

func (h *Hancock) Add(name, server, user, password, key string) error {
	s := &SSH{
		Name:     name,
		Server:   server,
		User:     user,
		Password: password,
	}
	if key != "" {
		b, err := os.ReadFile(key)
		if err != nil {
			return err
		}
		s.Key = b
	}
	err := h.DB.Update(func(tx *bbolt.Tx) error {
		b, err := json.Marshal(s)
		if err != nil {
			return err
		}
		t, err := tx.CreateBucketIfNotExists([]byte("ssh"))
		if err != nil {
			return err
		}
		if err := t.Put([]byte(s.Name), b); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (h *Hancock) Remove(name string) error {
	err := h.DB.Update(func(tx *bbolt.Tx) error {
		t, err := tx.CreateBucketIfNotExists([]byte("ssh"))
		if err != nil {
			return err
		}
		if err := t.Delete([]byte(name)); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (h *Hancock) PrintAll() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Server", "User", "Password", "Key"})
	err := h.DB.Update(func(tx *bbolt.Tx) error {
		t, err := tx.CreateBucketIfNotExists([]byte("ssh"))
		if err != nil {
			return err
		}
		err = t.ForEach(func(k, v []byte) error {
			s := &SSH{}
			if err := json.Unmarshal(v, s); err != nil {
				return err
			}
			ks := "false"
			if s.Key != nil {
				ks = "true"
			}
			table.Append([]string{s.Name, s.Server, s.User, s.Password, ks})
			return nil
		})
		return nil
	})
	if err != nil {
		log.Println(err)
		os.Exit(1)
		return
	}
	table.Render()
}

func (h *Hancock) prepare(args []string) (*Instance, string, error) {
	s := &SSH{}
	err := h.DB.Update(func(tx *bbolt.Tx) error {
		t, err := tx.CreateBucketIfNotExists([]byte("ssh"))
		if err != nil {
			return err
		}
		b := t.Get([]byte(args[0]))
		if b == nil {
			return errors.New("No instance named " + args[0])
		}
		if err := json.Unmarshal(b, s); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, "", err
	}
	l := make([]string, 0)
	for i, v := range args {
		if i == 0 {
			continue
		}
		if strings.Contains(v, " ") {
			v = fmt.Sprintf("\"%s\"", v)
			l = append(l, v)
			continue
		}
		if v == "&&" {
			v = "&& PATH=/root/.nami/bin:$PATH"
		}
		l = append(l, v)
	}
	i, err := NewInstance(s.Server, s.User, s.Password, s.Key)
	if err != nil {
		return nil, "", err
	}
	has, err := i.HasNami()
	if err != nil {
		return nil, "", err
	}
	if !has {
		if err := i.InstallNami(); err != nil {
			return nil, "", err
		}
	}
	has, err = i.HasJoker()
	if err != nil {
		return nil, "", err
	}
	if !has {
		if err := i.InstallJoker(); err != nil {
			return nil, "", err
		}
	}
	return i, strings.Join(l, " "), nil
}

func (h *Hancock) Run(args []string) error {
	i, s, err := h.prepare(args)
	if err != nil {
		return err
	}
	if err := i.Run(s); err != nil {
		return err
	}
	return nil
}

func (h *Hancock) Start(args []string) error {
	i, s, err := h.prepare(args)
	if err != nil {
		return err
	}
	return i.Start(s)
}

func (h *Hancock) Upload(name, file string) error {
	i, _, err := h.prepare([]string{name})
	if err != nil {
		return err
	}
	return i.Upload(file)
}

func (h *Hancock) Close() error {
	return h.DB.Close()
}
