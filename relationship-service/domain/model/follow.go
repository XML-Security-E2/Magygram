package model

type FollowRequest struct {
	SubjectId string `json:"subjectId"`
	ObjectId string `json:"objectId"`
}