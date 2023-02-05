package db

import (
	"BaigeiCode/yearbook_api/model"
	"fmt"
)

func MultiCreateEvents(events []*model.Event) ([]int64, error) {
	err := db.Omit("ID").Create(events).Error
	if err != nil {
		return nil, err
	}
	ids := make([]int64, 0, len(events))
	for _, event := range events {
		if event.ID == 0 {
			continue
		}
		ids = append(ids, event.ID)
	}
	return ids, nil
}

func DeleteAllEvents(year int64) error {
	return db.Where("publish_time like ?", fmt.Sprintf("%d-%-%", year)).Delete(&model.Event{}).Error
}
func MultiGetEvents(year int64) ([]*model.Event, error) {
	events := make([]*model.Event, 0)
	err := db.Where("publish_time like ?", fmt.Sprintf("%d-%-%", year)).Find(&events).Error
	return events, err
}
