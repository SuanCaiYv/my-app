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
    /*border: 1px solid black;*/
    /*box-sizing: border-box;*/
    border-radius: 16px;
    background-color: darkorange;
    opacity: 10%;
}

.content {
    width: 100%;
    height: 100%;
    grid-area: content;
    box-sizing: border-box;
    border-radius: 16px;
    display: inline-block;
    background-color: orange;
    opacity: 10%;
}

.update {
    width: 100%;
    height: 100%;
    grid-area: update;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    display: inline-block;
}

.delete {
    width: 100%;
    height: 100%;
    grid-area: delete;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    display: inline-block;
}

.visible {
    width: 100%;
    height: 100%;
    grid-area: visible;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    display: inline-block;
}

.button {
    width: 80px;
    height: 40px;
    margin-top: 20px;
    margin-bottom: 20px;
    border: none;
    border-radius: 18px;
    font-size: 1.2rem;
}
</style>