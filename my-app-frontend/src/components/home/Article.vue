<template>
    <div class="article">
        <PH1></PH1>
        <div class="title">{{title}}</div>
        <div class="content" v-html="contentRendered"></div>
        <PH2></PH2>
    </div>
</template>

<script setup lang="ts">
import {onMounted, ref, watch} from "vue"
import PH1 from "../placeholder/PH1.vue"
import PH2 from "../placeholder/PH2.vue"
import {marked} from "marked";

const name = ref<String>("Article")
const contentRendered = ref<string>('')

const props = defineProps({
    id: String,
    title: String,
    content: String,
})

onMounted(() => {
    // @ts-ignore
    contentRendered.value = marked.parse(props.content)
})
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
    margin-top: 25px;
}

.title {
    width: 600px;
    height: 100%;
    grid-area: title;
    border: 2px solid wheat;
    box-sizing: border-box;
    border-radius: 16px 16px 0 0;
    padding: 0 0 0 8px;
    font-size: 1.4rem;
    font-weight: bolder;
    text-align: left;
    background-color: #f5ecff;
}

.content {
    min-width: 600px;
    height: 100%;
    grid-area: content;
    border: 2px solid wheat;
    box-sizing: border-box;
    border-radius: 0 16px 16px 16px;
    display: inline-block;
    font-size: 1rem;
    text-align: left;
    background-color: #f5ecff;
}
</style>