// Go in Action
// @jeffotoni
// 2019-03-30

package server

import "github.com/jeffotoni/gologs/config"

func ServiceProducer(msgjson string) {

	switch config.SERVICE {

	case "postgres":
		PgProducer(msgjson)
		break
	case "redis":
		RdProducer(msgjson)
		break
	case "rabbitqm":
		RbProducer(msgjson)
		break
	case "mongo":
		MgProducer(msgjson)
		break

	default:
		PgProducer(msgjson)
	}
}
