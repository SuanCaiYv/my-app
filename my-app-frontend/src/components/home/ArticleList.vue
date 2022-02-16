<template>
    <div class="articleList">
        <div style="margin-bottom: -25px"></div>
        <div v-for="article in articleList">
            <Article :id="article.articleId" :title="article.articleName" :content="article.content"></Article>
        </div>
        <button class="button" @click="fetchArticles" :style="{ display: displayStr }">more</button>
    </div>
</template>

<script setup lang="ts">
import {reactive, ref} from "vue"
import Article from "./Article.vue"
import {ArticleLiteRaw} from "../../common/interface"
import {httpClient} from "../../net";
import {Response} from "../../common/interface";
import alertFunc from "../../util/alert";
import {toListResult} from "../../util/base";
import storage from "../../util/storage";
import {Constant} from "../../common/systemconstant";

const name = ref<string>("ArticleList")
// @ts-ignore
const articleList = reactive<Array<ArticleLiteRaw>>([])
const displayStr = ref<string>("none")

let pageNum = 1
let pageSize = 10

class ArticleRawClass implements ArticleLiteRaw {
    articleId: string
    articleName: string;
    content: string;

    constructor(articleId: string, title: string, body: string) {
        this.articleId = articleId
        this.articleName = title;
        this.content = body;
    }
}

const fetchArticles = function () {
    httpClient.get("/article/list/no_auth", {}, false, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        } else {
            const list = toListResult(resp.data)
            pageNum = list.pageNum
            for (let l of list.list) {
                // @ts-ignore
                articleList.push(new ArticleRawClass(l.article_id, l.article_name, l.content))
                storage.set(Constant.ARTICLE_ID + "_" + l.article_id, JSON.stringify(l))
            }
        }
    })
}

fetchArticles()

if (articleList.length > 0) {
    displayStr.value = "inline-block"
}
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