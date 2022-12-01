up: 
	echo "starting postgres container"
	docker compose up -d

down:
	docker compose down
# server:
# 	go build -o mainfile .	