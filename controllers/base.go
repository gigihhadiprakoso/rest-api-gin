package controllers

import (
	// . "rest-api-gin/helpers"
	"rest-api-gin/config"

	// . "rest-api-gin/models"
	// "encoding/json"
	// "fmt"
	"reflect"
)

var ID_Company uint
var db = config.ConnectDB()

func add(){

}

func edit(){

}

func get(tb interface{}, tipe reflect.Type, where_clause string){
	/*aa := reflect.New(tipe)
	fmt.Println(&tb)
	fmt.Println(reflect.TypeOf(&tb))
	fmt.Println(tb)
	fmt.Println(reflect.TypeOf(tb))
	// json.Unmarshal(*tb, &aa)
// abc := tb.(tipe)
	fmt.Println(reflect.TypeOf(&aa))*/

	/*var items []interface{}
	var records []Products
	object := reflect.ValueOf(tb)
	for i := 0; i < object.Len(); i++ {
		items = append(items, object.Index(i).Interface())
	}
	for _, v := range items {
		item := reflect.ValueOf(v)
		var record Products
		fmt.Println(item)

		// itm0 := item.Field(0).Interface()
		// itm1 := item.Field(1).Interface()
		// itm2 := item.Field(2).Interface()

		// record.Id = itm0.(int)
		// record.Key = itm1.(string)
		// record.Value = itm2.(int)

		records = append(records, record)
	}
	fmt.Println(reflect.TypeOf(&records))*/
// fmt.Println(reflect.TypeOf(&abc))
	/*result := db.Find(&tb)
	if result.Error != nil {
		ResponseJSON(998, result.Error)
	} else {
		ResponseJSON(100, result.Value)
	}*/
}

func delete(){

}