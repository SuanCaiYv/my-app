import axios from "axios";
import {useStore} from "vuex";

const baseUrl = "http://127.0.0.1:8190/v1"

axios.defaults.timeout = 2000

const store = useStore()

const httpClient = {
    // 路径参数一同放在uri里面了，所以不需要单独的参数
    get: function<T extends object> (uri: string, query: T, auth: boolean, callback: Function) {
        let str = ""
        for (let field in query) {
            str += (field + "=" + query[field])
        }
        let url = baseUrl + uri + "?" + str
        if (auth) {
            axios.get(url, {
                headers: {
                    "Authorization": "Bearer " + store.getters.accessToken,
                    "aaa": ""
                }
            })
        } else {
        }
    }
}

export default httpClient