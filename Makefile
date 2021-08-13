genpb:
	buf generate --path proto/greeting/v1 -o proto

run:
	go run main.go