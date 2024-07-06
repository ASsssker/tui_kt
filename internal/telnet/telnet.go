package telnet

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"strings"
	"unicode/utf8"
)

const (
	ECHO = 1

	WONT = 252
	DO   = 253
	DONT = 254
	IAC  = 255
)

type Client struct {
	conn net.Conn
}

func GetClient(host string, port uint) (*Client, error) {
	conn, err := connect(host, port)
	if err != nil {
		return nil, err
	}

	return &Client{conn: conn}, nil
}

func connect(host string, port uint) (net.Conn, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", addr)
	return conn, err
}

func (c *Client) ReadUntil(subbytes []byte) ([]byte, error) {
	var data []byte

	for {
		buf := make([]byte, 1)
		_, err := c.conn.Read(buf)
		if err != nil {
			return nil, err
		}
		data = append(data, buf...)

		if data[0] == IAC && len(data) == 3 {
			if data[2] == ECHO {
				data[1] = WONT
			}

			c.conn.Write(data)
			data = []byte{}
		}

		if bytes.Contains(data, subbytes) {
			return data, nil
		}
	}
}

func (c *Client) ReadString(substring string) (string, error) {
	subbytes := []byte(substring)
	if utf8.RuneCountInString(substring) != len(subbytes) {
		return "", errors.New("only ASCII characters are allowed")
	}

	data, err := c.ReadUntil(subbytes)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (c *Client) Write(data []byte) (int, error) {
	data = append(data, '\r', '\n')
	n, err := c.conn.Write(data)

	return n, err
}

func (c *Client) WriteString(data string) (int, error) {
	dataBytes := []byte(data)
	if utf8.RuneCountInString(data) != len(dataBytes) {
		return 0, errors.New("only ASCII characters are allowed")
	}
	n, err := c.Write(dataBytes)

	return n, err
}

func (c Client) Auth(username, password string) bool {
	key := [...]string{"login", "Password"}
	value := [...]string{username, password}

	for idx := range key {
		_, err := c.ReadString(key[idx])
		if err != nil {
			return false
		}

		_, err = c.WriteString(value[idx])
		if err != nil {
			return false
		}
	}

	_, err := c.ReadString("\r\n\r\n")
	if err != nil {
		return false
	}

	str, err := c.ReadString("\r\n")
	if err != nil || strings.Contains(str, "incorrect") {
		return false
	}
	c.ReadString("$")

	return true
}
