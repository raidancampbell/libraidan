package rcollections

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewStrMap(t *testing.T) {
	s := NewStrMap()
	assert.NotNil(t, s)
}

func TestStrMap_Add(t *testing.T) {
	s := NewStrMap()
	s = s.Add("foo", "bar")
	assert.Equal(t, s.innerMap()["foo"], "bar")
}

func TestStrMap_AddAll(t *testing.T) {
	s := NewStrMap()
	s = s.Add("fookey", "fooval")
	s = s.Add("barkey", "barval")
	s = s.Add("quxkey", "quxval")
	a, err := s.Get("fookey")
	assert.Equal(t, "fooval", a)
	assert.Nil(t, err)

	o := NewStrMap()
	o = o.Add("fookey", "other foo key")
	o = o.Add("barkey", "barval")

	actual := s.AddAll(o)
	assert.NotNil(t, actual)

	val := actual.GetWithDefault("fookey", "")
	assert.Equal(t, "other foo key", val)

	val = actual.GetWithDefault("barkey", "")
	assert.Equal(t, "barval", val)

	val = actual.GetWithDefault("quxkey", "")
	assert.Equal(t, "quxval", val)
}

func TestStrMap_Contains(t *testing.T) {
	s := NewStrMap()
	s = s.Add("foo", "bar")
	assert.True(t, s.Contains("foo"))
	assert.False(t, s.Contains("qux"))
	assert.False(t, s.Contains(""))

	s = s.Add("quux", nil)
	assert.True(t, s.Contains("quux"))
}

func TestStrMap_ContainsValue(t *testing.T) {
	s := NewStrMap()
	s = s.Add("foo", "bar")
	assert.True(t, s.ContainsValue("bar"))
	assert.False(t, s.ContainsValue("qux"))
	assert.False(t, s.ContainsValue(nil))

	s = s.Add("quux", nil)
	assert.True(t, s.ContainsValue(nil))
}

func TestStrMap_Get(t *testing.T) {
	s := NewStrMap()
	s = s.Add("foo", "bar")
	actual, err := s.Get("foo")
	assert.Equal(t, "bar", actual)
	assert.Nil(t, err)

	actual, err = s.Get("qux")
	assert.Nil(t, actual)
	assert.NotNil(t, err)
}

func TestStrMap_GetWithDefault(t *testing.T) {
	s := NewStrMap()
	s = s.Add("foo", "bar")
	actual := s.GetWithDefault("foo", "default")
	assert.Equal(t, "bar", actual)

	actual = s.GetWithDefault("qux", "default")
	assert.Equal(t, actual, "default")
}

func TestStrMap_Remove(t *testing.T) {
	s := NewStrMap()
	s = s.Add("foo", "bar")
	a, err := s.Get("foo")
	assert.Equal(t, "bar", a)
	assert.Nil(t, err)

	actual := s.Remove("foo")

	a, err = actual.Get("foo")
	assert.NotNil(t, err)
	assert.Nil(t, a)
}

func TestStrMap_RemoveIntersections(t *testing.T) {
	s := NewStrMap()
	s = s.Add("fookey", "fooval")
	s = s.Add("barkey", "barval")
	s = s.Add("quxkey", "quxval")
	a, err := s.Get("fookey")
	assert.Equal(t, "fooval", a)
	assert.Nil(t, err)

	o := NewStrMap()
	o = o.Add("fookey", "other foo key")
	o = o.Add("barkey", "barval")

	actual := s.RemoveIntersections(o)

	a, err = actual.Get("fookey")
	assert.NotNil(t, err)
	assert.Nil(t, a)

	a, err = actual.Get("barkey")
	assert.NotNil(t, err)
	assert.Nil(t, a)

	a, err = actual.Get("quxkey")
	assert.Nil(t, err)
	assert.Equal(t, a, "quxval")

}

func TestStrMap_RemoveAll(t *testing.T) {
	s := NewStrMap()
	s = s.Add("fookey", "fooval")
	s = s.Add("barkey", "barval")
	s = s.Add("quxkey", "quxval")
	a, err := s.Get("fookey")
	assert.Equal(t, "fooval", a)
	assert.Nil(t, err)

	actual := s.RemoveAll([]string{"fookey", "barkey"})

	assert.Equal(t, actual.Len(), 1)
	assert.False(t, actual.Contains("fookey"))
	assert.True(t, actual.Contains("quxkey"))

	actual = s.RemoveAll([]string{})

	assert.True(t, s.Equals(actual))

	actual = s.RemoveAll([]string{"fookey", "barkey", "quxkey"})
	assert.True(t, actual.IsEmpty())
}

func TestStrMap_Intersection(t *testing.T) {
	s := NewStrMap()
	s = s.Add("fookey", "fooval")
	s = s.Add("barkey", "barval")
	s = s.Add("quxkey", "quxval")
	a, err := s.Get("fookey")
	assert.Equal(t, "fooval", a)
	assert.Nil(t, err)

	o := NewStrMap()
	o = o.Add("fookey", "other foo val")
	o = o.Add("barkey", "barval")

	actual := s.Intersection(o)

	a, err = actual.Get("fookey")
	assert.Nil(t, err)
	assert.Equal(t, a, "other foo val")

	a, err = actual.Get("barkey")
	assert.Nil(t, err)
	assert.Equal(t, a, "barval")

	a, err = actual.Get("quxkey")
	assert.Nil(t, a)
	assert.NotNil(t, err)
}

func TestStrMap_Disjunction(t *testing.T) {
	s := NewStrMap()
	s = s.Add("fookey", "fooval")
	s = s.Add("barkey", "barval")
	a, err := s.Get("fookey")
	assert.Equal(t, "fooval", a)
	assert.Nil(t, err)

	o := NewStrMap()
	o = o.Add("fookey", "other foo val")
	o = o.Add("quxkey", "quxval")

	actual := s.Disjunction(o)

	a, err = actual.Get("fookey")
	assert.Nil(t, a)
	assert.NotNil(t, err)

	a, err = actual.Get("barkey")
	assert.Nil(t, err)
	assert.Equal(t, a, "barval")

	a, err = actual.Get("quxkey")
	assert.Equal(t, a, "quxval")
	assert.Nil(t, err)
}

func TestStrMap_ContainsAll(t *testing.T) {
	s := NewStrMap()
	s = s.Add("fookey", "fooval")
	s = s.Add("barkey", "barval")
	a, err := s.Get("fookey")
	assert.Equal(t, "fooval", a)
	assert.Nil(t, err)

	o := NewStrMap()
	o = o.Add("fookey", "other foo val")
	o = o.Add("quxkey", "quxval")

	actual := s.ContainsAll(o)
	assert.False(t, actual)

	s = s.Add("quxkey", "")
	actual = s.ContainsAll(o)
	assert.True(t, actual)

	assert.True(t, NewStrMap().ContainsAll(NewStrMap()))
	assert.True(t, s.ContainsAll(NewStrMap()))
}

func TestStrMap_ContainsAny(t *testing.T) {
	s := NewStrMap()
	s = s.Add("fookey", "fooval")
	s = s.Add("barkey", "barval")
	a, err := s.Get("fookey")
	assert.Equal(t, "fooval", a)
	assert.Nil(t, err)

	o := NewStrMap()
	o = o.Add("fookey", "other foo val")
	o = o.Add("quxkey", "quxval")

	actual := s.ContainsAny(o)
	assert.True(t, actual)

	s = s.Remove("fookey")
	actual = s.ContainsAny(o)
	assert.False(t, actual)

	assert.True(t, NewStrMap().ContainsAny(NewStrMap()))
	assert.True(t, s.ContainsAny(NewStrMap()))
}

func TestStrMap_Equals(t *testing.T) {
	s := NewStrMap()
	s = s.Add("fookey", "fooval")
	s = s.Add("barkey", "barval")
	a, err := s.Get("fookey")
	assert.Equal(t, "fooval", a)
	assert.Nil(t, err)

	o := NewStrMap()
	o = o.Add("fookey", "other foo val")
	o = o.Add("quxkey", "quxval")

	assert.False(t, s.Equals(o))

	o = o.Remove("quxkey")
	o = o.Add("barkey", "barval")
	assert.False(t, s.Equals(o))

	o = o.Add("fookey", "fooval")
	assert.True(t, s.Equals(o))
}

func TestStrMap_ValueFrequency(t *testing.T) {
	s := NewStrMap()
	s = s.Add("fookey", "fooval")
	s = s.Add("barkey", "barval")
	s = s.Add("quxkey", "quxval")
	a, err := s.Get("fookey")
	assert.Equal(t, "fooval", a)
	assert.Nil(t, err)

	assert.Equal(t, s.ValueFrequency("fooval"), 1)
	assert.Equal(t, s.ValueFrequency("nonexistant"), 0)
	s = s.Add("barkey", "fooval")
	assert.Equal(t, s.ValueFrequency("fooval"), 2)
}

func TestStrMap_Filter(t *testing.T) {
	s := NewStrMap()
	s = s.Add("fookey", "fooval")
	s = s.Add("barkey", "barval")
	s = s.Add("quxkey", "quxval")

	f := func(_ string, val interface{}) bool {
		s := val.(string)
		suffix := strings.HasSuffix(s, "val")
		return suffix
	}

	actual := s.Filter(f)
	assert.True(t, s.Equals(actual))

	s = s.Add("fookey", "val with different suffix")

	actual = s.Filter(f)
	assert.False(t, s.Equals(actual))
	assert.False(t, actual.Contains("fookey"))
}

func TestStrMap_Map(t *testing.T) {
	s := NewStrMap()
	s = s.Add("fookey", "fooval")
	s = s.Add("barkey", "barval")
	s = s.Add("quxkey", "quxval")

	f := func(_ string, val interface{}) interface{} {
		s := val.(string)
		if strings.HasSuffix(s, "val") {
			return s + "idated"
		}
		return s
	}

	actual := s.Map(f)
	assert.False(t, s.Equals(actual))
	assert.Equal(t, actual.GetWithDefault("fookey", ""), "foovalidated")
	assert.Equal(t, actual.GetWithDefault("barkey", ""), "barvalidated")
	assert.Equal(t, actual.GetWithDefault("quxkey", ""), "quxvalidated")

	s = s.Add("fookey", "val with different suffix")
	actual = s.Map(f)

	assert.False(t, s.Equals(actual))
	assert.Equal(t, actual.GetWithDefault("fookey", ""), "val with different suffix")
	assert.Equal(t, actual.GetWithDefault("barkey", ""), "barvalidated")
	assert.Equal(t, actual.GetWithDefault("quxkey", ""), "quxvalidated")
}

func TestStrMap_Reduce(t *testing.T) {
	s := NewStrMap()
	s = s.Add("fookey", "fooval")
	s = s.Add("barkey", "barval")
	s = s.Add("quxkey", "quxval")

	f := func(key string, val, acc interface{}) interface{} {
		valstr := val.(string)
		var accstr string
		if acc == nil {
			accstr = ""
		} else {
			accstr = acc.(string)
			accstr += ", "
		}
		return accstr + key + ":" + valstr
	}

	result := s.Reduce(f).(string)
	assert.True(t, strings.Contains(result, "fookey:fooval"))
	assert.True(t, strings.Contains(result, "barkey:barval"))
	assert.True(t, strings.Contains(result, "quxkey:quxval"))

	assert.Nil(t, NewStrMap().Reduce(f))
}

func TestNewStrMapFrom(t *testing.T) {
	m := make(map[string]interface{})
	m["foo"] = "bar"
	s := NewStrMapFrom(m)
	assert.NotNil(t, s)
	assert.True(t, s.Contains("foo"))
	assert.True(t, s.ContainsValue("bar"))
	assert.Equal(t, s.GetWithDefault("foo", ""), "bar")
}
