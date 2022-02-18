<template>
    <div class="body">
        <Today></Today>
        <PH1></PH1>
        <ArticleList></ArticleList>
    </div>
</template>

<script setup lang="ts">
import {provide, ref} from "vue"
import ArticleList from "./ArticleList.vue"
import Today from "./Options.vue"
import PH1 from "../placeholder/PH1.vue"

const name = ref<string>("Body")
provide("sort", ref<string>(''))
provide("page", ref<string>(''))
provide("desc", ref<string>(''))
</script>

<style scoped>
.body {
    width: 100%;
    height: 100%;
    grid-area: body;
    display: grid;
    grid-template-areas: "ph1 article-list";
    grid-template-columns: 200px 1fr;
    grid-template-rows: 1fr;
}
</style>