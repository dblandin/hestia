package notifiers

type Notifier interface {
	Log(message string)
}
