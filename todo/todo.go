package todo

import (
	"encoding/json"
	"io/ioutil"
)

type Item struct {
	Text     string
	Priority int
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 2:
		i.Priority = 2
	case 3:
		i.Priority = 3
	case 4:
		i.Priority = 4
	case 5:
		i.Priority = 5
	default:
		i.Priority = 3
	}
}

func SaveItems(filename string, items []Item) error {
	b, mErr := json.Marshal(items)
	if mErr != nil {
		return mErr
	}
	err := ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	return items, nil
}
