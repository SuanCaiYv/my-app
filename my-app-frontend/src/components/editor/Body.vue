<template>
    <div class="body">
        <Title></Title>
        <Edit></Edit>
        <Preview></Preview>
    </div>
</template>

<script setup lang="ts">
import {provide, ref} from "vue"
import Title from "./Title.vue"
import Edit from "./Edit.vue"
import Preview from "./Preview.vue"
import {httpClient} from "../../net";
import {Response} from "../../common/interface";
import alertFunc from "../../util/alert";
import {useStore} from "vuex";

const name = ref<string>("Body")
const store = useStore()

const id = ref<string>('')
const title = ref<string>('')
const content = ref<string>('')

const saveDraft = function () {
    httpClient.post("/article/draft", {}, {
        article_id: id.value,
        article_name: title.value,
        article_content: content.value
    }, true, function (resp: Response) {
        if (store.getters.draftArticleId === "" || store.getters.draftArticleId === null) {
            clearInterval(cancel)
        }
        if (!resp.ok) {
            clearInterval(cancel)
            alertFunc(resp.errMsg, function () {
            })
        } else {
            // @ts-ignore
            id.value = resp.data.article_id
            store.commit("updatedDraftArticleId", id.value)
        }
    })
}

const cancel = setInterval(saveDraft, 5000)

provide("id", id)
provide("title", title)
provide("content", content)
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