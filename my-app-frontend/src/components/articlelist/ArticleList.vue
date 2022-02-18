<template>
    <div class="article-list">
        <div style="margin-top: 80px"></div>
        <div v-for="article in articleList">
            <Article :id="article.articleId" :title="article.articleName" :summary="article.summary"></Article>
        </div>
        <button class="button" @click="fetchArticles" :style="{ display: displayStr }">more</button>
    </div>
</template>

<script setup lang="ts">
import {inject, reactive, Ref, ref, watch} from "vue"
import Article from "./Article.vue"
import {httpClient} from "../../net";
import {ArticleLiteRaw, Response} from "../../common/interface";
import alertFunc from "../../util/alert";
import {toListResult} from "../../util/base";
import storage from "../../util/storage";
import {Constant} from "../../common/systemconstant";

const name = ref<String>("ArticleList")
// @ts-ignore
const articleList = reactive<Array<ArticleLiteRaw>>([])
const displayStr = ref<string>("none")
const pageSize = inject("page") as Ref<string>
const sort = inject("sort") as Ref<string>
const desc = inject("desc") as Ref<string>
const searchKey = inject("searchKey") as Ref<string>
let pageNum = 1
let endPage = false

class ArticleRawClass implements ArticleLiteRaw {
    articleId: string
    articleName: string;
    summary: string;

    constructor(articleId: string, title: string, body: string) {
        this.articleId = articleId
        this.articleName = title;
        this.summary = body;
    }
}

watch(searchKey, () => {
    endPage = false
    articleList.splice(0, articleList.length)
    fetchArticles()
})
watch(sort, () => {
    endPage = false
    articleList.splice(0, articleList.length)
    fetchArticles()
})
watch(pageSize, () => {
    endPage = false
    articleList.splice(0, articleList.length)
    fetchArticles()
})
watch(desc, () => {
    endPage = false
    articleList.splice(0, articleList.length)
    fetchArticles()
})

const fetchArticles = function () {
    if (endPage) {
        return
    }
    let pageSizeValue
    let sortValue
    let descValue
    // @ts-ignore
    if (pageSize.value === "") {
        // @ts-ignore
        pageSizeValue = "10"
    } else {
        pageSizeValue = pageSize.value
    }
    // @ts-ignore
    if (sort.value === "") {
        // @ts-ignore
        sortValue = "release_time"
    } else {
        sortValue = sort.value
    }
    // @ts-ignore
    if (desc.value === "") {
        // @ts-ignore
        descValue = "false"
    } else {
        descValue = desc.value
    }
    httpClient.get("/article/list/no_auth", {
        page_num: pageNum,
        page_size: pageSizeValue,
        sort: sortValue,
        desc: descValue,
        search_key: searchKey.value
    }, false, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        } else {
            const list = toListResult(resp.data)
            endPage = list.endPage
            pageNum = list.pageNum
            for (let l of list.list) {
                // @ts-ignore
                articleList.push(new ArticleRawClass(l.article_id, l.article_name, l.summary))
                storage.set(Constant.ARTICLE_ID + "_" + l.article_id, JSON.stringify(l))
            }
            if (articleList.length > 0) {
                displayStr.value = "inline-block"
            }
        }
    })
}

articleList.splice(0, articleList.length)
endPage = false
fetchArticles()
</script>

<style scoped>
.article-list {
    width: 100%;
    height: 100%;
    grid-area: article-list;
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