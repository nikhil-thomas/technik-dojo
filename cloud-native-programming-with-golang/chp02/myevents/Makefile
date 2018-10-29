mongo-restart:
	@docker start myevent-mongo

mongo-start:
	@docker run -p 27017:27017 --name myevent-mongo -d mongo:4.1

mongo-stop:
	@docker stop myevent-mongo

mongo-clean:
	@docker rm myevent-mongo
