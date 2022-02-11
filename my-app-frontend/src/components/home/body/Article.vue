<template>
    <div class="article">
        <PH1></PH1>
        <div class="title"></div>
        <div class="content"></div>
        <PH2></PH2>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import PH1 from "../../placeholder/PH1.vue"
import PH2 from "../../placeholder/PH2.vue"

const name = ref<String>("Article")
</script>

<style scoped>
.article {
    width: 100%;
    height: 280px;
    display: grid;
    grid-template-areas:
        "ph1 title title"
        "ph1 content ph2"
        "ph1 content ph2"
        "ph1 content ph2";
    grid-template-columns: 25px 1fr 25px;
    grid-template-rows: 40px 80px 80px 80px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    margin-top: 25px;
}

.title {
    width: 320px;
    height: 100%;
    grid-area: title;
    border: 2px solid wheat;
    box-sizing: border-box;
    border-radius: 16px 16px 0 0;
    background-color: antiquewhite;
}

.content {
    width: 100%;
    height: 100%;
    grid-area: content;
    border: 2px solid wheat;
    box-sizing: border-box;
    border-radius: 0 16px 16px 16px;
    display: inline-block;
    background-color: antiquewhite;
}
</style>