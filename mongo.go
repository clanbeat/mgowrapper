package mgowrapper

import (
	"errors"
	"gopkg.in/mgo.v2"
)

type (
	Mongo struct {
		Session *mgo.Session
	}

	CollectionIndex struct {
		Name    string
		Indexes []mgo.Index
	}
)

const dbName = ""

func New(dbURL string) (*Mongo, error) {
	if len(dbURL) == 0 {
		return nil, errors.New("mongo db url missing")
	}
	s, err := mgo.Dial(dbURL)
	if err != nil {
		return nil, err
	}
	s.SetSafe(&mgo.Safe{})
	return &Mongo{Session: s}, nil
}

func (m *Mongo) Close() {
	m.Session.Close()
}

func (m *Mongo) C(collectionName string) *mgo.Collection {
	return m.Session.DB(dbName).C(collectionName)
}

func (m *Mongo) IsDown() bool {
	if err := m.Session.Ping(); err != nil {
		return true
	}
	return false
}

func (m *Mongo) EnsureIndex(colName string, ci mgo.Index) error {
	err := m.C(colName).EnsureIndex(ci)
	if err != nil {
		return err
	}
	return nil
}

func (m *Mongo) EnsureIndexes(ci CollectionIndex) error {
	for _, i := range ci.Indexes {
		if err := m.EnsureIndex(ci.Name, i); err != nil {
			return err
		}
	}
	return nil
}
