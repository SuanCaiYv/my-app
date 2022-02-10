<template>
    <div class="user">
        <div class="col1">
            <Img class="img" :url="avatar"></Img>
            <input class="nickname" v-model="nickname" @change="updateNickname"/>
        </div>
        <div class="col2">
            <div>
                <input class="input" type="email" v-model="email" @change="updateEmail">
                <div class="name">邮箱</div>
            </div>
            <div>
                <input class="input" type="text" v-model="phone" @change="updatePhone">
                <div class="name">手机</div>
            </div>
            <div>
                <input class="input" type="text" v-model="location" @change="updateLocation">
                <div class="name">地址</div>
            </div>
        </div>
        <div class="col3">
            <div>
                <input class="input" type="text" v-model="qq" @change="updateQQ">
                <div class="name">Q&nbsp;Q</div>
            </div>
            <div>
                <input class="input" type="text" v-model="weChat" @change="updateWeChat">
                <div class="name">微信</div>
            </div>
            <div>
                <input class="input" type="text" v-model="github" @change="updateGitHub">
                <div class="name">猫网</div>
            </div>
        </div>
        <div class="col4">
            <div>
                <textarea class="signature" v-model="signature" @change="updateSignature"/>
            </div>
        </div>
        <div class="col5"></div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Img from "../Img.vue"
import {httpClient} from "../../net";
import storage from "../../util/storage";
import {Constant} from "../../common/systemconstant";
import alertFunc from "../../util/alert";
import {useRouter} from "vue-router";
import {Response} from "../../net";

const name = ref<string>("User")
const router = useRouter()

const avatar = ref<string>('http://127.0.0.1:8190/v1/static/a/my-avatar.png')
const nickname = ref<string>('')
const email = ref<string>('')
const phone = ref<string>('')
const location =  ref<string>('')
const qq = ref<string>('')
const weChat = ref<string>('')
const github = ref<string>('')
const signature = ref<string>('')

const accessToken = storage.get(Constant.ACCESS_TOKEN)
if (accessToken === "") {
    alertFunc("请登录", function () {
        router.push("/sign")
    })
} else {
    httpClient.get("/user/info", {}, true, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        } else {
            avatar.value = resp.data.avatar
            nickname.value = resp.data.nickname
            email.value = resp.data.email
            phone.value = resp.data.phone
            location.value = resp.data.location
            qq.value = resp.data.qq
            weChat.value = resp.data.we_chat
            github.value = resp.data.github
            signature.value = resp.data.signature
        }
    })
}
const updateNickname = function () {
    httpClient.put("/user/info", {}, {
        nickname: nickname.value
    }, true, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        }
    })
}
const updateEmail = function () {
    httpClient.put("/user/info", {}, {
        email: email.value
    }, true, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        }
    })
}
const updatePhone = function () {
    httpClient.put("/user/info", {}, {
        phone: phone.value
    }, true, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        }
    })
}
const updateLocation = function () {
    httpClient.put("/user/info", {}, {
        location: location.value
    }, true, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        }
    })
}
const updateQQ = function () {
    httpClient.put("/user/info", {}, {
        qq: qq.value
    }, true, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        }
    })
}
const updateWeChat = function () {
    httpClient.put("/user/info", {}, {
        we_chat: weChat.value
    }, true, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        }
    })
}
const updateGitHub = function () {
    httpClient.put("/user/info", {}, {
        github: github.value
    }, true, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        }
    })
}
const updateSignature = function () {
    httpClient.put("/user/info", {}, {
        signature: signature.value
    }, true, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        }
    })
}
</script>

<style scoped>
.user {
    width: 100%;
    height: 100%;
    grid-area: user;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    display: grid;
    grid-template-areas: "col1 col2 col3 col4 col5";
    grid-template-columns: 200px 330px 330px 330px 1fr;
}

.col1 {
    width: 100%;
    height: 100%;
    grid-area: col1;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
}

.col2 {
    width: 100%;
    height: 100%;
    grid-area: col2;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
}

.col3 {
    width: 100%;
    height: 100%;
    grid-area: col3;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
}

.col4 {
    width: 100%;
    height: 100%;
    grid-area: col4;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
}

.col5 {
    width: 100%;
    height: 100%;
    grid-area: col5;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
}

.img {
    width: 150px;
    height: 150px;
    margin: 25px;
}

.img:active {
    opacity: 70%;
}

.nickname {
    width: 150px;
    height: 50px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    border: none;
    border-radius: 16px;
    margin-top: 15px;
    margin-left: 25px;
    margin-right: 25px;
    overflow-x: auto;
    font-size: 1rem;
    text-align: center;
    line-height: 50px;
    background-color: rgba(255,0,49,0.05);
}

.input {
    width: 220px;
    height: 50px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    border: none;
    border-radius: 16px 0 0 16px;
    margin-top: 20px;
    margin-bottom: 20px;
    padding: 0;
    font-size: 1rem;
    display: inline-block;
    vertical-align: top;
    background-color: rgba(0,167,255,0.05);
}

.name {
    width: 60px;
    height: 50px;
    border: none;
    border-radius: 0 16px 16px 0;
    margin-top: 20px;
    margin-bottom: 20px;
    font-size: 1.2rem;
    display: inline-block;
    vertical-align: top;
    line-height: 50px;
    background-color: rgba(0,63,121,0.11);
}

.signature {
    width: 280px;
    height: 145px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    border: none;
    border-radius: 16px;
    margin-top: 20px;
    margin-bottom: 20px;
    font-size: 1.2rem;
    display: inline-block;
    vertical-align: bottom;
    background-color: rgba(29,0,255,0.05);
}
</style>