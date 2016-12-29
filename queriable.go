package mongo

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Queryable interface {
		CollectionName() string
		Load(data bson.M) error
	}

	QueryableSlice interface {
		CollectionName() string
		InitItem() Queryable
		Append(Queryable)
	}

	Result interface{}

	QueryFunc func(m *Mongo) Result
)

func (m *Mongo) Find(qfunc QueryFunc) Result {
	return qfunc(m)
}

func (m *Mongo) FindAndLoadOne(item Queryable, q bson.M, sort string) error {
	var data bson.M
	if err := m.C(item.CollectionName()).Find(q).Sort(sort).One(&data); err != nil {
		return err
	}
	err := item.Load(data)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mongo) FindAndLoadMultiple(sl QueryableSlice, q bson.M, sort string, limit int) error {
	var data []bson.M
	query := m.C(sl.CollectionName()).Find(q).Sort(sort)
	if limit > 0 {
		query = query.Limit(20)
	}
	if err := query.All(&data); err != nil {
		return err
	}
	for _, d := range data {
		i := sl.InitItem()
		if err := i.Load(d); err != nil {
			return err
		}
		sl.Append(i)
	}

	return nil
}
