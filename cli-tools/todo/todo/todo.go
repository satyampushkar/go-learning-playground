package todo

import (
	//"fmt"
	"os"
	"encoding/json"
)

type Item struct {
	Text string
	Priority int
	//position int
	Done bool
}

func (i *Item) SetPriority(priority int) {
	switch priority {
		case 1: 
			i.Priority = 1
		case 3: 
			i.Priority = 3
		default:
			i.Priority = 2
	}
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	}
	return " "
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	//err = ioutil.WriteFile(filename, b, 0644)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = f.WriteString(string(b))
	if err != nil {
		//fmt.Println(err)
        f.Close()
		return err
	}

	//fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		//fmt.Println(err)
		return err
	}

	//fmt.Println(string(b))

	return nil
}

func ReadItems(filename string) ([]Item, error) {

	items := []Item{}
	
	data, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}

	//fmt.Println(string(data))
	err = json.Unmarshal(data, &items)
	if err != nil {
		return []Item{}, err
	}

	return items, nil
}