config:
	mkdir -p /data/software-theory
	cp config/config.example.yaml /data/software-theory/config.yaml

deploy: config
	docker-compose up -d