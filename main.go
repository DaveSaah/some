package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/DaveSaah/some/repl"
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

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println(logo)
	fmt.Printf("Hello %s, This is the some programming language.\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
