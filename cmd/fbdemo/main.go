package main

import (
	"fmt"

	"github.com/ethanfrey/flatbuffers/users"
	flatbuffers "github.com/google/flatbuffers/go"
)

func MakeUser(name string, id uint64) []byte {
	b := flatbuffers.NewBuilder(256)

	// create the name object and get its offset:
	fbName := b.CreateString(name)

	// write the User object:
	users.UserStart(b)
	users.UserAddName(b, fbName)
	users.UserAddId(b, id)
	userPosition := users.UserEnd(b)

	// finish the write operations by our User the root object:
	b.Finish(userPosition)

	// return the byte slice containing encoded data:
	return b.FinishedBytes()
}

func main() {
	data := MakeUser("Johnathan Bullwinkle", 1736)
	user := users.GetRootAsUser(data, 0)
	fmt.Printf("Read user #%d: %s\n",
		user.Id(), string(user.Name()))
}
