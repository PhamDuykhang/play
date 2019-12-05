package usercrud

import (
	"context"

	"github.com/PhamDuyKhang/userplayboar/internal/app/pkg/glog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var logger = glog.New().WithPrefix("hello")

type (
	//MongoDB is a implementation's UserManager IF with mongo db style
	MongoDB struct {
		Cl *mongo.Client
	}
)

//NewCrudMongo create instance of mongo db repository
func NewCrudMongo(c *mongo.Client) *MongoDB {
	return &MongoDB{
		Cl: c,
	}
}

//InsertUser add new user into database
func (r *MongoDB) InsertUser(ctx context.Context, e Employee) (emp Employee, err error) {
	logger.Debugc(ctx, "inserting employee")

	_, err = r.Cl.Database("play").Collection("employee").InsertOne(context.Background(), e)
	if err != nil {
		logger.Errorc(ctx, "can't insert employee")
		return
	}
	err = r.Cl.Database("play").Collection("employee").FindOne(ctx, bson.M{"emp_id": e.EmpID}).Decode(&emp)
	if err != nil {
		return emp, err
	}
	return emp, nil
}

//UpdateUser add new user into database
func (r *MongoDB) UpdateUser(ctx context.Context, e Employee) (emp Employee, err error) {
	logger.Debugc(ctx, "updating employee")
	filter := bson.M{
		"emp_id": e.EmpID,
	}
	e.EmpID = ""
	update := bson.M{
		"$set": e,
	}
	ops := options.FindOneAndUpdateOptions{}
	ops.SetReturnDocument(options.After)
	upResult := r.Cl.Database("play").Collection("employee").FindOneAndUpdate(context.Background(), filter, update, &ops)
	if err != nil {
		logger.Errorc(ctx, "can't update user", err)
		return emp, err
	}
	err = upResult.Decode(&emp)
	if err != nil {
		return emp, err
	}
	return emp, nil

}

//DeleteUser user form database
func (r *MongoDB) DeleteUser(ctx context.Context, emID string) (err error) {
	logger.Debugc(ctx, "deleting employee")
	rs := r.Cl.Database("play").Collection("employee").FindOneAndDelete(context.Background(), bson.M{"emp_id": emID})
	if rs.Err() != nil {
		logger.Errorc(ctx, "can't delete user", err)
		return rs.Err()
	}
	return
}

//Find user from database
func (r *MongoDB) Find(ctx context.Context, emID string) (emp Employee, err error) {
	logger.Debugc(ctx, "finding employee")
	rs := r.Cl.Database("play").Collection("employee").FindOne(ctx, bson.M{"emp_id": emID})
	var e Employee
	if rs.Err() != nil {
		logger.Errorc(ctx, "can't find employee %s", rs.Err().Error())
		return e, rs.Err()
	}
	err = rs.Decode(&e)
	if err != nil {
		logger.Errorc(ctx, "can't decode result")
		return e, err
	}
	return e, nil
}

//FindAll get all document still not suport sorting and pagination
func (r *MongoDB) FindAll(ctx context.Context) (emps []Employee, err error) {
	rs, err := r.Cl.Database("play").Collection("employee").Find(ctx, bson.M{}, nil)
	if err != nil {
		return emps, err
	}
	logger.Infoc(ctx, "finding employee")
	err = rs.All(ctx, &emps)
	if err != nil {
		return emps, err
	}
	return emps, nil
}
