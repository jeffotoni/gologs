// Back-End in Go server
// @jeffotoni
// 2019-01-04

package repo

import "sync"

var m sync.Map

func Map(key int, value string) bool {

	m.Store(key, value)
	//v, ok := m.Load(k)
	//m.Delete(k)
	return true
}
