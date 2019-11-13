PROJECT=software-theory-final
VERSION=1.0

config:
	mkdir -p /data/software-theory
	cp config/config.example.yaml /data/software-theory/config.yaml

build: config
	docker-compose build -d

deploy: config
	docker-compose up -d