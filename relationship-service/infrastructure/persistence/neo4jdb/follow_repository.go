package neo4jdb

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"io"
	"log"
	"relationship-service/domain/model"
)

type FollowRepository interface {
	CreateFollow(followRequest *model.FollowRequest) error
	CreateFollowRequest(followRequest *model.FollowRequest) error
	CreateUser(user *model.User) error
	Unfollow(followRequest *model.FollowRequest)  error
	ReturnFollowedUsers(user *model.User) (interface{}, error)
	ReturnFollowingUsers(user *model.User) (interface{}, error)
	ReturnFollowRequests(user *model.User) (interface{}, error)
	ReturnFollowRequestsForUser(user *model.User, loggedId string) (interface{}, error)
	AcceptFollowRequest(followRequest *model.FollowRequest) error
	IsUserFollowed(followRequest *model.FollowRequest) (interface{}, error)
	IsMuted(mute *model.Mute) (interface{}, error)
	Mute(mute *model.Mute) error
}

type followRepository struct {
	Driver neo4j.Driver
}

func NewFollowRepository(driver neo4j.Driver) FollowRepository {
	return &followRepository{driver}
}

func (f *followRepository) Mute(mute *model.Mute) error {
	session := f.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer unsafeClose(session)
	if _, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := "MATCH (subject:User)-[f:FOLLOW]->(object:User) where subject.id = $subjectId and object.id = $objectId SET f.muted = true"
		parameters := map[string]interface{}{
			"subjectId": mute.SubjectId,
			"objectId": mute.ObjectId,
		}
		_, err := tx.Run(query, parameters)
		return nil, err
	}); err != nil {
		return err
	}
	return nil
}

func (f *followRepository) Unmute(mute *model.Mute) error {
	session := f.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer unsafeClose(session)
	if _, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := "MATCH (subject:User)-[f:FOLLOW]->(object:User) where subject.id = $subjectId and object.id = $objectId REMOVE f.muted"
		parameters := map[string]interface{}{
			"subjectId": mute.SubjectId,
			"objectId": mute.ObjectId,
		}
		_, err := tx.Run(query, parameters)
		return nil, err
	}); err != nil {
		return err
	}
	return nil
}

func (f *followRepository) CreateFollow(followRequest *model.FollowRequest) (err error) {
	session := f.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer unsafeClose(session)

	if _, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){
		query := "MATCH (subject:User), (object:User) WHERE subject.id = $subjectId AND object.id = $objectId CREATE (subject)-[:FOLLOW]->(object)"
		parameters := map[string]interface{}{
			"subjectId": followRequest.SubjectId,
			"objectId": followRequest.ObjectId,
		}
		_, err := tx.Run(query, parameters)
		return nil, err
	}); err != nil {
		return err
	}
	return nil
}

func (f *followRepository) Unfollow(followRequest *model.FollowRequest) (err error) {
	session := f.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer unsafeClose(session)

	if _, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){
		query := "MATCH (subject:User)-[f:FOLLOW]->(object:User) WHERE subject.id = $subjectId AND object.id = $objectId DELETE f"
		parameters := map[string]interface{}{
			"subjectId": followRequest.SubjectId,
			"objectId": followRequest.ObjectId,
		}
		_, err := tx.Run(query, parameters)
		return nil, err
	}); err != nil {
		return err
	}
	return nil
}

func (f *followRepository) CreateFollowRequest(followRequest *model.FollowRequest) (err error) {
	session := f.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer unsafeClose(session)

	if _, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){
		query := "MATCH (subject:User), (object:User) WHERE subject.id = $subjectId AND object.id = $objectId CREATE (subject)-[:FOLLOWREQUEST]->(object)"
		parameters := map[string]interface{}{
			"subjectId": followRequest.SubjectId,
			"objectId": followRequest.ObjectId,
		}
		_, err := tx.Run(query, parameters)
		return nil, err
	}); err != nil {
		return err
	}
	return nil
}

func (f *followRepository) IsUserFollowed(followRequest *model.FollowRequest) (interface{}, error) {
	session := f.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer unsafeClose(session)

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){
		query := "MATCH (subject:User)-[f:FOLLOW]->(object:User) WHERE subject.id = $subjectId AND object.id = $objectId RETURN f"
		parameters := map[string]interface{}{
			"subjectId": followRequest.SubjectId,
			"objectId": followRequest.ObjectId,
		}
		records, err := tx.Run(query, parameters)
		for records.Next() {
			return true, err
		}
		return false, err
	});
	if err != nil {
		return false, err
	}
	return result, nil
}

func (f *followRepository) IsMuted(mute *model.Mute) (interface{}, error) {
	session := f.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer unsafeClose(session)

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error){
		query := "MATCH (subject:User)-[f:FOLLOW]->(object:User) WHERE subject.id = $subjectId AND object.id = $objectId RETURN f.mute"
		parameters := map[string]interface{}{
			"subjectId": mute.SubjectId,
			"objectId": mute.ObjectId,
		}
		records, err := tx.Run(query, parameters)
		for records.Next() {
			return true, err
		}
		return false, err
	});
	if err != nil {
		return false, err
	}
	return result, nil
}

func (f *followRepository) AcceptFollowRequest(followRequest *model.FollowRequest) error {
	session := f.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer unsafeClose(session)

	if _, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error){
		query := "MATCH (object:User)<-[f:FOLLOWREQUEST]-(subject:User) WHERE object.id = $objectId AND subject.id = $subjectId DELETE f"
		parameters := map[string]interface{}{
			"subjectId": followRequest.SubjectId,
			"objectId": followRequest.ObjectId,
		}
		_, err := tx.Run(query, parameters)

		query = "MATCH (object:User), (subject:User) WHERE object.id = $objectId AND subject.id = $subjectId CREATE (subject)-[:FOLLOW]->(object)"
		_, err = tx.Run(query, parameters)

		return nil, err
	}); err != nil {
		return err
	}
	return nil
}

func (f *followRepository) CreateUser(user *model.User) (err error) {
	session := f.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer unsafeClose(session)
	if _, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := "CREATE (:User{id:$uId})"
		parameters := map[string]interface{}{
			"uId": user.Id,
		}
		_, err := tx.Run(query, parameters)
		return nil, err
	}); err != nil {
		return err
	}
	return err
}

func (f *followRepository) ReturnFollowedUsers(user *model.User) (interface{}, error) {
	session := f.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer unsafeClose(session)
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := "MATCH (s:User)-[f:FOLLOW]->(o:User) WHERE s.id = $sId return o.id as id"
		parameters := map[string]interface{}{
			"sId": user.Id,
		}
		records, err := tx.Run(query, parameters)
		if err != nil {
			return nil, err
		}
		users := model.Users{}
		for records.Next() {
			record := records.Record()
			id, _ := record.Get("id")
			users.Users = append(users.Users, id.(string))
		}
		return users, nil
	})
	if err != nil {
		log.Println("error querying graph:", err)
		return nil, err
	}
	return result, nil
}

func (f *followRepository) ReturnFollowingUsers(user *model.User) (interface{}, error) {
	session := f.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer unsafeClose(session)
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := "MATCH (s:User)<-[f:FOLLOW]-(o:User) WHERE s.id = $sId return o.id as id"
		parameters := map[string]interface{}{
			"sId": user.Id,
		}
		records, err := tx.Run(query, parameters)
		if err != nil {
			return nil, err
		}
		users := model.Users{}
		for records.Next() {
			record := records.Record()
			id, _ := record.Get("id")
			users.Users = append(users.Users, id.(string))
		}
		return users, nil
	})
	if err != nil {
		log.Println("error querying graph:", err)
		return nil, err
	}
	return result, nil
}

func (f *followRepository) ReturnFollowRequests(user *model.User) (interface{}, error) {
	session := f.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer unsafeClose(session)
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := "MATCH (o:User)<-[f:FOLLOWREQUEST]-(s:User) WHERE o.id = $oId return s.id as id"
		parameters := map[string]interface{}{
			"oId": user.Id,
		}
		records, err := tx.Run(query, parameters)
		if err != nil {
			return nil, err
		}
		users := model.Users{}
		for records.Next() {
			record := records.Record()
			id, _ := record.Get("id")
			users.Users = append(users.Users, id.(string))
		}
		return users, nil
	})
	if err != nil {
		log.Println("error querying graph:", err)
		return nil, err
	}
	log.Println(result)
	return result, nil
}

func (f *followRepository) ReturnFollowRequestsForUser(user *model.User, loggedId string) (interface{}, error) {
	session := f.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer unsafeClose(session)
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := "MATCH (o:User)<-[f:FOLLOWREQUEST]-(s:User) WHERE o.id = $oId AND s.id = $loggedId RETURN f"
		parameters := map[string]interface{}{
			"oId": user.Id,
			"loggedId": loggedId,
		}
		records, err := tx.Run(query, parameters)
		for records.Next() {
			return true, err
		}
		return false, err
	})
	if err != nil {
		log.Println("error querying graph:", err)
		return false, err
	}
	return result.(bool), nil
}

func unsafeClose(closeable io.Closer) {
	if err := closeable.Close(); err != nil {
		log.Fatal(fmt.Errorf("could not close resource: %w", err))
	}
}
