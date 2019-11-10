config:
	mkdir -p /data/software-theory
	cp config/config.example.yaml /data/software-theory/config.yaml

build: config
	docker-compose build

deploy: config
	docker-compose up -d