prepare:
	go mod vendor
	cd frontend && npm install
	copy go.mod backend
	copy go.sum backend