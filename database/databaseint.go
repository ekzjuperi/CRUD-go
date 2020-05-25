package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"sync"

	"kvdb/kochetkov/consts"
)

// DatabaseInt struct database storage
type DatabaseInt struct {
	mutex *sync.RWMutex  // RWMutex for lock concurrency read and write
	base  map[string]int // Key-value storage map[key]value
}

// NewDatabaseInt init new database struct
func NewDatabaseInt() *DatabaseInt {
	return &DatabaseInt{
		mutex: new(sync.RWMutex),    // init mutex
		base:  make(map[string]int), // init empty map

	}
}

// Create function create key with value.
func (o *DatabaseInt) Create(key string, value int) (bool, error) {
	ok, err := o.IsExist(key)
	if v := o.base[key]; ok && v == value {
		return false, err
	}
	o.mutex.Lock()
	o.base[key] = value
	o.mutex.Unlock()
	return true, nil
}

// Read function read key by value
func (o *DatabaseInt) Read(key string) (int, error) {
	ok, err := o.IsExist(key)
	if !ok {
		return 0, err
	}
	o.mutex.RLock()
	value := o.base[key]
	o.mutex.RUnlock()

	return value, nil
}

// Exist function
func (o *DatabaseInt) IsExist(key string) (bool, error) {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	_, ok := o.base[key]
	if !ok {
		errMsg := fmt.Errorf("key %v not found", key)
		return false, errMsg
	}
	return true, nil
}

// Update function update value by key
func (o *DatabaseInt) Update(key string, value int) (bool, error) {
	ok, err := o.IsExist(key)
	if !ok {
		return false, err
	}
	o.mutex.Lock()
	o.base[key] = value
	o.mutex.Unlock()
	return true, nil
}

// Delete function delete key:value
func (o *DatabaseInt) Delete(key string) (bool, error) {
	ok, err := o.IsExist(key)
	if !ok {
		return false, err
	}
	o.mutex.Lock()
	delete(o.base, key)
	o.mutex.Unlock()
	return true, nil
}

// sum - returns the sum of all values
func (o *DatabaseInt) Sum() int {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	sum := 0
	for _, value := range o.base {
		sum = sum + value
	}
	return sum
}

// avg - returns the arithmetic mean of all values
func (o *DatabaseInt) Avg() float64 {
	sum := o.Sum()
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	if len(o.base) != 0 {
		avg := float64(sum) / float64(len(o.base))
		return avg
	}
	return 0
}

// med - returns the median for all values
func (o *DatabaseInt) Med() ([]int, float64) {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	s := make([]int, 0)
	for _, value := range o.base {
		s = append(s, value)
	}
	switch len(s) {
	case 0:
		return nil, 0
	case 1:
		return s, 0
	default:
		sort.Ints(s)
		median := 0.0
		middleValue := len(s) / 2
		remainder := len(s) % 2
		if remainder == 0 {
			median = float64((s[middleValue-1] + s[middleValue]) / 2)
			return s, median
		}

		median = float64(s[middleValue])
		return s, median
	}
}

// gt val - returns all key-value elements whose value is greater than val
func (o *DatabaseInt) GtVal(val int) string {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	result := ""
	for key, value := range o.base {
		if value > val {
			result = result + key + ":" + strconv.Itoa(value) + " "
		}
	}
	return result
}

//lt val - returns all key-value elements whose value is less than val
func (o *DatabaseInt) LtVal(val int) string {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	result := ""
	for key, value := range o.base {
		if value < val {
			result = result + key + ":" + strconv.Itoa(value) + " "
		}
	}
	return result
}

// eq val - returns all key-value elements whose value is val
func (o *DatabaseInt) EqVal(val int) string {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	result := ""
	for key, value := range o.base {
		if value == val {
			result = result + key + ":" + strconv.Itoa(value) + " "
		}
	}
	return result
}

//count - returns the number of elements in the database
func (o *DatabaseInt) Count() int {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	return len(o.base)
}

// Save function save data to file in map
func (o *DatabaseInt) Save() error {
	b, err := json.Marshal(o.base)
	if err != nil {
		return err
	}
	ioutil.WriteFile(consts.Dumpint, b, 0777)
	return nil
}

// Load function load data from file in map
func (o *DatabaseInt) Load() error {
	data, err := ioutil.ReadFile(consts.Dumpint)
	if err != nil {
		return err
	}

	var m map[string]int

	err = json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	o.base = m

	return nil
}
