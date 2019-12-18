package db

import (
	"context"

	"github.com/PhamDuyKhang/userplayboar/internal/app/conf"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	//ServiceConnection hold all out side connection of service
	ServiceConnection struct {
		MongoDBClient *mongo.Client
	}
)

//EstablishInfrastructure start all connection of service
func EstablishInfrastructure(cf *conf.Config) (*ServiceConnection, error) {
	if cf.Infrastructure.MongoDB.Enable {
		logger.Debug("connecting mongo db")
		c, err := ConnectMongo(cf)
		if err != nil {
			return nil, err
		}
		ensureIndexes(c, cf)
		return &ServiceConnection{
			MongoDBClient: c,
		}, nil
	}
	return nil, nil
}
func ensureIndexes(c *mongo.Client, cf *conf.Config) {
	index := mongo.IndexModel{
		Keys:    bson.M{"emp_id": 1},
		Options: options.Index().SetUnique(true),
	}
	c.Database(cf.Infrastructure.MongoDB.DatabaseName).
		Collection("employee").
		Indexes().
		CreateOne(context.Background(), index)
	indexDepartment := mongo.IndexModel{
		Keys:    bson.M{"id": 1},
		Options: options.Index().SetUnique(true),
	}
	c.Database(cf.Infrastructure.MongoDB.DatabaseName).
		Collection("departments").
		Indexes().
		CreateOne(context.Background(), indexDepartment)
	indexSkill := mongo.IndexModel{
		Keys:    bson.M{"skill_id": 1},
		Options: options.Index().SetUnique(true),
	}
	c.Database(cf.Infrastructure.MongoDB.DatabaseName).
		Collection("skills").
		Indexes().
		CreateOne(context.Background(), indexSkill)
}
