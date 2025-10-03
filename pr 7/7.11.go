//Корчагин Евгений 363

//у меня голова болит(

package main

import (
	"errors"
	"fmt"
)

type ContactInfo struct {
	ID    int
	Type  string
	Value string
}

type Contact struct {
	ID      int
	Name    string
	infoMap map[int]*ContactInfo
}

type ContactManager struct {
	contacts map[int]*Contact
	nextID   int
}

func (c *Contact) AddInfo(newInfo *ContactInfo) {
	if c.infoMap == nil {
		c.infoMap = make(map[int]*ContactInfo)
	}
	fmt.Printf("%s: %s добавлен в контакт %s\n", newInfo.Type, newInfo.Value, c.Name)
	c.infoMap[newInfo.ID] = newInfo
}

func (c *Contact) RemoveInfo(infoID int) error {
	value, ok := c.infoMap[infoID]
	if ok {
		fmt.Printf("%s: %s удален из контактов %s\n", value.Type, value.Value, c.Name)
		delete(c.infoMap, infoID)
		return nil
	} else {
		return errors.New("с таким ID нет")
	}
}

func (c *Contact) GetInfoByType(infoType string) []*ContactInfo {
	var result []*ContactInfo
	for _, info := range c.infoMap {
		if info.Type == infoType {
			result = append(result, info)
		}
	}
	return result
}

func (cm *ContactManager) AddContact(contact *Contact) {
	if cm.contacts == nil {
		cm.contacts = make(map[int]*Contact)
		cm.nextID = 1
	}
	contact.ID = cm.nextID
	cm.contacts[cm.nextID] = contact
	fmt.Printf("ID %d %s добавлен \n", cm.nextID, contact.Name)
	cm.nextID++
}

func (cm *ContactManager) RemoveContact(contactID int) error {
	contact, ok := cm.contacts[contactID]
	if ok {
		fmt.Printf("%s  удален\n", contact.Name)
		delete(cm.contacts, contactID)
		return nil
	} else {
		return errors.New("С таким ID нет")
	}
}

func (cm *ContactManager) FindContactByName(name string) []*Contact {
	var result []*Contact
	for _, contact := range cm.contacts {
		if contact.Name == name {
			result = append(result, contact)
		}
	}
	return result
}

func (cm *ContactManager) FindContactByInfo(value string) []*Contact {
	var result []*Contact
	//для каждого контакта
	for _, contact := range cm.contacts {
		//все его данные
		for _, info := range contact.infoMap {
			if info.Value == value {
				result = append(result, contact)
				break
			}
		}
	}
	return result
}

func main() {
	manager := &ContactManager{}

	contact1 := &Contact{Name: "1"}
	contact2 := &Contact{Name: "2"}

	manager.AddContact(contact1)
	manager.AddContact(contact2)

	phoneInfo := &ContactInfo{ID: 1, Type: "phone", Value: "81235652637"}
	emailInfo := &ContactInfo{ID: 2, Type: "email", Value: "JDH@mail.ru"}
	addressInfo := &ContactInfo{ID: 3, Type: "address", Value: "Москва, ул. Пушкина, 10"}

	contact1.AddInfo(phoneInfo)
	contact1.AddInfo(emailInfo)
	contact1.AddInfo(addressInfo)

	foundByName := manager.FindContactByName("2")
	for _, contact := range foundByName {
		fmt.Printf("Найден: %s (ID: %d)\n", contact.Name, contact.ID)
	}

	foundByPhone := manager.FindContactByInfo("81235652637")
	for _, contact := range foundByPhone {
		fmt.Printf("Найден: %s (ID: %d)\n", contact.Name, contact.ID)
	}

	fmt.Println()
	fmt.Println("Обладатель крутого номера:")
	phones := contact1.GetInfoByType("phone")
	for _, phone := range phones {
		fmt.Printf("%s\n", phone.Value)
	}
}
