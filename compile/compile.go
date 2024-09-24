package compile

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DaveSaah/some/lexer"
	"github.com/DaveSaah/some/token"
)

func Start(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
