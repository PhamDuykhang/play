package organization

import (
	"context"

	"github.com/palantir/stacktrace"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	//DeliveryCenterCollections is name of dc colletions
	DeliveryCenterCollections = "deliverycenter"
	//DeliveryGroupCollections is name of dg colletions
	DeliveryGroupCollections = "deliverygroup"
	//SkillCollections is name of skill colletions
	SkillCollections = "skills"
	//DepartmentCollections the collection name where we store department data
	DepartmentCollections = "departments"
)

type (
	//MongoDB is a implementation's RepoI IF with mongo db style
	MongoDB struct {
		Cl *mongo.Client
	}
)

//NewOrganizationMongo create a instance for mongo db interactive
func NewOrganizationMongo(c *mongo.Client) *MongoDB {
	return &MongoDB{
		Cl: c,
	}
}

//InsertDepartment add new department to database
func (r *MongoDB) InsertDepartment(ctx context.Context, d Organization) (Organization, error) {
	_, err := r.Cl.Database("play").Collection(DepartmentCollections).InsertOne(ctx, d)
	if err != nil {
		logger.Errorc(ctx, "Insert department got err: %v", err)
		err := stacktrace.Propagate(err, "inserting department got err")
		return Organization{}, err
	}
	rs := r.Cl.Database("play").Collection(DepartmentCollections).FindOne(ctx, bson.M{"id": d.ID})
	if rs.Err() != nil {
		logger.Errorc(ctx, "Get department after inserting error")
		err := stacktrace.Propagate(rs.Err(), "can't get department")
		return Organization{}, err
	}
	var o Organization
	err = rs.Decode(&o)
	if err != nil {
		logger.Errorc(ctx, "Decode got err: %v", err)
		return o, stacktrace.Propagate(err, "can't get decode department")
	}
	return o, nil
}

//FindChildrentByParentID the childrent department
func (r *MongoDB) FindChildrentByParentID(ctx context.Context, pID string) ([]Organization, error) {
	var ors []Organization
	rs, err := r.Cl.Database("play").Collection(DepartmentCollections).Find(ctx, bson.M{"parent_id": pID})
	if err != nil {
		logger.Errorc(ctx, "Get child department base on parent id error")
		err := stacktrace.Propagate(err, "can't get department")
		return ors, err
	}
	err = rs.All(ctx, &ors)
	if err != nil {
		logger.Errorc(ctx, "can't decode result err:%v", err)
		err := stacktrace.Propagate(err, "can't get child department")
		return ors, err
	}
	return ors, nil
}

//FindDepartmentByID the department base on id
func (r *MongoDB) FindDepartmentByID(ctx context.Context, id string) (Organization, error) {
	var ors Organization
	rs := r.Cl.Database("play").Collection(DepartmentCollections).FindOne(ctx, bson.M{"id": id})
	if rs.Err() != nil {
		logger.Errorc(ctx, "Get department base id error %v", rs.Err())
		err := stacktrace.Propagate(rs.Err(), "can't get department")
		return ors, err
	}
	err := rs.Decode(&ors)
	if err != nil {
		logger.Errorc(ctx, "can't decode result err:%v", err)
		err := stacktrace.Propagate(err, "can't get department")
		return ors, err
	}
	return ors, nil
}

//UpdateDepartment update department information
func (r *MongoDB) UpdateDepartment(ctx context.Context, d Organization) (Organization, error) {
	filter := bson.M{
		"id": d.ID,
	}
	d.ID = ""
	update := bson.M{
		"$set": d,
	}
	after := options.After
	opt := &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	var or Organization
	sr := r.Cl.Database("play").Collection(DepartmentCollections).FindOneAndUpdate(ctx, filter, update, opt)
	if sr.Err() != nil {
		return or, stacktrace.Propagate(sr.Err(), "can't get update department")
	}
	err := sr.Decode(&or)
	if err != nil {
		logger.Errorc(ctx, "can't decode result err:%v", err)
		err := stacktrace.Propagate(err, "can't get department")
		return or, err
	}
	return or, nil
}

//InsertSkill add new skill to mongo database
func (r *MongoDB) InsertSkill(ctx context.Context, sk Skill) (Skill, error) {
	var skill Skill
	_, err := r.Cl.Database("play").Collection(SkillCollections).InsertOne(ctx, sk)
	if err != nil {
		err := errors.Wrap(err, " can't insert data to data base")
		return skill, err
	}
	rs := r.Cl.Database("play").Collection(SkillCollections).FindOne(ctx, bson.M{"skill_id": sk.SkillID})
	if rs.Err() != nil {
		return skill, errors.Wrap(err, " can't get new data from data base")
	}
	err = rs.Decode(&skill)
	if err != nil {
		return skill, errors.Wrap(err, " can't decode new data from data base")
	}
	return skill, nil
}

//FindAllSkill find all skill we have in mongo database
func (r *MongoDB) FindAllSkill(ctx context.Context) ([]Skill, error) {
	rs, err := r.Cl.Database("play").Collection(SkillCollections).Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, " can't get data from data base")
	}
	var skills []Skill
	err = rs.All(ctx, &skills)
	if err != nil {
		return nil, errors.Wrap(err, " can't decode new data from data base")
	}
	return skills, nil
}
