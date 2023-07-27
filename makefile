sources = sources/main.go \
		  sources/lex.go \
		  sources/utils.go \
		  sources/pars.go

final = ayak.exe


build: $(sources)
	go build -o $(final) $(sources)