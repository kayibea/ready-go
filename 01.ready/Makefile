GO = go build
RM = rm -f

SRC = ready.go
BIN = $(SRC:.go=)

.PHONY: all clean

all: $(BIN)

$(BIN): $(SRC)
	$(GO) $<

clean:
	$(RM) $(BIN)
