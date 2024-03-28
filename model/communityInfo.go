package model

type CommunityInfo struct {
	ID   uint
	Name string
	Img  string
	Desc string
}

func NewCommunityInfo(community Community) CommunityInfo {
	return CommunityInfo{
		ID:   community.ID,
		Name: community.Name,
		Img:  community.Img,
		Desc: community.Desc,
	}
}
