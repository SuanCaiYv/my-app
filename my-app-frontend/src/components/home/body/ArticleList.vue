<template>
    <div class="articleList">
        <div style="margin-bottom: -25px"></div>
        <div v-for="article in articleList">
            <Article :title="article.title" :body="article.body"></Article>
        </div>
        <button class="button" @click="fetchArticles">more</button>
    </div>
</template>

<script setup lang="ts">
import {reactive, ref} from "vue"
import Article from "./Article.vue"
import {ArticleRaw} from "../../../common/interface"
import {httpClient} from "../../../net";
import {Response} from "../../../common/interface";
import alertFunc from "../../../util/alert";

const name = ref<string>("ArticleList")
// @ts-ignore
const articleList = reactive<Array<ArticleRaw>>([])

let pageNum = 1
let pageSize = 10

class ArticleRawClass implements ArticleRaw {
    body: string;
    title: string;

    constructor(title: string, body: string) {
        this.title = title;
        this.body = body;
    }
}

const fetchArticles = function () {
    httpClient.get("/article/list/no_auth", {}, false, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        } else {
            pageNum = resp.data.next_page_num
            for (let item of resp.data.list) {
                console.log(item)
                articleList.push(new ArticleRawClass(item.article_name, item.content))
            }
        }
    })
}

fetchArticles()
</script>

<style scoped>
.articleList {
    width: 100%;
    height: 100%;
    grid-area: articleList;
}

.button {
    width: 60px;
    height: 30px;
    border: none;
    padding: 0;
    font-size: 1.4rem;
    font-weight: bolder;
    line-height: 30px;
    color: #b4d4ff;
    background-color: white;
}

.button:hover {
    color: #9dbbff;
}
</style>