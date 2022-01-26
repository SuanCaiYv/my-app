<template>
    <div class="user">
        <div class="col1">
            <Img class="img" url="http://127.0.0.1:8190/v1/static/a/default-avatar.png"></Img>
            <input class="nickname" :value="nickname" />
        </div>
        <div class="col2">
            <div>
                <input class="input" type="text" placeholder="更新邮箱" @keydown.enter.down="testFunc">
            </div>
            <div>
                <input class="input">
            </div>
            <div>
                <input class="input">
            </div>
        </div>
        <div class="col3">
            <div>
                <input class="input">
            </div>
            <div>
                <input class="input">
            </div>
            <div>
                <input class="input">
            </div>
        </div>
        <div class="col4">
            <div>
                <textarea class="signature" />
            </div>
        </div>
        <div class="col5"></div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Img from "../Img.vue"

const name = ref<String>("User")

const nickname = ref<String>("nickname")

const testFunc = function () {
    console.log("aaa")
}
</script>

<style scoped>
.user {
    width: 100%;
    height: 100%;
    grid-area: user;
    border: 1px solid silver;
    box-sizing: border-box;
    display: grid;
    grid-template-areas: "col1 col2 col3 col4 col5";
    grid-template-columns: 200px 330px 330px 330px 1fr;
}

.col1 {
    width: 100%;
    height: 100%;
    grid-area: col1;
    border: 1px solid silver;
    box-sizing: border-box;
}

.col2 {
    width: 100%;
    height: 100%;
    grid-area: col2;
    border: 1px solid silver;
    box-sizing: border-box;
}

.col3 {
    width: 100%;
    height: 100%;
    grid-area: col3;
    border: 1px solid silver;
    box-sizing: border-box;
}

.col4 {
    width: 100%;
    height: 100%;
    grid-area: col4;
    border: 1px solid silver;
    box-sizing: border-box;
}

.col5 {
    width: 100%;
    height: 100%;
    grid-area: col5;
    border: 1px solid silver;
    box-sizing: border-box;
}

.img {
    width: 150px;
    height: 150px;
    margin: 25px;
    border-radius: 6px;
}

.img:active {
    opacity: 70%;
}

.nickname {
    width: 150px;
    height: 50px;
    border: 1px solid silver;
    box-sizing: border-box;
    margin-top: 15px;
    margin-left: 25px;
    margin-right: 25px;
    border-radius: 8px;
    overflow-x: auto;
    font-size: 1.2rem;
    text-align: center;
    line-height: 50px;
}

.input {
    width: 280px;
    height: 50px;
    border: 1px solid silver;
    box-sizing: border-box;
    border-radius: 6px;
    margin-top: 20px;
    margin-bottom: 20px;
    font-size: 1.2rem;
    display: inline-block;
    vertical-align: bottom;
}

.signature {
    width: 280px;
    height: 145px;
    border: 1px solid silver;
    box-sizing: border-box;
    border-radius: 6px;
    margin-top: 20px;
    margin-bottom: 20px;
    font-size: 1.2rem;
    display: inline-block;
    vertical-align: bottom;
}
</style>