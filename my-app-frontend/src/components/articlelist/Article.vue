<template>
    <div class="article">
        <PH1></PH1>
        <div class="title">{{title}}</div>
        <div class="content">{{summary}}</div>
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
import {ref, watch} from "vue"
import PH1 from "../placeholder/PH1.vue"
import PH2 from "../placeholder/PH2.vue"
import {marked} from "marked";

const name = ref<string>("Article")
const props = defineProps({
    title: String,
    summary: String,
})
let visibly = ref<string>("公开")

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
    margin-top: 25px;
}

.title {
    width: 40%;
    height: 100%;
    grid-area: title;
    border: none;
    box-sizing: border-box;
    border-radius: 16px 16px 0 0;
    padding: 0 0 0 8px;
    font-size: 1.4rem;
    font-weight: bolder;
    line-height: 40px;
    text-align: left;
    background-color: rgba(0,0,0,0.05);
}

.content {
    min-width: 600px;
    height: 100%;
    grid-area: content;
    border: none;
    box-sizing: border-box;
    border-radius: 0 16px 16px 16px;
    display: inline-block;
    font-size: 1.2rem;
    text-align: left;
    word-break: break-word;
    overflow: hidden;
    background-color: rgba(0,0,0,0.03);
}

.update {
    width: 100%;
    height: 100%;
    grid-area: update;
    display: inline-block;
}

.delete {
    width: 100%;
    height: 100%;
    grid-area: delete;
    display: inline-block;
}

.visible {
    width: 100%;
    height: 100%;
    grid-area: visible;
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
    font-weight: bolder;
    background-color: white;
}

.button:hover {
    background-color: lightgray;
}

.button:active {
    background-color: gainsboro;
}
</style>