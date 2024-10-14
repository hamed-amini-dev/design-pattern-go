package main

import "fmt"

type Database struct{}

func (d *Database) Insert(s string) error {
	fmt.Println("Inserted ", s)
	return nil
}

type Logger struct{}

func (d *Logger) Log(s string) {
	fmt.Println("Logging Info: ", s)
}

type Cache struct{}

func (d *Cache) Store(key, value string) error {
	fmt.Printf("Stored %s with value %s \n", key, value)
	return nil
}

type Facade struct {
	database Database
	cache    Cache
	logger   Logger
}

func NewFacade(database Database, cache Cache, logger Logger) *Facade {
	return &Facade{
		database: database,
		cache:    cache,
		logger:   logger,
	}
}

func (f *Facade) Save() error {
	err := f.database.Insert("Sample")
	if err != nil {
		return err
	}
	err = f.cache.Store("SampleKey", "SampleValue")
	if err != nil {
		return err
	}
	f.logger.Log("Inserted Value and cached")
	return nil
}

func main() {
	facade := NewFacade(Database{}, Cache{}, Logger{})
	err := facade.Save()
	if err != nil {
		panic(err)
	}

}
