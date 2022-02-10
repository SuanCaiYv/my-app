<template>
    <div class="body">
        <SortAndPage></SortAndPage>
        <UserInfo></UserInfo>
        <ArticleList></ArticleList>
        <TagList></TagList>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import SortAndPage from "./SortAndPage.vue"
import UserInfo from "./UserInfo.vue"
import ArticleList from "./ArticleList.vue";
import TagList from "./TagList.vue";

const name = ref<string>("Body")
</script>

<style scoped>
.body {
    width: 100%;
    height: 100%;
    grid-area: body;
    display: grid;
    grid-template-areas:
        "sortAndPage sortAndPage sortAndPage"
        "userInfo articleList tagList";
    grid-auto-columns: 200px 1fr 200px;
    grid-template-rows: 80px 1fr;
}
</style>