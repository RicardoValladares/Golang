
### generar binario como variable de entorno
```batch
go install 
Git.exe
```


### obtener libreria actualizada de git
```batch
go get github.com/Morty-debug/librerias@latest
```


### obtener binarios de la libreria en git y usarla como variable de entorno
```batch
go install github.com/Morty-debug/librerias@latest
librerias.exe
```


### crear el proyecto desde cero
```batch
mkdir Git 
cd Git 
notepad main.go
go mod init github.com/RicardoValladares/Golang/10-Git
go mod tidy
```
