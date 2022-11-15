
### probar webservice desde consola del navegador
```javascript

var ws = new WebSocket('ws://localhost:8080');

ws.onopen = function() { console.log("Conectado"); };
ws.onmessage = function (evento) { console.log(evento.data); };
ws.onclose = function() { console.log("Desconectado"); };

ws.send("hola");
```
