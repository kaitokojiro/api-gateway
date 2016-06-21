/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"

	"github.com/nats-io/nats"
)

var (
	mockServices = []Service{
		Service{
			ID:           "1",
			Name:         "test",
			GroupID:      1,
			DatacenterID: 1,
		},
		Service{
			ID:           "2",
			Name:         "test2",
			GroupID:      2,
			DatacenterID: 3,
		},
	}
)

func getServiceSubcriber() {
	n.Subscribe("service.get", func(msg *nats.Msg) {
		if len(msg.Data) != 0 {
			qs := Service{}
			json.Unmarshal(msg.Data, &qs)

			for _, service := range mockServices {
				if qs.GroupID != 0 && service.GroupID == qs.GroupID && service.ID == qs.ID {
					data, _ := json.Marshal(service)
					n.Publish(msg.Reply, data)
					return
				} else if qs.GroupID == 0 && service.ID == qs.ID {
					data, _ := json.Marshal(service)
					n.Publish(msg.Reply, data)
					return
				}
			}
		}
		n.Publish(msg.Reply, []byte(`{"error":"not found"}`))
	})
}

func findServiceSubcriber() {
	n.Subscribe("service.find", func(msg *nats.Msg) {
		data, _ := json.Marshal(mockServices)
		n.Publish(msg.Reply, data)
	})
}

func createServiceSubcriber() {
	n.Subscribe("service.set", func(msg *nats.Msg) {
		var s Service

		json.Unmarshal(msg.Data, &s)
		s.ID = "3"
		data, _ := json.Marshal(s)

		n.Publish(msg.Reply, data)
	})
}

func deleteServiceSubcriber() {
	n.Subscribe("service.del", func(msg *nats.Msg) {
		var s Service

		json.Unmarshal(msg.Data, &s)

		n.Publish(msg.Reply, []byte{})
	})
}