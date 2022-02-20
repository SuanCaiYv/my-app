<template>
    <div class="header">
        <HomePage></HomePage>
        <PH1></PH1>
        <Search :search-func="click"></Search>
        <PH2></PH2>
        <SignButton></SignButton>
    </div>
</template>

<script setup lang="ts">
import {inject, ref} from "vue"
import HomePage from "./HomePage.vue"
import PH1 from "../placeholder/PH1.vue"
import Search from "./Search.vue"
import PH2 from "../placeholder/PH2.vue"
import SignButton from "./SignButton.vue";

const name = ref<string>("HeaderWithSearch")
const searchKey = inject("searchKey")
const click = function (key: string) {
    // @ts-ignore
    searchKey.value = key
}
</script>

<style scoped>
.header {
    width: 100%;
    height: 60px;
    display: grid;
    grid-template-areas: "home-page ph1 search ph2 sign-button";
    grid-auto-columns: 200px 1fr 200px 25px 200px;
    position: fixed;
    background-color: white;
}
</style>