<template>
    <div class="articleList">
        <!--占位符，勿动！-->
        <div style="margin-bottom: 50px;box-sizing: border-box"></div>
        <div v-for="l in list">
            <Article :title="l.title" :body="l.body"></Article>
        </div>
    </div>
</template>

<script setup lang="ts">
import {inject, ref} from "vue"
import Article from "./Article.vue"

const name = ref<String>("ArticleList")
const list = inject("articleList")
</script>

<style scoped>
.articleList {
    width: auto;
    height: 100%;
    grid-area: articleList;
    margin-left: 50px;
    margin-right: 50px;
}
</style>