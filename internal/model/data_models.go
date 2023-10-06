package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BeatSheetData struct {
	Id    string `bson:"_id"`
	Title string `bson:"title"`
	Beats []Beat `bson:"beats"`
}

type BeatData struct {
	Id          string
	Description string
	Timestamp   primitive.Timestamp
	Acts        []Act
}

type ActData struct {
	Id          string
	Description string
	Timestamp   primitive.Timestamp
	Duration    int64
	CameraAngle string
}
