GO = go build
RM = rm -f

SRC = analyze.go
BIN = $(SRC:.go=)

.PHONY: all clean

all: $(BIN)

$(BIN): $(SRC)
	$(GO) $<

clean:
	$(RM) $(BIN)
