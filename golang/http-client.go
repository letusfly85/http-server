/**
 * refs: http://whispering-gophers.appspot.com/talk.slide#16
 * refs: https://gist.github.com/iwanbk/2295233
 *
 *
 *
 */

package main

import "net"
import "os"

func main() {
	strEcho := "hello! world!"
	servAddr := "localhost:3333"

	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("Resolve TCP address failed:", err.Error())
		os.Exit(1)
	}
	println(tcpAddr.IP)
	println(tcpAddr.Port)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("write to server = ", strEcho)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server=", string(reply))

	conn.Close()
}
