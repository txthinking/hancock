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
	"net"
	"time"
)

type Conn struct {
	net.Conn
	Timeout int
}

func (c *Conn) Read(b []byte) (int, error) {
	if err := c.Conn.SetDeadline(time.Now().Add(time.Duration(c.Timeout) * time.Second)); err != nil {
		return 0, err
	}
	return c.Conn.Read(b)
}

func (c *Conn) Write(b []byte) (int, error) {
	if err := c.Conn.SetDeadline(time.Now().Add(time.Duration(c.Timeout) * time.Second)); err != nil {
		return 0, err
	}
	return c.Conn.Write(b)
}
