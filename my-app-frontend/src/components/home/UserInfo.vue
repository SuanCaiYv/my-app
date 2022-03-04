<template>
    <div class="userInfo">
        <div class="img" @click="router.push('/about')">
            <img :src="avatar" style="border-radius: 150px; width: 100%; height: 100%; object-fit: cover"/>
        </div>
        <!--最多七个字或10个大写字母-->
        <div class="nickname">
            {{nickname}}
        </div>
        <!--最多20个字符-->
        <div class="email">
            <a :href="mailto" style="text-decoration: none; color: inherit">{{email}}</a>
        </div>
        <!--同上-->
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
import {useRouter} from "vue-router";
import {BASE_URL, httpClient} from "../../net";
import {Response} from "../../common/interface";
import alertFunc from "../../util/alert";

const name = ref<string>("UserInfo")

const avatar = ref<string>(BASE_URL + '/static/a/default-avatar.png')
const nickname = ref<string>('小白白白')
const email = ref<string>('codewithbuff@163.com')
const github = ref<string>('https://github.com/SuanCaiYv')
const signature = ref<string>('Gin+Vue3')
const mailto = ref<string>("mailto:" + email.value)

httpClient.get("/user/info/no_auth", {}, false, function (resp: Response) {
    if (!resp.ok) {
        alertFunc(resp.errMsg, function () {})
    } else {
        // @ts-ignore
        avatar.value = resp.data.avatar
        // @ts-ignore
        nickname.value = resp.data.nickname
        // @ts-ignore
        email.value = resp.data.email
        // @ts-ignore
        github.value = resp.data.github
        // @ts-ignore
        signature.value = resp.data.signature
    }
})

const router = useRouter()
</script>

<style scoped>
.userInfo {
    width: 200px;
    height: 100%;
    border: none;
    box-sizing: border-box;
    border-radius: 20px;
    position: fixed;
    top: 140px;
    background-color: rgba(0,0,0,0.05);
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
}

.img:active {
    opacity: 70%;
}

.nickname {
    width: 100%;
    height: auto;
    margin-top: 25px;
    margin-left: auto;
    margin-right: auto;
    padding: 0;
    border-radius: 6px;
    font-size: 1rem;
    line-height: 20px;
    font-weight: bolder;
    color: slateblue;
}

.email {
    width: 100%;
    height: auto;
    margin-top: 25px;
    margin-left: auto;
    margin-right: auto;
    padding: 0;
    border-radius: 6px;
    font-size: 1rem;
    font-weight: bolder;
    line-height: 20px;
    word-break: break-all;
}

.github {
    width: 200px;
    height: auto;
    margin-top: 25px;
    margin-left: auto;
    margin-right: auto;
    padding: 0;
    border-radius: 6px;
    font-size: 1rem;
    color: black;
    font-weight: bolder;
    line-height: 20px;
    word-break: break-all;
}

.signature {
    width: 100%;
    height: auto;
    margin-top: 25px;
    margin-left: auto;
    margin-right: auto;
    padding: 0;
    border-radius: 4px;
    font-size: 1rem;
    font-weight: bolder;
}
</style>