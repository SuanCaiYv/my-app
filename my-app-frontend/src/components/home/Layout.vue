<template>
    <div class="layout">
        <HeaderWithSearch></HeaderWithSearch>
        <PH1></PH1>
        <Body></Body>
        <Footer></Footer>
    </div>
</template>

<script setup lang="ts">
import {provide, ref} from "vue"
import HeaderWithSearch from "../common/HeaderWithSearch.vue"
import PH1 from "../placeholder/PH1.vue"
import Body from "./Body.vue"
import Footer from "./Footer.vue"

const name = ref<string>("Layout")
provide("searchKey", ref<string>(''))
</script>

<style scoped>
.layout {
    width: 100%;
    height: 100%;
    display: grid;
    grid-template-areas:
        "ph1"
        "body"
        "footer";
    grid-template-rows: 60px 1fr 50px;
}
</style>