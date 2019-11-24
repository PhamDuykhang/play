package db

import (
	"context"

	"github.com/PhamDuyKhang/userplayboar/internal/app/conf"
	"github.com/PhamDuyKhang/userplayboar/internal/app/pkg/glog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var logger = glog.New().WithPrefix("db")

//ConnectMongo to connect mongodb
func ConnectMongo(cf *conf.Config) (*mongo.Client, error) {
	ops := &options.ClientOptions{
		Hosts: cf.Infrastructure.MongoDB.Address,
		Auth: &options.Credential{
			Username:      cf.Infrastructure.MongoDB.Username,
			Password:      cf.Infrastructure.MongoDB.Password,
			AuthMechanism: "SCRAM-SHA-1",
			AuthSource:    cf.Infrastructure.MongoDB.DatabaseName,
		},
		ReadPreference: readpref.SecondaryPreferred(),
	}
	c, err := mongo.NewClient(ops)
	if err != nil {
		return nil, err
	}
	if err := c.Connect(context.Background()); err != nil {
		return nil, err
	}

	logger.WithField("database", cf.Infrastructure.MongoDB.DatabaseName).Info("connecting to mongo db is successfully ")
	return c, nil

}
