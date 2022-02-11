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

const name = ref<string>("Body")

const title = ref<string>('')
const content = ref<string>('')

provide("content", content)
provide("title", title)
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