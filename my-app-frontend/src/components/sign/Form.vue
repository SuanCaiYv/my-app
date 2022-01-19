<template>
    <div class="form">
        <div class="l1">
            <div class="name-show">用户名</div>
            <input class="value-input" type="email" v-model="username"/>
        </div>
        <div class="l2">
            <div class="name-show">密&nbsp;&nbsp;&nbsp;&nbsp;码</div>
            <input class="value-input" type="password" v-model="password"/>
        </div>
        <div class="l3">
            <div class="name-show name-show-click" @click="sendVerCode">验证码</div>
            <input class="value-input" type="text" v-model="verCode" />
        </div>
        <div class="l4">
            <button class="sign-button-in">登录</button>
            <button class="sign-button-up">注册</button>
        </div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Form")

let username = ref<String>("")
let password = ref<String>("")
let verCode = ref<String>("")

const sendVerCode = function () {
    console.log(username.value)
}
</script>

<style scoped>
.form {
    width: 450px;
    height: 315px;
    border: 1px solid silver;
    box-sizing: border-box;
    margin: 230px auto auto;
}

.l1 {
    width: 420px;
    height: 50px;
    border: 1px solid silver;
    box-sizing: border-box;
    border-radius: 8px;
    margin: 20px auto auto;
}

.l2 {
    width: 420px;
    height: 50px;
    border: 1px solid silver;
    box-sizing: border-box;
    border-radius: 8px;
    margin-top: 25px;
    margin-left: auto;
    margin-right: auto;
    margin-bottom: 0;
}

.l3 {
    width: 420px;
    height: 50px;
    border: 1px solid silver;
    box-sizing: border-box;
    border-radius: 8px;
    margin-top: 25px;
    margin-left: auto;
    margin-right: auto;
    margin-bottom: 0;
}

.l4 {
    width: 420px;
    height: 50px;
    border: 1px solid silver;
    box-sizing: border-box;
    border-radius: 8px;
    margin-top: 25px;
    margin-left: auto;
    margin-right: auto;
    margin-bottom: auto;
}

.name-show {
    width: 100px;
    height: 50px;
    margin-left: 10px;
    margin-right: 10px;
    border: 1px solid silver;
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
    margin-left: -2px;
    border: 1px solid silver;
    box-sizing: border-box;
    border-radius: 8px;
    display: inline-block;
    vertical-align: bottom;
    line-height: 50px;
}

.name-show-click:hover {
    background-color: lightskyblue;
}

.name-show-click:active {
    background-color: deepskyblue;
}

.sign-button-in {
    width: 150px;
    height: 50px;
    border: 1px solid silver;
    box-sizing: border-box;
    border-radius: 8px;
    margin-left: auto;
    margin-right: 25px;
    background-color: plum;
    font-size: 1.2rem;
}

.sign-button-up {
    width: 150px;
    height: 50px;
    border: 1px solid silver;
    box-sizing: border-box;
    border-radius: 8px;
    margin-left: 25px;
    margin-right: auto;
    background-color: wheat;
    font-size: 1.2rem;
}
</style>