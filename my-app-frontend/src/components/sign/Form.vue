<template>
    <div class="form">
        <div class="l1">
            <div class="name-show">用户名</div>
            <input class="value-input" type="email" v-model="username"/>
        </div>
        <div class="l2">
            <button class="name-show click" @click="switchFunc">密&nbsp;&nbsp;&nbsp;&nbsp;码</button>
            <input class="value-input" type="password" v-model="password"/>
        </div>
        <div class="l3">
            <button class="name-show click" @click="sendVerCode">验证码</button>
            <input class="value-input" type="text" v-model="verCode" placeholder='点击"验证码"以发送'/>
        </div>
        <div class="l4">
            <button class="name-show click" @click="login" :class="{disabled: disableSignIn}">登录</button>
            <button class="name-show click" @click="signup">{{ btnName }}</button>
        </div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {httpClient} from "../../net";
import {Response} from "../../common/interface";
import alertFunc from "../../util/alert";
import router from "../../router";
import storage from "../../util/storage";
import {Constant} from "../../common/systemconstant";
import {useStore} from "vuex";
import {useRoute} from "vue-router";

const name = ref<string>("Form")

let username = ref<string>("")
let password = ref<string>("")
let verCode = ref<string>("")
let btnName = ref<string>("注册")
let disableSignIn = ref<boolean>(false)
let alertMsg = "注册成功"

const operation = useStore().getters.operation
let flag = false
if (operation === "update_password" || useRoute().params.operation === "update_password") {
    btnName.value = "重置密码"
    disableSignIn.value = true
    alertMsg = "重置密码成功"
    flag = false
}
const switchFunc = function () {
    flag = !flag
    if (flag) {
        btnName.value = "注册"
        disableSignIn.value = false
        alertMsg = "注册成功"
    } else {
        btnName.value = "重置密码"
        disableSignIn.value = true
        alertMsg = "重置密码成功"
    }
}

storage.setOnce(Constant.LAST_VERIFY_CODE_SEND_TIMESTAMP, (new Date().getTime() - 120 * 1000) + "")

const sendVerCode = function () {
    if (new Date().getTime() - Number(storage.get(Constant.LAST_VERIFY_CODE_SEND_TIMESTAMP)) < 120 * 1000) {
        alertFunc("请" + Math.trunc(120 - (new Date().getTime() - Number(storage.get(Constant.LAST_VERIFY_CODE_SEND_TIMESTAMP))) / 1000) + "秒后重试", function () {
        })
        return
    }
    if (username.value === "") {
        alertFunc("请输入用户名", function () {
        })
        return
    }
    setTimeout(() => {
        storage.set(Constant.LAST_VERIFY_CODE_SEND_TIMESTAMP, (new Date().getTime() - 120 * 1000) + "")
    }, 120 * 1000)
    httpClient.post("/sign/ver_code", {}, {
        username: username.value
    }, false, function (resp: Response) {
        if (resp.ok) {
            alertFunc("验证码发送成功!", function () {
            })
            storage.set(Constant.LAST_VERIFY_CODE_SEND_TIMESTAMP, new Date().getTime() + "")
        }
    })
}

const jumpHome = function () {
    router.push("/home")
}

const login = function () {
    if (username.value === "") {
        alertFunc("用户名为空", function () {
        })
        return
    }
    if (password.value === "") {
        alertFunc("密码为空", function () {
        })
        return
    }
    httpClient.put("/sign", {}, {
        username: username.value,
        credential: password.value,
    }, false, function (resp: Response) {
        if (resp.ok) {
            // @ts-ignore
            storage.set(Constant.ACCESS_TOKEN, resp.data.access_token)
            // @ts-ignore
            storage.set(Constant.REFRESH_TOKEN, resp.data.refresh_token)
            storage.set(Constant.AUTHENTICATED, "true")
            alertFunc("登录成功", function () {
                jumpHome()
            })
        } else {
            alertFunc(resp.errMsg, function () {
            })
        }
    })
}

const signup = function () {
    if (username.value === "") {
        alertFunc("用户名为空", function () {
        })
        return
    }
    if (password.value === "") {
        alertFunc("密码为空", function () {
        })
        return
    }
    if (verCode.value === "") {
        alertFunc("验证码为空", function () {
        })
        return
    }
    httpClient.post("/sign", {}, {
        username: username.value,
        credential: password.value,
        ver_code: verCode.value,
        operation: operation
    }, false, function (resp: Response) {
        if (resp.ok) {
            alertFunc(alertMsg, function () {
            })
        } else {
            alertFunc(resp.errMsg, function () {
            })
        }
    })
}
</script>

<style scoped>
.form {
    width: 450px;
    height: 325px;
    grid-area: form;
    border: none;
    border-radius: 18px;
    background-color: rgba(0, 0, 0, 0.1);
}

.l1 {
    width: 100%;
    height: 50px;
    border-radius: 8px;
    margin-top: 25px;
}

.l2 {
    width: 100%;
    height: 50px;
    border-radius: 8px;
    margin-top: 25px;
}

.l3 {
    width: 100%;
    height: 50px;
    border-radius: 8px;
    margin-top: 25px;
}

.l4 {
    width: 100%;
    height: 50px;
    border-radius: 8px;
    margin-top: 25px;
    margin-bottom: 25px;
}

.name-show {
    width: 100px;
    height: 100%;
    margin-left: 10px;
    margin-right: 10px;
    padding: 0;
    display: inline-block;
    font-size: 1.2rem;
    vertical-align: bottom;
    line-height: 50px;
    font-weight: bolder;
}

.value-input {
    width: calc(100% - 138px);
    height: 100%;
    border: none;
    border-radius: 16px;
    padding: 0 0 0 8px;
    margin-right: 10px;
    display: inline-block;
    vertical-align: bottom;
    line-height: 50px;
    font-size: 1.2rem;
    outline: none;
}

.click {
    line-height: 50px;
    border: none;
    border-radius: 30px;
    background-color: white;
}

.click:hover {
    background-color: lightgray;
}

.click:active {
    background-color: gainsboro;
}

.disabled {
    pointer-events: none;
    cursor: default;
    opacity: 0.4;
}
</style>