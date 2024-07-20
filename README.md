# Go WebSocket Project Documentation
## Table of Contents

1. [Project Overview](#project-overview)
2. [Installation Instructions](#installation-instructions)
3. [Usage](#usage)
4. [API Endpoints](#api-endpoints)
5. [Dependencies](#dependencies)
6. [Example](#examples)

## <a name="project-overview">Project Overview</a>
This project is a WebSocket server written in Go. It allows real-time communication between clients and the server.

## <a name="installation-instructions">Installation Instructions</a>
### Prerequisites

- Go version 1.16 or higher
- Git

### Steps

1. **Clone the Repository**

    ```sh
    git clone https://github.com/yossev/go-websocket-chat.git
    cd go-websocket-chat
    ```

2. **Install Dependencies**

    ```sh
    go mod tidy
    ```

3. **Run the Server**
## <a name="usage">Usage</a>
    ```sh
    go run main.go
    ```
## <a name="usage">Usage</a>
1 - Run the `server.go` file. <br />
2 - Run the HTML File using any `live preview` extension, [example](https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer)

## <a name="api-endpoints">Api Endpoints </a>
### Websocket Endpoint
- Enpoint '/chat'
- Method `GET`
- Description: Establishes a WebSocket connection with the server.
  
## <a name="dependencies">Dependencies</a>
- [gorilla/websocket](https://github.com/gorilla/websocket) -  WebSocket Library for Go.

## <a name="examples"> Video Example </a>
https://github.com/user-attachments/assets/2d5f9dc5-514c-4353-8eda-5b8c8d08ab34







