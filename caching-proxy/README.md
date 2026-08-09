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