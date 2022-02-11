<template>
    <div class="edit">
        <textarea class="edit-area" v-model="content"></textarea>
    </div>
</template>

<script setup lang="ts">
import {inject, ref} from "vue"
import {baseUrl, httpClient} from "../../net";
import {Response} from "../../net";

const name = ref<string>("Edit")

const content = inject("content")

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
            // @ts-ignore
            let filename = resp.data.result
            let str = "![" + filename + "]"
            let url = baseUrl + "/static/a/" + filename
            str += "(" + url + ")"
            navigator.clipboard.writeText(str)
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
}

.edit-area {
    width: 100%;
    height: 100%;
    border: none;
    box-sizing: border-box;
    font-size: 1rem;
    font-weight: normal;
    outline: none;
    resize: none;
}
</style>