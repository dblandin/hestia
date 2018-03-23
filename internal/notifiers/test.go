package notifiers

type Test struct {
	Messages []string
}

func (n *Test) Log(message string) {
	n.Messages = append(n.Messages, message)

}
