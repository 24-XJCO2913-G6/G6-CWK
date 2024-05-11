package models

import (
	"fmt"
	"strconv"
)

type Track struct {
	Tid         int64 `xorm:"pk autoincr"` // 轨迹的唯一标识符，被标记为主键
	Uid         int64 // 用户的唯一标识符
	StrDate     string
	StrTime     string // 轨迹的起始时间
	EndDate     string
	EndTime     string // 轨迹的结束时间
	Duration    string
	Speed       string
	Distance    string // 轨迹的距离
	Coordinates string // 轨迹的坐标信息
	// 如[x,y],[x,y]
}
type Records struct {
	Uid         int64
	Distance    string
	Name        string
	Photo       string
	Coordinates string
}

func BuildModelTrack() {
	err := Db.Sync2(new(Track))
	if err != nil {
		panic(err)
	}
}

func AddTrack(Uid int64, StrDate string, StrTime string, EndDate string, EndTime string, Duration string, Distance string, Coordinates string) int64 {
	// 将 Duration 和 Distance 转换为 float64
	durationFloat, err := strconv.ParseFloat(Duration, 64)
	if err != nil {
		panic(err)
	}

	distanceFloat, err := strconv.ParseFloat(Distance, 64)
	if err != nil {
		panic(err)
	}

	// 计算速度
	speed := distanceFloat / durationFloat

	// 将速度转换为保留两位小数的 string 格式
	speedString := strconv.FormatFloat(speed, 'f', 2, 64)

	track := &Track{
		Uid:         Uid,
		StrDate:     StrDate,
		StrTime:     StrTime,
		EndDate:     EndDate,
		EndTime:     EndTime,
		Duration:    Duration,
		Distance:    Distance,
		Speed:       speedString,
		Coordinates: Coordinates,
	}

	fmt.Println("Data:")
	fmt.Println(track)

	Tid, err := Db.Insert(track)
	if err != nil {
		return -1
	}
	return Tid
}
