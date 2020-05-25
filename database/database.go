package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"kvdb/kochetkov/consts"
)

// Database struct database storage
type Database struct {
	mutex *sync.RWMutex     // RWMutex for lock concurrency read and write
	base  map[string]string // Key-value storage map[key]value
}

// NewDatabase init new database struct
func NewDatabase() *Database {
	return &Database{
		mutex: new(sync.RWMutex),       // init mutex
		base:  make(map[string]string), // init empty map
	}
}

// Create function create key with value.
func (o *Database) Create(key, value string) (bool, error) {
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
func (o *Database) Read(key string) (string, error) {
	ok, err := o.IsExist(key)
	if !ok {
		return "", err
	}
	o.mutex.RLock()
	value := o.base[key]
	o.mutex.RUnlock()
	return value, nil
}

// Exist function
func (o *Database) IsExist(key string) (bool, error) {
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
func (o *Database) Update(key, valueNew string) (bool, error) {
	ok, err := o.IsExist(key)
	if !ok {
		return false, err
	}
	o.mutex.Lock()
	o.base[key] = valueNew
	o.mutex.Unlock()
	return true, nil
}

// Delete function delete key:value
func (o *Database) Delete(key string) (bool, error) {
	ok, err := o.IsExist(key)
	if !ok {
		return false, err
	}
	o.mutex.Lock()
	delete(o.base, key)
	o.mutex.Unlock()
	return true, nil
}

// Save function save in file map key-values before quit
func (o *Database) Save() error {
	b, err := json.Marshal(o.base)
	if err != nil {
		return err
	}
	ioutil.WriteFile(consts.Dump, b, 0777)
	return nil
}

// Load function load data from file in map
func (o *Database) Load() error {
	data, err := ioutil.ReadFile(consts.Dump)
	if err != nil {
		return err
	}
	var m map[string]string
	err = json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	o.base = m
	return nil
}
