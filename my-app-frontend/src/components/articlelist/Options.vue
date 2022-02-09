<template>
    <div class="options">
        <PH1></PH1>
        <div class="publish">
            <button class="button-publish" @click="router.push('/more_article')">发布文章</button>
        </div>
        <div class="manager">
            <button class="button-manager" @click="router.push('/state_manager')">其他管理</button>
        </div>
        <div class="collector">
            <button class="button-collector" @click="alertFunc('暂未开通服务', function() {})">回收站</button>
        </div>
        <div class="page-limit">
            <select class="select-page-limit">
                <option selected value="10">分页数量</option>
                <option value="5">5</option>
                <option value="10">10</option>
                <option value="15">15</option>
                <option value="20">20</option>
                <option value="50">50</option>
            </select>
        </div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import PH1 from "../placeholder/PH1.vue"
import {useRouter} from "vue-router";
import alertFunc from "../../util/alert";

const name = ref<String>("Option")

const router = useRouter()
</script>

<style scoped>
.options {
    width: 200px;
    height: 800px;
    display: grid;
    grid-template-areas:
        "ph1"
        "publish"
        "manager"
        "collector"
        "page-limit";
    grid-template-rows: 80px 60px 60px 60px 60px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    position: fixed;
    top: 50px;
}

.publish {
    width: 100%;
    height: 100%;
    grid-area: publish;
}

.manager {
    width: 100%;
    height: 100%;
    grid-area: manager;
}

.collector {
    width: 100%;
    height: 100%;
    grid-area: collector;
}

.page-limit {
    width: 100%;
    height: 100%;
    grid-area: page-limit;
}

.button-publish {
    width: 120px;
    height: 50px;
    border: none;
    border-radius: 18px;
    font-size: 1.2rem;
    margin-top: 5px;
    margin-bottom: 5px;
    background-color: lavender;
}

.button-manager {
    width: 120px;
    height: 50px;
    border: none;
    border-radius: 18px;
    font-size: 1.2rem;
    margin-top: 5px;
    margin-bottom: 5px;
    background-color: lightpink;
}

.button-collector {
    width: 120px;
    height: 50px;
    border: none;
    border-radius: 18px;
    font-size: 1.2rem;
    margin-top: 5px;
    margin-bottom: 5px;
    background-color: darkseagreen;
}

.select-page-limit {
    width: 120px;
    height: 50px;
    border: none;
    border-radius: 18px;
    font-size: 1.2rem;
    margin-top: 5px;
    margin-bottom: 5px;
    appearance: none;
    -webkit-appearance: none;
    -moz-appearance: none;
    outline: none;
    padding-left: 18px;
    background-color: bisque;
}

.select-show {
}
</style>