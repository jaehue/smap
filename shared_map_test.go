package smap

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"time"
)

func TestStringer(t *testing.T) {
	m := New()
	m.Set("one", 1)
	m.Set("two", 2)
	fmt.Println(m)
}

func TestSet(t *testing.T) {
	m := New()
	m.Set("one", 1)
	m.Set("two", 2)
	if m.Count() != 2 {
		t.Error("map should contain exactly two elements.")
	}
}

func TestGet(t *testing.T) {
	m := New()

	v, ok := m.Get("one")
	if ok {
		t.Error("ok should be false.")
	}

	if v != nil {
		t.Error("v should be nil.")
	}

	m.Set("one", 1)

	one, ok := m.Get("one")

	if !ok {
		t.Error("ok should be true.")
	}

	if one == nil {
		t.Error("one should be not nil.")
	}

	if one != 1 {
		t.Error("item was modified.")
	}
}

func TestRemove(t *testing.T) {
	m := New()

	m.Set("one", 1)

	m.Remove("one")

	if m.Count() != 0 {
		t.Error("Count should be zero.")
	}

	one, ok := m.Get("one")

	if ok {
		t.Error("ok should be false.")
	}

	if one != nil {
		t.Error("one should be nil.")
	}
}

func TestCount(t *testing.T) {
	m := New()
	for i := 0; i < 100; i++ {
		m.Set(strconv.Itoa(i), i)
	}

	if m.Count() != 100 {
		t.Error("Count of map should be 100.")
	}
}
func TestConcurrent(t *testing.T) {
	rand.Seed(time.Now().Unix())
	const size = 1000
	ch := make(chan int)
	var a [size]int

	m := New()

	for i := 0; i < size; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
			m.Set(strconv.Itoa(i), i)
			v, _ := m.Get(strconv.Itoa(i))
			ch <- v.(int)
		}(i)
	}

	for i := 0; i < size; i++ {
		v := <-ch
		a[i] = v
	}

	if m.Count() != size {
		t.Errorf("Count should be %d.", size)
	}

	sort.Ints(a[0:size])

	for i := 0; i < size; i++ {
		if i != a[i] {
			t.Error("missing value", i)
		}
	}
}
