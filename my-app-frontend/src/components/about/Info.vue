<template>
    <div class="info">
        <div class="statistic"></div>
        <div class="brief">
            <div>
                <div class="name article-list" @click="router.push('/article_list')">文章总数</div>
                <div class="value"></div>
            </div>
            <div>
                <div class="name">上次更新</div>
                <div class="value"></div>
            </div>
            <div>
                <div class="name">最多标签</div>
                <div class="value"></div>
            </div>
            <div>
                <div class="name">最新文章</div>
                <div class="value"></div>
            </div>
        </div>
        <div class="third-party"></div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {useRouter} from "vue-router";

const name = ref<String>("Info")

const router = useRouter()
</script>

<style scoped>
.info {
    width: 100%;
    height: 100%;
    grid-area: info;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    display: grid;
    grid-template-areas: "statistic brief social";
    grid-template-columns: 530px 660px 1fr;
}

.statistic {
    width: 100%;
    height: 300px;
    grid-area: statistic;
    border: 2px solid lightsalmon;
    box-sizing: border-box;
    border-radius: 20px;
}

.brief {
    width: 100%;
    height: 300px;
    grid-area: brief;
    border: 2px solid salmon;
    box-sizing: border-box;
    border-radius: 20px;
}

.third-party {
    width: 100%;
    height: 300px;
    grid-area: social;
    border: 2px solid slateblue;
    box-sizing: border-box;
    border-radius: 20px;
}

.name {
    width: 100px;
    height: 50px;
    margin-left: 20px;
    margin-top: 20px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    border: none;
    border-radius: 18px;
    display: inline-block;
    vertical-align: bottom;
    font-size: 1.2rem;
    text-align: center;
    line-height: 50px;
    background-color: rgba(255,140,0,0.05);
}

.value {
    width: 510px;
    height: 50px;
    margin-top: 20px;
    margin-left: 26px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    border: none;
    border-radius: 16px;
    display: inline-block;
    vertical-align: bottom;
    font-size: 1.2rem;
    line-height: 50px;
    background-color: rgba(255,0,188,0.05);
}

.article-list {
}

.article-list:hover {
    opacity: 60%;
}

.article-list:active {
    opacity: 90%;
}
</style>