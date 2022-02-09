<template>
    <div class="article">
        <div class="title">
            {{title}}
        </div>
        <div class="content">
            {{body}}
        </div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Article")

let props = defineProps({
    title: String,
    body: String
})
</script>

<style scoped>
.article {
}

.title {
    width: 320px;
    height: 40px;
    margin: 0 auto 0 25px;
    /*border: 1px solid black;*/
    /*box-sizing: border-box;*/
    border-radius: 16px;
    text-align: left;
    font-size: 1.4rem;
    line-height: 40px;
    background-color: rgba(255,76,0,0.05);
}

.content {
    height: 200px;
    width: auto;
    margin: 0 25px 20px;
    /*border: 1px solid black;*/
    /*box-sizing: border-box;*/
    border-radius: 16px;
    text-align: left;
    font-size: 1.2rem;
    background-color: rgba(255,109,0,0.05);
}
</style>