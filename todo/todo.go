package todo

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	position int
	Done     bool
}

// Priorty of the todo item
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

// Saving the todo items to a file
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

// Reading the todo items from a file
func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	for i, _ := range items {
		items[i].position = i + 1
	}

	return items, nil
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

// Pretty print the todo items
func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(1)"
	}
	if i.Priority == 2 {
		return "(2)"
	}
	if i.Priority == 3 {
		return "(3)"
	}
	if i.Priority == 4 {
		return "(4)"
	}
	if i.Priority == 5 {
		return "(5)"
	}

	return " "
}

func (i *Item) PrettyDone() bool {
	if i.Done {
		return true
	}
	return false
}

// ByPrio implements sort.Interface for []Item based on
// the Priority and position field.
type ByPrio []Item

func (s ByPrio) Len() int {
	return len(s)
}
func (s ByPrio) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByPrio) Less(i int, j int) bool {
	if s[i].Done != s[j].Done {
		return s[i].Done
	}
	if s[i].Priority == s[j].Priority {
		return s[i].position < s[j].position
	}
	return s[i].Priority < s[j].Priority
}
