<template>
    <div class="body">
        <Title></Title>
        <Edit></Edit>
        <Preview></Preview>
    </div>
</template>

<script setup lang="ts">
import {inject, ref} from "vue"
import Title from "./Title.vue"
import Edit from "./Edit.vue"
import Preview from "./Preview.vue"
import {httpClient} from "../../net";
import {Response} from "../../common/interface";
import alertFunc from "../../util/alert";
import storage from "../../util/storage";
import {Constant} from "../../common/systemconstant";

const name = ref<string>("Body")

const id = inject("id")
const title = inject("title")
const content = inject("content")

const saveDraft = function () {
    storage.set(Constant.ARTICLE_ID, id.value)
    storage.set(Constant.ARTICLE_TITLE, title.value)
    storage.set(Constant.ARTICLE_CONTENT, content.value)
    httpClient.post("/article/draft", {}, {
        article_id: id.value,
        article_name: title.value,
        article_content: content.value
    }, true, function (resp: Response) {
        if (!resp.ok) {
            clearInterval(cancel)
            console.log("http failed")
            alertFunc(resp.errMsg, function () {
            })
        } else {
            // @ts-ignore
            id.value = resp.data.article_id
        }
    })
}

saveDraft()

const cancel = setInterval(saveDraft, 5000)
storage.set(Constant.DRAFT_INTERVAL_CANCEL, cancel + "")
</script>

<style scoped>
.body {
    width: 100%;
    height: 100%;
    grid-area: body;
    display: grid;
    grid-template-areas:
        "title title"
        "edit preview";
    grid-template-rows: 40px 1fr;
    grid-template-columns: 1fr 1fr;
}
</style>