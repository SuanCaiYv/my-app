import axios, {AxiosResponse} from "axios";
import storage from "../util/storage";
import {Constant} from "../common/systemconstant";
import {Response} from "../common/interface";

export const BASE_URL = "http://127.0.0.1:8190/v1"

axios.defaults.timeout = 2000

class ResponseClass implements Response {
    ok: boolean
    errCode: number
    errMsg: string
    data: object
    timestamp: Date

    constructor() {
        this.ok = false
        this.errCode = 0
        this.errMsg = ""
        this.data = {}
        this.timestamp = new Date()
    }
}

const httpClient = {
    // 路径参数一同放在uri里面了，所以不需要单独的参数
    get: function <T extends object>(uri: string, query: T, auth: boolean, callback: Function) {
        let str = ""
        for (let field in query) {
            str += (field + "=" + query[field] + "&")
        }
        let url = BASE_URL + uri
        if (str.length > 0) {
            url += "?" + str.substring(0, str.length-1)
        }
        if (auth) {
            axios.get(url, {
                headers: {
                    "Authorization": "Bearer " + storage.get(Constant.ACCESS_TOKEN),
                }
            }).then(function (resp) {
                callback(dealResp(resp))
            }).catch(err => {
                console.log(err)
            })
        } else {
            axios.get(url)
                .then(function (resp) {
                    callback(dealResp(resp))
                }).catch(err => {
                console.log(err)
            })
        }
    },
    post: function <T extends object>(uri: string, query: T, params: T, auth: boolean, callback: Function) {
        let str = ""
        for (let field in query) {
            str += (field + "=" + query[field] + "&")
        }
        let url = BASE_URL + uri
        if (str.length > 0) {
            url += "?" + str.substring(0, str.length-1)
        }
        if (auth) {
            axios.post(url, params, {
                headers: {
                    "Authorization": "Bearer " + storage.get(Constant.ACCESS_TOKEN),
                    "Content-Type": "application/json"
                }
            }).then(function (resp) {
                callback(dealResp(resp))
            }).catch(err => {
                console.log(err)
            })
        } else {
            axios.post(url, params, {
                headers: {
                    "Content-Type": "application/json"
                }
            }).then(function (resp) {
                callback(dealResp(resp))
            }).catch(err => {
                console.log(err)
            })
        }
    },
    put: function <T extends object>(uri: string, query: T, params: T, auth: boolean, callback: Function) {
        let str = ""
        for (let field in query) {
            str += (field + "=" + query[field] + "&")
        }
        let url = BASE_URL + uri
        if (str.length > 0) {
            url += "?" + str.substring(0, str.length-1)
        }
        if (auth) {
            axios.put(url, params, {
                headers: {
                    "Authorization": "Bearer " + storage.get(Constant.ACCESS_TOKEN),
                    "Content-Type": "application/json"
                }
            }).then(function (resp) {
                callback(dealResp(resp))
            }).catch(err => {
                console.log(err)
            })
        } else {
            axios.put(url, params, {
                headers: {
                    "Content-Type": "application/json"
                }
            }).then(function (resp) {
                callback(dealResp(resp))
            }).catch(err => {
                console.log(err)
            })
        }
    },
    sign: function <T extends object>(uri: string, callback: Function) {
        let url = BASE_URL + uri
        axios.put(url, {}, {
            headers: {
                "Authorization": "Bearer " + storage.get(Constant.REFRESH_TOKEN),
                "Content-Type": "application/json"
            }
        }).then(function (resp) {
            callback(dealResp(resp))
        }).catch(err => {
            console.log(err)
        })
    },
    delete: function <T extends object>(uri: string, query: T, auth: boolean, callback: Function) {
        let str = ""
        for (let field in query) {
            str += (field + "=" + query[field] + "&")
        }
        let url = BASE_URL + uri
        if (str.length > 0) {
            url += "?" + str.substring(0, str.length-1)
        }
        if (auth) {
            axios.delete(url, {
                headers: {
                    "Authorization": "Bearer " + storage.get(Constant.ACCESS_TOKEN),
                }
            }).then(function (resp) {
                callback(dealResp(resp))
            }).catch(err => {
                console.log(err)
            })
        } else {
            axios.delete(url)
                .then(function (resp) {
                    callback(dealResp(resp))
                }).catch(err => {
                console.log(err)
            })
        }
    },
    // 此方法必须携带Token
    upload: function<T extends object> (uri: string, query: T, formData: FormData, callback: Function) {
        let str = ""
        for (let field in query) {
            str += (field + "=" + query[field] + "&")
        }
        let url = BASE_URL + uri
        if (str.length > 0) {
            url += "?" + str.substring(0, str.length-1)
        }
        axios.post(url, formData, {
            headers: {
                "Authorization": "Bearer " + storage.get(Constant.ACCESS_TOKEN),
                "Content-Type": "multipart/form-data"
            }
        }).then(resp => {
            callback(dealResp(resp))
        }).catch(err => {
            console.log(err)
        })
    }
}

const dealResp = function (resp: AxiosResponse): ResponseClass {
    let r = new ResponseClass()
    if (resp.status === 200) {
        let rawData = resp.data
        r.ok = rawData.code === 200
        r.errCode = rawData.code
        r.errMsg = rawData.msg
        r.data = rawData.data
        r.timestamp = new Date(rawData.timestamp)
    } else {
        r.ok = false
        r.errCode = 500
        r.errMsg = "Server Error!"
        r.data = {}
        r.timestamp = new Date()
    }
    return r
}

export {httpClient, ResponseClass}