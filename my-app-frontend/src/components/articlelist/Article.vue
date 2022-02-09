<template>
    <div class="article">
        <div class="title"></div>
        <div class="content"></div>
        <div class="update">
            <button class="button">更新</button>
        </div>
        <div class="delete">
            <button class="button">删除</button>
        </div>
        <div class="visible">
            <button class="button">{{visibly}}</button>
        </div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Article")

let visibly = ref<String>("公开")
</script>

<style scoped>
.article {
    width: 100%;
    height: 320px;
    grid-area: article;
    display: grid;
    grid-template-areas:
        "title title"
        "content update"
        "content delete"
        "content visible";
    grid-template-columns: 1fr 200px;
    grid-template-rows: 40px 80px 80px 80px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    margin-top: 40px;
}

.title {
    width: 350px;
    height: 100%;
    grid-area: title;
    /*border: 1px solid black;*/
    /*box-sizing: border-box;*/
    border-radius: 16px;
    background-color: darkorange;
    opacity: 10%;
}

.content {
    width: 100%;
    height: 100%;
    grid-area: content;
    box-sizing: border-box;
    border-radius: 16px;
    display: inline-block;
    background-color: orange;
    opacity: 10%;
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
    width: 80px;
    height: 40px;
    margin-top: 20px;
    margin-bottom: 20px;
    border: none;
    border-radius: 18px;
    font-size: 1.2rem;
}
</style>