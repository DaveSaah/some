package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/DaveSaah/some/lexer"
	"github.com/DaveSaah/some/token"
	"github.com/spf13/cobra"
)

const (
	PROMPT = ">> "
	logo   = `
-++++++-----==-----=++++=-
*%####+-+*#****##+=*####%*
#%*---+%+:  ..  .+%*---#%*
#%*--*%-  =##*#+  :%*--#%*
*#=-+%#  .%*..=%- :##--**=
=+=-+%%.  =%#*%#+#%#=---==
#%%-=#%%=.  :+#%%#=---=%%+
++====+#%%*=.  :+##+---==-
+%#==--+#%%#%*=.  -%*--#%*
**+=-+%%*=+*+=#%:  =%==**=
=*+=+%%. :%*..=%=  =%+-=*=
#%*=+#%+  -####=  :%#-=#%*
#%*==+#%#=. .  .-*%*=-=#%*
#%#*****#%%%###%#*=+***%%*
+#####*==++++++====####**=

`
)

// replCmd represents the repl command
var replCmd = &cobra.Command{
	Use:   "repl",
	Short: "Start a repl",
	Run: func(cmd *cobra.Command, args []string) {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}

		fmt.Print(logo)
		fmt.Printf(
			"Hello %s, This is the some programming language.\n",
			user.Username,
		)
		startRepl(os.Stdin)
	},
}

func init() {
	rootCmd.AddCommand(replCmd)
}

func startRepl(in io.Reader) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
