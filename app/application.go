package app

func StartApp() {
	go startMessageConsumer()
	select {}
}
