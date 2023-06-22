package funcDBserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PhoneNumber struct {
	Id          uint `gorm:"primary key;autoIncrement"`
	Name        string
	Surname     string
	Phonenumber string
}

func GetPhoneNumbers(_ interface{}) (interface{}, error) {
	fmt.Println(PhoneNumbers)
	return PhoneNumbers, nil
}

func GetPhoneNumber(key interface{}) (interface{}, error) {
	fmt.Println("sono qua merde")
	phoneNumberReturn := PhoneNumber{}
	fmt.Println("questa e la key ---->", key)

	for _, number := range PhoneNumbers {
		if fmt.Sprintf("%v", number.Id) == key {
			phoneNumberReturn = number
		}
	}

	fmt.Println(phoneNumberReturn)

	return phoneNumberReturn, nil

}

func DeletePhoneNumber(key interface{}) (interface{}, error) {
	db.Find(&PhoneNumbers)
	for _, number := range PhoneNumbers {
		if fmt.Sprintf("%v", number.Id) == key {
			db.Delete(&PhoneNumbers, number.Id)
			db.Find(&PhoneNumbers)
		}
	}

	fmt.Println(PhoneNumbers)

	return PhoneNumbers, nil
}

func addPhoneNumber(numberInput interface{}) (interface{}, error) {

	number, ok := numberInput.(PhoneNumber)

	if !ok {
		return nil, fmt.Errorf("errore")
	}

	db.Save(&number)
	PhoneNumbers = append(PhoneNumbers, number)

	return PhoneNumbers, nil
}

func getBody(r *http.Request) (interface{}, error) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var number PhoneNumber
	err = json.Unmarshal(reqBody, &number)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	return number, nil
}
