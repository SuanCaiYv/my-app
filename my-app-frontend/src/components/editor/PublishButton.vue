<template>
    <div class="publish">
        <button class="button" @click="publish">发布</button>
    </div>
</template>

<script setup lang="ts">
import {createApp, ref} from "vue"
import PublishConfirm from "./PublishConfirm.vue";

const name = ref<string>("PublishButton")

const publish = function () {
    let divElement = document.createElement("div");
    const instance = createApp(PublishConfirm, {
        divNode: divElement,
        content: "aaa",
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
    margin: 10px 120px 10px 0;
    padding: 0;
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