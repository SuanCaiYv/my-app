<template>
    <div class="articleList">
        <!--占位符，勿动！-->
        <div style="margin-bottom: 50px;box-sizing: border-box"></div>
        <div ref="list"></div>
    </div>
</template>

<script setup lang="ts">
import {createApp, h, onMounted, ref} from "vue"
import Article from "./Article.vue"

const name = ref<String>("ArticleList")

let article1 = createApp(Article)
article1._component.props.title = "aaa"
let node = h(article1)

let list = ref<HTMLElement | null>(null)

onMounted(() => {
    list.value.appendChild(node)
})
</script>

<style scoped>
.articleList {
    width: 100%;
    height: 100%;
    grid-area: articleList;
}
</style>