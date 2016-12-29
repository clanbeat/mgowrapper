package mgowrapper

import (
	"gopkg.in/mgo.v2/bson"
)

type Updatable interface {
	CollectionName() string
	ItemID() bson.ObjectId
}

func (m *Mongo) Update(u Updatable) error {
	return m.C(u.CollectionName()).Update(bson.M{"_id": u.ItemID()}, u)
}

func (m *Mongo) UpdateAll(collectionName string, params bson.M, vals bson.M) error {
	_, err := m.C(collectionName).UpdateAll(params, bson.M{"$set": vals})
	return err
}
