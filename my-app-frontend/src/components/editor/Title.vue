<template>
    <div class="title">
        <input type="text" class="input" v-model="title" placeholder="文章标题">
    </div>
</template>

<script setup lang="ts">
import {inject, ref} from "vue"
import storage from "../../util/storage";
import {Constant} from "../../common/systemconstant";
import {useRoute} from "vue-router";

const name = ref<string>("Title")
const title = inject("title")
const id = inject("id")
const route = useRoute()
if (route.params.type === "update") {
    // @ts-ignore
    const article = JSON.parse(storage.get(Constant.ARTICLE_ID_PREFIX + id.value))
    // @ts-ignore
    title.value = article.article_name
}
</script>

<style scoped>
.title {
    width: 100%;
    height: 100%;
    grid-area: title;
}

.input {
    height: 100%;
    width: 100%;
    border: none;
    padding: 0 0 0 4px;
    font-size: 1.6rem;
    font-weight: bolder;
    outline: none;
}
</style>