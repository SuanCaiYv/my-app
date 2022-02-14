<template>
    <div class="publish">
        <div class="notification"><span :style="{ color: notificationColor }">{{notification}}</span></div>
        <div class="cov-img">
            <label class="label" for="coverImg"><span>选择封面</span></label>
            <input class="input" type="file" id="coverImg" @change="commitCovImg($event)">
            <div class="show-value">{{ covImgPath }}</div>
        </div>
        <div class="visibility">
            <button class="button" @click="commitVisibility"><span>{{ visibilityFlag }}</span></button>
            <div class="show-value">{{ visibilityMsg }}</div>
        </div>
        <div class="kind-list">
            <select class="select" v-model="chosenKind" @change="commitKindList">
                <option value="" selected>选择分类</option>
                <option v-for="item in kindList" :value="item.id">{{ item.value }}</option>
            </select>
            <input class="input" type="text" placeholder="键入并回车以新建分类" v-model="newKind"
                   @keydown.enter.down="createKind"/>
        </div>
        <div class="chosen-kind-list">
            <div class="full-show-value">
                <ChoiceItem v-for="item in chosenKindList" :id="item.id" :value="item.value"
                            :delete-func="cancelKindList"></ChoiceItem>
            </div>
        </div>
        <div class="tag-list">
            <select class="select" v-model="chosenTag" @change="commitTagList">
                <option value="" selected>选择标签</option>
                <option v-for="item in tagList" :value="item.id">{{ item.value }}</option>
            </select>
            <input class="input" type="text" placeholder="键入并回车以新建标签" v-model="newTag" @keydown.enter.down="createTag"/>
        </div>
        <div class="chosen-tag-list">
            <div class="full-show-value">
                <ChoiceItem v-for="item in chosenTagList" :id="item.id" :value="item.value"
                            :delete-func="cancelTagList"></ChoiceItem>
            </div>
        </div>
        <div class="rollback">
            <select class="select" v-model="chosenRollback">
                <option value="" selected>选择回滚</option>
                <option v-for="item in rollbackList" :value="item.id">{{ item.value }}</option>
            </select>
            <div class="fill-all-show-value">{{ chosenRollback }}</div>
        </div>
        <div class="brief">
            <textarea class="textarea" placeholder="来点文章简介吧！"></textarea>
        </div>
        <div class="done">
            <button class="button done-button"><span>完成</span></button>
        </div>
    </div>
    <div class="publish-mask" @click="close"></div>
</template>

<script setup lang="ts">
import {inject, reactive, ref, watch} from "vue"
import ChoiceItem from "./ChoiceItem.vue"
import {IdValue} from "../../common/interface";
import {httpClient} from "../../net";
import {Response} from "../../common/interface";
import alertFunc from "../../util/alert";

const name = ref<string>("Publish")

const props = defineProps({
    divNode: Node
})

const id = inject("id")
const title = inject("title")
const content = inject("content")

const notification = ref<string>("")
const notificationColor = ref<string>("blue")
let covImg = null
let covImgPath = ref<string>("")

let visibility = ref<boolean>(false)
let visibilityFlag = ref<string>("私密文章")
let visibilityMsg = ref<string>("此文章仅允许作者阅览")

const kindList = reactive<Array<IdValue>>([])
const chosenKindList = reactive<Array<IdValue>>([])
const chosenKind = ref<string>('')
const kindMap = new Map<string, string>()
const newKind = ref<string>('')

const tagList = reactive<Array<IdValue>>([])
const chosenTagList = reactive<Array<IdValue>>([])
const chosenTag = ref<string>('')
const tagMap = new Map<string, string>()
const newTag = ref<string>('')

const rollbackList = reactive<Array<IdValue>>([])
const chosenRollback = ref<string>('')

class IdAndValue implements IdValue {
    id: string
    value: string

    constructor(id0: string, value0: string | undefined) {
        this.id = id0
        if (value0 === undefined) {
            this.value = ""
        } else {
            this.value = value0
        }
    }
}

const clearNotification = function () {
    setTimeout(() => {
        notification.value = ""
    }, 2000)
}

const commitCovImg = function (event: Event) {
    // @ts-ignore
    covImg = event.target.files[0]
    // @ts-ignore
    covImgPath.value = covImg.name
}
const commitVisibility = function () {
    visibility.value = !visibility.value
    if (visibility.value) {
        visibilityFlag.value = "公开文章"
        visibilityMsg.value = "此文章允许所有人阅览"
    } else {
        visibilityFlag.value = "私密文章"
        visibilityMsg.value = "此文章仅允许作者阅览"
    }
}
const commitKindList = function () {
    chosenKindList.splice(0, chosenKindList.length)
    chosenKindList.push(new IdAndValue(chosenKind.value, kindMap.get(chosenKind.value)))
}
const cancelKindList = function (id: string) {
    chosenKindList.splice(chosenKindList.findIndex(item => item.id === id), 1)
}
const createKind = function () {
    if (newKind.value === "") {
        notificationColor.value = "red"
        notification.value = "分类不可为空"
        clearNotification()
        return
    }
    httpClient.post("/article/kind", {}, {
        kind_name: newKind.value
    }, true, function (resp: Response) {
        if (!resp.ok) {
            notificationColor.value = "red"
            notification.value = resp.errMsg
            clearNotification()
        } else {
            kindList.splice(0, kindList.length)
            notificationColor.value = "dodgerblue"
            notification.value = "创建分类成功"
            clearNotification()
            getKindList()
        }
    })
}
const commitTagList = function () {
    chosenTagList.push(new IdAndValue(chosenTag.value, tagMap.get(chosenTag.value)))
}
const cancelTagList = function (id: string) {
    chosenTagList.splice(chosenTagList.findIndex(item => item.id === id), 1)
}
const createTag = function () {
    if (newTag.value === "") {
        notificationColor.value = "red"
        notification.value = "标签不可为空"
        clearNotification()
        return
    }
    httpClient.post("/article/tag", {}, {
        tag_name: newTag.value
    }, true, function (resp: Response) {
        if (!resp.ok) {
            notificationColor.value = "red"
            notification.value = resp.errMsg
            clearNotification()
        } else {
            tagList.splice(0, tagList.length)
            notificationColor.value = "dodgerblue"
            notification.value = "创建标签成功"
            clearNotification()
            getTagList()
        }
    })
}

const getKindList = function () {
    httpClient.get("/article/kind_list", {
        page_num: -1,
        page_size: 10
    }, true, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {
            })
        } else {
            // @ts-ignore
            let list = resp.data.list
            for (let item of list) {
                kindList.push(new IdAndValue(item.kind_id, item.kind_name))
            }
        }
    })
}
const getTagList = function () {
    httpClient.get("/article/tag_list", {
        page_num: -1,
        page_size: 10
    }, true, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {
            })
        } else {
            // @ts-ignore
            let list = resp.data.list
            for (let item of list) {
                tagList.push(new IdAndValue(item.tag_id, item.tag_name))
            }
        }
    })
}

const getRollBackList = function () {}

const syncKindList = function () {
    for (let item of kindList) {
        kindMap.set(item.id, item.value)
    }
}

const syncTagList = function () {
    for (let item of tagList) {
        tagMap.set(item.id, item.value)
    }
}

watch(kindList, () => {
    syncKindList()
})
watch(tagList, () => {
    syncTagList()
})

getKindList()
getTagList()

console.log(tagList)

const close = function () {
    // @ts-ignore
    document.getElementById("app").removeChild(props.divNode)
}
</script>

<style scoped>
.publish {
    width: 800px;
    height: 500px;
    position: absolute;
    top: 20%;
    left: 50%;
    transform: translate(-50%, 0);
    border-radius: 20px;
    z-index: 1001;
    display: grid;
    grid-template-areas:
        "notification notification"
        "cov-img visibility"
        "kind-list chosen-kind-list"
        "tag-list chosen-tag-list"
        "rollback rollback"
        "brief brief"
        "done done";
    grid-template-columns: 1fr 1fr;
    grid-template-rows: 40px 60px 60px 60px 60px 160px 60px;
    background-color: white;
}

.notification {
    width: 100%;
    height: 100%;
    border-top-left-radius: 20px;
    border-top-right-radius: 20px;
    grid-area: notification;
    display: inline-block;
    font-size: 1rem;
    font-weight: bolder;
    line-height: 40px;
}

.cov-img {
    width: 100%;
    height: 100%;
    grid-area: cov-img;
    display: inline-block;
    font-size: 1rem;
    font-weight: bolder;
}

.visibility {
    width: 100%;
    height: 100%;
    grid-area: visibility;
    display: inline-block;
    font-size: 1rem;
    font-weight: bolder;
}

.kind-list {
    width: 100%;
    height: 100%;
    grid-area: kind-list;
    display: inline-block;
    font-size: 1rem;
    font-weight: bolder;
}

.chosen-kind-list {
    width: 100%;
    height: 100%;
    grid-area: chosen-kind-list;
    display: inline-block;
    font-size: 1rem;
    font-weight: bolder;
}

.tag-list {
    width: 100%;
    height: 100%;
    grid-area: tag-list;
    display: inline-block;
    font-size: 1rem;
    font-weight: bolder;
}

.chosen-tag-list {
    width: 100%;
    height: 100%;
    grid-area: chosen-tag-list;
    display: inline-block;
    font-size: 1rem;
    font-weight: bolder;
}

.rollback {
    width: 100%;
    height: 100%;
    grid-area: rollback;
    display: inline-block;
    font-size: 1rem;
    font-weight: bolder;
}

.brief {
    width: 100%;
    height: 100%;
    grid-area: brief;
    display: inline-block;
    font-size: 1rem;
    font-weight: bolder;
}

.done {
    width: 100%;
    height: 100%;
    border-bottom-left-radius: 20px;
    border-bottom-right-radius: 20px;
    grid-area: done;
    display: inline-block;
    font-size: 1rem;
    font-weight: bolder;
}

.label {
    width: 80px;
    height: 44px;
    margin-left: 5px;
    margin-top: 8px;
    margin-bottom: 8px;
    padding: 0;
    display: inline-block;
    border: 2px solid lightgray;
    box-sizing: border-box;
    border-radius: 30px;
    vertical-align: bottom;
    line-height: 44px;
    background-color: white;
}

span {
    width: 100%;
    text-align: justify;
    text-align-last: justify;
    font-size: 1rem;
    font-weight: bolder;
}

span:after {
    width: 100%;
    height: 0;
    display: inline-block;
    visibility: hidden;
    content: '';
}

.label:hover {
    background-color: lightgray;
}

.input {
    width: 300px;
    height: 44px;
    margin: 8px 5px 8px 10px;
    padding: 0 0 0 4px;
    border: none;
    border-radius: 30px;
    box-sizing: border-box;
    display: inline-block;
    vertical-align: bottom;
    font-size: 1rem;
    font-weight: bolder;
    line-height: 44px;
    outline: none;
    background-color: rgba(0, 0, 0, 0.07);
}

.input[type=file] {
    display: none;
}

.button {
    width: 80px;
    height: 44px;
    margin-left: 5px;
    margin-top: 8px;
    margin-bottom: 8px;
    padding: 0;
    display: inline-block;
    border: 2px solid lightgray;
    box-sizing: border-box;
    border-radius: 30px;
    vertical-align: bottom;
    line-height: 44px;
    background-color: white;
}

.button:hover {
    background-color: lightgray;
}

.button:active {
    background-color: gainsboro;
}

.select {
    width: 80px;
    height: 44px;
    margin-left: 5px;
    margin-top: 8px;
    margin-bottom: 8px;
    padding: 0;
    display: inline-block;
    border: 2px solid lightgray;
    box-sizing: border-box;
    border-radius: 30px;
    vertical-align: bottom;
    line-height: 44px;
    font-size: 1rem;
    font-weight: bolder;
    outline: none;
    appearance: none;
    -webkit-appearance: none;
    -moz-appearance: none;
    background-color: white;
}

.show-value {
    width: 300px;
    height: 44px;
    border-radius: 30px;
    box-sizing: border-box;
    margin: 8px 5px 8px 10px;
    padding-left: 8px;
    display: inline-block;
    vertical-align: bottom;
    text-align-last: left;
    line-height: 44px;
    background-color: rgba(0, 0, 0, 0.07);
}

.full-show-value {
    width: 390px;
    height: 44px;
    margin: 8px 5px 8px 5px;
    border-radius: 30px;
    display: inline-block;
    vertical-align: bottom;
    text-align: left;
    overflow-x: auto;
    overflow-y: hidden;
    background-color: rgba(0, 0, 0, 0.07);
}

.fill-all-show-value {
    width: 700px;
    height: 44px;
    margin: 8px 5px 8px 10px;
    border-radius: 30px;
    display: inline-block;
    vertical-align: bottom;
    background-color: rgba(0, 0, 0, 0.07);
}

.publish-mask {
    width: 100%;
    height: 100%;
    position: absolute;
    top: 0;
    left: 0;
    opacity: 0.5;
    z-index: 1000;
    background-color: black;
}

.textarea {
    width: 784px;
    height: 144px;
    border: 2px solid antiquewhite;
    box-sizing: border-box;
    border-radius: 8px;
    margin: 8px;
    padding: 4px;
    font-size: 1rem;
    font-weight: normal;
    outline: none;
    resize: none;
}

.done-button {
    width: 120px;
    height: 44px;
}
</style>