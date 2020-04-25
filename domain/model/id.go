package model

type UserID string

func (id UserID) String() string {
	return string(id)
}

type GroupID string

func (id GroupID) String() string {
	return string(id)
}

func (id GroupID) VillageID() VillageID {
	return VillageID(id)
}
