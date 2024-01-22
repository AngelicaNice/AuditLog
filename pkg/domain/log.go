package audit

import (
	"errors"
	"time"
)

const (
	ENTITY_ACTOR = "ACTOR"
	ENTITY_USER  = "USER"

	ACTION_REGISTER = "REGISTER"
	ACTION_LOGIN    = "LOGIN"
	ACTION_CREATE   = "CREATE"
	ACTION_GET      = "GET"
	ACTION_UPDATE   = "UPDATE"
	ACTION_DELETE   = "DELETE"
)

var (
	entities = map[string]LogRequest_Entities{
		ENTITY_ACTOR: LogRequest_ACTOR,
		ENTITY_USER:  LogRequest_USER,
	}
	actions = map[string]LogRequest_Actions{
		ACTION_REGISTER: LogRequest_REGISTER,
		ACTION_LOGIN:    LogRequest_LOGIN,
		ACTION_CREATE:   LogRequest_CREATE,
		ACTION_GET:      LogRequest_GET,
		ACTION_UPDATE:   LogRequest_UPDATE,
		ACTION_DELETE:   LogRequest_DELETE,
	}
)

type LogItem struct {
	Entity    string    `bson:"entity"`
	Action    string    `bson:"action"`
	EntityID  int64     `bson:"entity_id"`
	Timestamp time.Time `bson:"timestamp"`
}

func ToPbEntity(entity string) (LogRequest_Entities, error) {
	val, exist := entities[entity]
	if !exist {
		return 0, errors.New("invalid entity")
	}

	return val, nil
}

func ToPbAction(action string) (LogRequest_Actions, error) {
	val, exist := actions[action]
	if !exist {
		return 0, errors.New("invalid action")
	}

	return val, nil
}
