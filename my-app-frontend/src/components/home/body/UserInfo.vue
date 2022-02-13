<template>
    <div class="userInfo">
        <div class="img" @click="router.push('/about')">
            <Img :url="avatar" />
        </div>
        <div class="nickname">
            {{nickname}}
        </div>
        <div class="email">
            <a :href="mailto" style="text-decoration: none; color: inherit">{{email}}</a>
        </div>
        <div class="github">
            <a :href="github" style="text-decoration: none; color: inherit">{{github}}</a>
        </div>
        <div class="signature">
            {{signature}}
        </div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Img from "../../Img.vue"
import {useRouter} from "vue-router";
import {httpClient, Resp} from "../../../net";
import alertFunc from "../../../util/alert";

const name = ref<string>("UserInfo")

const avatar = ref<string>('http://127.0.0.1:8190/v1/static/a/my-avatar.png')
const nickname = ref<string>('小白白白')
const email = ref<string>('codewithbuff@163.com')
const github = ref<string>('https://github.com/SuanCaiYv')
const signature = ref<string>('Gin+Vue3')
const mailto = ref<String>("mailto:" + email.value)

httpClient.get("/user/info", {}, true, function (resp: Resp) {
    if (!resp.ok) {
        alertFunc(resp.errMsg, function () {})
    } else {
        avatar.value = resp.data.avatar
        nickname.value = resp.data.nickname
        email.value = resp.data.email
        location.value = resp.data.location
        github.value = resp.data.github
        signature.value = resp.data.signature
    }
})

const router = useRouter()
</script>

<style scoped>
.userInfo {
    width: 100%;
    height: 100%;
    grid-area: userInfo;
    border: 2px solid salmon;
    box-sizing: border-box;
    border-radius: 20px;
    background-color: mistyrose;
}

.img {
    width: 150px;
    height: 150px;
    padding: 0;
    border: none;
    margin-top: 23px;
    margin-left: auto;
    margin-right: auto;
    box-sizing: border-box;
    border-radius: 6px;
}

.img:active {
    opacity: 70%;
}

.nickname {
    width: 150px;
    height: auto;
    margin: 25px 25px 0;
    padding: 0;
    /*border: 1px solid black;*/
    /*box-sizing: border-box;*/
    border-radius: 6px;
    font-size: 1.2rem;
    line-height: 20px;
    font-weight: bolder;
    color: slateblue;
}

.email {
    width: 200px;
    height: auto;
    margin: 25px 0 0;
    padding: 0;
    /*border: 1px solid black;*/
    /*box-sizing: border-box;*/
    border-radius: 6px;
    font-size: 1rem;
    font-weight: bolder;
    line-height: 20px;
}

.github {
    width: 200px;
    height: auto;
    margin: 25px 0 0;
    padding: 0;
    /*border: 1px solid black;*/
    /*box-sizing: border-box;*/
    border-radius: 6px;
    font-size: 1rem;
    color: dodgerblue;
    font-weight: bolder;
    line-height: 20px;
}

.signature {
    width: 180px;
    height: auto;
    margin: 25px 10px 0;
    padding: 0;
    /*border: 1px solid black;*/
    /*box-sizing: border-box;*/
    border-radius: 4px;
    font-size: 1rem;
    font-weight: bolder;
}
</style>