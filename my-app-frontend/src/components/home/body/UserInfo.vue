<template>
    <div class="userInfo">
        <div class="img" @click="router.push('/about')">
            <Img url="http://127.0.0.1:8190/v1/static/a/my-avatar.png" />
        </div>
        <div class="nickname">
            {{nickname}}
        </div>
        <div class="email">
            <a :href="mailto" style="text-decoration: none; color: inherit">{{email}}</a>
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

const name = ref<String>("UserInfo")

const props = defineProps({
    nickname: String,
    email: String,
    signature: String
})

const mailto = ref<String>("mailto:" + props.email)

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
    border-radius: 6px;
    font-size: 1rem;
    line-height: 20px;
}

.email {
    width: 200px;
    height: 20px;
    margin: 25px 0 0;
    padding: 0;
    /*border: 1px solid black;*/
    /*box-sizing: border-box;*/
    border-radius: 6px;
    font-size: 1rem;
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
}
</style>