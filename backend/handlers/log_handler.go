package handlers

import (
	. "main/backend/models"
)

func GetLogs() ([]Log, error) {
	var logs []Log
	err := Db.Find(&logs)
	if err != nil {
		return []Log{}, nil
	}
	if len(logs) == 0 {
		return []Log{}, nil
	}
	return logs, nil
}
