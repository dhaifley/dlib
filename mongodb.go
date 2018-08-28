package dlib

import mgo "gopkg.in/mgo.v2"

// MongoDBIter types represent the functionality of a MongoDB iterator.
type MongoDBIter interface {
	Close() error
	Done() bool
	Next(result interface{}) bool
	Err() error
}

// MongoIter values implement the MongoDBIter interface.
type MongoIter struct {
	*mgo.Iter
}

// MongoDBQuery types represent the functionality of a MongoDB query.
type MongoDBQuery interface {
	Sort(fields ...string) MongoDBQuery
	Limit(n int) MongoDBQuery
	Iter() MongoDBIter
}

// MongoQuery values implement the MongoDBQuery interface.
type MongoQuery struct {
	*mgo.Query
}

// Sort shadows *mgo Sort function.
func (q MongoQuery) Sort(fields ...string) MongoDBQuery {
	return &MongoQuery{Query: q.Query.Sort(fields...)}
}

// Limit shadows *mgo Limit function.
func (q MongoQuery) Limit(n int) MongoDBQuery {
	return &MongoQuery{Query: q.Query.Limit(n)}
}

// Iter shadows *mgo Iter function.
func (q MongoQuery) Iter() MongoDBIter {
	return &MongoIter{Iter: q.Query.Iter()}
}

// MongoDBCollection types represent the functionality of a MongoDB collection.
type MongoDBCollection interface {
	Find(query interface{}) MongoDBQuery
	FindID(id interface{}) MongoDBQuery
	Insert(docs ...interface{}) error
	RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error)
	RemoveID(id interface{}) error
	EnsureIndex(index mgo.Index) error
}

// MongoCollection wraps a mgo.Collection to embed methods in models.
type MongoCollection struct {
	*mgo.Collection
}

// Find shadows *mgo Find function.
func (c MongoCollection) Find(query interface{}) MongoDBQuery {
	return &MongoQuery{Query: c.Collection.Find(query)}
}

// FindID shadows *mgo FindId function.
func (c MongoCollection) FindID(id interface{}) MongoDBQuery {
	return &MongoQuery{Query: c.Collection.FindId(id)}
}

// RemoveID shadows *mgo RemoveId function.
func (c MongoCollection) RemoveID(id interface{}) error {
	return c.Collection.RemoveId(id)
}

// MongoDBDatabase types are used to access a MongoDB database.
type MongoDBDatabase interface {
	C(name string) MongoDBCollection
}

// MongoDatabase wraps a mgo.Database to embed methods in models.
type MongoDatabase struct {
	*mgo.Database
}

// C shadows *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
func (d MongoDatabase) C(name string) MongoDBCollection {
	return &MongoCollection{Collection: d.Database.C(name)}
}

// MongoDBSession types are used to access a MongoDB session.
type MongoDBSession interface {
	DB(name string) MongoDBDatabase
	Close()
	Ping() error
}

// MongoSession values are MongoDB sessions.
type MongoSession struct {
	*mgo.Session
}

// DB shadows *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
func (s MongoSession) DB(name string) MongoDBDatabase {
	return &MongoDatabase{Database: s.Session.DB(name)}
}
