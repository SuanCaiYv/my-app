<template>
    <div class="layout">
        <Header></Header>
        <PH1></PH1>
        <Body></Body>
    </div>
</template>

<script setup lang="ts">
import {provide, ref} from "vue"
import Header from "./Header.vue"
import PH1 from "../placeholder/PH1.vue"
import Body from "./Body.vue"

const name = ref<string>("Layout")
provide("searchKey", ref<string>(''))
</script>

<style scoped>
.layout {
    width: 100%;
    height: 100%;
    display: grid;
    grid-template-areas:
        "ph1"
        "body";
    grid-template-rows: 60px 1fr;
}
</style>