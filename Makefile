help:
	@cat Makefile | grep '^\w'

aozora.zip:
	wget -O aozora.zip https://github.com/aozorabunko/aozorabunko/archive/master.zip

extract: aozora.zip
	unzip -f aozora.zip

run:
	go run main.go

