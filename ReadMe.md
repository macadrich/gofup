# Send and receive file 
## Objective is to accelerate upload large files

An example that demonstrate client and server for uploading and receiving 
file via TCP Protocol, Also, it can handle multiple upload file from client

## Usage
    * Client
        * To build a client, In /gofup folder, build a client app 
          using command `go build -o client cli/main.go`
        * Copy the client app to your prefered location folder (e.g Desktop)
        * Create a folder `myfile` and put some files that you want to upload in server
        * Run a client using command `./client send myfile`
    * Server
        * In /gofup folder, you can run directly using command `go run cli/main.go recv myfile`
        * Make sure myfile folder exist on the same directory.

## Using Docker
    * run as server `docker run -p 15223:15223 gofup recv myfile`
    * run as client `docker run gofup send myfile`

## How It Works?
    * Client
        * On initiate client side, client will hash all the file listed 
          in directory `myfile` and send it to a server. list of Hash file 
          will use in server side to make sure that it received the exact file uploaded by client.
        * Client send a chunk of bytes to the server for uploading process. 

    * Server
        * On server side, server will receive first the list of hash file, 
          hash file will use to check if incoming files are correct
        * Receiving file will write by chunk of bytes


