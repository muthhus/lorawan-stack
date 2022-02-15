// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/lib/pq"
	"go.thethings.network/lorawan-stack/v3/pkg/jsonpb"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

type Notification struct {
	Model

	EntityID   string `gorm:"type:UUID;index:notification_entity_index;not null"`
	EntityType string `gorm:"type:VARCHAR(32);index:notification_entity_index;not null"`

	NotificationType string `gorm:"not null"`

	Data []byte `gorm:"type:JSONB"`

	Sender   *User
	SenderID string `gorm:"type:UUID;index:notification_sender_index;not null"`

	Receivers pq.Int32Array `gorm:"type:INT ARRAY"`
}

type NotificationReceiver struct {
	Notification   *Notification
	NotificationID string `gorm:"type:UUID;unique_index:notification_receiver_index;index:notification_receiver_notification_id_index;not null"`

	Receiver   *User
	ReceiverID string `gorm:"type:UUID;unique_index:notification_receiver_index;index:notification_receiver_user_index;not null"`

	Status          int32     `gorm:"not null"`
	StatusUpdatedAt time.Time `gorm:"not null"`
}

func init() {
	registerModel(&Notification{}, &NotificationReceiver{})
}

type notificationWithStatus struct {
	Notification

	FriendlyApplicationID          string
	FriendlyClientID               string
	FriendlyEndDeviceID            string
	FriendlyEndDeviceApplicationID string
	FriendlyGatewayID              string
	FriendlyOrganizationID         string
	FriendlyUserID                 string

	FriendlySenderID string

	Status          int32
	StatusUpdatedAt time.Time
}

func (n notificationWithStatus) toPB(pb *ttnpb.Notification) error {
	pb.Id = n.ID
	pb.CreatedAt = ttnpb.ProtoTimePtr(cleanTime(n.CreatedAt))
	if pb.EntityIds == nil {
		switch n.EntityType {
		case "application":
			pb.EntityIds = (&ttnpb.ApplicationIdentifiers{
				ApplicationId: n.FriendlyApplicationID,
			}).GetEntityIdentifiers()
		case "client":
			pb.EntityIds = (&ttnpb.ClientIdentifiers{
				ClientId: n.FriendlyClientID,
			}).GetEntityIdentifiers()
		case "end_device":
			pb.EntityIds = (&ttnpb.EndDeviceIdentifiers{
				ApplicationIds: &ttnpb.ApplicationIdentifiers{
					ApplicationId: n.FriendlyEndDeviceApplicationID,
				},
				DeviceId: n.FriendlyEndDeviceID,
			}).GetEntityIdentifiers()
		case "gateway":
			pb.EntityIds = (&ttnpb.GatewayIdentifiers{
				GatewayId: n.FriendlyGatewayID,
			}).GetEntityIdentifiers()
		case "organization":
			pb.EntityIds = (&ttnpb.OrganizationIdentifiers{
				OrganizationId: n.FriendlyOrganizationID,
			}).GetEntityIdentifiers()
		case "user":
			pb.EntityIds = (&ttnpb.UserIdentifiers{
				UserId: n.FriendlyUserID,
			}).GetEntityIdentifiers()
		}
	}
	pb.NotificationType = n.NotificationType
	if n.Data != nil && pb.Data == nil {
		var anyPB types.Any
		err := jsonpb.TTN().Unmarshal(n.Data, &anyPB)
		if err != nil {
			return err
		}
		pb.Data = &anyPB
	}
	if n.FriendlySenderID != "" {
		pb.SenderIds = &ttnpb.UserIdentifiers{UserId: n.FriendlySenderID}
	} else if n.Sender != nil {
		pb.SenderIds = &ttnpb.UserIdentifiers{UserId: n.Sender.Account.UID}
	}
	pb.Receivers = make([]ttnpb.NotificationReceiver, len(n.Receivers))
	for i, receiver := range n.Receivers {
		pb.Receivers[i] = ttnpb.NotificationReceiver(receiver)
	}
	pb.Status = ttnpb.NotificationStatus(n.Status)
	pb.StatusUpdatedAt = ttnpb.ProtoTimePtr(cleanTime(n.StatusUpdatedAt))
	return nil
}
