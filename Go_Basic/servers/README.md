# Server

This project uses concurrency and go routines to check if the following servers are running:

- http://google.com
- http://facebook.com
- http://instagram.com

## How to use it

Move to the project folder and run the project on the terminal with the command: `go run main.go` it will print if the server is available or not. 

You can add your own servers adding the url to the `servers` variable as follows:

    servers := []string{
        "http://google.com",
        "http://facebook.com",
        "http://instagram.com",
        "http://myServer.com",
    }
