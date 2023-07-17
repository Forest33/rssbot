package entity

import "strings"

const (
	CommandStart              = "/start"
	CommandAddSubscription    = "/add"
	CommandListSubscriptions  = "/list"
	CommandDeleteSubscription = "/delete"
	CommandHelp               = "/help"
)

func ParseCommand(msg string) (string, []string, error) {
	msg = strings.TrimSpace(msg)
	args := strings.Split(msg, " ")

	if !strings.HasPrefix(args[0], "/") {
		if ValidateURL(msg) == nil {
			return CommandAddSubscription, []string{msg}, nil
		}
		return "", nil, ErrNotCommand
	}

	args[0] = strings.ToLower(args[0])

	if len(args) > 1 {
		return args[0], args[1:], nil
	}

	return args[0], nil, nil
}
