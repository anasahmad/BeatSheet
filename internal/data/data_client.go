package data

import "BeatSheet/internal/model"

type DataClient interface {
	CreateBeatSheet(beatSheet model.BeatSheet) (string, error)
	RetrieveBeatSheet(id string) (beatSheet *model.BeatSheet, err error)
	UpdateBeatSheet(id string, sheet *model.BeatSheet) (*model.BeatSheet, error)
	DeleteBeatSheet(id string) error
	RetrieveBeatSheets() ([]model.BeatSheet, error)

	AddBeat(beatSheetId string, beat model.Beat) (beatId string, err error)
	UpdateBeat(beatSheetId string, beatId string, beat model.Beat) (*model.Beat, error)
	DeleteBeat(beatSheetId string, beatId string) error

	AddAct(beatSheetId string, beatId string, act model.Act) (actId string, err error)
	UpdateAct(beatSheetId string, beatId string, actId string, act model.Act) (*model.Act, error)
	DeleteAct(beatSheetId string, beatId string, actId string) error
}
