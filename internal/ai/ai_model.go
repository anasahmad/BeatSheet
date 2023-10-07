package ai

import (
	"BeatSheet/internal/model"
	"errors"
	"github.com/rxwycdh/rxhash"
	"sync"
	"time"
)

var beatSheetModel *AIModel

type AIModelSpec interface {
	AddBeat(beat model.Beat)
	AddAct(act model.Act)
	GetNextBeatSuggestion() (model.Beat, error)
	GetNextActSuggestion() (model.Act, error)
}

type AIModel struct {
	beats []model.Beat
	acts  []model.Act
	mu    sync.Mutex
}

// NewAIModel creates a new instance of AIModel or return the instance which was already created.
func NewAIModel() *AIModel {
	if beatSheetModel == nil {
		beatSheetModel = &AIModel{
			beats: make([]model.Beat, 0),
			acts:  make([]model.Act, 0),
		}
	}
	return beatSheetModel
}

// Train method adds new beats to the AI model.
func (ai *AIModel) Train(beatsData []model.Beat, actsData []model.Act) {
	ai.mu.Lock()
	defer ai.mu.Unlock()

	// Append the new beats to the model's beats slice
	ai.beats = append(ai.beats, beatsData...)
	ai.acts = append(ai.acts, actsData...)
}

func (ai *AIModel) AddBeat(beat model.Beat) {
	ai.mu.Lock()
	defer ai.mu.Unlock()

	// Append the new beat to the model's beats slice
	ai.beats = append(ai.beats, beat)
}

func (ai *AIModel) AddAct(act model.Act) {
	ai.mu.Lock()
	defer ai.mu.Unlock()

	// Append the new act to the model's act slice
	ai.acts = append(ai.acts, act)
}

// GetNextBeatSuggestion  goes through the list and checks what's the beat which comes after the ones matching with the previous beat
// we can count which were the most & we can return that
// error if the latest beat was the only time it was used
// return nothing if the collection is empty
func (ai *AIModel) GetNextBeatSuggestion() (model.Beat, error) {
	ai.mu.Lock()
	defer ai.mu.Unlock()

	if len(ai.beats) == 0 {
		return model.Beat{}, errors.New("beats == 0 - no previous history of beats")
	}

	latestBeat := ai.beats[len(ai.beats)-1]

	// map of hash to recurrence count
	recurringMap := make(map[string]int)
	// map of hash to Beat
	hashToBeat := make(map[string]model.Beat)

	var mostChanceCount = -1
	var mostChanceHash string

	// Find the instances of the latest beat and find out what was the beat after it & if this occurs more than ones then
	// the beat coming after the latest the most times should be returned
	for i, beat := range ai.beats {
		if beat.Description == latestBeat.Description {
			if i+1 < len(ai.beats) {
				//id and time are always unique
				beat.Id = ""
				beat.Timestamp = time.Time{}
				h, _ := rxhash.HashStruct(ai.beats[i+1])
				hashToBeat[h] = ai.beats[i+1]
				recurringMap[h] = recurringMap[h] + 1
			}
		}
	}

	// find out which hash has the highest recurring count
	for k, v := range recurringMap {
		if v > mostChanceCount {
			mostChanceCount = v
			mostChanceHash = k
		}
	}

	// if there wasn't beat after the latest beat, return an error
	if mostChanceHash == "" {
		return model.Beat{}, errors.New("no previous history of beats")
	}

	return hashToBeat[mostChanceHash], nil
}

// GetNextActSuggestion  goes through the list and checks what's the act which comes after the ones matching with the previous act
// we can count which were the most & we can return that
// error if the latest act was the only time it was used
// return nothing if the collection is empty
func (ai *AIModel) GetNextActSuggestion() (model.Act, error) {
	ai.mu.Lock()
	defer ai.mu.Unlock()

	if len(ai.acts) == 0 {
		return model.Act{}, errors.New("acts == 0 - no previous history of acts")
	}

	latestAct := ai.acts[len(ai.acts)-1]

	// map of hash to recurrence count
	recurringMap := make(map[string]int)
	// map of hash to Act
	hashToAct := make(map[string]model.Act)

	var mostChanceCount = -1
	var mostChanceHash string

	// Find older occurrences of the latest Act (it's attributes) and check what came after it
	// for we will be adding points for each attribute map i.e. description, camera angle, duration
	for i, act := range ai.acts {
		if i+1 < len(ai.acts) {
			//id and time are always unique
			act.Id = ""
			act.Timestamp = time.Time{}
			h, _ := rxhash.HashStruct(ai.acts[i+1])
			if act.Description == latestAct.Description {
				hashToAct[h] = ai.acts[i+1]
				recurringMap[h] = recurringMap[h] + 1
			}

			if act.CameraAngle == latestAct.CameraAngle {
				hashToAct[h] = ai.acts[i+1]
				recurringMap[h] = recurringMap[h] + 1
			}

			if act.Duration == latestAct.Duration {
				hashToAct[h] = ai.acts[i+1]
				recurringMap[h] = recurringMap[h] + 1
			}
		}
	}

	// find out which hash has the highest recurring count
	for k, v := range recurringMap {
		if v > mostChanceCount {
			mostChanceCount = v
			mostChanceHash = k
		}
	}

	// if there was act after the latest act, return an error
	if mostChanceHash == "" {
		return model.Act{}, errors.New("no previous history of acts")
	}

	return hashToAct[mostChanceHash], nil
}
