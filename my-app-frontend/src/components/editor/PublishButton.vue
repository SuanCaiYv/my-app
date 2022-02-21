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

const publish = function () {
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
    border: none;
    box-sizing: border-box;
    border-radius: 18px;
    font-size: 1.0rem;
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