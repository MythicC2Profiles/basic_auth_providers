package main

import (
	"GoServices/my_auth"
	"github.com/MythicMeta/MythicContainer"
)

func main() {
	// load up the agent functions directory so all the init() functions execute
	//httpfunctions.Initialize()
	//basicAgent.Initialize()
	//mytranslatorfunctions.Initialize()
	//my_webhooks.Initialize()
	//my_logger.Initialize()
	//my_event_processor.Initialize()
	//customAugmentFunctions.Initialize()
	my_auth.Initialize()
	// sync over definitions and listen
	MythicContainer.StartAndRunForever([]MythicContainer.MythicServices{
		//MythicContainer.MythicServiceC2,
		//MythicContainer.MythicServiceTranslationContainer,
		//MythicContainer.MythicServiceWebhook,
		//MythicContainer.MythicServiceLogger,
		//MythicContainer.MythicServicePayload,
		//MythicContainer.MythicServiceEventing,
		MythicContainer.MythicServiceAuth,
	})
}
