package dlib

import (
	"encoding/json"
	"net/url"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Log values represent entries in the log database.
type Log struct {
	ID    bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	TS    *time.Time    `json:"ts,string,omitempty" bson:"ts,omitempty"`
	Entry interface{}   `json:"entry,omitempty" bson:"entry,omitempty"`
}

// LogAccess values are used for accessing log entries in the database.
type LogAccess struct {
	DB MongoDBDatabase
}

// LogAccessor is an interface describing types capable of providing
// access to log records in the database.
type LogAccessor interface {
	GetLogs(filter interface{}, limit int) <-chan Result
	GetLogByID(id bson.ObjectId) <-chan Result
	SaveLog(l *Log) <-chan Result
}

// Equals tests for deep equality between log values.
func (l Log) Equals(b Log) bool {
	if &l == &b {
		return true
	}
	if l.ID != b.ID {
		return false
	}
	if l.Entry != b.Entry {
		return false
	}
	if (l.TS == nil && b.TS != nil) || (l.TS != nil && b.TS == nil) {
		return false
	}
	if l.TS != nil && b.TS != nil && *l.TS != *b.TS {
		return false
	}

	return true
}

// String formats a log entry as a JSON format string.
func (l Log) String() string {
	str, err := json.Marshal(l)
	if err != nil {
		return ""
	}

	return string(str)
}

// Save persists the log entry to the database.
// Save ensures that primary key uniqueness is preserved.
// It will return a Result channel.
func (l *Log) Save(dbs MongoDBDatabase) <-chan Result {
	c := make(chan Result, 256)

	go func() {
		defer close(c)
		if l.ID == bson.ObjectId("") || !l.ID.Valid() {
			l.ID = bson.NewObjectId()
		} else {
			err := dbs.C("logs").RemoveID(l.ID)
			if err != nil {
				c <- Result{Err: err}
				return
			}
		}

		err := dbs.C("logs").Insert(l)
		if err != nil {
			c <- Result{Err: err}
			return
		}

		c <- Result{Val: l, Err: nil}
	}()

	return c
}

// FromQueryValues populates a point value from a query string values map.
func (l *Log) FromQueryValues(vals url.Values) error {
	if vals.Get("id") != "" {
		l.ID = bson.ObjectIdHex(vals.Get("id"))
	}
	if vals.Get("ts") != "" {
		pt, err := time.ParseInLocation("2006-01-02", vals.Get("ts"), time.Local)
		if err != nil {
			return err
		}
		l.TS = &pt
	}

	return nil
}

// NewLogAccessor creates a new LogAccess instance and returns
// a pointer to it.
func NewLogAccessor(dbs MongoDBDatabase) LogAccessor {
	la := LogAccess{DB: dbs}
	return &la
}

// GetLogs finds log entries in the database.
// It searches using a filtering function or a generic invoice interface value.
// It returns the results of the operation in an Result channel.
func (la *LogAccess) GetLogs(filter interface{}, limit int) <-chan Result {
	c := make(chan Result, 256)

	go func() {
		defer close(c)
		q := la.DB.C("logs").Find(filter).Sort("-$natural")
		if limit > 0 {
			q = q.Limit(limit)
		}

		cur := q.Iter()
		defer cur.Close()

		if cur.Done() {
			c <- Result{
				Err: &Error{Code: 404, Msg: "Resource not found"},
			}
			return
		}

		var l Log
		for cur.Next(&l) {
			lv := l
			c <- Result{Val: &lv, Err: nil}
		}

		if cur.Err() != nil {
			c <- Result{Err: cur.Err()}
		}
	}()

	return c
}

// GetLogByID finds a log entry in the database by ID.
// It returns the results of the operation in an Result channel.
func (la *LogAccess) GetLogByID(id bson.ObjectId) <-chan Result {
	c := make(chan Result, 256)

	go func() {
		defer close(c)
		cur := la.DB.C("logs").FindID(id).Iter()
		defer cur.Close()

		if cur.Done() {
			c <- Result{
				Err: &Error{Code: 404, Msg: "Resource not found"},
			}
			return
		}

		var l Log
		cur.Next(&l)
		c <- Result{Val: &l, Err: nil}
	}()

	return c
}

// SaveLog saves a log entry to the database.
// It returns the results of the operation in an Result channel.
func (la *LogAccess) SaveLog(l *Log) <-chan Result {
	return l.Save(la.DB)
}
