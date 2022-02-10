<template>
    <div class="userInfo">
        <div class="img" @click="router.push('/about')">
            <Img :url="avatar" />
        </div>
        <div class="nickname" :style="{fontSize: nicknameFontSize + 'px'}">{{nickname}}</div>
        <div class="email">
            <a :href="mailto" style="text-decoration: none; color: inherit" :style="{fontSize: emailFontSize + 'px'}">{{email}}</a>
        </div>
        <div class="github" :style="{fontSize: githubFontSize + 'px'}">
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
import {httpClient, Response} from "../../../net";
import alertFunc from "../../../util/alert";

const name = ref<string>("UserInfo")
const router = useRouter()

const avatar = ref<string>('http://127.0.0.1:8190/v1/static/a/my-avatar.png')
const nickname = ref<string>('小白白白')
const email = ref<string>('codewithbuff@163.com')
const github = ref<string>('https://github.com/SuanCaiYv')
const signature = ref<string>('Gin+Vue3')
const mailto = ref<String>("mailto:" + email.value)

const nicknameFontSize = ref<number>(12)
const emailFontSize = ref<number>(12)
const githubFontSize = ref<number>(12)

httpClient.get("/user/info", {}, true, function (resp: Response) {
    if (!resp.ok) {
        alertFunc(resp.errMsg, function () {})
    } else {
        avatar.value = resp.data.avatar
        nickname.value = resp.data.nickname
        email.value = resp.data.email
        location.value = resp.data.location
        github.value = resp.data.github
        signature.value = resp.data.signature
        nicknameFontSize.value = Math.min(Math.floor(150 * 1.6 / nickname.value.length), 20)
        emailFontSize.value = Math.min(Math.floor(196 * 1.6 / email.value.length), 20)
        githubFontSize.value = Math.min(Math.floor(196 * 1.8 / github.value.length), 20)
    }
})
</script>

<style scoped>
.userInfo {
    width: 100%;
    height: 100%;
    grid-area: userInfo;
    border: 2px solid salmon;
    box-sizing: border-box;
    border-radius: 20px;
}

.img {
    width: 150px;
    height: 150px;
    margin: 25px 25px 0;
    padding: 0;
    border: 1px solid white;
    box-sizing: border-box;
    border-radius: 6px;
}

.img:active {
    opacity: 70%;
}

.nickname {
    width: 150px;
    height: 20px;
    margin: 25px 25px 0;
    padding: 0;
    /*border: 1px solid black;*/
    /*box-sizing: border-box;*/
    /*border-radius: 6px;*/
    line-height: 20px;
}

.email {
    width: 196px;
    height: 20px;
    margin: 25px 0 0;
    padding: 0;
    /*border: 1px solid black;*/
    /*box-sizing: border-box;*/
    /*border-radius: 6px;*/
    line-height: 20px;
}

.github {
    width: 196px;
    height: 20px;
    margin: 25px 0 0;
    padding: 0;
    /*border: 1px solid black;*/
    /*box-sizing: border-box;*/
    /*border-radius: 6px;*/
    line-height: 20px;
}

.signature {
    width: 196px;
    height: auto;
    margin: 25px 0 0;
    padding: 0;
    /*border: 1px solid black;*/
    /*box-sizing: border-box;*/
    border-radius: 4px;
    font-size: 1rem;
    text-align: center;
}
</style>