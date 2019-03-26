package mongo

import (
        "log"

        "gopkg.in/mgo.v2"
)

type MsgBroken struct {
        Json  string
        Count int
}

func InsertOne(key int, value string) bool {
        session, err := mgo.Dial("localhost:27017")
        if err != nil {
                log.Println(err)
                return false
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)
        c := session.DB("gologs").C("gologs")

        err = c.Insert(&MsgBroken{Json: value, Count: key})
        if err != nil {
                log.Println(err)
                return false
        }
        return true
}
