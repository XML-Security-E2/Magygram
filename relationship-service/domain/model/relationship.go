package model

type FollowRequest struct {
	SubjectId string `json:"subjectId"`
	ObjectId string `json:"objectId"`
}

type Mute struct {
	SubjectId string `json:"subjectId"`
	ObjectId string `json:"objectId"`
}