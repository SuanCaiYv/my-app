<template>
    <div class="body">
        <PH1></PH1>
        <ArticleList></ArticleList>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import PH1 from "../placeholder/PH1.vue"
import Options from "./Options.vue"
import ArticleList from "./ArticleList.vue"

const name = ref<String>("Body")
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
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
}
</style>