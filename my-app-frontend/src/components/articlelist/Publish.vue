<template>
    <div class="publish">
        <button class="button" @click="publish">发布文章</button>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {useRouter} from "vue-router";
import {httpClient} from "../../net";
import {Response} from "../../common/interface";
import alertFunc from "../../util/alert";
import storage from "../../util/storage";
import {Constant} from "../../common/systemconstant";

const name = ref<string>("Publish")
const router = useRouter()

const publish = function () {
    httpClient.post("/article/draft", {}, {
        article_id: "",
        article_name: new Date(),
        article_content: ""
    }, true, function (resp: Response) {
        if (!resp.ok) {
            console.log("http failed")
            alertFunc(resp.errMsg, function () {
            })
        } else {
            // @ts-ignore
            storage.set(Constant.DRAFT_ARTICLE_ID, resp.data.article_id)
            storage.set(Constant.ARTICLE_TITLE, "")
            storage.set(Constant.ARTICLE_CONTENT, "")
            storage.set(Constant.EDITOR_TYPE, "draft")
            router.push("/editor/draft")
        }
    })
}
</script>

<style scoped>
.publish {
    width: 120px;
    height: 100%;
    grid-area: publish;
    margin-right: 20px;
}

.button {
    width: 100%;
    height: 40px;
    margin-top: 10px;
    margin-bottom: 10px;
    padding: 0;
    border: none;
    border-radius: 20px;
    font-size: 1rem;
    font-weight: bolder;
    background-color: rgba(0,0,0,0.1);
}

.button:hover {
    background-color: rgba(0,0,0,0.15);
}

.button:active {
    background-color: rgba(0,0,0,0.2);
}
</style>