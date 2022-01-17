import axios from "axios";
import {Ref} from "vue";

const baseUrl = "http://127.0.0.1:8190/v1"

axios.defaults.timeout = 2000

const httpClient = {
    get: function<T extends object> (uri: string, query: T, func: Function) {
        let str = ""
        for (let field in query) {
            str += (field + "=" + query[field])
        }
        let url = baseUrl + uri + "?" + str
        axios.get(url)
            .then(resp => {
                func(resp)
            })
            .catch(err => {
                console.log(err)
            })
    },
    post: function<T extends object> (uri: string, params: T, func: Function) {
        let url = baseUrl + uri
        axios.post(url, params)
            .then(resp => {
                func(resp)
            })
            .catch(err => {
                console.log(err)
            })
    },
    put: function<T extends object> (uri: string, params: T, func: Function) {
        let url = baseUrl + uri
        axios.put(url, params)
            .then(resp => {
                func(resp)
            })
            .catch(err => {
                console.log(err)
            })
    }
}

export default httpClient