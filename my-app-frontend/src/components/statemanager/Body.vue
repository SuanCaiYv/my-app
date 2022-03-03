<template>
    <div class="body">
        <div style="margin-top: 80px"></div>
        <div class="block-name">分类管理</div>
        <div class="kind-list">
            <div v-for="item in kindList" style="display: inline-block">
                <KindOrTagItem :id="item.id" :name="item.name" :kind="'kind'" :onDeleted="deleteKind"></KindOrTagItem>
            </div>
        </div>
        <div class="block-name">标签管理</div>
        <div class="tag-list">
            <div v-for="item in tagList" style="display: inline-block">
                <KindOrTagItem :id="item.id" :name="item.name" :kind="'tag'" :onDeleted="deleteTag"></KindOrTagItem>
            </div>
        </div>
        <div class="block-name">图片管理</div>
        <div class="img-list">
            <div v-for="item in imgList" style="display: inline-block">
                <img class="img" :src="item" @click="deleteImg(item)">
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import {reactive, ref} from "vue"
import KindOrTagItem from "../common/KindOrTagItem.vue"
import {IdName, Response} from "../../common/interface";
import {BASE_URL, httpClient} from "../../net";
import {toListResult} from "../../util/base";
import alertFunc from "../../util/alert";
import {confirmFunc} from "../../util/confirm";

const name = ref<string>("Body")
const kindList = reactive<Array<IdName>>([])
const tagList = reactive<Array<IdName>>([])
const imgList = reactive<Array<string>>([])

class IdNameClass implements IdName {
    id: string;
    name: string;

    constructor(id: string, name: string) {
        this.id = id;
        this.name = name;
    }
}

const fetchKindList = function () {
    httpClient.get("/article/kind_list", {
        page_num: -1,
        page_size: 10
    },false, function (resp: Response) {
        if (!resp.ok) {
            console.log(resp.errMsg)
        } else {
            const list = toListResult(resp.data)
            for (let l of list.list) {
                kindList.push(new IdNameClass(l.kind_id, l.kind_name))
            }
        }
    })
}

const deleteKind = function (id: string) {
    kindList.splice(kindList.findIndex(item => item.id === id), 1)
}

const deleteTag = function (id: string) {
    tagList.splice(tagList.findIndex(item => item.id === id), 1)
}

const fetchTagList = function () {
    httpClient.get("/article/tag_list", {
        page_num: -1,
        page_size: 10
    },false, function (resp: Response) {
        if (!resp.ok) {
            console.log(resp.errMsg)
        } else {
            const list = toListResult(resp.data)
            for (let l of list.list) {
                tagList.push(new IdNameClass(l.tag_id, l.tag_name))
            }
        }
    })
}

const fetchImgList = function () {
    httpClient.get("/static/file/list", {
        page_num: -1,
        page_size: 10,
        archive: "doc_img",
    },true, function (resp: Response) {
        if (!resp.ok) {
            console.log(resp.errMsg)
        } else {
            const list = toListResult(resp.data)
            for (let l of list.list) {
                imgList.push(BASE_URL + "/static/a/" + l)
            }
        }
    })
    httpClient.get("/static/file/list", {
        page_num: -1,
        page_size: 10,
        archive: "avatar",
    },true, function (resp: Response) {
        if (!resp.ok) {
            console.log(resp.errMsg)
        } else {
            const list = toListResult(resp.data)
            for (let l of list.list) {
                imgList.push(BASE_URL + "/static/a/" + l)
            }
        }
    })
}

const deleteImg = function (id: string) {
    confirmFunc("确认删除？", function () {}, function () {
        id = id.replace(BASE_URL + "/static/a/", "")
        httpClient.delete("/static/file/" + id, {},true, function (resp: Response) {
            if (!resp.ok) {
                console.log(resp.errMsg)
            } else {
                alertFunc("删除成功", function () {})
            }
        })
        imgList.splice(imgList.findIndex(item => item === id), 1)
    })
}

fetchKindList()
fetchTagList()
fetchImgList()
</script>

<style scoped>
.body {
    width: 100%;
    height: 100%;
    grid-area: body;
}

.block-name {
    width: 200px;
    height: 40px;
    border-radius: 18px;
    margin-top: 20px;
    text-align: left;
    line-height: 40px;
    font-size: 1.6rem;
    font-weight: bolder;
    background-color: white;
}

.kind-list {
    width: 100%;
    height: auto;
    border-radius: 16px;
    text-align: left;
    background-color: rgba(0,0,0,0.05);
}

.tag-list {
    width: 100%;
    height: auto;
    border-radius: 16px;
    text-align: left;
    background-color: rgba(0,0,0,0.05);
}

.img-list {
    width: 100%;
    height: auto;
    border-radius: 16px;
    text-align: left;
    background-color: rgba(0,0,0,0.05);
}

.img {
    width: auto;
    max-height: 200px;
    margin-left: 10px;
    margin-right: 10px;
    border: 0;
    object-fit: cover;
    display: inline-block;
}
</style>