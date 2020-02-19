package fhandler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

// Send connect to server and upload file
func (u *FileHandler) Send() error {
	log.Println("Folder:", u.FileFolder, "Address:", u.Address)
	server, err := net.Dial("tcp", u.Address)
	if err != nil {
		return err
	}
	defer server.Close()

	log.Println("Connected.")

	files, err := CHashFile(u.FileFolder)
	if err != nil {
		return err
	}

	filesInByte, err := json.Marshal(files)
	if err != nil {
		log.Println("File in bytes:", err)
		return err
	}

	_, err = server.Write(filesInByte)
	if err != nil {
		log.Println("write server :", err)
		return err
	}

	fdec := json.NewDecoder(server) // fdec buffering
	for {
		var req File
		// decode to File object (from fdec buffering)
		if err := fdec.Decode(&req); err != nil {
			return errors.New("Decoding error: EOF")
		}

		// check if no files found in (fdec buffering)
		if req.Name == "" {
			log.Println("done!")
			return nil
		}

		var found bool
		for _, f := range files {
			if f == req {
				found = true
				break
			}
		}

		if !found {
			return errors.New("requested file not found")
		}

		log.Println("sending: ", req.Name)
		fi, err := os.Open(filepath.Join(u.FileFolder, req.Name)) // get specific path file name
		if err != nil {
			return err
		}

		bytechunk := make([]byte, BUFFERSIZE)
		// run to goroutine
		go func(server net.Conn, bc []byte) {
			for {
				_, err = fi.Read(bytechunk)
				if err != nil {
					if err == io.EOF {
						log.Println("end:", err) // indicate read end of file.
						break
					}
					log.Fatal(err)
				}
				server.Write(bytechunk)
			}
			log.Println("File has been sent!")
		}(server, bytechunk)
	}
}
