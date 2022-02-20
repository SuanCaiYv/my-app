<template>
    <div class="options">
        <div class="sort">
            <select class="select" v-model="sort">
                <option disabled value="">排序方式</option>
                <option value="release_time">时间</option>
                <option value="kind.name">分类</option>
                <option value="name">名称</option>
            </select>
        </div>
        <div class="page">
            <select class="select" v-model="page">
                <option disabled value="">分页数量</option>
                <option value="5">5</option>
                <option value="10">10</option>
                <option value="15">15</option>
                <option value="20">20</option>
                <option value="50">50</option>
            </select>
        </div>
        <div class="desc">
            <select class="select" v-model="desc">
                <option disabled value="">显示顺序</option>
                <option value="true">正序</option>
                <option value="false">倒序</option>
            </select>
        </div>
    </div>
</template>

<script setup lang="ts">
import {inject, ref, watch} from "vue"

const name = ref<string>("Today")
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
.options {
    width: 200px;
    height: 100%;
    position: fixed;
    top: 60px;
    padding-top: 120px;
}

.page {
    width: 100%;
    height: 80px;
}

.sort {
    width: 100%;
    height: 80px;
}

.desc {
    width: 100%;
    height: 80px;
}

.select {
    width: auto;
    height: 40px;
    border: none;
    box-sizing: border-box;
    border-radius: 18px;
    margin-top: 20px;
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