package listener

import "fair-ticket-be-tutorial/service"

func Start() {
	e := service.NewEventService()
	go e.HandleProjectCreated()
	go e.HandleProjectStarted()
	go e.HandleProjectFinished()
	go e.HandleMagicNumberPublished()
}
