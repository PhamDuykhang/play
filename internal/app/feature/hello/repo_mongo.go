package hello

import (
	"context"

	"github.com/PhamDuyKhang/userplayboar/internal/app/pkg/glog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var logger = glog.New().WithPrefix("hello")

type (
	//Richer is a implementation with mongo db style
	Richer struct {
		Cl *mongo.Client
	}
)

//NewRicher return a hello repository
func NewRicher(cl *mongo.Client) *Richer {
	return &Richer{
		Cl: cl,
	}
}

//SayHello return a phase
func (r *Richer) SayHello(mgs string) string {
	rs, err := r.Cl.Database("play").Collection("hello").InsertOne(context.Background(), bson.M{"k": mgs})
	if err != nil {
		logger.Error("can't insert message", err)
		return ""
	}
	logger.Info(rs.InsertedID)
	return mgs
}
