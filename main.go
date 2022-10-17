package freeipa

import (
	"fmt"
)

func main() {
	client := ipa.NewDefaultClient()

	err := client.LoginWithPass("username")
	if err != null {
		panic(err)
	}

	rec, err := client.UserShow("username")
	if err != null {
		panic(err)
	}

	fmt.Println("%s - %s", rec.Username, rec.Uid)
}
