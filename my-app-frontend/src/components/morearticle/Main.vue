<template>
    <div class="main">
        <PH1></PH1>
        <div class="l1">
            <div class="inline-block">
                <label class="select" for="documentFile">选择文档</label>
                <input style="display: none" type="file" id="documentFile">
                <div class="show"></div>
            </div>
            <div class="inline-block">
                <label class="select" for="coverImg">选择封面</label>
                <input style="display: none" type="file" id="coverImg">
                <div class="show"></div>
            </div>
        </div>
        <div class="l2">
            <div class="inline-block">
                <select class="select">
                    <option selected>分类</option>
                    <option  v-for="item in kindList" :value="item">{{item}}</option>
                </select>
                <input class="input" type="text" placeholder="键入并回车以新建分类"/>
            </div>
            <div class="inline-block">
                <select class="select">
                    <option selected>标签</option>
                    <option  v-for="item in kindList" :value="item">{{item}}</option>
                </select>
                <input class="input" type="text" placeholder="键入并回车以新建标签"/>
            </div>
        </div>
        <div class="l3">
            <div class="inline-block">
                <select class="select">
                    <option selected>公开</option>
                    <option>私密</option>
                </select>
                <div class="show" style="border: 0"/>
            </div>
            <div class="inline-block">
                <select class="select">
                    <option selected>回退</option>
                    <option  v-for="item in rollbackList" :value="item">{{item}}</option>
                </select>
                <div class="show" style="border: 0"/>
            </div>
        </div>
        <div class="l4">
            <textarea class="textarea"></textarea>
        </div>
        <div class="l5">
            <button class="select" style="width: 300px;margin-left: auto;margin-right: auto">Done</button>
        </div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import PH1 from "../placeholder/PH1.vue";

const name = ref<String>("Main")

const kindList = ref<Array<String>>([])
const tagList = ref<Array<String>>([])
const rollbackList = ref<Array<String>>([])

kindList.value.push("bbb")
kindList.value.push("aaa")
kindList.value.push("ccc")
tagList.value.push("ddd")
tagList.value.push("eee")
tagList.value.push("fff")
rollbackList.value.push("aaabbbcccdddeeefffggghhhiii")
rollbackList.value.push("aaabbbcccdddeeefffggghhhiii")
rollbackList.value.push("aaabbbcccdddeeefffggghhhiii")
</script>

<style scoped>
.main {
    width: 100%;
    height: 100%;
    grid-area: main;
    border: 1px solid silver;
    box-sizing: border-box;
    display: grid;
    grid-template-areas:
        "ph1"
        "l1"
        "l2"
        "l3"
        "l4"
        "l5";
    grid-template-rows: 120px 60px 60px 60px 200px 60px;
}

.l1 {
    width: 100%;
    height: 100%;
    border: 1px solid silver;
    box-sizing: border-box;
}

.l2 {
    width: 100%;
    height: 100%;
    border: 1px solid silver;
    box-sizing: border-box;
}

.l3 {
    width: 100%;
    height: 100%;
    border: 1px solid silver;
    box-sizing: border-box;
}

.l4 {
    width: 100%;
    height: 100%;
    border: 1px solid silver;
    box-sizing: border-box;
}

.l5 {
    width: 100%;
    height: 100%;
    border: 1px solid silver;
    box-sizing: border-box;
}

.inline-block {
    display: inline-block;
}

.select {
    width: 80px;
    height: 40px;
    border: 1px solid silver;
    box-sizing: border-box;
    border-radius: 30px;
    margin-top: 10px;
    margin-bottom: 10px;
    display: inline-block;
    vertical-align: bottom;
    font-size: 1rem;
    line-height: 40px;
    background-color: white;
}

.select[type=file] {
    display: none;
}

.show {
    width: 400px;
    height: 40px;
    border: 1px solid silver;
    box-sizing: border-box;
    border-radius: 16px;
    margin-top: 10px;
    margin-bottom: 10px;
    margin-left: 8px;
    margin-right: 10px;
    display: inline-block;
    vertical-align: bottom;
    text-align: left;
    line-height: 40px;
}

.input {
    width: 400px;
    height: 40px;
    border: 1px solid silver;
    box-sizing: border-box;
    border-radius: 16px;
    margin-top: 10px;
    margin-bottom: 10px;
    margin-left: 8px;
    margin-right: 10px;
    display: inline-block;
    vertical-align: bottom;
    text-align: left;
    line-height: 40px;
}

.textarea {
    width: 980px;
    height: 180px;
    border: 1px solid silver;
    box-sizing: border-box;
    border-radius: 8px;
    margin-top: 10px;
    margin-bottom: auto;
    margin-left: auto;
    margin-right: auto;
    font-size: 0.8rem;
}
</style>