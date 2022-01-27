<template>
    <div class="body">
        <div class="block-name">分类管理</div>
        <div class="kind-list">
        </div>
        <div class="block-name">标签管理</div>
        <div class="tag-list"></div>
        <div class="block-name">图片管理</div>
        <div class="img-list"></div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import KindOrTagItem from "./KindOrTagItem.vue"

const name = ref<String>("Body")
</script>

<style scoped>
.body {
    width: 100%;
    height: 100%;
    grid-area: body;
    border: 1px solid silver;
    box-sizing: border-box;
}

.block-name {
    width: 200px;
    height: 40px;
    border: 1px solid silver;
    box-sizing: border-box;
    margin-top: 20px;
    text-align: center;
    line-height: 40px;
    font-size: 1.4rem;
}

.kind-list {
    width: 100%;
    min-height: 200px;
    border: 1px solid silver;
    box-sizing: border-box;
    text-align: left;
}

.tag-list {
    width: 100%;
    min-height: 200px;
    border: 1px solid silver;
    box-sizing: border-box;
    text-align: left;
}

.img-list {
    width: 100%;
    min-height: 400px;
    border: 1px solid silver;
    box-sizing: border-box;
}
</style>