<template>
    <div class="article-list">
        <div style="margin-top: 80px"></div>
        <Article></Article>
        <Article></Article>
        <Article></Article>
    </div>
</template>

<script setup lang="ts">
import {reactive, ref} from "vue"
import Article from "./Article.vue"
import {httpClient} from "../../net";
import {Response} from "../../common/interface";
import alertFunc from "../../util/alert";

const name = ref<String>("ArticleList")

let pageNum = 1
let pageSize = 10

const articleList = reactive<Array<{}>>([])

const fetchArticles = function (pageNum: number, pageSize: number) {
    httpClient.get("/article/list", {
        page_num: pageNum,
        page_size: pageSize
    }, true, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        } else {
            pageNum = resp.data.next_page_num
        }
    })
}
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