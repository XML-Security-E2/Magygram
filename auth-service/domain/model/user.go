package model

import (
	"errors"
	"html"
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        string `bson:"_id,omitempty"`
	Active    bool   `bson:"active"`
	Email     string `bson:"email" validate:"required,email"`
	Password  string `bson:"password"`
	Roles     []Role `bson:"roles"`
	TotpToken string `bson:"totp_token"`
}

type UserRequest struct {
	Id               string `json:"id"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	RepeatedPassword string `json:"repeatedPassword"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginTwoFactoryRequest struct {
	Email    string
	Password string
	Token    string
}

type ActivatedRequest struct {
	Email string `json:"email"`
}

type PasswordChangeRequest struct {
	UserId         string `json:"userId"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"passwordRepeat"`
}

func NewUser(userRequest *UserRequest, token string) (*User, error) {
	hashAndSalt, err := HashAndSaltPasswordIfStrongAndMatching(userRequest.Password, userRequest.RepeatedPassword)
	if err != nil {
		return nil, err
	}
	return &User{Id: userRequest.Id,
		Active:   false,
		Email:    html.EscapeString(userRequest.Email),
		Password: hashAndSalt,
		Roles: []Role{{Name: "user", Permissions: []Permission{
			{"create_post"},
			{"edit_profile"},
			{"edit_profile_photo"},
			{"get_logged_info"},
			{"view_user_profile"},
			{"view_followers"},
			{"view_following"},
			{"get_follow_request"},
			{"accept_follow_request"},
			{"follow"},
			{"unfollow"},
			{"mute"},
			{"unmute"},
			{"search"},
			{"create_highlights"},
			{"get_profile_highlights"},
			{"create_collection"},
			{"add_post_to_collection"},
			{"get_post_from_collection"},
			{"detele_collection"},
			{"get_user_collection"},
			{"check_post_favourites"},
			{"get_timeline_post"},
			{"create_post"},
			{"update_post"},
			{"get_post"},
			{"like_post"},
			{"dislike_post"},
			{"comment_post"},
			{"search_posts"},
			{"get_storyline_stories"},
			{"create_story"},
			{"get_story_highlights"},
			{"get_user_stories"},
			{"get_personal_stories"},
			{"visit_user_story"},
			{"get_liked_posts"},
			{"get_disliked_posts"},
			{"check_if_verified"},
			{"get_follow_recommendation"},
		}}},
		TotpToken: token}, err
}

func NewAgent(userRequest *UserRequest, token string) (*User, error) {
	return &User{Id: userRequest.Id,
		Active:   false,
		Email:    html.EscapeString(userRequest.Email),
		Password: userRequest.Password,
		Roles: []Role{{Name: "user", Permissions: []Permission{
			{"create_post"},
			{"edit_profile"},
			{"edit_profile_photo"},
			{"get_logged_info"},
			{"view_user_profile"},
			{"view_followers"},
			{"view_following"},
			{"get_follow_request"},
			{"accept_follow_request"},
			{"follow"},
			{"unfollow"},
			{"mute"},
			{"unmute"},
			{"search"},
			{"create_highlights"},
			{"get_profile_highlights"},
			{"create_collection"},
			{"add_post_to_collection"},
			{"get_post_from_collection"},
			{"detele_collection"},
			{"get_user_collection"},
			{"check_post_favourites"},
			{"get_timeline_post"},
			{"create_post"},
			{"update_post"},
			{"get_post"},
			{"like_post"},
			{"dislike_post"},
			{"comment_post"},
			{"search_posts"},
			{"get_storyline_stories"},
			{"create_story"},
			{"get_story_highlights"},
			{"get_user_stories"},
			{"get_personal_stories"},
			{"visit_user_story"},
			{"get_liked_posts"},
			{"get_disliked_posts"},
			{"check_if_verified"},
			{"get_follow_recommendation"},
		}},{Name: "agent", Permissions: []Permission{
			{"create_campaign"},
		}}},
		TotpToken: token}, nil
}



func HashAndSaltPasswordIfStrongAndMatching(password string, repeatedPassword string) (string, error) {
	isMatching := password == repeatedPassword
	if !isMatching {
		return "", errors.New("passwords are not matching")
	}
	isWeak, _ := regexp.MatchString("^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[^!@#$%^&*(),.?\":{}|<>~'_+=]*)$", password)

	if isWeak {
		return "", errors.New("password must contain minimum eight characters, at least one capital letter, one number and one special character")
	}
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash), err
}
