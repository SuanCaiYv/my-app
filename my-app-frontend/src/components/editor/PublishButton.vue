<template>
    <div class="publish">
        <button class="button" @click="publish">发布</button>
    </div>
</template>

<script setup lang="ts">
import {createApp, inject, ref} from "vue"
import Publish from "./Publish.vue";
import storage from "../../util/storage";
import {Constant} from "../../common/systemconstant";

const name = ref<string>("PublishButton")

const id = inject("id")
const title = inject("title")
const content = inject("content")

const publish = function () {
    storage.set(Constant.ARTICLE_ID, id.value)
    storage.set(Constant.ARTICLE_TITLE, title.value)
    storage.set(Constant.ARTICLE_CONTENT, content.value)
    let divElement = document.createElement("div");
    const instance = createApp(Publish, {
        divNode: divElement,
    })
    instance.mount(divElement)
    // @ts-ignore
    document.getElementById("app").appendChild(divElement)
}
</script>

<style scoped>
.publish {
    width: 100%;
    height: 100%;
    grid-area: publish-button;
}

.button {
    width: 80px;
    height: 40px;
    margin-top: 10px;
    margin-left: 0;
    margin-right: calc(100% - 80px);
    padding: 0 20px;
    border: 2px solid lightgray;
    box-sizing: border-box;
    border-radius: 18px;
    font-size: 1.0rem;
    font-weight: bolder;
    background-color: white;
}

.button:hover {
    background-color: lightgray;
}

.button:active {
    background-color: gainsboro;
}
</style>