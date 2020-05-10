package connectionHandler

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// Handle
// Process the connection from the tcp server
func Handle(connection net.Conn) {
	defer connection.Close()

	// Receive and process the connection
	uri := receiver(connection)

	// Process and respond to connection based on uri

}

func receiver(conn net.Conn) string {
	// Create a new scanner to read from the connection
	scanner := bufio.NewScanner(conn)
	var uri string
	i := 0
	for scanner.Scan() {
		textLine := scanner.Text()

		// Write all the header lines to the standard output.
		io.WriteString(os.Stdout, "\n"+textLine+"\n")

		// If it is an http connection, first line contains methods and resource
		if i == 0 {
			resources := strings.Fields(textLine)[0]
		}

		// Stop scanner when blank line after headers reached. (As per IETF standards)
		if textLine == "" {
			break
		}
	}
	return uri
}
