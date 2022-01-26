<template>
    <div class="layout">
        <Header></Header>
        <Body></Body>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Header from "./Header.vue"
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
    grid-template-rows: 50px 1fr;
}
</style>