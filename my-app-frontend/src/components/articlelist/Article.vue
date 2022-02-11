<template>
    <div class="article">
        <PH1></PH1>
        <div class="title"></div>
        <div class="content"></div>
        <PH2></PH2>
        <div class="update">
            <button class="button">更新</button>
        </div>
        <div class="delete">
            <button class="button">删除</button>
        </div>
        <div class="visible">
            <button class="button" @click="setVisibly">{{visibly}}</button>
        </div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import PH1 from "../placeholder/PH1.vue"
import PH2 from "../placeholder/PH2.vue"

const name = ref<String>("Article")

let visibly = ref<String>("公开")

const setVisibly = function () {
    visibly.value = "私密"
}
</script>

<style scoped>
.article {
    width: 100%;
    height: 280px;
    display: grid;
    grid-template-areas:
        "ph1 title title title"
        "ph1 content ph2 update"
        "ph1 content ph2 delete"
        "ph1 content ph2 visible";
    grid-template-columns: 25px 1fr 25px 200px;
    grid-template-rows: 40px 80px 80px 80px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    margin-top: 25px;
}

.title {
    width: 320px;
    height: 100%;
    grid-area: title;
    border: 2px solid wheat;
    box-sizing: border-box;
    border-radius: 16px 16px 0 0;
    background-color: antiquewhite;
}

.content {
    width: 100%;
    height: 100%;
    grid-area: content;
    border: 2px solid wheat;
    box-sizing: border-box;
    border-radius: 0 16px 16px 16px;
    display: inline-block;
    background-color: antiquewhite;
}

.update {
    width: 100%;
    height: 100%;
    grid-area: update;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    display: inline-block;
}

.delete {
    width: 100%;
    height: 100%;
    grid-area: delete;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    display: inline-block;
}

.visible {
    width: 100%;
    height: 100%;
    grid-area: visible;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    display: inline-block;
}

.button {
    width: 120px;
    height: 40px;
    margin-top: 20px;
    margin-bottom: 20px;
    margin-left: -80px;
    border: 2px solid lightgray;
    padding: 0;
    border-radius: 18px;
    font-size: 1.2rem;
    background-color: white;
}

.button:hover {
    background-color: lightgray;
}

.button:active {
    background-color: gainsboro;
}
</style>