<template>
    <div class="sortAndPage">
        <PH1 class="ph1"></PH1>
        <div class="sort block">
            <select class="select" v-model="sort">
                <option disabled value="">排序方式</option>
                <option value="release_time">时间</option>
                <option value="kind.name">分类</option>
                <option value="name">名称</option>
            </select>
        </div>
        <div class="page block">
            <select class="select" v-model="page">
                <option disabled value="">分页数量</option>
                <option value="5">5</option>
                <option value="10">10</option>
                <option value="15">15</option>
                <option value="20">20</option>
                <option value="50">50</option>
            </select>
        </div>
        <div class="desc block">
            <select class="select" v-model="desc">
                <option disabled value="">显示顺序</option>
                <option value="true">正序</option>
                <option value="false">倒序</option>
            </select>
        </div>
        <PH2 class="ph2"></PH2>
    </div>
</template>

<script setup lang="ts">
import {inject, ref, watch} from "vue"
import PH1 from "../placeholder/PH1.vue"
import PH2 from "../placeholder/PH2.vue"

const name = ref<string>("SortAndPage")
const sort = ref<string>('')
const page = ref<string>('')
const desc = ref<string>('')
const s = inject("sort")
const p = inject("page")
const d = inject("desc")

watch(sort, () => {
    // @ts-ignore
    s.value = sort.value
})
watch(page, () => {
    // @ts-ignore
    p.value = page.value
})
watch(desc, () => {
    // @ts-ignore
    d.value = desc.value
})
</script>

<style scoped>
.sortAndPage {
    width: 100%;
    height: 80px;
    display: grid;
    grid-template-areas: "ph1 sort page desc ph2";
    grid-template-columns: 225px 120px 120px 120px 1fr;
    position: fixed;
    top: 60px;
    background-color: white;
}

.ph1 {
    grid-area: ph1;
}

.ph2 {
    grid-area: ph2;
}

.block {
    width: 80px;
    height: 40px;
    margin-top: auto;
    margin-bottom: auto;
    margin-right: 20px;
    text-align: center;
}

.sort {
    grid-area: sort;
}

.page {
    grid-area: page;
}

.desc {
    grid-area: desc;
}

.select {
    width: auto;
    height: 100%;
    border: none;
    box-sizing: border-box;
    border-radius: 18px;
    padding: 0 8px;
    font-size: 1rem;
    font-weight: bolder;
    text-align: -webkit-center;
    text-align-last: center;
    appearance: none;
    -webkit-appearance: none;
    -moz-appearance: none;
    outline: none;
    background-color: rgba(0,0,0,0.1);
}

.select:hover {
    background-color: rgba(0,0,0,0.15);
}

.select:active {
    background-color: rgba(0,0,0,0.2);
}
</style>