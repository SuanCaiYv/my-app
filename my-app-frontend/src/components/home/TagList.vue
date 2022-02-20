<template>
    <div class="tagList">
        <!--占位符，勿动！-->
        <div style="margin-top: 25px;box-sizing: border-box"></div>
        <div v-for="t in tagList" style="display: inline-block">
            <Tag :id="t.id" :name="t.name" :click-func="clicked"></Tag>
        </div>
    </div>
</template>

<script setup lang="ts">
import {inject, reactive, ref} from "vue"
import Tag from "./Tag.vue";
import {IdName, Response} from "../../common/interface";
import {httpClient} from "../../net";
import {toListResult} from "../../util/base";

const name = ref<string>("TagList")
const tagList = reactive<Array<IdName>>([])
const chosenTagList = inject("chosenTagList")

class IdNameClass implements IdName {
    id: string;
    name: string;

    constructor(id: string, name: string) {
        this.id = id;
        this.name = name;
    }
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

fetchTagList()

const clicked = function (id: string, isActive: boolean) {
    if (isActive) {
        // @ts-ignore
        chosenTagList.push(id)
    } else {
        // @ts-ignore
        chosenTagList.splice(chosenTagList.findIndex(item => item === id), 1)
    }
}
</script>

<style scoped>
.tagList {
    width: 200px;
    height: 100%;
    border: none;
    box-sizing: border-box;
    border-radius: 20px;
    position: fixed;
    top: 140px;
    right: 0;
    background-color: rgba(0,0,0,0.05);
}
</style>