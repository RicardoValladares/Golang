package main
 
import (
	"fmt"
	"log"
	"net/http" 
	"github.com/gorilla/websocket"
)

// arreglo de clientes
var clients = make(map[*websocket.Conn]bool)

// variables necesaria para nuestro WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}


// funcion donde abrimos una nueva conexion 
func onopen(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	clients[ws] = true //agregamos la conexion al arreglo de clientes
	log.Println("Cliente Conectado")
	
	send(ws, "Bienvenido al WebSocket Golang")
	//ws.WriteMessage(websocket.TextMessage, []byte())
	/*if err != nil {
		log.Println(err)
	}*/
	onmessage(ws)
}

// funcion donde recibimos los mensajes 
func onmessage(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))
		//send(conn, string(p))
		onbroadcast(string(p))
	}
}

// funcion donde enviamos el mensaje a un cliente
func send(client *websocket.Conn, data string) {
	err := client.WriteMessage(websocket.TextMessage, []byte(string(data)))
	if err != nil {
		client.Close()
		delete(clients, client)
	}
}

// funcion donde enviamos el mensaje a todos los cliente conectados
func onbroadcast(data string) {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(string(data)))
		if err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}

// funcion donde cerramos un conexion 
func onclose(client *websocket.Conn) {
	client.Close()
	delete(clients, client)
}

//funcion del html y javascript que simula un chat 
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<!DOCTYPE html>
	<html lang="es">
		<head>
			<title>Chat Example</title>
			<script type="text/javascript">
				window.onload = function () {
				    var conn;
				    var msg = document.getElementById("msg");
				    var log = document.getElementById("log");
	
				    function appendLog(item) {
				        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
				        log.appendChild(item);
				        if (doScroll) {
				            log.scrollTop = log.scrollHeight - log.clientHeight;
				        }
				    }
	
				    document.getElementById("form").onsubmit = function () {
				        if (!conn) {
				            return false;
				        }
				        if (!msg.value) {
				            return false;
				        }
				        conn.send(msg.value);
				        msg.value = "";
				        return false;
				    };
	
				    if (window["WebSocket"]) {
				        conn = new WebSocket("ws://" + document.location.host + "/ws");
				        conn.onclose = function (evt) {
				            var item = document.createElement("div");
				            item.innerHTML = "<b>Conexion cerrada.</b>";
				            appendLog(item);
				        };
				        conn.onmessage = function (evt) {
				            var messages = evt.data.split('\n');
				            for (var i = 0; i < messages.length; i++) {
				                var item = document.createElement("div");
				                item.innerText = messages[i];
				                appendLog(item);
				            }
				        };
	    			} else {
				        var item = document.createElement("div");
				        item.innerHTML = "<b>Tu navegador no soporta WebSockets.</b>";
				        appendLog(item);
				    }
				};
			</script>
			<style type="text/css">
				html { overflow: hidden; }
				body { overflow: hidden; padding: 0; margin: 0; width: 100%; height: 100%; background: gray; }
				#log { background: white; margin: 0; padding: 0.5em 0.5em 0.5em 0.5em; position: absolute; top: 0.5em; left: 0.5em; right: 0.5em; bottom: 3em; overflow: auto; }
				#form { padding: 0 0.5em 0 0.5em; margin: 0; position: absolute; bottom: 1em; left: 0px; width: 100%; overflow: hidden; }
			</style>
		</head>
		<body><div id="log"></div><form id="form"><input type="submit" value="Send" /><input type="text" id="msg" size="64" autofocus /></form></body>
	</html>`)
}
 

//metodo principal 
func main() {
	fmt.Println("localhost:8080")
	http.HandleFunc("/", homepage) //http
	http.HandleFunc("/ws", onopen) //ws
	log.Fatal(http.ListenAndServe(":8080", nil))
}
