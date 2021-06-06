import Axios from "axios";
import { authHeader } from "../helpers/auth-header";

export const searchService = {
	guestSearchUsers,
    guestSearchHashtagPosts,
};

function guestSearchUsers(value,callback){
    var options;
    
    Axios.get(`/api/users/search/${value}/guest`, { validateStatus: () => true, headers: authHeader() })
                .then((res) => {
                    console.log(res.data);
                    if (res.status === 200) {
                        options = res.data.map(option => ({ value: option.Username, label: option.Username, id: option.Id}))
                        callback(options);
                    }})
                .catch((err) => {
                    console.log(err)
                });
}

function guestSearchHashtagPosts(value,callback){
    
}