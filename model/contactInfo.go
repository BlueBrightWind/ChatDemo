package model

type ContactInfo struct {
	ID       uint
	OwnerId  uint
	TargetId uint
	Type     int
	Desc     string
}

func NewContactInfo(contact Contact) ContactInfo {
	return ContactInfo{
		ID:       contact.ID,
		OwnerId:  contact.OwnerId,
		TargetId: contact.TargetId,
		Type:     contact.Type,
		Desc:     contact.Desc,
	}
}

func NewContactInfoList(contacts []Contact) []ContactInfo {
	contactInfos := make([]ContactInfo, 0)
	for _, contact := range contacts {
		contactInfos = append(contactInfos, NewContactInfo(contact))
	}
	return contactInfos
}
