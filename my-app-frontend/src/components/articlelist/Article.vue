<template>
    <div class="article">
        <PH1></PH1>
        <div class="title" @click="router.push('/view/' + id)">{{title}}</div>
        <PH3></PH3>
        <div class="content" @click="router.push('/view/' + id)">{{summary}}</div>
        <PH2></PH2>
        <div class="update">
            <button class="button" @click="upt">更新</button>
        </div>
        <div class="delete">
            <button class="button" @click="del">删除</button>
        </div>
        <div class="visible">
            <button class="button" @click="setVisibility">{{visibly}}</button>
        </div>
    </div>
</template>

<script setup lang="ts">
import {onMounted, ref, watch} from "vue"
import PH1 from "../placeholder/PH1.vue"
import PH2 from "../placeholder/PH2.vue"
import PH3 from "../placeholder/PH3.vue";
import storage from "../../util/storage";
import {Constant} from "../../common/systemconstant";
import {useRouter} from "vue-router";
import {confirmFunc} from "../../util/confirm";
import {httpClient} from "../../net";
import {Response} from "../../common/interface";
import alert from "../../util/alert";

const name = ref<string>("Article")
const router = useRouter()
const props = defineProps({
    id: String,
    title: String,
    summary: String,
    visibility: Number,
})
const visibly = ref<string>()
const flag = ref<boolean>()

watch(flag,  () => {
    if (flag.value) {
        visibly.value = "公开"
    } else {
        visibly.value = "私密"
    }
    httpClient.put("/article", {}, {
        article_id: props.id + "",
        visibility: flag.value ? 2 : 1
    }, true, function (resp: Response) {})
})

onMounted(() => {
    if (Number(props.visibility) === 1) {
        flag.value = false
    } else if (Number(props.visibility) === 2) {
        flag.value = true
    }
})

const setVisibility = function () {
    flag.value = !flag.value
}

const upt = function () {
    storage.set(Constant.UPDATE_ARTICLE_ID, props.id + "")
    storage.set(Constant.EDITOR_TYPE, "update")
    router.push("/editor/update")
}

const del = function () {
    confirmFunc("确认删除吗？", function () {}, function () {
        httpClient.delete("/article/" + props.id, {}, true, function (resp: Response) {
            if (!resp.ok) {
                console.log(resp.errMsg)
            } else {
                alert("已删除！", function () {})
            }
        })
    })
}
</script>

<style scoped>
.article {
    width: 100%;
    height: 280px;
    display: grid;
    grid-template-areas:
        "ph1 title title ph3"
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
    padding: 8px;
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
    border: none;
    padding: 0;
    border-radius: 18px;
    font-size: 1.2rem;
    font-weight: bolder;
    background-color: rgba(0,0,0,0.1);
}

.button:hover {
    background-color: rgba(0,0,0,0.15);
}

.button:active {
    background-color: rgba(0,0,0,0.2);
}
</style>