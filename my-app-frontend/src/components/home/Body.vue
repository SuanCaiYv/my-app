<template>
    <div class="body">
        <PH1></PH1>
        <SortAndPage></SortAndPage>
        <PH2></PH2>
        <UserInfo></UserInfo>
        <ArticleList></ArticleList>
        <PH3></PH3>
        <TagList></TagList>
    </div>
</template>

<script setup lang="ts">
import {provide, reactive, ref} from "vue"
import SortAndPage from "./SortAndPage.vue"
import UserInfo from "./UserInfo.vue"
import ArticleList from "./ArticleList.vue";
import TagList from "./TagList.vue";
import PH1 from "../placeholder/PH1.vue"
import PH2 from "../placeholder/PH2.vue"
import PH3 from "../placeholder/PH3.vue"

const name = ref<string>("Body")
provide("sort", ref<string>(''))
provide("page", ref<string>(''))
provide("desc", ref<string>(''))
provide("chosenTagList", reactive<Array<string>>([]))
</script>

<style scoped>
.body {
    width: 100%;
    height: 100%;
    grid-area: body;
    display: grid;
    grid-template-areas:
        "ph1 ph1 ph1"
        "ph2 articleList ph3";
    grid-auto-columns: 200px 1fr 200px;
    grid-template-rows: 80px 1fr;
}
</style>