<template>
    <div class="article-list">
        <Article></Article>
        <Article></Article>
        <Article></Article>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Article from "./Article.vue"

const name = ref<String>("ArticleList")
</script>

<style scoped>
.article-list {
    width: 100%;
    height: 100%;
    grid-area: article-list;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
}
</style>