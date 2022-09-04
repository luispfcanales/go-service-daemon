# Web http golang

```go
func main() {
  http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
    w.Write([]byte("Hi Gophers!"))
  })
  http.ListenAndServe(":80",nil)
}
```

## List to learn
- [x] create file to execute as a service ``goweb.service``
  
  ```
  [Unit]
  Description=descriptiton to application service
  ConditionPathExists=/ApplicationDirectoryPath
  After=network.target

  [Service]
  WorkingDirectory=/ApplicationDirectoryPath
  ExecStart=/ApplicationDirectoryPath/BinaryApp
  Restart=on-failure
  RestartSec=10

  [Install]
  WantedBy=multi-user.target
  ```
- [x] http web server running in daemon mode
  
  Commands to execute as a root user

  ```bash
  systemctl start nameService
  systemctl stop nameService
  systemctl enable nameService
  ```
  ```bash
  systemctl daemon-reload
  ```
