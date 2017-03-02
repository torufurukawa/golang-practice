package main

import (
	"bufio"
	"fmt"
	"io"
)

func main() {
	//conn, err := net.Dial("tcp", "golang.org:80")
	conn := NewMockConn()
	conn.Inbuf = []byte("HTTP/1.0 302 Found\r\n")

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	fmt.Println("outbuf:", string(conn.Outbuf))
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("status:", status)
	fmt.Println("err:", err)
}

type MockConn struct {
	Inbuf  []byte
	read   bool
	Outbuf []byte
	reader io.PipeReader
	writer io.PipeWriter
}

func NewMockConn() *MockConn {
	c := MockConn{}
	c.Inbuf = make([]byte, 0)
	c.read = false
	c.Outbuf = make([]byte, 0)
	return &c
}

// Read reads data from the connection.
// Read can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetReadDeadline.
func (c *MockConn) Read(b []byte) (n int, err error) {
	if c.read {
		return 0, nil
	}

	for i := 0; i < len(c.Inbuf); i++ {
		b[i] = c.Inbuf[i]
	}
	return len(c.Inbuf), nil
}

// Write writes data to the connection.
// Write can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetWriteDeadline.
func (c *MockConn) Write(b []byte) (n int, err error) {
	n = len(b)
	c.Outbuf = append(c.Outbuf, b...)
	return n, nil
}
