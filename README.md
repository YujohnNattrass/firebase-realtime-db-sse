# firebase-realtime-db-sse

Steps to reproduce:

1. Run `go run main.go` and we successfully stream an event from firebase with a small data payload.
2. terminate the terminal.
3. Comment line `11` and uncomment line `14` in `main.go`.
4. Run `go run main.go` and wait..... nothing will happen.