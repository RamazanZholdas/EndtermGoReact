run:
	sudo docker-compose up -d --build

stop:
	sudo docker-compose stop

runClient:
	cd client/ && npm start

runServer: 
	cd server/ && go run main.go