<template>
    <div class="sortAndPage">
        <PH1 class="ph1"></PH1>
        <div class="sort select">
            <select class="select-input">
                <option selected value="release_time">排序方式</option>
                <option value="release_time">时间</option>
                <option value="tag">标签</option>
                <option value="kind">分类</option>
                <option value="name">名称</option>
            </select>
        </div>
        <div class="page select">
            <select class="select-input">
                <option selected value="10">分页数量</option>
                <option value="5">5</option>
                <option value="10">10</option>
                <option value="15">15</option>
                <option value="20">20</option>
                <option value="50">50</option>
            </select>
        </div>
        <div class="desc select">
            <select class="select-input">
                <option selected value="true">显示顺序</option>
                <option value="true">正序</option>
                <option value="false">倒序</option>
            </select>
        </div>
        <PH2 class="ph2"></PH2>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import PH1 from "../../placeholder/PH1.vue"
import PH2 from "../../placeholder/PH2.vue"

const name = ref<String>("SortAndPage")
</script>

<style scoped>
.sortAndPage {
    width: 100%;
    height: 100%;
    grid-area: sortAndPage;
    /*border: 1px solid lightgrey;*/
    /*box-sizing: border-box;*/
    display: grid;
    grid-template-areas: "ph1 sort page desc ph2";
    grid-template-columns: 300px 120px 120px 120px 1fr;
}

.ph1 {
    grid-area: ph1;
}

.ph2 {
    grid-area: ph2;
}

.select {
    width: 80px;
    height: 40px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    margin: auto 40px 0 0;
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

.select-input {
    width: 100%;
    height: 100%;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    border: none;
    border-radius: 18px;
    font-size: 0.8rem;
    text-align: center;
    background-color: whitesmoke;
}
</style>