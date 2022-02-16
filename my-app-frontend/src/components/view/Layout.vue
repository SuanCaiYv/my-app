<template>
    <div class="layout">
        <HeaderWithoutSearch></HeaderWithoutSearch>
        <Body></Body>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import HeaderWithoutSearch from "../common/HeaderWithoutSearch.vue";
import Body from "./Body.vue"

const name = ref<String>("Layout")
</script>

<style scoped>
.layout {
    width: 100%;
    height: 100%;
    display: grid;
    grid-template-areas:
        "header"
        "body";
    grid-template-rows: 60px 1fr;
}
</style>