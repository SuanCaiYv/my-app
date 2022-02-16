<template>
    <div class="preview" v-html="content"></div>
</template>

<script setup lang="ts">
import {inject, ref, watch} from "vue"
import {marked} from "marked";

const name = ref<string>("Preview")

const contentRaw = inject("content")
const content = ref<string>('')
const title = inject("title")

// @ts-ignore
watch(contentRaw, (n, o) => {
    content.value = marked.parse(n)
})
</script>

<style scoped>
@import "../../components/common/render.css";

.preview {
    width: 100%;
    height: calc(100vh - 101px);
    grid-area: preview;
    border-top: 1px solid lightgray;
    border-left: 2px solid lightgray;
    margin-left: -2px;
    border-top-left-radius: 16px;
    box-sizing: border-box;
    padding-left: 3px;
    padding-right: 10px;
    padding-top: 10px;
    text-align: left;
    word-break: break-all;
    overflow-y: scroll;
    overflow-x: auto;
}

</style>