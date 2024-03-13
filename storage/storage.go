package storage

import "RoadmapCalendar/types"

type Storage interface {
	PostEvents(string, types.Events) error
	PostUser(types.User) error
    GetEvents(string) (*[]types.Events, error)
}
