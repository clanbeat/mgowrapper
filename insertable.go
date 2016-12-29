package mgowrapper

type Insertable interface {
	CollectionName() string
}

func (m *Mongo) Insert(i Insertable) error {
	return m.C(i.CollectionName()).Insert(i)
}
