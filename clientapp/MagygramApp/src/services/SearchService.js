import Axios from "axios";
import { authHeader } from "../helpers/auth-header";

export const searchService = {
	guestSearchUsers,
    guestSearchHashtagPosts,
    userSearchUsers,
};

function guestSearchUsers(value,callback){
    var options;
    
    Axios.get(`/api/users/search/${value}/guest`, { validateStatus: () => true, headers: authHeader() })
        .then((res) => {
            console.log(res.data);
            if (res.status === 200) {
                options = res.data.map(option => ({ value: option.Username, label: option.Username, id: option.Id}))
                callback(options);
        }}).catch((err) => {
            console.log(err)
    });
}

function userSearchUsers(value,callback){
    var options;
    
    Axios.get(`/api/users/search/${value}/user`, { validateStatus: () => true, headers: authHeader() })
        .then((res) => {
            console.log(res.data);
            if (res.status === 200) {
                options = res.data.map(option => ({ value: option.Username, label: option.Username, id: option.Id}))
                callback(options);
        }}).catch((err) => {
            console.log(err)
    });
}

function guestSearchHashtagPosts(value,callback){
    var options;
    value = value.substring(1);
    Axios.get(`/api/posts/hashtag-search/${value}/guest`, { validateStatus: () => true, headers: authHeader() })
        .then((res) => {
            console.log(res.data);
            if (res.status === 200) {
                options = res.data.map(option => ({ value: option.Hashtag, label: '#'+option.Hashtag + " ("+ option.NumberOfPosts+")", id: option.Hashtag, searchType:"hashtag"}))
                callback(options);
        }}).catch((err) => {
            console.log(err)
    });
}