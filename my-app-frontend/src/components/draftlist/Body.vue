<template>
    <div class="body">
        <PH1></PH1>
        <PH2></PH2>
        <DraftList></DraftList>
        <PH3></PH3>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import DraftList from "./DraftList.vue";
import PH1 from "../placeholder/PH1.vue";
import PH2 from "../placeholder/PH2.vue";
import PH3 from "../placeholder/PH3.vue";

const name = ref<string>("Body")
</script>

<style scoped>
.body {
    width: 100%;
    height: 100%;
    grid-area: body;
    display: grid;
    grid-template-areas:
        "ph1 ph1 ph1"
        "ph2 articleList ph3";
    grid-auto-columns: 200px 1fr 200px;
    grid-template-rows: 80px 1fr;
}
</style>