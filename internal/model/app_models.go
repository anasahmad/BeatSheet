package model

import (
	"time"
)

type BeatSheet struct {
	Id    string `bson:"id" json:"id"`
	Title string `bson:"title" json:"title"`
	Beats []Beat `bson:"beats" json:"beats"`
}

type Beat struct {
	Id          string    `bson:"id" json:"id"`
	Description string    `bson:"description" json:"description"`
	Timestamp   time.Time `bson:"timestamp" json:"timestamp"`
	Acts        []Act     `bson:"acts" json:"acts"`
}

type Act struct {
	Id          string    `bson:"id" json:"id"`
	Description string    `bson:"description" json:"description"`
	Timestamp   time.Time `json:"timestamp" json:"timestamp"`
	Duration    int64     `bson:"duration" json:"duration"`
	CameraAngle string    `bson:"cameraAngle" json:"cameraAngle"`
}
