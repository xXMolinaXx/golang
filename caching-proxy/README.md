- init 
```bash
go mod init caching-proxy
```

- run proyect 
```bash
go run . --port 3000 --origin http://example.com
```

- create binary file
```bash
go build -o caching-proxy
```

- execute binary file
```bash
./caching-proxy --port 3000 --origin http://example.com
```

-- Install System-Wide
```bash
sudo mv caching-proxy /usr/local/bin/
go install
```

## Tareas

Lo que aún falta o está incompleto
- El cache usa sólo r.URL.Path como clave. No consideras query strings ni otros posibles identificadores de la petición, por lo que requests distintos pueden colisionar.
- El slice global alreadyCached no es seguro para concurrencia. Si llegan varias peticiones simultáneas tu proxy puede tener condiciones de carrera.
- En un cache hit siempre devuelves http.StatusOK y Content-Type: application/json, incluso si el origen devolvió otro código o contenido diferente.
- No preservas los headers originales del origen ni respaldas el StatusCode en el cache.
- No eliminas realmente las entradas expiradas: las dejas en el slice y las ignoras, pero el almacenamiento crece sin límite.