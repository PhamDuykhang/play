package db

import (
	"github.com/PhamDuyKhang/userplayboar/internal/app/conf"

	"go.mongodb.org/mongo-driver/mongo"
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
		return &ServiceConnection{
			MongoDBClient: c,
		}, nil
	}
	return nil, nil
}
