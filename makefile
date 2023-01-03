all: rest kafka
rest:
	go build -ldflags '-w -s -extldflags "-static"' -o usersearch-rest ./cmd/rest
kafka:
	go build -ldflags '-w -s -extldflags "-static"' -o usersearch-kafka ./cmd/kafka
