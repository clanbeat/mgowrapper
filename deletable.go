package mgowrapper

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (m *Mongo) DeleteAll(collectionName string, params map[string]interface{}) error {
	_, err := m.C(collectionName).UpdateAll(params, bson.M{"$set": bson.M{"deletedAt": time.Now()}})
	return err
}
