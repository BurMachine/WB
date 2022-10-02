package LRU

import (
	"container/list"
	"fmt"
	"log"
)

type Item struct {
	Key   string      // Клюс
	Value interface{} // Значение кешируемошо элемента
}

type LRU struct {
	capacity int                      // Количество яеек
	items    map[string]*list.Element // Хэш таблица
	queue    *list.List               // Поле двухсвязного списка
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		queue:    list.New(),
	}
}

// Методы работы с очередью(кеш)

func (c *LRU) Delete(key string) error {
	element, exists := c.items[key]
	if !exists {
		err := fmt.Errorf("error in mapping: Element does not exist in cache")
		return err
	}
	c.queue.MoveToBack(element)
	item := c.queue.Remove(element).(*Item)
	delete(c.items, item.Key)
	return nil
}

func (c LRU) purge() {
	if element := c.queue.Back(); element != nil { // Смотрим является ли элемент последним в списке
		item := c.queue.Remove(element).(*Item) // Удаляем узел в двухсвязном списке
		delete(c.items, item.Key)               // Удаляем пару из хэш таблицы
	}
}

func (c *LRU) Set(key string, value interface{}) bool {
	if element, exists := c.items[key]; exists { // Проверка элемента на наличие в хэш таблице
		c.queue.MoveToFront(element)        // Если элемент имеется, то мы в листе сдвигаем его на первое место
		element.Value.(*Item).Value = value // Меняем значение на новое
		return true
	}
	if c.queue.Len() == c.capacity { // Если достигли cap
		c.purge() // Удаляем послледний элемент
	}
	log.Println(c.queue.Len())

	item := &Item{
		Key:   key,
		Value: value,
	}
	element := c.queue.PushFront(item) // Если элемента нет в мапе и кап в норме, то закидываю его в лист
	c.items[item.Key] = element        // Закидываю в мапу актуальное значние
	return true
}

func (c LRU) Get(key string) interface{} {
	element, exists := c.items[key]
	if !exists {
		return nil
	}
	c.queue.MoveToFront(element)
	return element.Value.(*Item).Value
}

func (c LRU) FindForDaY(param string) []interface{} {
	/*
		Как искать в кеше по дате
		Надстройка в сигнатуре функции, где мжет быь один параметр из трез возможных?
		day/month/year
	*/
	if param == "day" {
		/*
			Прохожусь по
		*/
	} else if param == "month" {

	} else if param == "year" {

	}
}
