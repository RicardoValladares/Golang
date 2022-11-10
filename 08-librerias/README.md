### crear proyecto
```batch
mkdir go_ejemplo
cd go_ejemplo
notepad main.go
mkdir libreria
cd libreria
notepad salidas.go
cd ..
```

### inicializar proyecto
```batch
go mod init github.com/RicardoValladares/Golang/08-librerias
go mod tidy
```

### interpretar codigo
```batch
go run main.go
```

### compilar codigo
```batch
go build main.go
```

