<template>
    <div class="home">
        <Layout></Layout>
    </div>
</template>

<script setup lang="ts">
import {provide, reactive, ref} from "vue"
import Layout from "../components/home/layout/Layout.vue"
import {ArticleRaw} from "../common/interface";
import {uploadLocalFile} from "../util/file";

const name = ref<String>("Home")

class ArticleRawClass implements ArticleRaw {
    title: String
    body: String
    constructor(t: String, b: String) {
        this.title = t
        this.body = b
    }
}

const articleList = reactive<Array<ArticleRaw>>([])

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