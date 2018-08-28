package dlib

import (
	"net/url"
	"testing"
	"time"

	"reflect"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FakeMongoDBIterLog struct {
	counter int
}

func (fk *FakeMongoDBIterLog) Close() error {
	return nil
}

func (fk *FakeMongoDBIterLog) Done() bool {
	return false
}

func (fk *FakeMongoDBIterLog) Next(result interface{}) bool {
	fk.counter++
	d := time.Date(1983, 2, 2, 0, 0, 0, 0, time.Local)
	l := Log{
		ID: bson.ObjectIdHex("4d88e15b60f486e428412dc9"),
		TS: &d,
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(l))

	if fk.counter > 1 {
		return false
	}

	return true
}

func (fk *FakeMongoDBIterLog) Err() error {
	return nil
}

type FakeMongoDBQueryLog struct{}

func (fk *FakeMongoDBQueryLog) Sort(fields ...string) MongoDBQuery {
	return fk
}

func (fk *FakeMongoDBQueryLog) Limit(n int) MongoDBQuery {
	return fk
}

func (fk *FakeMongoDBQueryLog) Iter() MongoDBIter {
	return &FakeMongoDBIterLog{}
}

type FakeMongoDBCollectionLog struct{}

func (fk *FakeMongoDBCollectionLog) Find(query interface{}) MongoDBQuery {
	return &FakeMongoDBQueryLog{}
}

func (fk *FakeMongoDBCollectionLog) FindID(id interface{}) MongoDBQuery {
	return &FakeMongoDBQueryLog{}
}

func (fk *FakeMongoDBCollectionLog) Insert(docs ...interface{}) error {
	return nil
}

func (fk *FakeMongoDBCollectionLog) RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error) {
	return &mgo.ChangeInfo{Removed: 1}, nil
}

func (fk *FakeMongoDBCollectionLog) RemoveID(id interface{}) error {
	return nil
}

func (fk *FakeMongoDBCollectionLog) EnsureIndex(index mgo.Index) error {
	return nil
}

type FakeMongoDBDatabaseLog struct{}

func (fk *FakeMongoDBDatabaseLog) C(name string) MongoDBCollection {
	return &FakeMongoDBCollectionLog{}
}

func TestLogEquals(t *testing.T) {
	cases := []struct {
		a        Log
		b        Log
		expected bool
	}{
		{
			a: Log{
				ID:    bson.ObjectIdHex("4d88e15b60f486e428412dc9"),
				Entry: "testentry",
			},
			b: Log{
				ID:    bson.ObjectIdHex("4d88e15b60f486e428412dc9"),
				Entry: "testentry",
			},
			expected: true,
		},
		{
			a: Log{
				ID:    bson.ObjectIdHex("4d88e15b60f486e428412dc9"),
				Entry: "testentry",
			},
			b: Log{
				ID:    bson.ObjectIdHex("4d88e15b60f486e428412dc9"),
				Entry: "testentry2",
			},
			expected: false,
		},
	}

	for _, c := range cases {
		result := c.a.Equals(c.b)
		if result != c.expected {
			t.Errorf("Expected bool: %v, got: %v", c.expected, result)
		}
	}
}

func TestLogString(t *testing.T) {
	d := time.Date(1983, 2, 2, 0, 0, 0, 0, time.Local)
	a := Log{
		ID:    bson.ObjectIdHex("4d88e15b60f486e428412dc9"),
		TS:    &d,
		Entry: "testentry",
	}

	expected := `{"id":"4d88e15b60f486e428412dc9","ts":"1983-02-02T00:00:00-05:00","entry":"testentry"}`
	result := a.String()
	if result != expected {
		t.Errorf("Expected string: %s, got: %s", expected, result)
	}
}

func TestLogSave(t *testing.T) {
	a := Log{ID: bson.ObjectIdHex("4d88e15b60f486e428412dc9")}
	c := a.Save(&FakeMongoDBDatabaseLog{})
	for r := range c {
		if r.Err != nil {
			t.Error(r.Err)
		}
	}

	expected := bson.ObjectIdHex("4d88e15b60f486e428412dc9")
	if a.ID != expected {
		t.Errorf("ID expected: %q, got: %q", expected, a.ID)
	}
}

func TestLogFromQueryValues(t *testing.T) {
	vals := url.Values{}
	vals.Add("id", "4d88e15b60f486e428412dc9")
	dv := Log{}
	dv.FromQueryValues(vals)
	expected := bson.ObjectIdHex("4d88e15b60f486e428412dc9")
	if dv.ID != expected {
		t.Errorf("ID expected: %q, got: %q", expected, dv.ID)
	}
}

func TestNewLogAccessor(t *testing.T) {
	la := NewLogAccessor(&FakeMongoDBDatabaseLog{})
	_, ok := la.(LogAccessor)
	if !ok {
		t.Errorf("Type expected: LogAccessor, got: %T", la)
	}
}

func TestLogAccessGetLogs(t *testing.T) {
	af := bson.M{"_id": "4d88e15b60f486e428412dc9"}
	la := NewLogAccessor(&FakeMongoDBDatabaseLog{})

	var as []*Log
	c := la.GetLogs(af, 10)
	for r := range c {
		if r.Err != nil {
			t.Error(r.Err)
		}

		switch v := r.Val.(type) {
		case *Log:
			as = append(as, v)
		default:
			t.Errorf("Invalid data type returned")
		}
	}

	expected := bson.ObjectIdHex("4d88e15b60f486e428412dc9")
	if as[0].ID != expected {
		t.Errorf("ID expected: %q, got: %q", expected, as[0].ID)
	}
}

func TestLogAccessGetLogByID(t *testing.T) {
	id := bson.ObjectIdHex("4d88e15b60f486e428412dc9")
	la := NewLogAccessor(&FakeMongoDBDatabaseLog{})

	var a *Log
	c := la.GetLogByID(id)
	for r := range c {
		if r.Err != nil {
			t.Error(r.Err)
		}

		switch v := r.Val.(type) {
		case *Log:
			a = v
		default:
			t.Errorf("Invalid data type returned")
		}
	}

	expected := bson.ObjectIdHex("4d88e15b60f486e428412dc9")
	if a.ID != expected {
		t.Errorf("ID expected: %q, got: %q", expected, a.ID)
	}
}

func TestLogAccessSaveLog(t *testing.T) {
	a := Log{ID: bson.ObjectIdHex("4d88e15b60f486e428412dc9")}
	la := NewLogAccessor(&FakeMongoDBDatabaseLog{})

	c := la.SaveLog(&a)
	for r := range c {
		if r.Err != nil {
			t.Error(r.Err)
		}
	}

	expected := bson.ObjectIdHex("4d88e15b60f486e428412dc9")
	if a.ID != expected {
		t.Errorf("ID expected: %q, got: %q", expected, a.ID)
	}
}
