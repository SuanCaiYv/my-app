<template>
    <div class="user">
        <PH1></PH1>
        <div class="col1">
            <label for="avatar">
                <img class="img" :src="avatar"/>
            </label>
            <input type="file" id="avatar" hidden @change="updateAvatar($event)">
            <input class="nickname" v-model="nickname" @keydown.enter.down="updateNickname"/>
        </div>
        <div class="col2">
            <div>
                <input class="input" type="email" v-model="email" @keydown.enter.down="updateEmail">
                <div class="name">邮箱</div>
            </div>
            <div>
                <input class="input" type="text" v-model="phone" @keydown.enter.down="updatePhone">
                <div class="name">手机</div>
            </div>
            <div>
                <input class="input" type="text" v-model="location" @keydown.enter.down="updateLocation">
                <div class="name">地址</div>
            </div>
        </div>
        <div class="col3">
            <div>
                <input class="input" type="text" v-model="qq" @keydown.enter.down="updateQQ">
                <div class="name">Q&nbsp;Q</div>
            </div>
            <div>
                <input class="input" type="text" v-model="weChat" @keydown.enter.down="updateWeChat">
                <div class="name">微信</div>
            </div>
            <div>
                <input class="input" type="text" v-model="github" @keydown.enter.down="updateGitHub">
                <div class="name">猫网</div>
            </div>
        </div>
        <div class="col4">
            <div>
                <textarea class="signature" v-model="signature" @change="updateSignature"/>
            </div>
        </div>
        <PH2></PH2>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {baseUrl, httpClient} from "../../net";
import storage from "../../util/storage";
import {Constant} from "../../common/systemconstant";
import alertFunc from "../../util/alert";
import {useRouter} from "vue-router";
import {Response} from "../../common/interface";
import PH1 from "../placeholder/PH1.vue"
import PH2 from "../placeholder/PH2.vue"
import {parseBoolean} from "../../util/base";

const name = ref<string>("User")
const router = useRouter()

const avatar = ref<string>('http://127.0.0.1:8190/v1/static/a/default-avatar.png')
const nickname = ref<string>('')
const email = ref<string>('')
const phone = ref<string>('')
const location =  ref<string>('')
const qq = ref<string>('')
const weChat = ref<string>('')
const github = ref<string>('')
const signature = ref<string>('')

if (parseBoolean(storage.get(Constant.AUTHENTICATED))) {
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

const updateAvatar = function (event: Event) {
    // @ts-ignore
    let file = event.target.files[0]
    let form = new FormData()
    form.append("file", file)
    httpClient.upload("/static/file", {
        archive: "avatar",
    }, form, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        } else {
            let url = baseUrl + "/static/a/" + resp.data.filename
            httpClient.put("/user/info", {
                archive: "avatar"
            }, {
                avatar: url
            }, true, function (resp: Response) {
                if (!resp.ok) {
                    alertFunc(resp.errMsg, function () {})
                } else {
                    router.go(0)
                }
            })
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
    display: grid;
    grid-template-areas: "ph1 col1 col2 col3 col4 ph2";
    grid-template-columns: 1fr 200px 330px 330px 330px 1fr;
}

.col1 {
    width: 100%;
    height: 100%;
    grid-area: col1;
}

.col2 {
    width: 100%;
    height: 100%;
    grid-area: col2;
}

.col3 {
    width: 100%;
    height: 100%;
    grid-area: col3;
}

.col4 {
    width: 100%;
    height: 100%;
    grid-area: col4;
}

.img {
    width: 150px;
    height: 150px;
    margin: 25px;
    border: none;
    border-radius: 150px;
    box-sizing: border-box;
    object-fit: cover;
}

.img:active {
    opacity: 70%;
}

.nickname {
    width: 150px;
    height: 50px;
    border: none;
    box-sizing: border-box;
    border-radius: 16px;
    margin-top: 0;
    margin-left: 25px;
    margin-right: 25px;
    overflow-x: auto;
    font-size: 1rem;
    font-weight: bolder;
    text-align: center;
    line-height: 50px;
    outline: none;
    background-color: rgba(0,0,0,0.05);
}

.input {
    width: 220px;
    height: 50px;
    border: none;
    box-sizing: border-box;
    border-radius: 16px 0 0 16px;
    margin-top: 20px;
    margin-bottom: 20px;
    padding-left: 8px;
    font-size: 1rem;
    font-weight: bold;
    display: inline-block;
    vertical-align: top;
    background-color: rgba(0,0,0,0.05);
    outline: none;
}

.name {
    width: 60px;
    height: 50px;
    border: none;
    border-radius: 0 16px 16px 0;
    margin-top: 20px;
    margin-bottom: 20px;
    font-size: 1.2rem;
    font-weight: bolder;
    display: inline-block;
    vertical-align: top;
    line-height: 50px;
    background-color: rgba(0,0,0,0.1);
}

.signature {
    width: 280px;
    height: 145px;
    border: none;
    box-sizing: border-box;
    border-radius: 16px;
    margin-top: 20px;
    margin-bottom: 20px;
    padding: 8px;
    font-size: 1.2rem;
    font-weight: bolder;
    font-style: oblique;
    display: inline-block;
    vertical-align: bottom;
    outline: none;
    resize: none;
    background-color: rgba(0,0,0,0.05);
}
</style>