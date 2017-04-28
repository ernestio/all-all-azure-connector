/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"log"
	"os"
	"runtime"

	ecc "github.com/ernestio/ernest-config-client"
	"github.com/ernestio/ernestprovider"
	"github.com/nats-io/nats"
)

var nc *nats.Conn
var natsErr error
var err error

func main() {
	nc = ecc.NewConfig(os.Getenv("NATS_URI")).Nats()
	_, err := nc.Subscribe("*.*.azure", func(m *nats.Msg) {
		key := os.Getenv("ERNEST_CRYPTO_KEY")
		subject, body := ernestprovider.GetAndHandle(m.Subject, m.Data, key)
		if err := nc.Publish(subject, body); err != nil {
			log.Println("Can't connect to nats")
		}
	})

	if err != nil {
		log.Println("Couldn't publish to nats")
	}

	runtime.Goexit()
}
