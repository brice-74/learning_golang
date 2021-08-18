start:
	go run ./cmd/web

create-coverprofile:
	go test -coverprofile=tmp/profile.out ./...

watch-coverprofile:
	go tool cover -html=tmp/profile.out