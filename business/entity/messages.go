package entity

const (
	MessageStart = `Hello, I'm a RSS bot! I will receive news from the sites you subscribe to`
	MessageHelp  = `/add [site or RSS feed URL] or just send the site URL in a message - add subscription 
/list - get a list of my subscriptions
/delete [subscription number] - delete subscription
/help - get this help`
	MessageNoSubscriptions = `You have not yet subscribed to more than one feed, to do this, enter:
/add [site or RSS feed URL]`
	MessageSubscriptionDeleted       = `Subscription successful deleted`
	MessageSubscriptionCreated       = `Subscription successful created`
	MessageUnknownCommand            = `Invalid command, enter /help to get a list of commands`
	MessageSubscriptionAlreadyExists = `You are already subscribed to this channel`
	MessageWrongSubscriptionNumber   = `You do not have a subscription with this number`
	MessageWrongURL                  = `Incorrect website url`
	MessageNoFeed                    = `The site does not contain RSS feed`
	MessageErrorLoadURL              = `Failed to load this website`
	MessageInternalError             = `Oops! something went wrong, please try again later`
)
