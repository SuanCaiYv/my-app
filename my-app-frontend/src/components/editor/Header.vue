<template>
    <div class="header">
        <HomePage></HomePage>
        <PH1></PH1>
        <PublishButton></PublishButton>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import HomePage from "../common/HomePage.vue"
import PH1 from "../placeholder/PH1.vue"
import PublishButton from "./PublishButton.vue"

const name = ref<String>("Header")
</script>

<style scoped>
.header {
    width: 100%;
    height: 60px;
    grid-area: header;
    display: grid;
    grid-template-areas: "home-page ph1 publish-button";
    grid-auto-columns: 200px 1fr 200px;
    position: fixed;
    background-color: white;
}
</style>