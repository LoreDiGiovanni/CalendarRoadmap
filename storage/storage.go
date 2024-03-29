package storage

import "RoadmapCalendar/types"

type Storage interface {
	PostEvents(types.Events) error
	PostUser(types.User) error
    GetEvents(string) (*[]types.Events, error)
    GetUser(types.User) (*types.User,error)
}
