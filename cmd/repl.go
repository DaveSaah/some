/*
Copyright © 2024 David Saah <davesaah@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/user"

	"github.com/DaveSaah/some/repl"
	"github.com/spf13/cobra"
)

const logo = `
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
		fmt.Printf("Hello %s, This is the some programming language.\n", user.Username)
		repl.Start(os.Stdin)
	},
}

func init() {
	rootCmd.AddCommand(replCmd)
}
