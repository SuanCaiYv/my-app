<template>
    <div class="form">
        <div class="l1">
            <div class="name-show" style="background-color: #f0bbff">用户名</div>
            <input class="value-input" type="email" v-model="username"/>
        </div>
        <div class="l2">
            <div class="name-show" style="background-color: #ffbabf">密&nbsp;&nbsp;&nbsp;&nbsp;码</div>
            <input class="value-input" type="password" v-model="password"/>
        </div>
        <div class="l3">
            <div class="name-show name-show-click ver-code-button" @click="sendVerCode">验证码</div>
            <input class="value-input" type="text" v-model="verCode" placeholder='点击"验证码"以发送'/>
        </div>
        <div class="l4">
            <button class="sign-button-in" @click="login">登录</button>
            <button class="sign-button-up" @click="signup">注册</button>
        </div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {httpClient, Response} from "../../net";
import alertFunc from "../../util/alert";
import router from "../../router";
import storage from "../../util/storage";

const name = ref<string>("Form")

let username = ref<string>("")
let password = ref<string>("")
let verCode = ref<string>("")

storage.setOnce("lastVerCodeSendTimestamp", (new Date().getTime() - 120 * 1000) + "")

const sendVerCode = function () {
    if (new Date().getTime() - Number(storage.get("lastVerCodeSendTimestamp")) < 120 * 1000) {
        alertFunc("请" + Math.trunc(120 - (new Date().getTime() - Number(storage.get("lastVerCodeSendTimestamp"))) / 1000) + "秒后重试", function () {})
        return
    }
    setTimeout(() => {
        storage.set("lastVerCodeSendTimestamp", (new Date().getTime() - 120 * 1000) + "")
    }, 120 * 1000)
    httpClient.post("/sign/ver_code", {}, {
        username: username.value
    }, false, function (resp: Response) {
        if (resp.ok) {
            alertFunc("验证码发送成功!", function () {})
            storage.set("lastVerCodeSendTimestamp", new Date().getTime() + "")
        }
    })
}

const jumpHome = function () {
    router.push("/home")
}

const login = function () {
    httpClient.put("/sign", {}, {
        username: username.value,
        credential: password.value,
    }, false, function (resp: Response) {
        if (resp.ok) {
            // @ts-ignore
            storage.set("accessToken", resp.data.access_token)
            // @ts-ignore
            storage.set("refreshToken", resp.data.refresh_token)
           storage.set("authed", "true")
            alertFunc("登录成功", function () {
                jumpHome()
            })
        } else {
            alertFunc(resp.errMsg, function () {})
        }
    })
}

const signup = function () {
    httpClient.post("/sign", {}, {
        username: username.value,
        credential: password.value,
        ver_code: verCode.value
    }, false, function (resp: Response) {
        if (resp.ok) {
            alertFunc("注册成功", function () {})
        } else {
            alertFunc(resp.errMsg, function () {})
        }
    })
}
</script>

<style scoped>
.form {
    width: 450px;
    height: 315px;
    /*border: 1px solid silver;*/
    box-sizing: border-box;
    margin: 35% auto auto;
}

.l1 {
    width: 420px;
    height: 50px;
    /*border: 1px solid silver;*/
    box-sizing: border-box;
    border-radius: 8px;
    margin: 20px auto auto;
}

.l2 {
    width: 420px;
    height: 50px;
    /*border: 1px solid silver;*/
    box-sizing: border-box;
    border-radius: 8px;
    margin: 25px auto 0;
}

.l3 {
    width: 420px;
    height: 50px;
    /*border: 1px solid silver;*/
    box-sizing: border-box;
    border-radius: 8px;
    margin: 25px auto 0;
}

.l4 {
    width: 420px;
    height: 50px;
    /*border: 1px solid silver;*/
    box-sizing: border-box;
    border-radius: 8px;
    margin: 25px auto auto;
}

.name-show {
    width: 100px;
    height: 50px;
    opacity: 75%;
    margin-left: 10px;
    margin-right: 10px;
    /*border: 1px solid silver;*/
    box-sizing: border-box;
    border-radius: 8px;
    display: inline-block;
    font-size: 1.2rem;
    vertical-align: bottom;
    line-height: 50px;
}

.value-input {
    width: 300px;
    height: 50px;
    opacity: 75%;
    margin-left: -2px;
    /*border: 1px solid silver;*/
    border: none;
    box-sizing: border-box;
    border-radius: 8px;
    display: inline-block;
    vertical-align: bottom;
    line-height: 50px;
    font-size: 1.2rem;
}

.ver-code-button {
    background-color: #ffe0ad;
}

.sign-button-in {
    width: 150px;
    height: 50px;
    opacity: 85%;
    background-color: #b4d4ff;
    /*border: 1px solid silver;*/
    border: none;
    box-sizing: border-box;
    border-radius: 8px;
    margin-left: auto;
    margin-right: 25px;
    font-size: 1.2rem;
}

.sign-button-up {
    width: 150px;
    height: 50px;
    opacity: 65%;
    background-color: #c5ffea;
    /*border: 1px solid silver;*/
    border: none;
    box-sizing: border-box;
    border-radius: 8px;
    margin-left: 25px;
    margin-right: auto;
    font-size: 1.2rem;
}

.ver-code-button:hover {
    background-color: #ffd792;
}

.ver-code-button:active {
    background-color: #ffcf6e;
}

.sign-button-in:hover {
    background-color: #9dbbff;
}

.sign-button-in:active {
    background-color: #809dff;
}

.sign-button-up:hover {
    background-color: #82ffaf;
}

.sign-button-up:active {
    background-color: #5dff8a;
}
</style>