<template>
    <div class="body">
        <Today></Today>
        <ArticleList></ArticleList>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import ArticleList from "./ArticleList.vue"
import Today from "./Today.vue"

const name = ref<String>("Body")
</script>

<style scoped>
.body {
    width: 100%;
    height: 100%;
    grid-area: body;
    display: grid;
    grid-template-areas: "today article-list";
    grid-template-columns: 200px 1fr;
    grid-template-rows: 1fr;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
}
</style>