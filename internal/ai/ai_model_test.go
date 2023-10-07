package ai

import (
	"BeatSheet/internal/model"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestAIModel_AddAct(t *testing.T) {
	type fields struct {
		acts []model.Act
	}
	type args struct {
		act model.Act
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "exists", fields: fields{acts: []model.Act{{}}}, args: args{act: model.Act{
			Id:          "12312312321312",
			Description: "first act",
			Timestamp:   time.Time{},
			Duration:    1000,
			CameraAngle: "Left",
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ai := &AIModel{
				acts: tt.fields.acts,
			}
			ai.AddAct(tt.args.act)
			if !containsAct(tt.fields.acts, tt.args.act) {
				t.Failed()
			}
		})
	}
}

func containsAct(s []model.Act, e model.Act) bool {
	for _, a := range s {
		if reflect.DeepEqual(a, e) {
			return true
		}
	}
	return false
}

func containsBeat(s []model.Beat, e model.Beat) bool {
	for _, a := range s {
		if reflect.DeepEqual(a, e) {
			return true
		}
	}
	return false
}

func TestAIModel_AddBeat(t *testing.T) {
	type fields struct {
		beats []model.Beat
	}
	type args struct {
		beat model.Beat
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		exists bool
	}{
		{name: "exists", fields: fields{beats: []model.Beat{{}}}, args: args{beat: model.Beat{
			Id:          "12312312321312",
			Description: "first beat",
			Timestamp:   time.Time{},
		}}, exists: true},
		{name: "exists", fields: fields{beats: []model.Beat{{}}}, args: args{beat: model.Beat{
			Id:          "12312312321312",
			Description: "first beat",
			Timestamp:   time.Time{},
		}}, exists: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ai := &AIModel{
				beats: tt.fields.beats,
			}
			if tt.exists {
				ai.AddBeat(tt.args.beat)
				if !containsBeat(ai.beats, tt.args.beat) {
					t.Error()
				}
			} else {
				if containsBeat(ai.beats, tt.args.beat) {
					t.Error()
				}
			}
		})
	}
}

func TestAIModel_GetNextActSuggestion(t *testing.T) {
	type fields struct {
		acts []model.Act
	}
	tests := []struct {
		name    string
		fields  fields
		want    model.Act
		wantErr bool
	}{
		{name: "8", fields: fields{acts: []model.Act{
			{
				Id:          "asjdlkasjdlkas",
				Description: "first act",
				Timestamp:   time.Time{},
				Duration:    100,
				CameraAngle: "Left",
			},
			{
				Id:          "asjdlkassf32423jdlkas",
				Description: "second act",
				Timestamp:   time.Time{},
				Duration:    200,
				CameraAngle: "Right",
			},
			{
				Id:          "sdfdsfssdfs",
				Description: "first act",
				Timestamp:   time.Time{},
				Duration:    100,
				CameraAngle: "Left",
			},
			{
				Id:          "cxvxwerewrwerwe",
				Description: "second act",
				Timestamp:   time.Time{},
				Duration:    200,
				CameraAngle: "Right",
			},
			{
				Id:          "asjdlkasjdlkas",
				Description: "first act",
				Timestamp:   time.Time{},
				Duration:    100,
				CameraAngle: "Left",
			},
			{
				Id:          "asjdlkassf32423jdlkas",
				Description: "third act",
				Timestamp:   time.Time{},
				Duration:    200,
				CameraAngle: "Right",
			},
			{
				Id:          "sdfdsfssdfs",
				Description: "first act",
				Timestamp:   time.Time{},
				Duration:    100,
				CameraAngle: "Left",
			},
			{
				Id:          "cxvxwerewrwerwe",
				Description: "third act",
				Timestamp:   time.Time{},
				Duration:    200,
				CameraAngle: "Right",
			},
		}}, want: model.Act{
			Id:          "asjdlkasjdlkas",
			Description: "first act",
			Timestamp:   time.Time{},
			Duration:    200,
			CameraAngle: "Left",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ai := &AIModel{
				acts: tt.fields.acts,
			}
			got, err := ai.GetNextActSuggestion()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNextActSuggestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !((got.Description == tt.want.Description) || (got.CameraAngle == tt.want.CameraAngle) || (got.Duration == tt.want.Duration)) {
				t.Errorf("GetNextActSuggestion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAIModel_GetNextBeatSuggestion(t *testing.T) {
	type fields struct {
		beats []model.Beat
	}
	tests := []struct {
		name    string
		fields  fields
		want    model.Beat
		wantErr bool
	}{
		{name: "8", fields: fields{beats: []model.Beat{
			{
				Id:          "asjdlkasjdlkas",
				Description: "first beat",
				Timestamp:   time.Time{},
			},
			{
				Id:          "asjdlkassf32423jdlkas",
				Description: "second beat",
				Timestamp:   time.Time{},
			},
			{
				Id:          "sdfdsfssdfs",
				Description: "first beat",
				Timestamp:   time.Time{},
			},
			{
				Id:          "cxvxwerewrwerwe",
				Description: "second beat",
				Timestamp:   time.Time{},
			},
			{
				Id:          "asjdlkasjdlkas",
				Description: "first beat",
				Timestamp:   time.Time{},
			},
			{
				Id:          "asjdlkassf32423jdlkas",
				Description: "third beat",
				Timestamp:   time.Time{},
			},
			{
				Id:          "sdfdsfssdfs",
				Description: "first beat",
				Timestamp:   time.Time{},
			},
			{
				Id:          "cxvxwerewrwerwe",
				Description: "third beat",
				Timestamp:   time.Time{},
			},
		}}, want: model.Beat{
			Id:          "asjdlkasjdlkas",
			Description: "first beat",
			Timestamp:   time.Time{},
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ai := &AIModel{
				beats: tt.fields.beats,
			}
			got, err := ai.GetNextBeatSuggestion()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNextBeatSuggestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got.Id = ""
			tt.want.Id = ""
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNextBeatSuggestion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAIModel_Train(t *testing.T) {
	type fields struct {
		beats []model.Beat
		acts  []model.Act
		mu    sync.Mutex
	}
	type args struct {
		beatsData []model.Beat
		actsData  []model.Act
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ai := &AIModel{
				beats: tt.fields.beats,
				acts:  tt.fields.acts,
				mu:    tt.fields.mu,
			}
			ai.Train(tt.args.beatsData, tt.args.actsData)
		})
	}
}

func TestNewAIModel(t *testing.T) {
	tests := []struct {
		name string
		want *AIModel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAIModel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAIModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
