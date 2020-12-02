package controller

import (
    "encoding/json"
    "project/global"
    "project/model"
)


func SaveUser(name string, password string) error {
    return global.DB.Put(name, password)
}


func GetUser(name string) (model.User, error) {
    var user model.User
    password, e := global.DB.Get(name)
    if e != nil {
        return user, e
    }
    user.Pass = password
    return user, nil
}

func EncodeUser(user model.User) []byte {
    out, _ := json.Marshal(user)
    return out
}
