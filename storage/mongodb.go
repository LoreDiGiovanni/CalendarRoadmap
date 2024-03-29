package storage

import (
	"RoadmapCalendar/types"
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStorage struct {
	db *mongo.Client
}

func NewMongoStore() (*mongoStorage, error) {
	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	} else {
		return &mongoStorage{db: client}, nil
	}
}

func StringIdToObjectId(id string) (*primitive.ObjectID,error){
    objectID, err := primitive.ObjectIDFromHex(id); if err != nil {
		return nil,errors.New("Invalid ID") 
	}else{
        return &objectID,nil
    }
}

func (s mongoStorage) PostEvents( event types.Events) error {
	coll := s.db.Database("RoadmapCalendar").Collection("events")
	_, err := coll.InsertOne(context.TODO(), event)
	if err != nil {
		return err
	} else {
		return nil
	}
}
func (s mongoStorage) PostUser(user types.User) error{
    coll := s.db.Database("RoadmapCalendar").Collection("users")
	_, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (s mongoStorage) GetUser(user types.User) (*types.User,error){
    coll := s.db.Database("RoadmapCalendar").Collection("users")
    query  := bson.D{{"email",user.Email}}
    var dbuser types.User
	err := coll.FindOne(context.TODO(), query).Decode(&dbuser);
    if err!= nil{
        return nil,err
    }else{
        return &dbuser,nil
    }
}

func (s mongoStorage) GetEvents(user string) (*[]types.Events, error) {
	coll := s.db.Database("RoadmapCalendar").Collection("events")
    var results []types.Events
    query := bson.D{{"owner",user}}
	cursor, err := coll.Find(context.TODO(), query)
    cursor.All(context.TODO(), &results)
	return &results,err
}

