<template>
    <div class="body">
        <div class="block-name">分类管理</div>
        <div class="kind-list">
            <KindOrTagItem></KindOrTagItem>
        </div>
        <div class="block-name">标签管理</div>
        <div class="tag-list"></div>
        <div class="block-name">图片管理</div>
        <div class="img-list"></div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import KindOrTagItem from "../common/KindOrTagItem.vue"

const name = ref<String>("Body")
</script>

<style scoped>
.body {
    width: 100%;
    height: 100%;
    grid-area: body;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
}

.block-name {
    width: 200px;
    height: 40px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    border-radius: 18px;
    margin-top: 20px;
    text-align: center;
    line-height: 40px;
    font-size: 1.4rem;
    background-color: rgba(255,0,154,0.11);
}

.kind-list {
    width: 100%;
    min-height: 200px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    border-radius: 16px;
    text-align: left;
    background-color: rgba(255,0,12,0.05);
}

.tag-list {
    width: 100%;
    min-height: 200px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    border-radius: 16px;
    text-align: left;
    background-color: rgba(168,0,255,0.11);
}

.img-list {
    width: 100%;
    min-height: 400px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    border-radius: 16px;
    background-color: rgba(121,121,121,0.11);
}
</style>