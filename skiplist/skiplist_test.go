package skiplist

import (
	"testing"
)

func checkLevel(t *testing.T, s *SkipList, level int) {
	l := s.Level()
	if l != level {
		t.Errorf("checkLevel expect:%v, got:%v", level, l)
	}
}

func checkLength(t *testing.T, s *SkipList, length int) {
	l := s.Length()
	if l != length {
		t.Errorf("checkLength expect:%v, got:%v", length, l)
	}
}

func checkRank(t *testing.T, s *SkipList, key interface{}, rank int) {
	r := s.GetRank(key)
	if r != rank {
		t.Errorf("checkRank key:%v, expect:%v, got:%v", key, rank, r)
	}
}

func checkGet(t *testing.T, s *SkipList, key interface{}, value interface{}) {
	v := s.Get(key)
	if v != value {
		t.Errorf("checkGet key:%v，expect:%v, got:%v", key, value, v)
	}
}

func checkGetDataByRank(t *testing.T, s *SkipList, rank int, value interface{}) {
	v := s.GetDataByRank(rank)
	if v != value {
		t.Errorf("checkGetDataByRank rank:%v，expect:%v, got:%v", rank, value, v)
	}
}

func checkExist(t *testing.T, s *SkipList, key interface{}, exists bool) {
	e := s.Exist(key)
	if e != exists {
		t.Errorf("checkExist key:%v, expect:%v, got:%v", key, exists, e)
	}
}

func checkTop(t *testing.T, s *SkipList, rank int, num int) {
	list := s.Top(rank)

	if len(list) != num {
		t.Errorf("checkTop rank:%v, expect:%v, got:%v", rank, num, len(list))
	}
}

func checkBottom(t *testing.T, s *SkipList, rank int, num int) {
	list := s.Bottom(rank)
	if len(list) != num {
		t.Errorf("checkBottom rank:%v, expect:%v, got:%v", rank, num, len(list))
	}
}

func less(d1 interface{}, d2 interface{}) bool {
	data1 := d1.(*LadderData)
	data2 := d2.(*LadderData)

	if data1.value != data2.value {
		return data1.value > data2.value
	}

	return data1.id < data2.id
}

type LadderData struct {
	id    int
	value int
}

func TestAll(t *testing.T) {
	// init
	s := NewSkipList()
	if s == nil {
		t.Errorf("skiplist is nil")
	}
	s.Less = less

	checkLevel(t, s, 1)
	checkLength(t, s, 0)

	v1 := &LadderData{1, 1}
	s.Set("k1", v1)
	v2 := &LadderData{2, 2}
	s.Set("k2", v2)
	v3 := &LadderData{1, 2}
	s.Set("k1", v3)
	v4 := &LadderData{4, 4}
	s.Set("k4", v4)

	checkLength(t, s, 3)
	checkRank(t, s, "k1", 2)
	checkRank(t, s, "k2", 3)
	checkRank(t, s, "k4", 1)

	checkGetDataByRank(t, s, 1, v4)
	checkGetDataByRank(t, s, 2, v3)
	checkGetDataByRank(t, s, 3, v2)

	checkGet(t, s, "k1", v3)
	checkGet(t, s, "k2", v2)
	checkGet(t, s, "k4", v4)

	checkExist(t, s, "k1", true)
	checkExist(t, s, "k3", false)

	checkTop(t, s, 3, 3)
	checkTop(t, s, 0, 0)

	checkBottom(t, s, 2, 2)
	checkBottom(t, s, 0, 0)

	s.Delete("k1")
	checkLength(t, s, 2)
	checkRank(t, s, "k2", 2)
	checkExist(t, s, "k2", true)
	checkExist(t, s, "k1", false)

	checkRank(t, s, "k1", 0)
	checkGetDataByRank(t, s, 3, nil)

	checkGet(t, s, "k1", nil)

	s.Clear()

	checkLevel(t, s, 1)
	checkLength(t, s, 0)

	checkExist(t, s, "k2", false)
}
