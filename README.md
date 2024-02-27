# Chater

## 1. Introduction
This is a simple chat application that allows users to send and receive messages in real time using tcp sockets. The application is written in Go and uses the `net` package to create a tcp server and client. The server listens for incoming connections and broadcasts messages to all connected clients. The client connects to the server and sends messages to the server. The application uses ssl to encrypt the communication between the server and the client. The application also use  [BubbleTea](https://github.com/charmbracelet/bubbletea) to create a simple terminal UI.

### Run the client
```$ ./chater client -d <server-ip> -p <port>```

### Run the server
```$ ./chater server -d <server-ip> -p <port>```

### Create ssl certificate
```$ mkdir server && cd server```

```$ openssl ecparam -genkey -name secp384r1 -out server.key```

```$ openssl req -new -x509 -sha256 -key server.key -out server.crt -days 365```


#### TODO
- [x] Create a simple chat server/client with ssl 
- [x] build a simple terminal UI using BubbleTea
- [ ] Add more features to the chat application
- [ ] Add tests
- [ ] Add simple database to store chat messages
- [ ] Add authentication
- [ ] Add more security features
- [ ] Server just handle clients to connect to each other and send messages with encryption and data storage
