config:
	mkdir /data
	cp config/config.example.yaml /data/config.yaml

deploy: config
	docker-compose up -d