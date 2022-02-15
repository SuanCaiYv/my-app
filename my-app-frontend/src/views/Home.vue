<template>
    <div class="home">
        <Layout></Layout>
    </div>
</template>

<script setup lang="ts">
import {provide, reactive, ref} from "vue"
import Layout from "../components/home/Layout.vue"
import {ArticleLiteRaw} from "../common/interface";
import {uploadLocalFile} from "../util/file";

const name = ref<String>("Home")

class ArticleRawClass implements ArticleLiteRaw {
    articleName: String
    content: String
    constructor(t: String, b: String) {
        this.articleName = t
        this.content = b
    }
}

const articleList = reactive<Array<ArticleLiteRaw>>([])

articleList.push(new ArticleRawClass("aaa", "bbb"))
articleList.push(new ArticleRawClass("ccc", "ddd"))
articleList.push(new ArticleRawClass("eee", "fff"))

provide("articleList", articleList)
</script>

<style scoped>
.home {
    width: 100%;
    height: 100%;
}
</style>