package repository

import (
	"reflect"
	"studentgit.kata.academy/eldar/review_4_2/app"
	"studentgit.kata.academy/eldar/review_4_2/model"
	"testing"
)

func TestCreateDeleteUpdateList(t *testing.T) {
	dao, err := (&app.App{}).CreateDAO()
	if err != nil {
		t.Errorf("CreateDAO() err: %s", err)
		return
	}

	rep := New(dao)

	city := &model.City{Name: "Вайоминг", State: "Аляска"}
	id, err := rep.Create(city.Name, city.State)
	if err != nil {
		t.Errorf("rep.Create err: %s", err)
		return
	}
	city.Id = id

	cities, err := rep.List()
	if err != nil {
		t.Errorf("rep.Create err: %s", err)
		return
	}
	got := searchCity(id, cities)
	if got == nil || !reflect.DeepEqual(city, got) {
		t.Errorf("after create: city: %#v not found, found: %#v", city, got)
		return
	}

	city.State += city.State
	city.Name += city.Name
	err = rep.Update(id, got.Name+got.Name, got.State+got.State)
	cities, _ = rep.List()
	got = searchCity(id, cities)
	if got == nil || !reflect.DeepEqual(city, got) {
		t.Errorf("after update: city: %#v not found, found: %#v", city, got)
		return
	}

	err = rep.Delete(id)
	if err != nil {
		t.Errorf("rep.Delete(id) err: %s", err)
		return
	}
	cities, _ = rep.List()
	got = searchCity(id, cities)
	if got != nil {
		t.Errorf("after delete: found %#v, but don't must", got)
		return
	}
}

func searchCity(id int, cities []*model.City) *model.City {
	for i := range cities {
		if cities[i].Id == id {
			return cities[i]
		}
	}
	return nil
}
