package controller

import (
	"BeatSheet/internal/ai"
	"BeatSheet/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestActController_DELETEAct(t *testing.T) {
	hp := HappyPathClient{}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
	ts := httptest.NewServer(handler)
	req, _ := http.NewRequest("DELETE", ts.URL, nil)
	c.Request = req
	con := NewActController(&hp, ai.NewAIModel())
	con.DELETEAct(c)
	assert.Equal(t, 204, c.Writer.Status())
}

func TestActController_POSTAct(t *testing.T) {
	hp := HappyPathClient{}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
	ts := httptest.NewServer(handler)
	req, _ := http.NewRequest("DELETE", ts.URL, nil)
	c.Request = req
	con := NewActController(&hp, ai.NewAIModel())
	con.POSTAct(c)

	assert.Equal(t, 201, c.Writer.Status())
}

func TestActController_PUTAct(t *testing.T) {

	hp := HappyPathClient{}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
	ts := httptest.NewServer(handler)
	req, _ := http.NewRequest("DELETE", ts.URL, nil)
	c.Request = req
	con := NewActController(&hp, ai.NewAIModel())
	con.PUTAct(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestNewActController(t *testing.T) {
	hp := &HappyPathClient{}
	c := NewActController(hp, nil)
	acctId, err := c.dataClient.AddAct("", "", model.Act{})
	accIdHp, errHp := hp.AddAct("", "", model.Act{})
	assert.Equal(t, acctId, accIdHp)
	assert.Equal(t, err, errHp)
}

type HappyPathAI struct {
}

type HappyPathClient struct {
}

func (m *HappyPathClient) CreateBeatSheet(beatSheet model.BeatSheet) (string, error) {
	return "", nil
}
func (m *HappyPathClient) RetrieveBeatSheet(id string) (beatSheet *model.BeatSheet, err error) {
	return nil, err
}

func (m *HappyPathClient) UpdateBeatSheet(id string, sheet *model.BeatSheet) (*model.BeatSheet, error) {
	return nil, nil
}

func (m *HappyPathClient) DeleteBeatSheet(id string) (err error) {
	return err
}

func (m *HappyPathClient) RetrieveBeatSheets() (beatSheets []model.BeatSheet, err error) {
	return nil, err
}

func (m *HappyPathClient) AddBeat(beatSheetId string, beat model.Beat) (beatId string, err error) {
	return "", nil
}

func (m *HappyPathClient) UpdateBeat(beatSheetId string, beatId string, beat model.Beat) (*model.Beat, error) {
	return nil, nil
}

func (m *HappyPathClient) DeleteBeat(beatSheetId string, beatId string) error {
	return nil
}

func (m *HappyPathClient) AddAct(beatSheetId string, beatId string, act model.Act) (actId string, err error) {
	return "123", nil
}

func (m *HappyPathClient) UpdateAct(beatSheetId string, beatId string, actId string, act model.Act) (*model.Act, error) {
	return &model.Act{
		Id:          "123",
		Description: "des",
		Timestamp:   time.Time{},
		Duration:    100,
		CameraAngle: "Up",
	}, nil
}

func (m *HappyPathClient) DeleteAct(beatSheetId string, beatId string, actId string) error {
	return nil
}
