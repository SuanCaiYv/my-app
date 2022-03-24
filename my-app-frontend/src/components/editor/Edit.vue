<template>
    <div class="edit">
        <textarea class="edit-area" v-model="content"></textarea>
    </div>
</template>

<script setup lang="ts">
import {inject, Ref, ref} from "vue"
import {BASE_URL, httpClient} from "../../net";
import {Response} from "../../common/interface";
import storage from "../../util/storage";
import {Constant} from "../../common/systemconstant";
import {useRoute} from "vue-router";

const name = ref<string>("Edit")
const content = inject("content") as Ref<string>
const id = inject("id") as Ref<string>
const route = useRoute()
if (route.params.type === "update") {
    const article = JSON.parse(storage.get(Constant.ARTICLE_ID_PREFIX + id.value))
    content.value = article.content
}

document.addEventListener('paste', function (event: ClipboardEvent) {
    const items = event.clipboardData && event.clipboardData.items
    let file: File | null = null
    if (items && items.length) {
        // 检索剪切板items
        for (let i = 0; i < items.length; i++) {
            if (items[i].type.indexOf('image') !== -1) {
                file = items[i].getAsFile()
                break;
            }
        }
    }
    if (file !== null) {
        console.log(file.size)
        let formData = new FormData()
        // @ts-ignore
        formData.append("file", file)
        httpClient.upload("/static/file", {
            archive: "doc_img"
        }, formData, function (resp: Response) {
            if (!resp.ok) {
                console.log(resp.errMsg)
            } else {
                // @ts-ignore
                let filename = resp.data.filename
                let str = "![" + filename + "]"
                let url = BASE_URL + "/static/a/" + filename
                str += "(" + url + ")"
                content.value += str
            }
        })
    }
});
</script>

<style scoped>
.edit {
    width: 100%;
    height: 100%;
    grid-area: edit;
    border-top: 1px solid lightgray;
    border-right: 2px solid lightgray;
    margin-right: -2px;
    border-top-right-radius: 16px;
    box-sizing: border-box;
    padding-left: 10px;
    padding-right: 3px;
    padding-top: 10px;
    overflow-y: auto;
}

.edit-area {
    width: 100%;
    height: calc(100vh - 111px);
    border: none;
    box-sizing: border-box;
    font-size: 1rem;
    font-weight: normal;
    outline: none;
    resize: none;
}
</style>