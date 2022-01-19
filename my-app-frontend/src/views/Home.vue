<template>
    <div class="home">
        <Layout></Layout>
    </div>
</template>

<script setup lang="ts">
import {provide, reactive, ref} from "vue"
import Layout from "../components/home/layout/Layout.vue"

const name = ref<String>("Home")

class ArticleRaw {
    title: String
    body: String
    constructor(t: String, b: String) {
        this.title = t
        this.body = b
    }
}

const articleList = reactive<Array<ArticleRaw>>([])

// 测试数据
articleList.push(new ArticleRaw("1", "aaa"))
articleList.push(new ArticleRaw("2", "bbb"))
articleList.push(new ArticleRaw("3", "ccc"))

provide("searchKey", function (key: String) {
    console.log(key)
})
provide("articleList", articleList)
</script>

<style scoped>
.home {
    width: 100%;
    height: 100%;
}
</style>