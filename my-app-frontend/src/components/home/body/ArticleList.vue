<template>
    <div class="articleList">
        <div style="margin-bottom: -25px"></div>
        <div v-for="article in articleList">
            <Article :title="article.title" :body="article.body"></Article>
        </div>
    </div>
</template>

<script setup lang="ts">
import {inject, ref} from "vue"
import Article from "./Article.vue"
import {ArticleRaw} from "../../../common/interface"

const name = ref<string>("ArticleList")

// @ts-ignore
const articleList: Array<ArticleRaw> = inject("articleList")
</script>

<style scoped>
.articleList {
    width: 100%;
    height: 100%;
    grid-area: articleList;
}
</style>