package data

import (
	"BeatSheet/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var mClient *MongoClient

type MongoClient struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewMongoClient(ctx context.Context) *MongoClient {
	if mClient != nil {
		return mClient
	}
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	mClient = &MongoClient{
		ctx:        ctx,
		collection: client.Database("beatsheetdb").Collection("beatsheets"),
	}
	return mClient
}

func (m *MongoClient) CreateBeatSheet(beatSheet model.BeatSheet) (string, error) {
	_, err := m.collection.InsertOne(m.ctx, beatSheet)

	if err != nil {
		return "", err
	}

	return beatSheet.Id, nil
}
func (m *MongoClient) RetrieveBeatSheet(id string) (beatSheet *model.BeatSheet, err error) {
	result := m.collection.FindOne(m.ctx, bson.M{"id": id})
	result.Decode(&beatSheet)
	return
}

func (m *MongoClient) UpdateBeatSheet(id string, sheet *model.BeatSheet) (*model.BeatSheet, error) {
	filter := bson.D{{"id", id}}
	update := bson.M{
		"$set": sheet,
	}
	_, err := m.collection.UpdateOne(m.ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return sheet, nil
}

func (m *MongoClient) DeleteBeatSheet(id string) (err error) {
	_, err = m.collection.DeleteOne(m.ctx, bson.M{"id": id})
	if err != nil {
		log.Println(err)
	}
	return
}

func (m *MongoClient) RetrieveBeatSheets() (beatSheets []model.BeatSheet, err error) {
	var results *mongo.Cursor
	results, err = m.collection.Find(m.ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return
	}
	if err = results.All(m.ctx, &beatSheets); err != nil {
		log.Println(err)
		return
	}

	if beatSheets == nil {
		return []model.BeatSheet{}, nil
	}

	return
}

func (m *MongoClient) AddBeat(beatSheetId string, beat model.Beat) (beatId string, err error) {
	filter := bson.D{{"id", beatSheetId}}
	update := bson.D{
		{"$push", bson.D{{"beats", beat}}},
	}
	_, err = m.collection.UpdateOne(m.ctx, filter, update)
	if err != nil {
		return
	}
	return beat.Id, nil
}

func (m *MongoClient) UpdateBeat(beatSheetId string, beatId string, beat model.Beat) (*model.Beat, error) {
	filter := bson.M{
		"id":       beatSheetId,
		"beats.id": beatId,
	}

	update := bson.M{
		"$set": bson.M{
			"beats.$": beat,
		},
	}
	_, err := m.collection.UpdateOne(m.ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return &beat, nil
}

func (m *MongoClient) DeleteBeat(beatSheetId string, beatId string) error {
	filter := bson.M{
		"id": beatSheetId,
	}

	update := bson.M{
		"$pull": bson.M{
			"beats": bson.M{
				"id": beatId,
			},
		},
	}

	_, err := m.collection.UpdateOne(m.ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoClient) AddAct(beatSheetId string, beatId string, act model.Act) (actId string, err error) {
	filter := bson.M{
		"id":       beatSheetId,
		"beats.id": beatId,
	}
	update := bson.M{
		"$push": bson.M{"beats.$.acts": act},
	}

	_, err = m.collection.UpdateOne(m.ctx, filter, update)
	if err != nil {
		return
	}
	return act.Id, nil
}

func (m *MongoClient) UpdateAct(beatSheetId string, beatId string, actId string, act model.Act) (*model.Act, error) {
	filter := bson.M{
		"id": beatSheetId,
	}
	upsert := true
	arrayFilters := options.ArrayFilters{Filters: bson.A{bson.M{"x.id": beatId}, bson.M{"y.id": actId}}}
	opts := options.UpdateOptions{
		ArrayFilters: &arrayFilters,
		Upsert:       &upsert,
	}
	update := bson.M{
		"$set": bson.M{
			"beats.$[x].acts.$[y]": act,
		},
	}
	_, err := m.collection.UpdateOne(m.ctx, filter, update, &opts)
	if err != nil {
		return nil, err
	}
	return &act, nil
}

func (m *MongoClient) DeleteAct(beatSheetId string, beatId string, actId string) error {
	filter := bson.M{
		"id":       beatSheetId,
		"beats.id": beatId,
	}

	upsert := true
	arrayFilters := options.ArrayFilters{Filters: bson.A{bson.M{"x.id": beatId}}}
	opts := options.UpdateOptions{
		ArrayFilters: &arrayFilters,
		Upsert:       &upsert,
	}

	update := bson.M{
		"$pull": bson.M{
			"beats.$[x].acts": bson.M{
				"id": actId,
			},
		},
	}
	_, err := m.collection.UpdateOne(m.ctx, filter, update, &opts)
	if err != nil {
		return err
	}
	return nil
}
