package notifiers

import (
	"fmt"
)

type Stdout struct {
}

func (n Stdout) Log(message string) {
	fmt.Println(message)
}
