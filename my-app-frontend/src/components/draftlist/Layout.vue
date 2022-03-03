<template>
    <div class="layout">
        <HeaderWithoutSearch></HeaderWithoutSearch>
        <PH1></PH1>
        <Body></Body>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import HeaderWithoutSearch from "../common/HeaderWithoutSearch.vue"
import PH1 from "../placeholder/PH1.vue"
import Body from "./Body.vue"

const name = ref<string>("Layout")
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