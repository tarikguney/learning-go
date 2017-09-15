## Run the application

```
go run main.go &
```

## Stop

```
# Finding which process ID using port 8080
lsof -i tcp:8080
kill process_id
```