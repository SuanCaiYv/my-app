<template>
    <div class="draft-list">
        <div style="margin-bottom: -25px"></div>
        <div v-for="draft in draftList">
            <Draft :id="draft.articleId" :title="draft.articleName" :content="draft.content"></Draft>
        </div>
    </div>
</template>

<script setup lang="ts">
import {reactive, ref} from "vue"
import {ArticleRaw, Response} from "../../common/interface";
import Draft from "./Draft.vue"
import {httpClient} from "../../net";
import alertFunc from "../../util/alert";
import {toListResult} from "../../util/base";
import {toArticleRawWithObject} from "../../util/base";
import storage from "../../util/storage";
import {Constant} from "../../common/systemconstant";

const name = ref<string>("DraftList")
const draftList = reactive<Array<ArticleRaw>>([])

const fetchDraftList = function () {
    httpClient.get("/article/draft/list", {
        page_num: -1,
        page_size: 10,
    }, true, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        } else {
            const list = toListResult(resp.data)
            for (let l of list.list) {
                // @ts-ignore
                draftList.push(toArticleRawWithObject(l))
                storage.set(Constant.ARTICLE_ID_PREFIX + l.article_id, JSON.stringify(l))
                console.log(l)
            }
        }
    })
}

fetchDraftList()
</script>

<style scoped>
.draft-list {
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
    color: darkgray;
    background-color: white;
}

.button:hover {
    color: gray;
}
</style>