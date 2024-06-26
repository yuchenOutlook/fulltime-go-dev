package main

import (
	"fmt"
	"go-module-project/types"
	"go-module-project/utls"
)

/**
-- program --
1. something that runs / executes the code, an entry code
that's => main()
=> package main

2. package / library / module
it's something can be exported / imported for others to run in their programs.
-> package <whatever>

3. User -> Public access everywhere
	 	-> Private access BUT public in its own package
**/
func main() {
	user := types.User{Username: utls.GetUserName(), Age: utls.GetAge()}
	
	fmt.Printf("THe user is %s, age is %d\n", user.Username, user.Age)
}