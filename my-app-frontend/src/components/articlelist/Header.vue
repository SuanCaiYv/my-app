<template>
    <div class="header">
        <HomePage></HomePage>
        <PH1></PH1>
        <Publish></Publish>
        <Draft></Draft>
        <Manager></Manager>
        <Search :search-func="click"></Search>
        <PH2></PH2>
        <SignButton></SignButton>
    </div>
</template>

<script setup lang="ts">
import {inject, ref} from "vue"
import HomePage from "../common/HomePage.vue"
import PH1 from "../placeholder/PH1.vue"
import Search from "../common/Search.vue"
import PH2 from "../placeholder/PH2.vue"
import SignButton from "../common/SignButton.vue";
import Publish from "./Publish.vue";
import Draft from "./Draft.vue";
import Manager from "./Manager.vue";

const name = ref<String>("Header")
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
    grid-template-areas: "home-page ph1 publish draft manager ph2 search ph3 sign-button";
    grid-auto-columns: 200px 25px 140px 140px 140px 1fr 200px 25px 200px;
    position: fixed;
    background-color: white;
}
</style>