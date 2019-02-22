package local

import (
	"reflect"
)

type LocalStore struct {
	list map[string][]interface{}
}

// var userList []usr.UserInfo
// var list = make(map[string][]interface{})

// type User struct {
// 	Name string
// 	Id   int
// }

// type User usr.UserInfo

func NewLocalStore() *LocalStore {
	store := &LocalStore{
		list: make(map[string][]interface{}),
	}
	return store
}

func (store *LocalStore) Read(key string, id string, data interface{}) error {
	_, d := getTypeAndVal(data)
	for _, val := range store.list[key] {
		_, v := getTypeAndVal(val)
		var _v reflect.Value
		// 指针类型处理
		if v.Kind() == reflect.Ptr {
			_v = v.Elem()
		} else {
			_v = v
		}
		status := findById(_v, id)
		if !status {
			continue
		}
		d.Set(v)
		break
	}
	return nil
}

// func (store *LocalStore) ReadAll(key string, data interface{}) error {
// 	v := reflect.ValueOf(data)
// 	v = v.Elem()
// 	// list
// 	userList := store.list[key]
// 	if len(userList) < 1 {
// 		return nil
// 	}
// 	v.Set(reflect.MakeSlice(v.Type(), len(userList), len(userList)))
// 	for index := 0; index < len(userList); index++ {

// 		v.Index(index).Set(reflect.ValueOf(userList[index]))
// 	}
// 	return nil
// }

func (store *LocalStore) ReadAll(key string, data interface{}) error {
	v := reflect.ValueOf(data)
	v = v.Elem()
	userList := store.list[key]
	if len(userList) < 1 {
		return nil
	}
	v.Set(reflect.MakeSlice(v.Type(), len(userList), len(userList)))
	for index := 0; index < len(userList); index++ {
		_, val := getTypeAndVal(userList[index])
		v.Index(index).Set(val)
	}
	return nil
}

func (store *LocalStore) Write(key string, x interface{}) error {
	// var Id int64
	if store.list[key] == nil {
		store.list[key] = []interface{}{}
	}
	store.list[key] = append(store.list[key], x)
	return nil
}
func (store *LocalStore) Update(key string, id string, x interface{}, result interface{}) {
	for k, val := range store.list[key] {
		_, v := getTypeAndVal(val)
		var _v reflect.Value
		// 指针类型处理
		if v.Kind() == reflect.Ptr {
			_v = v.Elem()
		} else {
			_v = v
		}
		f := _v.FieldByName("Id")
		if !f.IsValid() {
			continue
		}
		if f.Kind() == reflect.String {
			if f.String() == id {
				store.list[key][k] = x
				_, rs := getTypeAndVal(result)
				_, d := getTypeAndVal(store.list[key][k])
				rs.Set(d)
				break
			}
		}

	}
}
func (store *LocalStore) Delete(key string, id string) {
	for k, val := range store.list[key] {
		_, v := getTypeAndVal(val)
		var _v reflect.Value
		// 指针类型处理
		if v.Kind() == reflect.Ptr {
			_v = v.Elem()
		} else {
			_v = v
		}
		f := _v.FieldByName("Id")
		if !f.IsValid() {
			continue
		}
		if f.Kind() == reflect.String {
			if f.String() == id {
				store.list[key] = append(store.list[key][:k], store.list[key][k+1:]...)
				break
			}
		}

	}
}

func getTypeAndVal(s interface{}) (reflect.Type, reflect.Value) {
	v := reflect.ValueOf(s)
	v = reflect.Indirect(v)
	t := v.Type()
	return t, v
}

func findById(v reflect.Value, id string) bool {
	f := v.FieldByName("Id")
	if !f.IsValid() {
		return false
	}
	if f.Kind() == reflect.String {
		if f.String() == id {
			return true
		}
	}
	return false
}
