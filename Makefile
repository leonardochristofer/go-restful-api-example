goCMD=go

dockerCMD=docker-compose

run:
	$(dockerCMD) up --build

stop:
	$(dockerCMD) down

fmt:
	$(goCMD)fmt -s -w .

mock-gen:
	$(goCMD) generate ./...

test:
	$(goCMD) test ./...
