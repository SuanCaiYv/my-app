<template>
    <div class="main">
        <PH1></PH1>
        <div class="l1">
            <div class="inline-block">
                <label class="select" for="coverImg">选择封面</label>
                <input style="display: none" type="file" id="coverImg" @change="uploadCovImg($event)">
                <div class="show">{{covImgPath}}</div>
            </div>
            <div class="inline-block">
                <button class="select" style="text-align: left" @click="updateVisibly">{{visiblyButton}}</button>
                <div class="show">{{visiblyMsg}}</div>
            </div>
        </div>
        <div class="l2">
            <div class="inline-block">
                <select class="select">
                    <option selected>分类</option>
                    <option  v-for="item in kindList" :value="item">{{item}}</option>
                </select>
                <input class="input" type="text" placeholder="键入并回车以新建分类" v-model="newKind" @keydown.enter.down="createKind"/>
            </div>
            <div class="inline-block">
                <select class="select">
                    <option selected>标签</option>
                    <option  v-for="item in kindList" :value="item">{{item}}</option>
                </select>
                <input class="input" type="text" placeholder="键入并回车以新建标签" v-model="newTag" @keydown.enter.down="createTag"/>
            </div>
        </div>
        <div class="l3">
            <div class="inline-block">
                <select class="select">
                    <option selected>回退</option>
                    <option  v-for="item in rollbackList" :value="item">{{item}}</option>
                </select>
                <div class="show" style="width: 900px"></div>
            </div>
        </div>
        <div class="l4">
            <textarea class="textarea" placeholder="来点文章简介吧！"></textarea>
        </div>
        <div class="l5">
            <button class="select" style="width: 300px;margin-left: auto;margin-right: auto" @click="publish">Done</button>
        </div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import PH1 from "../placeholder/PH1.vue";
import alertFunc from "../../util/alert";

const name = ref<string>("Main")

const props = defineProps({
    close: Function,
})

const kindList = ref<Array<string>>([])
const tagList = ref<Array<string>>([])
const rollbackList = ref<Array<string>>([])

kindList.value.push("bbb")
kindList.value.push("aaa")
kindList.value.push("ccc")
tagList.value.push("ddd")
tagList.value.push("eee")
tagList.value.push("fff")
rollbackList.value.push("aaabbbcccdddeeefffggghhhiii")
rollbackList.value.push("aaabbbcccdddeeefffggghhhiii")
rollbackList.value.push("aaabbbcccdddeeefffggghhhiii")

let covImg = ''
let covImgPath = ref<string>('')
let visibly = ref<boolean>(false)
let visiblyMsg = ref<string>('此文章仅作者可见')
let visiblyButton = ref<string>('私密')
let newKind = ref<string>('')
let newTag = ref<string>('')

const uploadCovImg = function (event: Event) {
    // @ts-ignore
    covImg = event.target.files[0].fullPath
    // @ts-ignore
    covImgPath.value = covImg.mozFullpath
}

const updateVisibly = function () {
    visibly.value = !visibly.value
    if (visibly.value === true) {
        visiblyButton.value = "公开"
        visiblyMsg.value = "此文章将会公开供所有人可见"
    } else if (visibly.value === false) {
        visiblyMsg.value = "此文章仅作者可见"
        visiblyButton.value = "私密"
    }
}

const createKind = function () {
    console.log(newKind)
    alertFunc(newKind.value, function () {})
}

const createTag = function () {}

const publish = function () {
    // @ts-ignore
    props.close()
}
</script>

<style scoped>
.main {
    width: 100%;
    height: 460px;
    grid-area: main;
    border: 2px solid mediumpurple;
    box-sizing: border-box;
    border-radius: 20px;
    display: grid;
    grid-template-areas:
        "ph1"
        "l1"
        "l2"
        "l3"
        "l4"
        "l5";
    grid-template-rows: 20px 60px 60px 60px 200px 60px;
}

.l1 {
    width: 100%;
    height: 100%;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
}

.l2 {
    width: 100%;
    height: 100%;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
}

.l3 {
    width: 100%;
    height: 100%;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
}

.l4 {
    width: 100%;
    height: 100%;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
}

.l5 {
    width: 100%;
    height: 100%;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
}

.inline-block {
    display: inline-block;
}

.select {
    width: 80px;
    height: 40px;
    border: none;
    box-sizing: border-box;
    border-radius: 30px;
    margin-top: 10px;
    margin-bottom: 10px;
    display: inline-block;
    vertical-align: bottom;
    font-size: 1rem;
    line-height: 40px;
    appearance: none;
    -webkit-appearance: none;
    -moz-appearance: none;
    outline: none;
    background-color: seashell;
}

.select[type=file] {
    display: none;
}

.select:hover {
    background-color: silver;
}

.select:active {
    background-color: gainsboro;
}

.show {
    width: 400px;
    height: 40px;
    border: 1px solid lightpink;
    box-sizing: border-box;
    border-radius: 16px;
    margin: 10px 10px 10px 8px;
    display: inline-block;
    vertical-align: bottom;
    text-align: left;
    font-size: 1rem;
    line-height: 40px;
}

.input {
    width: 400px;
    height: 40px;
    border: 1px solid lightpink;
    box-sizing: border-box;
    border-radius: 16px;
    margin: 10px 10px 10px 8px;
    padding: 0;
    display: inline-block;
    vertical-align: bottom;
    text-align: left;
    line-height: 40px;
}

.textarea {
    width: 980px;
    height: 180px;
    border: 1px solid lightcoral;
    box-sizing: border-box;
    border-radius: 8px;
    margin: 10px auto auto;
    padding: 0;
    font-size: 0.8rem;
}
</style>