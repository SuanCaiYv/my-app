<template>
    <div class="body">
        <PH1></PH1>
        <PH2></PH2>
        <div class="main">
            <Title :value="title"></Title>
            <div class="cover-img" :style="{display: displayCovImg}">
                <img class="img" :src="covImg">
            </div>
            <div class="content" v-html="content"></div>
            <div class="kind-and-tag"></div>
        </div>
        <PH3></PH3>
        <PH4></PH4>
    </div>
</template>

<script setup lang="ts">
import {reactive, ref} from "vue"
import PH1 from "../placeholder/PH1.vue"
import PH2 from "../placeholder/PH2.vue"
import PH3 from "../placeholder/PH3.vue"
import PH4 from "../placeholder/PH4.vue"
import {useRoute} from "vue-router";
import storage from "../../util/storage";
import {Constant} from "../../common/systemconstant";
import {httpClient} from "../../net";
import {ArticleRaw, IdName, Response} from "../../common/interface";
import {parseBoolean, toArticleRaw} from "../../util/base";
import alertFunc from "../../util/alert";
import Title from "./Title.vue";
import {marked} from "marked";

const name = ref<string>("Body")
const route = useRoute()

const covImg = ref<string>("")
const displayCovImg = ref<string>("none")
const title = ref<string>("")
const content = ref<string>("")
const kind = ref<IdName>()
const tagList = reactive<Array<IdName>>([])

const articleId = route.params.article_id

let articleCache = storage.get(Constant.ARTICLE_ID_PREFIX + articleId);
let article: ArticleRaw | null = null
if (articleCache === "") {
    if (!parseBoolean(storage.get(Constant.AUTHENTICATED))) {
        httpClient.get("/article/" + articleId, {}, false, function (resp: Response) {
            if (!resp.ok) {
                alertFunc(resp.errMsg, function () {})
            } else {
                article = toArticleRaw(JSON.stringify(resp.data))
                if (article.coverImg === "") {
                    displayCovImg.value = "none"
                } else {
                    displayCovImg.value = "block"
                }
                covImg.value = article.coverImg
                title.value = article.articleName
                content.value = marked.parse(article.content)
            }
        })
    } else {
        httpClient.get("/article/detail/" + articleId, {}, true, function (resp: Response) {
            if (!resp.ok) {
                alertFunc(resp.errMsg, function () {})
            } else {
                article = toArticleRaw(JSON.stringify(resp.data))
                if (article.coverImg === "") {
                    displayCovImg.value = "none"
                } else {
                    displayCovImg.value = "block"
                }
                covImg.value = article.coverImg
                title.value = article.articleName
                content.value = marked.parse(article.content)
            }
        })
    }
} else {
    article = toArticleRaw(articleCache)
    if (article.coverImg === "") {
        displayCovImg.value = "none"
    } else {
        displayCovImg.value = "block"
    }
    covImg.value = article.coverImg
    title.value = article.articleName
    content.value = marked.parse(article.content)
}
</script>

<style scoped>
@import "../../components/common/render.css";

.body {
    width: 100%;
    height: 100%;
    grid-area: body;
    display: grid;
    grid-template-areas:
        "ph1 ph2 main ph3 ph4";
    grid-template-columns: 1fr 2fr 8fr 2fr 1fr;
}

.main {
    width: 100%;
    height: 100%;
    grid-area: main;
    display: inline-block;
}

.cover-img {
    width: 100%;
    height: 300px;
    text-align: center;
}

.img {
    max-width: 100%;
    max-height: 100%;
    object-fit: contain;
}

.content {
    width: 100%;
    height: 100%;
    text-align: left;
}

.kind-and-tag {
    width: 100%;
    height: 60px;
}
</style>