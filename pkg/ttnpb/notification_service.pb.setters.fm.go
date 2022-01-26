// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *Notification) SetFields(src *Notification, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "id":
			if len(subs) > 0 {
				return fmt.Errorf("'id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Id = src.Id
			} else {
				var zero string
				dst.Id = zero
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				dst.CreatedAt = nil
			}
		case "entity_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EntityIdentifiers
				if (src == nil || src.EntityIds == nil) && dst.EntityIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EntityIds
				}
				if dst.EntityIds != nil {
					newDst = dst.EntityIds
				} else {
					newDst = &EntityIdentifiers{}
					dst.EntityIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EntityIds = src.EntityIds
				} else {
					dst.EntityIds = nil
				}
			}
		case "notification_type":
			if len(subs) > 0 {
				return fmt.Errorf("'notification_type' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.NotificationType = src.NotificationType
			} else {
				var zero string
				dst.NotificationType = zero
			}
		case "data":
			if len(subs) > 0 {
				return fmt.Errorf("'data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Data = src.Data
			} else {
				dst.Data = nil
			}
		case "sender_ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.SenderIds == nil) && dst.SenderIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.SenderIds
				}
				if dst.SenderIds != nil {
					newDst = dst.SenderIds
				} else {
					newDst = &UserIdentifiers{}
					dst.SenderIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.SenderIds = src.SenderIds
				} else {
					dst.SenderIds = nil
				}
			}
		case "receivers":
			if len(subs) > 0 {
				return fmt.Errorf("'receivers' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Receivers = src.Receivers
			} else {
				dst.Receivers = nil
			}
		case "status":
			if len(subs) > 0 {
				return fmt.Errorf("'status' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Status = src.Status
			} else {
				var zero NotificationStatus
				dst.Status = zero
			}
		case "status_updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'status_updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.StatusUpdatedAt = src.StatusUpdatedAt
			} else {
				dst.StatusUpdatedAt = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateNotificationRequest) SetFields(src *CreateNotificationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "entity_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EntityIdentifiers
				if (src == nil || src.EntityIds == nil) && dst.EntityIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EntityIds
				}
				if dst.EntityIds != nil {
					newDst = dst.EntityIds
				} else {
					newDst = &EntityIdentifiers{}
					dst.EntityIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EntityIds = src.EntityIds
				} else {
					dst.EntityIds = nil
				}
			}
		case "notification_type":
			if len(subs) > 0 {
				return fmt.Errorf("'notification_type' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.NotificationType = src.NotificationType
			} else {
				var zero string
				dst.NotificationType = zero
			}
		case "data":
			if len(subs) > 0 {
				return fmt.Errorf("'data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Data = src.Data
			} else {
				dst.Data = nil
			}
		case "sender_ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.SenderIds == nil) && dst.SenderIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.SenderIds
				}
				if dst.SenderIds != nil {
					newDst = dst.SenderIds
				} else {
					newDst = &UserIdentifiers{}
					dst.SenderIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.SenderIds = src.SenderIds
				} else {
					dst.SenderIds = nil
				}
			}
		case "receivers":
			if len(subs) > 0 {
				return fmt.Errorf("'receivers' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Receivers = src.Receivers
			} else {
				dst.Receivers = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateNotificationResponse) SetFields(src *CreateNotificationResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "id":
			if len(subs) > 0 {
				return fmt.Errorf("'id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Id = src.Id
			} else {
				var zero string
				dst.Id = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListNotificationsRequest) SetFields(src *ListNotificationsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "receiver_ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.ReceiverIds == nil) && dst.ReceiverIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.ReceiverIds
				}
				if dst.ReceiverIds != nil {
					newDst = dst.ReceiverIds
				} else {
					newDst = &UserIdentifiers{}
					dst.ReceiverIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ReceiverIds = src.ReceiverIds
				} else {
					dst.ReceiverIds = nil
				}
			}
		case "status":
			if len(subs) > 0 {
				return fmt.Errorf("'status' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Status = src.Status
			} else {
				dst.Status = nil
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListNotificationsResponse) SetFields(src *ListNotificationsResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "notifications":
			if len(subs) > 0 {
				return fmt.Errorf("'notifications' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Notifications = src.Notifications
			} else {
				dst.Notifications = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UpdateNotificationStatusRequest) SetFields(src *UpdateNotificationStatusRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "receiver_ids":
			if len(subs) > 0 {
				var newDst, newSrc *UserIdentifiers
				if (src == nil || src.ReceiverIds == nil) && dst.ReceiverIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.ReceiverIds
				}
				if dst.ReceiverIds != nil {
					newDst = dst.ReceiverIds
				} else {
					newDst = &UserIdentifiers{}
					dst.ReceiverIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ReceiverIds = src.ReceiverIds
				} else {
					dst.ReceiverIds = nil
				}
			}
		case "ids":
			if len(subs) > 0 {
				return fmt.Errorf("'ids' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Ids = src.Ids
			} else {
				dst.Ids = nil
			}
		case "status":
			if len(subs) > 0 {
				return fmt.Errorf("'status' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Status = src.Status
			} else {
				var zero NotificationStatus
				dst.Status = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
