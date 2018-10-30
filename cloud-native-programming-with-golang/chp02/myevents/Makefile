mongo-restart:
	@docker start myevent-mongo

mongo-start:
	@docker run -p 27017:27017 --name myevent-mongo -d mongo:4.1

mongo-stop:
	@docker stop myevent-mongo

mongo-clean:
	@docker rm myevent-mongo

cert_openssl:
	openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 365

cert_go:
	go run /usr/local/go/src/crypto/tls/generate_cert.go --host=localhost
