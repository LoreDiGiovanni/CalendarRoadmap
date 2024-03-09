package storage

import "RoadmapCalendar/types"

type Storage interface {
	PostEvents(types.User, types.Events) error
    GetEvents(types.User) (*[]types.Events, error)
}
