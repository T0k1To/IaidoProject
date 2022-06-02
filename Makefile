all: libprocesshider client server

libprocesshider: processhider.c
	gcc -Wall -fPIC -shared -o libprocesshider.so processhider.c -ldl

client:
	go build -o bin/ -ldflags "-s -w" Iaido.go

server:
	go build -o bin/ -ldflags "-s -w" IaidoSV.go

.PHONY clean:
	rm -f libprocesshider.so
