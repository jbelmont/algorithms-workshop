package dictionary

import "sync"

type DictionaryKey string
type DictionaryValue string

type Dictionary struct {
	Elements map[DictionaryKey]DictionaryValue
	Lock     sync.RWMutex
}

func (d *Dictionary) Put(key DictionaryKey, value DictionaryValue) {
	d.Lock.Lock()
	defer d.Lock.Unlock()

	if d.Elements == nil {
		d.Elements = make(map[DictionaryKey]DictionaryValue)
	}
	d.Elements[key] = value
}

func (d *Dictionary) Remove(key DictionaryKey) bool {
	d.Lock.Lock()
	defer d.Lock.Unlock()

	var exists bool
	_, exists = d.Elements[key]
	if exists {
		delete(d.Elements, key)
	}
	return exists
}

func (d *Dictionary) Contains(key DictionaryKey) bool {
	d.Lock.RLock()
	defer d.Lock.RUnlock()

	var exists bool
	_, exists = d.Elements[key]
	return exists
}

func (d *Dictionary) Find(key DictionaryKey) DictionaryValue {
	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return d.Elements[key]
}

func (d *Dictionary) Reset() {
	d.Lock.Lock()
	defer d.Lock.Unlock()

	d.Elements = make(map[DictionaryKey]DictionaryValue)
}

func (d *Dictionary) Count() int {
	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return len(d.Elements)
}

func (d *Dictionary) GetKeys() []DictionaryKey {
	d.Lock.RLock()
	defer d.Lock.RUnlock()

	dictionaryKeys := []DictionaryKey{}
	for key := range d.Elements {
		dictionaryKeys = append(dictionaryKeys, key)
	}
	return dictionaryKeys
}

func (d *Dictionary) GetValues() []DictionaryValue {
	d.Lock.RLock()
	defer d.Lock.RUnlock()

	dictionaryValues := []DictionaryValue{}

	for key := range d.Elements {
		dictionaryValues = append(dictionaryValues, d.Elements[key])
	}
	return dictionaryValues
}
