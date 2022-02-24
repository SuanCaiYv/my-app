<template>
    <div class="article" @click="router.push('/view/' + id)">
        <PH1></PH1>
        <div class="title">{{title}}</div>
        <div class="content">{{summary}}</div>
        <PH2></PH2>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import PH1 from "../placeholder/PH1.vue"
import PH2 from "../placeholder/PH2.vue"
import {useRouter} from "vue-router";

const name = ref<string>("Article")
const router = useRouter()

const props = defineProps({
    id: String,
    title: String,
    summary: String,
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
    width: 40%;
    height: 100%;
    grid-area: title;
    border: none;
    box-sizing: border-box;
    border-radius: 16px 16px 0 0;
    padding: 0 0 0 8px;
    font-size: 1.4rem;
    font-weight: bolder;
    line-height: 40px;
    text-align: left;
    background-color: rgba(0,0,0,0.05);
}

.content {
    min-width: 600px;
    height: 100%;
    grid-area: content;
    border: none;
    padding: 8px;
    box-sizing: border-box;
    border-radius: 0 16px 16px 16px;
    display: inline-block;
    font-size: 1.2rem;
    text-align: left;
    word-break: break-word;
    overflow: hidden;
    background-color: rgba(0,0,0,0.03);
}
</style>