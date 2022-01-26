<template>
    <div class="body">
        <PH1></PH1>
        <Main></Main>
        <PH2></PH2>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import PH1 from "../placeholder/PH1.vue";
import Main from "./Main.vue";
import PH2 from "../placeholder/PH2.vue";

const name = ref<String>("Body")
</script>

<style scoped>
.body {
    width: 100%;
    height: 100%;
    grid-area: body;
    border: 1px solid silver;
    box-sizing: border-box;
    display: grid;
    grid-template-areas: "ph1 main ph2";
    grid-auto-columns: 1fr 1000px 1fr;
}
</style>