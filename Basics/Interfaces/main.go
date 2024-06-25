package main

import "fmt"

// Storage
// Stor(er)

type NumberStorer interface {
	GetAll()  			([]int, error)
	// Put(number int) [you can omit the number name]	error
	Put(int) 			error
}

type MongoDBNumberStore struct {
	// some values 
}

type PostgresDBStore struct {

}

func (p PostgresDBStore) GetAll() ([]int, error) {

	return []int {1, 2, 3, 5, 6}, nil
}

func (p PostgresDBStore) Put(numberToStore int) error {
	fmt.Println("This is the postgres db that put numbers to the numberStore.")
	return nil
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int {1, 2, 3}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	fmt.Println("Putting a number into the mongoDB storage.")
	return nil
}

type ApiServer struct {
	// numberStore MongoDBStorer (this is highly coupled to mongo, not recommeneed)
	numberStore NumberStorer
}

func main() {
	apiServer := ApiServer{
		numberStore: &MongoDBNumberStore{},
	}	

	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic("THere is an error")
	}
	fmt.Println(numbers)

	apiServer.numberStore = &PostgresDBStore{}
	numbers, err = apiServer.numberStore.GetAll()
	if err != nil {
		panic("THere is an error")
	}
	fmt.Println(numbers)
}