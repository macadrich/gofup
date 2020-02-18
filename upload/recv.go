package upload

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

// Receive wait to receive file from client
func (u *Upload) Receive() error {
	log.Println("listening on: ", u.Address)
	conn, err := net.Listen("tcp", u.Address)
	if err != nil {
		return err
	}

	files, err := CHashFile(u.Dir)
	if err != nil {
		return err
	}

	log.Println("waiting for connection...")

	for {
		s, err := conn.Accept()
		if err != nil {
			return err
		}

		go handleConn(s, files, u.Dir)
	}

	return nil
}

func handleConn(server net.Conn, files []File, dir string) {
	dec := json.NewDecoder(server)
	enc := json.NewEncoder(server)

	var available []File
	if err := dec.Decode(&available); err != nil {
		log.Println(err)
	}

	for _, a := range available {
		var found bool
		for _, f := range files {
			if f.Name == a.Name {
				if f.Hash != a.Hash {
					log.Printf("sender has file %s with different hash\n", a.Name)
					break
				} else {
					found = true
					break
				}
			}
		}
		if found {
			continue
		}

		log.Printf("Incoming File: %s\n", a.Name)
		if err := enc.Encode(&a); err != nil {
			log.Println(err)
		}

		fi, err := os.Create(filepath.Join(dir, a.Name))
		if err != nil {
			log.Println(err)
		}
		defer fi.Close()

		var receivedBytes int64 // receivedBytes = 0
		for {
			// receivedBytes = 0, receivedBytes = 1024, 2048...
			if (a.Size - receivedBytes) < BUFFERSIZE {
				io.CopyN(fi, server, (a.Size - receivedBytes))
				server.Read(make([]byte, (receivedBytes+BUFFERSIZE)-a.Size))
				log.Println("Received:", a.Name, receivedBytes)
				break
			}
			io.CopyN(fi, server, BUFFERSIZE)
			// increment 1024 to receivedBytes
			receivedBytes += BUFFERSIZE
		}
	}
	if err := enc.Encode(File{}); err != nil {
		log.Println(err)
	}
	server.Close()
}
